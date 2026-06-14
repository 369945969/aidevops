package agent

import (
	"fmt"
	"time"
)

// ============================================================
// BaseAgent — interface for all DevOps pipeline agents
// ============================================================

// DevOpsAgent is the Go interface mirroring BaseAgent(ABC) from
// daily_stock_analysis/src/agent/agents/base_agent.py.
//
// Every agent in the pipeline implements this interface.
// The orchestrator calls SystemPrompt() and BuildUserMessage() to
// construct the LLM request, then PostProcess() to extract structured
// output from the raw LLM response.
type DevOpsAgent interface {
	// AgentName returns the unique identifier for this agent.
	AgentName() string

	// MaxSteps returns the maximum LLM round-trips for this agent.
	MaxSteps() int

	// SystemPrompt builds the system prompt for this agent given the context.
	SystemPrompt(ctx *DevOpsContext) string

	// BuildUserMessage constructs the user message for the LLM.
	BuildUserMessage(ctx *DevOpsContext) string

	// PostProcess transforms the raw LLM text into a structured AgentOutput.
	// Returns nil if the agent doesn't produce structured output.
	PostProcess(ctx *DevOpsContext, rawText string) *AgentOutput

	// Run executes this agent and returns a StageResult.
	// The orchestrator calls this for each agent in the pipeline chain.
	Run(ctx *DevOpsContext, adapter *LLMAdapter) StageResult
}

// ============================================================
// RunAgent — shared execution logic for all agents
// ============================================================

// RunAgent executes a single DevOpsAgent using the shared pattern:
// 1. Build system prompt + user message
// 2. Call LLM adapter
// 3. Post-process into AgentOutput
// 4. Append output to context
// 5. Return StageResult
//
// This mirrors BaseAgent.run() from the Python reference project.
func RunAgent(agent DevOpsAgent, ctx *DevOpsContext, adapter *LLMAdapter) StageResult {
	t0 := time.Now()
	result := StageResult{
		StageName: agent.AgentName(),
		Status:    StageRunning,
		Meta:      make(map[string]interface{}),
	}

	systemPrompt := agent.SystemPrompt(ctx)
	userMessage := agent.BuildUserMessage(ctx)

	// Inject pre-fetched data from context into the user message,
	// mirroring BaseAgent._inject_cached_data() from the Python project.
	injectedData := buildInjectedDataSection(ctx)
	if injectedData != "" {
		userMessage = injectedData + "\n\n" + userMessage
	}

	// Call LLM
	content, tokens, err := adapter.Call(systemPrompt, userMessage)
	if err != nil {
		result.Status = StageFailed
		result.Error = err.Error()
		result.DurationMs = int64(time.Since(t0).Milliseconds())
		return result
	}

	result.Meta["raw_text"] = content
	result.Meta["models_used"] = []string{adapter.Model}
	result.TokensUsed = tokens

	// Post-process into structured output
	output := agent.PostProcess(ctx, content)
	if output != nil {
		output.AgentName = agent.AgentName()
		ctx.AddOutput(*output)
		result.Output = output

		// Persist key artifacts in the context for downstream agents
		persistArtifacts(ctx, output)
	}

	result.Status = StageCompleted
	result.DurationMs = int64(time.Since(t0).Milliseconds())
	return result
}

// persistArtifacts maps agent outputs to the appropriate context fields,
// mirroring how TechnicalAgent writes to ctx.data["realtime_quote"] etc.
func persistArtifacts(ctx *DevOpsContext, output *AgentOutput) {
	switch output.AgentName {
	case "requirement_analyst":
		ctx.Specification = output.Content
	case "architect":
		ctx.Architecture = output.Content
	case "developer":
		if rawData, ok := output.RawData["files"].(map[string]interface{}); ok {
			for k, v := range rawData {
				if s, ok := v.(string); ok {
					ctx.Code[k] = s
				}
			}
		}
		ctx.SetData("developer_output", output.Content)
	case "tester":
		ctx.TestResults = output.Content
		ctx.SetData("test_plan", output.RawData["test_plan"])
	case "deployer":
		ctx.DeploymentPlan = output.Content
	}
}

// buildInjectedDataSection serializes accumulated context data for
// injection into the user message, mirroring BaseAgent._inject_cached_data().
func buildInjectedDataSection(ctx *DevOpsContext) string {
	var parts []string

	if ctx.Specification != "" {
		parts = append(parts, "[Pre-fetched: 需求分析结果]\n"+ctx.Specification)
	}
	if ctx.Architecture != "" {
		parts = append(parts, "[Pre-fetched: 架构设计方案]\n"+ctx.Architecture)
	}
	if len(ctx.Code) > 0 {
		codeSummary := "[Pre-fetched: 已生成代码文件]\n"
		for filename := range ctx.Code {
			codeSummary += "- " + filename + "\n"
		}
		parts = append(parts, codeSummary)
	}
	if ctx.TestResults != "" {
		parts = append(parts, "[Pre-fetched: 测试结果]\n"+ctx.TestResults)
	}

	// Add prior agent opinions (mirroring ctx.opinions injection)
	if len(ctx.Outputs) > 0 {
		parts = append(parts, "[Pre-fetched: 前序Agent输出摘要]")
		for _, op := range ctx.Outputs {
			parts = append(parts, "  Agent: "+op.AgentName+" | Signal: "+op.Signal+" | Confidence: "+fmt.Sprintf("%.2f", op.Confidence))
			if op.Reasoning != "" {
				reasoning := op.Reasoning
				if len(reasoning) > 500 {
					reasoning = reasoning[:500] + "...(truncated)"
				}
				parts = append(parts, "  Reasoning: "+reasoning)
			}
		}
	}

	// Add risk flags
	if ctx.HasRiskFlags() {
		parts = append(parts, "[风险标记]")
		for _, flag := range ctx.RiskFlags {
			parts = append(parts, "  ["+flag.Severity+"] "+flag.Category+": "+flag.Description)
		}
	}

	result := ""
	for i, part := range parts {
		if i > 0 {
			result += "\n\n"
		}
		result += part
	}
	return result
}
