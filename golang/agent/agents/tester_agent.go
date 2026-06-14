package agents

import (
	"encoding/json"
	"fmt"

	"devops/agent"
)

// ============================================================
// TesterAgent
// ============================================================

// TesterAgent reads the generated code and produces test plans and results.
// Mirrors RiskAgent from the Python project but adapted for DevOps testing.
type TesterAgent struct{}

func (a *TesterAgent) AgentName() string { return "tester" }
func (a *TesterAgent) MaxSteps() int     { return 1 }
func (a *TesterAgent) Run(ctx *agent.DevOpsContext, adapter *agent.LLMAdapter) agent.StageResult {
	return agent.RunAgent(a, ctx, adapter)
}

func (a *TesterAgent) SystemPrompt(ctx *agent.DevOpsContext) string {
	return `你是一位专业的软件测试Agent，负责为生成的代码设计测试方案并评估代码质量。

## 工作流程

1. 阅读需求规格和架构设计，理解验收条件
2. 检查生成的代码文件
3. 设计测试用例覆盖每个功能需求
4. 评估代码质量和安全性
5. 输出结构化的测试报告

## 输出格式

你的最终响应必须是以下结构的有效 JSON 对象（直接输出JSON，不要使用 markdown 代码块包裹）：

{
  "test_plan": {
    "unit_tests": [
      {
        "target": "测试目标模块/函数",
        "cases": [
          {
            "name": "测试用例名称",
            "description": "测试描述",
            "input": "输入",
            "expected": "预期输出",
            "type": "正向|逆向|边界"
          }
        ]
      }
    ],
    "integration_tests": [
      {
        "name": "集成测试名称",
        "description": "测试描述",
        "endpoints": ["涉及的接口"]
      }
    ]
  },
  "code_quality": {
    "score": 0-100,
    "issues": [
      {
        "severity": "critical|warning|info",
        "description": "问题描述",
        "file": "所在文件",
        "suggestion": "修改建议"
      }
    ]
  },
  "security_scan": {
    "vulnerabilities": [
      {
        "type": "漏洞类型",
        "severity": "高|中|低",
        "description": "漏洞描述",
        "file": "所在文件",
        "remediation": "修复建议"
      }
    ]
  },
  "coverage_estimate": 0-100,
  "signal": "approve|needs_revision|block",
  "confidence": 0.0-1.0,
  "reasoning": "测试评估推理过程"
}

## 规则

1. **验收标准导向** — 每个测试用例必须对应需求规格中的验收条件
2. **安全优先** — 必须检查常见安全漏洞（XSS、注入、权限等）
3. **质量评分客观** — 基于实际代码问题评分，不编造问题
4. **覆盖完整** — 正向、逆向、边界测试都要覆盖`
}

func (a *TesterAgent) BuildUserMessage(ctx *agent.DevOpsContext) string {
	codeFiles := ""
	for filename, _ := range ctx.Code {
		codeFiles += "- " + filename + "\n"
	}
	if codeFiles == "" {
		codeFiles = "无代码文件（请基于需求直接生成测试方案）"
	}

	return fmt.Sprintf(`请对以下代码进行测试设计和质量评估：

需求规格：
%s

架构方案：
%s

代码文件列表：
%s

请输出完整的JSON测试报告。`, truncate(ctx.Specification, 2000), truncate(ctx.Architecture, 2000), codeFiles)
}

func (a *TesterAgent) PostProcess(ctx *agent.DevOpsContext, rawText string) *agent.AgentOutput {
	var testResult map[string]interface{}
	if err := json.Unmarshal([]byte(rawText), &testResult); err != nil {
		extracted := extractJSON(rawText)
		if extracted != nil {
			if err2 := json.Unmarshal(extracted, &testResult); err2 != nil {
				return &agent.AgentOutput{
					AgentName:  a.AgentName(),
					Signal:     "approve",
					Confidence: 0.5,
					Reasoning:  "无法解析为结构化JSON",
					Content:    rawText,
					RawData:    map[string]interface{}{"raw": rawText},
				}
			}
		} else {
			return &agent.AgentOutput{
				AgentName:  a.AgentName(),
				Signal:     "approve",
				Confidence: 0.5,
				Reasoning:  "无法解析为结构化JSON",
				Content:    rawText,
				RawData:    map[string]interface{}{"raw": rawText},
			}
		}
	}

	signal := getString(testResult, "signal", "approve")
	confidence := getFloat(testResult, "confidence", 0.7)
	reasoning := getString(testResult, "reasoning", "")

	// Extract code quality issues as risk flags
	if quality, ok := testResult["code_quality"].(map[string]interface{}); ok {
		if issues, ok := quality["issues"].([]interface{}); ok {
			for _, issue := range issues {
				if issueMap, ok := issue.(map[string]interface{}); ok {
					ctx.AddRiskFlag(
						"code_quality",
					 getString(issueMap, "description", ""),
					 getString(issueMap, "severity", "warning"),
					 a.AgentName(),
					)
				}
			}
		}
	}

	contentBytes, _ := json.MarshalIndent(testResult, "", "  ")
	content := string(contentBytes)

	return &agent.AgentOutput{
		AgentName:  a.AgentName(),
		Signal:     signal,
		Confidence: confidence,
		Reasoning:  reasoning,
		Content:    content,
		RawData: map[string]interface{}{
			"test_plan": testResult["test_plan"],
		},
		KeyLevels: map[string]string{
			"quality_score":   fmt.Sprintf("%v", testResult["code_quality"]),
			"coverage":        fmt.Sprintf("%v", testResult["coverage_estimate"]),
		},
	}
}

