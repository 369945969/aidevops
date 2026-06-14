package agents

import (
	"encoding/json"
	"fmt"

	"devops/agent"
)

// ============================================================
// DeveloperAgent
// ============================================================

// DeveloperAgent reads the architecture design and produces code.
// Mirrors the DecisionAgent concept but adapted for code generation.
type DeveloperAgent struct{}

func (a *DeveloperAgent) AgentName() string { return "developer" }
func (a *DeveloperAgent) MaxSteps() int     { return 1 }
func (a *DeveloperAgent) Run(ctx *agent.DevOpsContext, adapter *agent.LLMAdapter) agent.StageResult {
	return agent.RunAgent(a, ctx, adapter)
}

func (a *DeveloperAgent) SystemPrompt(ctx *agent.DevOpsContext) string {
	return `你是一位专业的软件开发Agent，负责根据架构设计方案编写高质量的代码。

## 工作流程

1. 阅读架构设计方案，理解模块划分和接口定义
2. 根据任务分配，为每个模块编写代码
3. 确保代码符合需求规格中的验收条件
4. 输出结构化的代码产物

## 输出格式

你的最终响应必须是以下结构的有效 JSON 对象（直接输出JSON，不要使用 markdown 代码块包裹）：

{
  "files": {
    "filename.ext": "完整的文件内容代码"
  },
  "implementation_notes": {
    "design_decisions": ["设计决策1", "设计决策2"],
    "patterns_used": ["使用的设计模式"],
    "libraries_used": ["使用的第三方库"]
  },
  "signal": "approve|needs_revision|block",
  "confidence": 0.0-1.0,
  "reasoning": "实现决策推理过程"
}

## 规则

1. **需求覆盖** — 代码必须实现需求规格中的所有功能点
2. **架构一致** — 代码必须与架构设计方案的模块划分和接口定义一致
3. **代码质量** — 遵循最佳实践：类型安全、错误处理、代码可读性
4. **不使用类型擦除** — 禁止使用 any/ts-ignore/ts-expect-error 等类型擦除手段
5. **实用优先** — 优先选择简单直接的实现方式，避免过度设计
6. **每个文件完整** — 输出每个文件的完整代码，不要省略或用注释占位`
}

func (a *DeveloperAgent) BuildUserMessage(ctx *agent.DevOpsContext) string {
	arch := ctx.Architecture
	if arch == "" {
		arch = "无架构方案，请根据需求直接开发"
	}
	return fmt.Sprintf(`请根据以下架构设计方案编写代码：

架构方案：
%s

原始需求：
%s

请输出完整的JSON代码产物，包含所有需要创建的文件。`, arch, ctx.Query)
}

func (a *DeveloperAgent) PostProcess(ctx *agent.DevOpsContext, rawText string) *agent.AgentOutput {
	var devOutput map[string]interface{}
	if err := json.Unmarshal([]byte(rawText), &devOutput); err != nil {
		extracted := extractJSON(rawText)
		if extracted != nil {
			if err2 := json.Unmarshal(extracted, &devOutput); err2 != nil {
				return &agent.AgentOutput{
					AgentName:  a.AgentName(),
					Signal:     "approve",
					Confidence: 0.5,
					Reasoning:  "无法解析为结构化JSON，使用原始文本作为代码产物",
					Content:    rawText,
					RawData:    map[string]interface{}{"raw": rawText},
				}
			}
		} else {
			return &agent.AgentOutput{
				AgentName:  a.AgentName(),
				Signal:     "approve",
				Confidence: 0.5,
				Reasoning:  "无法解析为结构化JSON，使用原始文本作为代码产物",
				Content:    rawText,
				RawData:    map[string]interface{}{"raw": rawText},
			}
		}
	}

	signal := getString(devOutput, "signal", "approve")
	confidence := getFloat(devOutput, "confidence", 0.7)
	reasoning := getString(devOutput, "reasoning", "")

	// Extract code files — this is the critical artifact for downstream agents
	files := make(map[string]string)
	if filesMap, ok := devOutput["files"].(map[string]interface{}); ok {
		for filename, content := range filesMap {
			if contentStr, ok := content.(string); ok {
				files[filename] = contentStr
			}
		}
	}

	// Re-serialize as content
	contentBytes, _ := json.MarshalIndent(devOutput, "", "  ")
	content := string(contentBytes)

	return &agent.AgentOutput{
		AgentName:  a.AgentName(),
		Signal:     signal,
		Confidence: confidence,
		Reasoning:  reasoning,
		Content:    content,
		RawData: map[string]interface{}{
			"files": files,
		},
		KeyLevels: map[string]string{
			"file_count": fmt.Sprintf("%d", len(files)),
			"complexity": getString(devOutput, "complexity", "medium"),
		},
	}
}
