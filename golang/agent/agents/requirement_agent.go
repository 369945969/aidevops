package agents

import (
	"encoding/json"
	"fmt"

	"devops/agent"
)

// ============================================================
// RequirementAnalystAgent
// ============================================================

// RequirementAnalystAgent receives a single-sentence requirement and
// produces a structured specification document. Mirrors TechnicalAgent
// from the Python project but adapted for DevOps requirements.
type RequirementAnalystAgent struct{}

func (a *RequirementAnalystAgent) AgentName() string { return "requirement_analyst" }
func (a *RequirementAnalystAgent) MaxSteps() int     { return 1 }
func (a *RequirementAnalystAgent) Run(ctx *agent.DevOpsContext, adapter *agent.LLMAdapter) agent.StageResult {
	return agent.RunAgent(a, ctx, adapter)
}

func (a *RequirementAnalystAgent) SystemPrompt(ctx *agent.DevOpsContext) string {
	return `你是一位专业的需求分析师Agent，负责将用户的一句话需求转化为结构化的需求规格文档。

## 工作流程

1. 仔细分析用户的原始需求描述
2. 澄清需求的边界条件、目标用户、功能范围
3. 识别潜在的技术难点和风险
4. 输出结构化的需求规格文档

## 输出格式

你的最终响应必须是以下结构的有效 JSON 对象（直接输出JSON，不要使用 markdown 代码块包裹）：

{
  "title": "需求标题（简洁明确）",
  "priority": "高|中|低",
  "description": "需求详细描述（200字以内）",
  "user_scenarios": [
    {
      "persona": "用户角色",
      "action": "用户操作",
      "goal": "用户目标"
    }
  ],
  "functional_requirements": [
    {
      "id": "FR-001",
      "title": "功能标题",
      "description": "功能描述",
      "acceptance_criteria": ["验收条件1", "验收条件2"]
    }
  ],
  "non_functional_requirements": [
    {
      "id": "NFR-001",
      "category": "性能|安全|可用性|兼容性",
      "description": "非功能需求描述",
      "metric": "衡量标准"
    }
  ],
  "assumptions": ["假设1", "假设2"],
  "risks": [
    {
      "description": "风险描述",
      "severity": "高|中|低",
      "mitigation": "缓解措施"
    }
  ],
  "dependencies": ["依赖1", "依赖2"],
  "signal": "approve|needs_revision|block",
  "confidence": 0.0-1.0,
  "reasoning": "分析推理过程"
}

## 规则

1. **不编造需求** — 只基于用户提供的描述进行分析和扩展
2. **边界清晰** — 明确标注需求范围内和范围外的内容
3. **验收标准可测试** — 每个功能需求必须有可验证的验收条件
4. **风险优先** — 必须识别潜在的技术风险和业务风险
5. **输出JSON** — 最终响应必须是有效的JSON对象`
}

func (a *RequirementAnalystAgent) BuildUserMessage(ctx *agent.DevOpsContext) string {
	return fmt.Sprintf(`请分析以下需求并生成结构化的需求规格文档：

原始需求："%s"

请输出完整的JSON需求规格文档。`, ctx.Query)
}

func (a *RequirementAnalystAgent) PostProcess(ctx *agent.DevOpsContext, rawText string) *agent.AgentOutput {
	var spec map[string]interface{}
	if err := json.Unmarshal([]byte(rawText), &spec); err != nil {
		// Try extracting JSON from markdown
		extracted := extractJSON(rawText)
		if extracted != nil {
			if err2 := json.Unmarshal(extracted, &spec); err2 != nil {
				// Fallback: treat entire text as specification
				return &agent.AgentOutput{
					AgentName:  a.AgentName(),
					Signal:     "approve",
					Confidence: 0.6,
					Reasoning:  "无法解析为结构化JSON，使用原始文本作为需求规格",
					Content:    rawText,
					RawData:    map[string]interface{}{"raw": rawText},
				}
			}
		} else {
			return &agent.AgentOutput{
				AgentName:  a.AgentName(),
				Signal:     "approve",
				Confidence: 0.6,
				Reasoning:  "无法解析为结构化JSON，使用原始文本作为需求规格",
				Content:    rawText,
				RawData:    map[string]interface{}{"raw": rawText},
			}
		}
	}

	signal := getString(spec, "signal", "approve")
	confidence := getFloat(spec, "confidence", 0.7)
	reasoning := getString(spec, "reasoning", "")
	title := getString(spec, "title", ctx.Query)
	priority := getString(spec, "priority", "中")

	// Re-serialize the cleaned spec as the content
	contentBytes, _ := json.MarshalIndent(spec, "", "  ")
	content := string(contentBytes)

	// Add risk flags to context
	if risks, ok := spec["risks"].([]interface{}); ok {
		for _, r := range risks {
			if riskMap, ok := r.(map[string]interface{}); ok {
				severity := getString(riskMap, "severity", "medium")
				description := getString(riskMap, "description", "")
				ctx.AddRiskFlag("requirement_risk", description, severity, a.AgentName())
			}
		}
	}

	return &agent.AgentOutput{
		AgentName:  a.AgentName(),
		Signal:     signal,
		Confidence: confidence,
		Reasoning:  reasoning,
		Content:    content,
		RawData: map[string]interface{}{
			"title":    title,
			"priority": priority,
			"spec":     spec,
		},
		KeyLevels: map[string]string{
			"priority": priority,
			"scope":    getString(spec, "scope", "standard"),
		},
	}
}

// ============================================================
// Helper functions (shared across agents)
// ============================================================

func extractJSON(content string) []byte {
	// Look for ```json ... ``` blocks
	jsonTag := "```json"
	plainTag := "```"

	for i := 0; i <= len(content)-len(jsonTag); i++ {
		if content[i:i+len(jsonTag)] == jsonTag {
			start := i + len(jsonTag)
			for j := start; j <= len(content)-len(plainTag); j++ {
				if content[j:j+len(plainTag)] == plainTag {
					return []byte(content[start:j])
				}
			}
		}
	}

	// Look for raw JSON { ... }
	startIdx := -1
	endIdx := -1
	for i := 0; i < len(content); i++ {
		if content[i] == '{' && startIdx == -1 {
			startIdx = i
		}
		if content[i] == '}' {
			endIdx = i
		}
	}
	if startIdx >= 0 && endIdx > startIdx {
		return []byte(content[startIdx : endIdx+1])
	}

	return nil
}

func getString(m map[string]interface{}, key, defaultVal string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return defaultVal
}

func getFloat(m map[string]interface{}, key string, defaultVal float64) float64 {
	if v, ok := m[key]; ok {
		switch n := v.(type) {
		case float64:
			return n
		case int:
			return float64(n)
		}
	}
	return defaultVal
}

func getInt(m map[string]interface{}, key string, defaultVal int) int {
	if v, ok := m[key]; ok {
		switch n := v.(type) {
		case int:
			return n
		case float64:
			return int(n)
		}
	}
	return defaultVal
}
