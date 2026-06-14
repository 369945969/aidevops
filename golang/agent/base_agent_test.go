package agent

import (
	"strings"
	"testing"
)

func TestPersistArtifactsRequirement(t *testing.T) {
	ctx := NewDevOpsContext("test")
	output := &AgentOutput{AgentName: "requirement_analyst", Content: "structured spec"}
	persistArtifacts(ctx, output)
	if ctx.Specification != "structured spec" {
		t.Errorf("Specification = %q, want 'structured spec'", ctx.Specification)
	}
}

func TestPersistArtifactsArchitect(t *testing.T) {
	ctx := NewDevOpsContext("test")
	output := &AgentOutput{AgentName: "architect", Content: "architecture design"}
	persistArtifacts(ctx, output)
	if ctx.Architecture != "architecture design" {
		t.Errorf("Architecture = %q, want 'architecture design'", ctx.Architecture)
	}
}

func TestPersistArtifactsDeveloper(t *testing.T) {
	ctx := NewDevOpsContext("test")
	output := &AgentOutput{
		AgentName: "developer",
		Content:   "code output",
		RawData: map[string]interface{}{
			"files": map[string]interface{}{
				"handler.go": "package handler",
				"main.go":    "package main",
			},
		},
	}
	persistArtifacts(ctx, output)
	if ctx.Code["handler.go"] != "package handler" {
		t.Errorf("Code[handler.go] = %q, want 'package handler'", ctx.Code["handler.go"])
	}
	if ctx.GetData("developer_output") != "code output" {
		t.Errorf("developer_output = %v, want 'code output'", ctx.GetData("developer_output"))
	}
}

func TestPersistArtifactsTester(t *testing.T) {
	ctx := NewDevOpsContext("test")
	output := &AgentOutput{
		AgentName: "tester",
		Content:   "test results",
		RawData:   map[string]interface{}{"test_plan": "plan data"},
	}
	persistArtifacts(ctx, output)
	if ctx.TestResults != "test results" {
		t.Errorf("TestResults = %q, want 'test results'", ctx.TestResults)
	}
	if ctx.GetData("test_plan") != "plan data" {
		t.Errorf("test_plan = %v, want 'plan data'", ctx.GetData("test_plan"))
	}
}

func TestPersistArtifactsDeployer(t *testing.T) {
	ctx := NewDevOpsContext("test")
	output := &AgentOutput{AgentName: "deployer", Content: "deploy plan"}
	persistArtifacts(ctx, output)
	if ctx.DeploymentPlan != "deploy plan" {
		t.Errorf("DeploymentPlan = %q, want 'deploy plan'", ctx.DeploymentPlan)
	}
}

func TestBuildInjectedDataSectionEmpty(t *testing.T) {
	ctx := NewDevOpsContext("test")
	result := buildInjectedDataSection(ctx)
	if result != "" {
		t.Errorf("empty context should produce empty section, got: %q", result)
	}
}

func TestBuildInjectedDataSectionFull(t *testing.T) {
	ctx := NewDevOpsContext("test")
	ctx.Specification = "spec content"
	ctx.Architecture = "arch content"
	ctx.Code = map[string]string{"handler.go": "code", "main.go": "code2"}
	ctx.TestResults = "test results"
	ctx.AddOutput(AgentOutput{AgentName: "analyst", Signal: "approve", Confidence: 0.9, Reasoning: "looks good"})
	ctx.AddRiskFlag("security", "plain text", "high", "tester")

	result := buildInjectedDataSection(ctx)
	if !strings.Contains(result, "需求分析结果") {
		t.Error("should contain specification header")
	}
	if !strings.Contains(result, "架构设计方案") {
		t.Error("should contain architecture header")
	}
	if !strings.Contains(result, "已生成代码文件") {
		t.Error("should contain code files header")
	}
	if !strings.Contains(result, "handler.go") {
		t.Error("should list code file names")
	}
	if !strings.Contains(result, "测试结果") {
		t.Error("should contain test results header")
	}
	if !strings.Contains(result, "前序Agent输出摘要") {
		t.Error("should contain agent output summary")
	}
	if !strings.Contains(result, "风险标记") {
		t.Error("should contain risk flags section")
	}
}

func TestBuildInjectedDataSectionLongReasoningTruncated(t *testing.T) {
	ctx := NewDevOpsContext("test")
	longReasoning := strings.Repeat("x", 600)
	ctx.AddOutput(AgentOutput{AgentName: "analyst", Signal: "approve", Confidence: 0.9, Reasoning: longReasoning})

	result := buildInjectedDataSection(ctx)
	if !strings.Contains(result, "...(truncated)") {
		t.Error("long reasoning should be truncated")
	}
}
