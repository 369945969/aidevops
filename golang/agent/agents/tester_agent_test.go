package agents

import (
	"strings"
	"testing"

	"devops/agent"
)

func TestTesterAgent_AgentName(t *testing.T) {
	a := &TesterAgent{}
	if name := a.AgentName(); name != "tester" {
		t.Errorf("AgentName = %q, want tester", name)
	}
}

func TestTesterAgent_MaxSteps(t *testing.T) {
	a := &TesterAgent{}
	if s := a.MaxSteps(); s != 1 {
		t.Errorf("MaxSteps = %d, want 1", s)
	}
}

func TestTesterAgent_SystemPrompt(t *testing.T) {
	a := &TesterAgent{}
	ctx := agent.NewDevOpsContext("test")
	prompt := a.SystemPrompt(ctx)
	if prompt == "" {
		t.Error("SystemPrompt should not be empty")
	}
	if !strings.Contains(prompt, "测试") {
		t.Error("SystemPrompt should contain Chinese identifier")
	}
}

func TestTesterAgent_BuildUserMessage(t *testing.T) {
	a := &TesterAgent{}
	ctx := agent.NewDevOpsContext("test payment")
	_ = ctx // BuildUserMessage uses ctx.Specification & ctx.Architecture, not ctx.Query directly
	msg := a.BuildUserMessage(ctx)
	if msg == "" {
		t.Error("BuildUserMessage should not be empty")
	}
}

func TestTesterAgent_PostProcess_ValidJSON(t *testing.T) {
	a := &TesterAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{
		"test_plan": {
			"unit_tests": [{"target": "auth", "cases": [{"name": "test_login", "type": "unit"}]}]
		},
		"code_quality": {"score": 80, "issues": []},
		"coverage_estimate": 75,
		"signal": "approve",
		"confidence": 0.8
	}`
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should return non-nil output")
	}
	if output.Signal != "approve" {
		t.Errorf("Signal = %q, want approve", output.Signal)
	}
	rd := output.RawData
	if rd["test_plan"] == nil {
		t.Error("test_plan should be present in RawData")
	}
}

func TestTesterAgent_PostProcess_Fallback(t *testing.T) {
	a := &TesterAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "Plain text test plan"
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess fallback should not return nil")
	}
	if output.Confidence != 0.5 {
		t.Errorf("fallback Confidence = %v, want 0.5", output.Confidence)
	}
}


