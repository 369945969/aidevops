package agents

import (
	"encoding/json"
	"fmt"

	"devops/agent"
)

// ============================================================
// DeployerAgent
// ============================================================

// DeployerAgent reads the test results and produces a deployment plan.
// This is the final stage in the pipeline, producing the deployment
// configuration and strategy.
type DeployerAgent struct{}

func (a *DeployerAgent) AgentName() string { return "deployer" }
func (a *DeployerAgent) MaxSteps() int     { return 1 }
func (a *DeployerAgent) Run(ctx *agent.DevOpsContext, adapter *agent.LLMAdapter) agent.StageResult {
	return agent.RunAgent(a, ctx, adapter)
}

func (a *DeployerAgent) SystemPrompt(ctx *agent.DevOpsContext) string {
	return `你是一位专业的DevOps部署Agent，负责为项目生成部署方案和配置。

## 工作流程

1. 阅读架构方案和代码产物
2. 根据技术栈选择合适的部署方式
3. 设计环境配置、数据库迁移、回滚策略
4. 输出结构化的部署方案

## 输出格式

你的最终响应必须是以下结构的有效 JSON 对象（直接输出JSON，不要使用 markdown 代码块包裹）：

{
  "deployment_strategy": "部署策略（如滚动部署、蓝绿部署等）",
  "environments": {
    "staging": {
      "platform": "部署平台",
      "config": {
        "base_url": "基础URL",
        "port": "端口",
        "env_vars": ["环境变量列表"]
      }
    },
    "production": {
      "platform": "部署平台",
      "config": {
        "base_url": "基础URL",
        "port": "端口",
        "env_vars": ["环境变量列表"]
      }
    }
  },
  "database_migrations": [
    {
      "version": "迁移版本",
      "description": "迁移描述",
      "steps": ["迁移步骤"]
    }
  ],
  "deployment_steps": [
    {
      "order": 1,
      "name": "步骤名称",
      "description": "步骤描述",
      "command": "执行命令",
      "estimated_time": "预估时间"
    }
  ],
  "rollback_strategy": {
    "method": "回滚方式",
    "steps": ["回滚步骤"],
    "estimated_time": "回滚预估时间"
  },
  "health_checks": [
    {
      "name": "检查名称",
      "type": "HTTP/数据库/服务",
      "endpoint": "检查端点",
      "expected": "预期结果"
    }
  ],
  "risks": [
    {
      "description": "部署风险",
      "severity": "高|中|低",
      "mitigation": "缓解措施"
    }
  ],
  "signal": "approve|needs_revision|block",
  "confidence": 0.0-1.0,
  "reasoning": "部署方案推理过程"
}

## 规则

1. **安全优先** — 生产环境配置必须安全（不暴露密钥等）
2. **回滚必须** — 任何部署方案都必须包含回滚策略
3. **健康检查** — 部署后必须有可验证的健康检查
4. **环境隔离** — staging和生产环境配置必须独立`
}

func (a *DeployerAgent) BuildUserMessage(ctx *agent.DevOpsContext) string {
	codeFiles := ""
	for filename := range ctx.Code {
		codeFiles += "- " + filename + "\n"
	}

	return fmt.Sprintf(`请为以下项目生成部署方案：

需求规格摘要：
%s

架构方案摘要：
%s

代码文件：
%s

测试结果摘要：
%s

请输出完整的JSON部署方案。`, truncate(ctx.Specification, 1000), truncate(ctx.Architecture, 1000), codeFiles, truncate(ctx.TestResults, 1000))
}

func (a *DeployerAgent) PostProcess(ctx *agent.DevOpsContext, rawText string) *agent.AgentOutput {
	var deployResult map[string]interface{}
	if err := json.Unmarshal([]byte(rawText), &deployResult); err != nil {
		extracted := extractJSON(rawText)
		if extracted != nil {
			if err2 := json.Unmarshal(extracted, &deployResult); err2 != nil {
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

	signal := getString(deployResult, "signal", "approve")
	confidence := getFloat(deployResult, "confidence", 0.7)
	reasoning := getString(deployResult, "reasoning", "")

	// Add deployment risk flags
	if risks, ok := deployResult["risks"].([]interface{}); ok {
		for _, r := range risks {
			if riskMap, ok := r.(map[string]interface{}); ok {
				ctx.AddRiskFlag(
					"deployment_risk",
				 getString(riskMap, "description", ""),
				 getString(riskMap, "severity", "medium"),
				 a.AgentName(),
				)
			}
		}
	}

	contentBytes, _ := json.MarshalIndent(deployResult, "", "  ")
	content := string(contentBytes)

	return &agent.AgentOutput{
		AgentName:  a.AgentName(),
		Signal:     signal,
		Confidence: confidence,
		Reasoning:  reasoning,
		Content:    content,
		RawData:    deployResult,
		KeyLevels: map[string]string{
			"strategy": getString(deployResult, "deployment_strategy", "unknown"),
		},
	}
}
