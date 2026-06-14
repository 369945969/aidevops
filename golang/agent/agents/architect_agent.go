package agents

import (
	"encoding/json"
	"fmt"

	"devops/agent"
)

// ============================================================
// ArchitectAgent
// ============================================================

// ArchitectAgent reads the structured specification and produces
// a technical architecture design. Mirrors IntelAgent from the
// Python project but adapted for DevOps architecture design.
type ArchitectAgent struct{}

func (a *ArchitectAgent) AgentName() string { return "architect" }
func (a *ArchitectAgent) MaxSteps() int     { return 1 }
func (a *ArchitectAgent) Run(ctx *agent.DevOpsContext, adapter *agent.LLMAdapter) agent.StageResult {
	return agent.RunAgent(a, ctx, adapter)
}

func (a *ArchitectAgent) SystemPrompt(ctx *agent.DevOpsContext) string {
	return `你是一位专业的技术架构师Agent，负责将需求规格转化为可执行的技术架构方案。

## 工作流程

1. 阅读需求规格文档，理解功能和非功能需求
2. 选择合适的技术栈和架构模式
3. 设计模块划分、接口定义和数据模型
4. 识别架构风险和性能瓶颈
5. 输出结构化的架构设计方案

## 输出格式

你的最终响应必须是以下结构的有效 JSON 对象（直接输出JSON，不要使用 markdown 代码块包裹）：

{
  "tech_stack": {
    "frontend": "技术选择及理由",
    "backend": "技术选择及理由",
    "database": "技术选择及理由",
    "infrastructure": "技术选择及理由"
  },
  "architecture_pattern": "架构模式（如微服务、单体、分层等）",
  "modules": [
    {
      "name": "模块名称",
      "responsibility": "模块职责",
      "dependencies": ["依赖的其他模块"],
      "interfaces": [
        {
          "name": "接口名称",
          "method": "HTTP方法",
          "path": "路径",
          "input": "输入参数描述",
          "output": "输出参数描述"
        }
      ]
    }
  ],
  "data_models": [
    {
      "name": "模型名称",
      "fields": [
        {"name": "字段名", "type": "数据类型", "description": "字段描述"}
      ]
    }
  ],
  "task_decomposition": [
    {
      "id": "TASK-001",
      "title": "任务标题",
      "description": "任务描述",
      "assignee": "前端开发|后端开发|测试|DevOps",
      "priority": "高|中|低",
      "estimated_hours": 2,
      "dependencies": ["前置任务ID"]
    }
  ],
  "risks": [
    {
      "description": "架构风险描述",
      "severity": "高|中|低",
      "mitigation": "缓解措施"
    }
  ],
  "signal": "approve|needs_revision|block",
  "confidence": 0.0-1.0,
  "reasoning": "架构设计推理过程"
}

## 规则

1. **需求驱动** — 架构必须完全覆盖需求规格中的每个功能点
2. **模块清晰** — 每个模块职责单一，接口明确
3. **技术栈合理** — 选择主流且适合项目规模的技术栈
4. **任务可分配** — 每个开发任务有明确的负责人和预估时间
5. **风险优先** — 必须标注架构层面的风险点`
}

func (a *ArchitectAgent) BuildUserMessage(ctx *agent.DevOpsContext) string {
	spec := ctx.Specification
	if spec == "" {
		spec = ctx.Query
	}
	return fmt.Sprintf(`请基于以下需求规格文档，设计完整的技术架构方案：

需求规格：
%s

请输出完整的JSON架构设计方案。`, spec)
}

func (a *ArchitectAgent) PostProcess(ctx *agent.DevOpsContext, rawText string) *agent.AgentOutput {
	var arch map[string]interface{}
	if err := json.Unmarshal([]byte(rawText), &arch); err != nil {
		extracted := extractJSON(rawText)
		if extracted != nil {
			if err2 := json.Unmarshal(extracted, &arch); err2 != nil {
				return &agent.AgentOutput{
					AgentName:  a.AgentName(),
					Signal:     "approve",
					Confidence: 0.6,
					Reasoning:  "无法解析为结构化JSON，使用原始文本作为架构方案",
					Content:    rawText,
					RawData:    map[string]interface{}{"raw": rawText},
				}
			}
		} else {
			return &agent.AgentOutput{
				AgentName:  a.AgentName(),
				Signal:     "approve",
				Confidence: 0.6,
				Reasoning:  "无法解析为结构化JSON，使用原始文本作为架构方案",
				Content:    rawText,
				RawData:    map[string]interface{}{"raw": rawText},
			}
		}
	}

	signal := getString(arch, "signal", "approve")
	confidence := getFloat(arch, "confidence", 0.7)
	reasoning := getString(arch, "reasoning", "")

	// Re-serialize as content
	contentBytes, _ := json.MarshalIndent(arch, "", "  ")
	content := string(contentBytes)

	// Extract task decomposition for downstream context
	if tasks, ok := arch["task_decomposition"].([]interface{}); ok {
		ctx.SetData("task_decomposition", tasks)
	}

	// Add risk flags
	if risks, ok := arch["risks"].([]interface{}); ok {
		for _, r := range risks {
			if riskMap, ok := r.(map[string]interface{}); ok {
				severity := getString(riskMap, "severity", "medium")
				description := getString(riskMap, "description", "")
				ctx.AddRiskFlag("architecture_risk", description, severity, a.AgentName())
			}
		}
	}

	return &agent.AgentOutput{
		AgentName:  a.AgentName(),
		Signal:     signal,
		Confidence: confidence,
		Reasoning:  reasoning,
		Content:    content,
		RawData:    arch,
		KeyLevels: map[string]string{
			"complexity": getString(arch, "complexity", "medium"),
			"pattern":    getString(arch, "architecture_pattern", "unknown"),
		},
	}
}
