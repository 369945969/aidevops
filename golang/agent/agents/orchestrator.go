package agents

import (
	"devops/agent"
	"log"
	"time"
)

// ============================================================
// DevOpsOrchestrator — multi-agent pipeline coordinator
// ============================================================

// DevOpsOrchestrator manages the lifecycle of specialised agents for
// a single DevOps pipeline run. Mirrors AgentOrchestrator from
// daily_stock_analysis/src/agent/orchestrator.py.
//
// Pipeline stages:
// - quick:    RequirementAnalyst → Developer
// - standard: RequirementAnalyst → Architect → Developer
// - full:     RequirementAnalyst → Architect → Developer → Tester → Deployer
// - review:   Full + human review nodes at architecture, code, deploy
//
// The orchestrator:
// 1. Seeds a DevOpsContext with the user query
// 2. Runs agents sequentially, passing the shared context
// 3. Collects StageResult from each agent
// 4. Produces a unified agent.OrchestratorResult
type DevOpsOrchestrator struct {
	LLMAdapter *agent.LLMAdapter
	Mode       agent.PipelineMode
	TimeoutMs  int64 // pipeline timeout in ms, 0 = disabled
	DB         interface{} // GORM DB for persisting pipeline state
}

// NewDevOpsOrchestrator creates an orchestrator with the given configuration.
func NewDevOpsOrchestrator(adapter *agent.LLMAdapter, mode agent.PipelineMode, timeoutMs int64) *DevOpsOrchestrator {
	if mode == "" {
		mode = agent.ModeFull
	}
	return &DevOpsOrchestrator{
		LLMAdapter: adapter,
		Mode:       mode,
		TimeoutMs:  timeoutMs,
	}
}

// Run executes the multi-agent pipeline for a given requirement.
// Returns an agent.OrchestratorResult with the final assembled output.
func (o *DevOpsOrchestrator) Run(query string, pipelineID uint) agent.OrchestratorResult {
	ctx := agent.NewDevOpsContext(query)
	ctx.PipelineID = pipelineID
	return o.ExecutePipeline(ctx)
}

// ExecutePipeline runs the agent pipeline according to the configured mode.
// Mirrors AgentOrchestrator._execute_pipeline() from the Python project.
func (o *DevOpsOrchestrator) ExecutePipeline(ctx *agent.DevOpsContext) agent.OrchestratorResult {
	stats := &agent.AgentRunStats{}
	allToolCalls := []map[string]interface{}{}
	modelsUsed := []string{}
	t0 := time.Now()

	agents := o.BuildAgentChain()

	for i := 0; i < len(agents); i++ {
		ag := agents[i]
		elapsedMs := time.Since(t0).Milliseconds()

		// Timeout/budget checks (mirroring Python orchestrator)
		if o.TimeoutMs > 0 && elapsedMs >= o.TimeoutMs {
			log.Printf("[Orchestrator] pipeline timed out before stage '%s'", ag.AgentName())
			return agent.OrchestratorResult{
				Success:      false,
				Error:        "Pipeline timed out",
				Stats:        stats,
				TotalTokens:  stats.TotalTokens,
				ToolCallsLog: allToolCalls,
			}
		}

		// Minimum budget for a stage to do useful work (mirroring _MIN_STAGE_BUDGET_S)
		minStageBudgetMs := int64(15000) // 15 seconds
		remainingMs := o.TimeoutMs - elapsedMs
		if o.TimeoutMs > 0 && i > 0 && remainingMs < minStageBudgetMs {
			log.Printf("[Orchestrator] insufficient budget before stage '%s'", ag.AgentName())
			return agent.OrchestratorResult{
				Success:      false,
				Error:        "Pipeline budget insufficient",
				Stats:        stats,
				TotalTokens:  stats.TotalTokens,
				ToolCallsLog: allToolCalls,
			}
		}

		log.Printf("[Orchestrator] starting stage: %s", ag.AgentName())

		// Execute the agent
		result := agent.RunAgent(ag, ctx, o.LLMAdapter)
		stats.RecordStage(result)

		if result.Meta != nil {
			if tcLog, ok := result.Meta["tool_calls_log"].([]map[string]interface{}); ok {
				allToolCalls = append(allToolCalls, tcLog...)
			}
			if models, ok := result.Meta["models_used"].([]string); ok {
				modelsUsed = append(modelsUsed, models...)
			}
		}

		// Handle stage failure (mirroring Python's critical vs non-critical logic)
		if result.Status == agent.StageFailed {
			nonCritical := isNonCriticalStage(ag.AgentName())
			if !nonCritical {
				log.Printf("[Orchestrator] critical stage '%s' failed: %s", ag.AgentName(), result.Error)
				return agent.OrchestratorResult{
					Success:      false,
					Error:        result.Error,
					Stats:        stats,
					TotalTokens:  stats.TotalTokens,
					ToolCallsLog: allToolCalls,
				}
			}
			log.Printf("[Orchestrator] non-critical stage '%s' failed, degrading: %s", ag.AgentName(), result.Error)
		}

		// Check post-stage timeout
		elapsedMs = time.Since(t0).Milliseconds()
		if o.TimeoutMs > 0 && elapsedMs >= o.TimeoutMs {
			log.Printf("[Orchestrator] pipeline timed out after stage '%s'", ag.AgentName())
			// Return partial result with whatever was collected
			return o.BuildPartialResult(ctx, stats, allToolCalls, modelsUsed, elapsedMs)
		}

		log.Printf("[Orchestrator] completed stage: %s (duration: %dms, status: %s)",
			ag.AgentName(), result.DurationMs, result.Status)
	}

	// Assemble final output
	totalDurationMs := time.Since(t0).Milliseconds()
	stats.TotalDurationMs = totalDurationMs
	stats.ModelsUsed = uniqueStrings(modelsUsed)

	pipelineData := o.ResolveFinalOutput(ctx)
	content := o.BuildContentSummary(ctx)

	return agent.OrchestratorResult{
		Success:      content != "",
		Content:      content,
		PipelineData: pipelineData,
		ToolCallsLog: allToolCalls,
		TotalSteps:   stats.TotalStages,
		TotalTokens:  stats.TotalTokens,
		Provider:     firstString(modelsUsed),
		Model:        joinStrings(modelsUsed, ", "),
		Stats:        stats,
	}
}

// BuildAgentChain instantiates the ordered agent list based on mode.
// Mirrors AgentOrchestrator._build_agent_chain() from the Python project.
func (o *DevOpsOrchestrator) BuildAgentChain() []agent.DevOpsAgent {
	analyst := &RequirementAnalystAgent{}
	architect := &ArchitectAgent{}
	developer := &DeveloperAgent{}
	tester := &TesterAgent{}
	deployer := &DeployerAgent{}

	switch o.Mode {
	case agent.ModeQuick:
		return []agent.DevOpsAgent{analyst, developer}
	case agent.ModeStandard:
		return []agent.DevOpsAgent{analyst, architect, developer}
	case agent.ModeFull:
		return []agent.DevOpsAgent{analyst, architect, developer, tester, deployer}
	case agent.ModeReview:
		// Review mode adds human review pause points — for now same chain as full
		// The review nodes will be handled via API callbacks (pause/resume)
		return []agent.DevOpsAgent{analyst, architect, developer, tester, deployer}
	default:
		return []agent.DevOpsAgent{analyst, architect, developer, tester, deployer}
	}
}

// ResolveFinalOutput assembles the structured pipeline output from context.
// Mirrors AgentOrchestrator._resolve_final_output() from the Python project.
func (o *DevOpsOrchestrator) ResolveFinalOutput(ctx *agent.DevOpsContext) map[string]interface{} {
	data := map[string]interface{}{
		"query":         ctx.Query,
		"pipeline_id":   ctx.PipelineID,
		"specification": ctx.Specification,
		"architecture":  ctx.Architecture,
		"test_results":  ctx.TestResults,
		"deploy_plan":   ctx.DeploymentPlan,
	}

	if len(ctx.Code) > 0 {
		data["code_files"] = ctx.Code
	}

	// Collect agent outputs
	outputs := make([]map[string]interface{}, 0, len(ctx.Outputs))
	for _, op := range ctx.Outputs {
		outputs = append(outputs, map[string]interface{}{
			"agent_name":  op.AgentName,
			"signal":      op.Signal,
			"confidence":  op.Confidence,
			"reasoning":   op.Reasoning,
		})
	}
	data["agent_outputs"] = outputs

	// Collect risk flags
	if ctx.HasRiskFlags() {
		flags := make([]map[string]interface{}, 0, len(ctx.RiskFlags))
		for _, flag := range ctx.RiskFlags {
			flags = append(flags, map[string]interface{}{
				"category":    flag.Category,
				"description": flag.Description,
				"severity":    flag.Severity,
			})
		}
		data["risk_flags"] = flags
	}

	return data
}

// BuildContentSummary creates a plaintext summary of the pipeline run.
func (o *DevOpsOrchestrator) BuildContentSummary(ctx *agent.DevOpsContext) string {
	if ctx.Specification == "" && ctx.Architecture == "" && len(ctx.Code) == 0 {
		return ""
	}

	summary := "=== DevOps Pipeline Result ===\n\n"
	summary += "原始需求: " + ctx.Query + "\n\n"

	if ctx.Specification != "" {
		summary += "## 需求分析\n" + truncate(ctx.Specification, 2000) + "\n\n"
	}
	if ctx.Architecture != "" {
		summary += "## 架构设计\n" + truncate(ctx.Architecture, 2000) + "\n\n"
	}
	if len(ctx.Code) > 0 {
		summary += "## 代码产出\n"
		for filename := range ctx.Code {
			summary += "  - " + filename + "\n"
		}
		summary += "\n"
	}
	if ctx.TestResults != "" {
		summary += "## 测试结果\n" + truncate(ctx.TestResults, 1000) + "\n\n"
	}
	if ctx.DeploymentPlan != "" {
		summary += "## 部署计划\n" + truncate(ctx.DeploymentPlan, 1000) + "\n"
	}

	return summary
}

// BuildPartialResult creates a result from whatever was collected before timeout.
func (o *DevOpsOrchestrator) BuildPartialResult(
	ctx *agent.DevOpsContext,
	stats *agent.AgentRunStats,
	toolCalls []map[string]interface{},
	modelsUsed []string,
	elapsedMs int64,
) agent.OrchestratorResult {
	stats.TotalDurationMs = elapsedMs
	stats.ModelsUsed = uniqueStrings(modelsUsed)

	pipelineData := o.ResolveFinalOutput(ctx)
	content := "[降级结果] 管道超时，以下内容基于已完成阶段自动降级生成。\n\n" + o.BuildContentSummary(ctx)

	return agent.OrchestratorResult{
		Success:      content != "",
		Content:      content,
		PipelineData: pipelineData,
		Error:        "Pipeline timed out",
		Stats:        stats,
		TotalTokens:  stats.TotalTokens,
		ToolCallsLog: toolCalls,
		Provider:     firstString(modelsUsed),
		Model:        joinStrings(modelsUsed, ", "),
	}
}

// ============================================================
// Helpers
// ============================================================

// isNonCriticalStage returns true if the stage can degrade gracefully.
// Mirrors Python: agent.agent_name in ("intel", "risk") or is_skill_agent.
func isNonCriticalStage(agentName string) bool {
	nonCritical := map[string]bool{
		"tester":  true,
		"deployer": true,
	}
	return nonCritical[agentName]
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...(truncated)"
}

func uniqueStrings(ss []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, s := range ss {
		if !seen[s] && s != "" {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}

func firstString(ss []string) string {
	if len(ss) > 0 {
		return ss[0]
	}
	return ""
}

func joinStrings(ss []string, sep string) string {
	unique := uniqueStrings(ss)
	result := ""
	for i, s := range unique {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
