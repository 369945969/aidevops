package agents

import (
	"strings"
	"testing"

	"devops/agent"
)

func TestDeployerAgent_AgentName(t *testing.T) {
	a := &DeployerAgent{}
	if name := a.AgentName(); name != "deployer" {
		t.Errorf("AgentName = %q, want deployer", name)
	}
}

func TestDeployerAgent_MaxSteps(t *testing.T) {
	a := &DeployerAgent{}
	if s := a.MaxSteps(); s != 1 {
		t.Errorf("MaxSteps = %d, want 1", s)
	}
}

func TestDeployerAgent_SystemPrompt(t *testing.T) {
	a := &DeployerAgent{}
	ctx := agent.NewDevOpsContext("test")
	prompt := a.SystemPrompt(ctx)
	if prompt == "" {
		t.Error("SystemPrompt should not be empty")
	}
	if !strings.Contains(prompt, "部署方案") {
		t.Error("SystemPrompt should contain deployment-related Chinese")
	}
}

func TestDeployerAgent_BuildUserMessage(t *testing.T) {
	a := &DeployerAgent{}
	ctx := agent.NewDevOpsContext("deploy payment v2")
	_ = ctx // BuildUserMessage uses ctx.Specification, Architecture, Code, TestResults — not ctx.Query
	msg := a.BuildUserMessage(ctx)
	if msg == "" {
		t.Error("BuildUserMessage should not be empty")
	}
}

func TestDeployerAgent_PostProcess_ValidJSON(t *testing.T) {
	a := &DeployerAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{
		"deploy_strategy": "rolling_update",
		"steps": [
			{"order": 1, "action": "stop_service"},
			{"order": 2, "action": "migrate_db"}
		],
		"rollback": {"strategy": "full_rollback", "estimated_time": "5min"},
		"env_config": {"REDIS_TTL": "600"},
		"monitoring": ["health_check", "error_rate"],
		"signal": "approve",
		"confidence": 0.9
	}`
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should return non-nil output")
	}
	if output.Signal != "approve" {
		t.Errorf("Signal = %q, want approve", output.Signal)
	}
	if !strings.Contains(output.Content, "rolling_update") {
		t.Errorf("Content should contain deploy strategy")
	}
}

func TestDeployerAgent_PostProcess_Fallback(t *testing.T) {
	a := &DeployerAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "Plain text deploy plan"
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess fallback should not return nil")
	}
	if output.Confidence != 0.5 {
		t.Errorf("fallback Confidence = %v, want 0.5", output.Confidence)
	}
}


