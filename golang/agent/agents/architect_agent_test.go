package agents

import (
	"strings"
	"testing"

	"devops/agent"
)

func TestArchitectAgent_AgentName(t *testing.T) {
	a := &ArchitectAgent{}
	if name := a.AgentName(); name != "architect" {
		t.Errorf("AgentName = %q, want architect", name)
	}
}

func TestArchitectAgent_MaxSteps(t *testing.T) {
	a := &ArchitectAgent{}
	if s := a.MaxSteps(); s != 1 {
		t.Errorf("MaxSteps = %d, want 1", s)
	}
}

func TestArchitectAgent_SystemPrompt(t *testing.T) {
	a := &ArchitectAgent{}
	ctx := agent.NewDevOpsContext("test")
	prompt := a.SystemPrompt(ctx)
	if prompt == "" {
		t.Error("SystemPrompt should not be empty")
	}
	if !strings.Contains(prompt, "架构师") {
		t.Error("SystemPrompt should contain Chinese identifier")
	}
}

func TestArchitectAgent_BuildUserMessage(t *testing.T) {
	a := &ArchitectAgent{}
	ctx := agent.NewDevOpsContext("build payment module")
	msg := a.BuildUserMessage(ctx)
	if !strings.Contains(msg, "build payment module") {
		t.Error("BuildUserMessage should include ctx.Query")
	}
}

func TestArchitectAgent_PostProcess_ValidJSON(t *testing.T) {
	a := &ArchitectAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{
		"architecture": "microservices",
		"tech_stack": ["Go", "Redis", "PostgreSQL"],
		"modules": [{"name": "api", "description": "API gateway"}],
		"signal": "approve",
		"confidence": 0.85,
		"reasoning": "solid architecture"
	}`
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should return non-nil output")
	}
	if output.Signal != "approve" {
		t.Errorf("Signal = %q, want approve", output.Signal)
	}
	if !strings.Contains(output.Content, "microservices") {
		t.Errorf("Content should contain architecture type")
	}
}

func TestArchitectAgent_PostProcess_Fallback(t *testing.T) {
	a := &ArchitectAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "This is plain text architecture description without JSON"
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should not return nil for fallback")
	}
	if output.Confidence != 0.6 {
		t.Errorf("fallback Confidence = %v, want 0.6", output.Confidence)
	}
}

func TestArchitectAgent_PostProcess_MarkdownJSON(t *testing.T) {
	a := &ArchitectAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "Here's the design:\n```json\n{\"architecture\":\"event-driven\",\"signal\":\"approve\",\"confidence\":0.9}\n```"
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should handle markdown JSON")
	}
	if output.Signal != "approve" {
		t.Errorf("Signal = %q, want approve", output.Signal)
	}
}
