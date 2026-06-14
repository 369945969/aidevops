package agents

import (
	"strings"
	"testing"

	"devops/agent"
)

func TestDeveloperAgent_AgentName(t *testing.T) {
	a := &DeveloperAgent{}
	if name := a.AgentName(); name != "developer" {
		t.Errorf("AgentName = %q, want developer", name)
	}
}

func TestDeveloperAgent_MaxSteps(t *testing.T) {
	a := &DeveloperAgent{}
	if s := a.MaxSteps(); s != 1 {
		t.Errorf("MaxSteps = %d, want 1", s)
	}
}

func TestDeveloperAgent_SystemPrompt(t *testing.T) {
	a := &DeveloperAgent{}
	ctx := agent.NewDevOpsContext("test")
	prompt := a.SystemPrompt(ctx)
	if prompt == "" {
		t.Error("SystemPrompt should not be empty")
	}
	if !strings.Contains(prompt, "软件开发") {
		t.Error("SystemPrompt should contain Chinese identifier")
	}
}

func TestDeveloperAgent_BuildUserMessage(t *testing.T) {
	a := &DeveloperAgent{}
	ctx := agent.NewDevOpsContext("implement SMS API")
	msg := a.BuildUserMessage(ctx)
	if !strings.Contains(msg, ctx.Query) {
		t.Error("BuildUserMessage should include ctx.Query")
	}
}

func TestDeveloperAgent_PostProcess_ValidJSON(t *testing.T) {
	a := &DeveloperAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{
		"files": {
			"main.go": "package main\nfunc main() {}",
			"handler.go": "package handler"
		},
		"signal": "approve",
		"confidence": 0.9,
		"explanation": "generated two files"
	}`
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should return non-nil output")
	}
	if output.Signal != "approve" {
		t.Errorf("Signal = %q, want approve", output.Signal)
	}
	if output.RawData["files"] == nil {
		t.Error("RawData should contain files map")
	}
}

func TestDeveloperAgent_PostProcess_Fallback(t *testing.T) {
	a := &DeveloperAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "Just some code text without JSON"
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess fallback should not return nil")
	}
	if output.Confidence != 0.5 {
		t.Errorf("fallback Confidence = %v, want 0.5", output.Confidence)
	}
}

func TestDeveloperAgent_PostProcess_MarkdownFiles(t *testing.T) {
	a := &DeveloperAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "Here are the files:\n```json\n{\"files\":{\"app.go\":\"package app\"},\"signal\":\"approve\",\"confidence\":0.8}\n```"
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should handle markdown JSON")
	}
	if output.RawData["files"] == nil {
		t.Error("should parse files from markdown JSON")
	}
}


