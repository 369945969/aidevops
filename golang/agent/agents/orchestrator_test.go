package agents

import (
	"testing"

	"devops/agent"
)

func TestBuildAgentChainQuick(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeQuick, 0)
	chain := o.BuildAgentChain()
	if len(chain) != 2 {
		t.Errorf("quick mode chain length = %d, want 2", len(chain))
	}
	if chain[0].AgentName() != "requirement_analyst" {
		t.Errorf("first agent = %q, want requirement_analyst", chain[0].AgentName())
	}
	if chain[1].AgentName() != "developer" {
		t.Errorf("second agent = %q, want developer", chain[1].AgentName())
	}
}

func TestBuildAgentChainStandard(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeStandard, 0)
	chain := o.BuildAgentChain()
	if len(chain) != 3 {
		t.Errorf("standard mode chain length = %d, want 3", len(chain))
	}
}

func TestBuildAgentChainFull(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 0)
	chain := o.BuildAgentChain()
	if len(chain) != 5 {
		t.Errorf("full mode chain length = %d, want 5", len(chain))
	}
	expected := []string{"requirement_analyst", "architect", "developer", "tester", "deployer"}
	for i, ag := range chain {
		if ag.AgentName() != expected[i] {
			t.Errorf("chain[%d].AgentName = %q, want %q", i, ag.AgentName(), expected[i])
		}
	}
}

func TestBuildAgentChainReview(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeReview, 0)
	chain := o.BuildAgentChain()
	if len(chain) != 5 {
		t.Errorf("review mode chain length = %d, want 5", len(chain))
	}
}

func TestBuildAgentChainEmptyMode(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, "", 0)
	chain := o.BuildAgentChain()
	if len(chain) != 5 {
		t.Errorf("empty mode defaults to full, chain length = %d, want 5", len(chain))
	}
}

func TestIsNonCriticalStage(t *testing.T) {
	if !isNonCriticalStage("tester") {
		t.Error("tester should be non-critical")
	}
	if !isNonCriticalStage("deployer") {
		t.Error("deployer should be non-critical")
	}
	if isNonCriticalStage("requirement_analyst") {
		t.Error("requirement_analyst should be critical")
	}
	if isNonCriticalStage("architect") {
		t.Error("architect should be critical")
	}
	if isNonCriticalStage("developer") {
		t.Error("developer should be critical")
	}
}

func TestTruncate(t *testing.T) {
	if truncate("short", 10) != "short" {
		t.Error("short string should not be truncated")
	}
	long := "this is a very long string that exceeds the limit"
	if truncate(long, 10) != long[:10]+"...(truncated)" {
		t.Errorf("long string truncation mismatch")
	}
	if truncate("", 10) != "" {
		t.Error("empty string should remain empty")
	}
}

func TestUniqueStrings(t *testing.T) {
	result := uniqueStrings([]string{"a", "b", "a", "c", "b"})
	if len(result) != 3 {
		t.Errorf("uniqueStrings length = %d, want 3", len(result))
	}
	seen := map[string]bool{}
	for _, s := range result {
		seen[s] = true
	}
	if !seen["a"] || !seen["b"] || !seen["c"] {
		t.Errorf("missing unique elements: %v", seen)
	}

	if len(uniqueStrings([]string{})) != 0 {
		t.Error("empty input should return empty output")
	}

	if len(uniqueStrings([]string{"x"})) != 1 {
		t.Error("single element should return single element")
	}
}

func TestFirstString(t *testing.T) {
	if firstString([]string{"a", "b"}) != "a" {
		t.Errorf("firstString = %q, want a", firstString([]string{"a", "b"}))
	}
	if firstString([]string{}) != "" {
		t.Error("firstString on empty slice should return empty string")
	}
}

func TestJoinStrings(t *testing.T) {
	if joinStrings([]string{"a", "b", "a"}, ", ") != "a, b" {
		t.Errorf("joinStrings with duplicates = %q, want 'a, b'", joinStrings([]string{"a", "b", "a"}, ", "))
	}
	if joinStrings([]string{"x"}, ", ") != "x" {
		t.Errorf("joinStrings single = %q, want x", joinStrings([]string{"x"}, ", "))
	}
	if joinStrings([]string{}, ", ") != "" {
		t.Error("joinStrings empty should return empty string")
	}
}

func TestResolveFinalOutputEmpty(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 0)
	ctx := agent.NewDevOpsContext("test query")
	ctx.PipelineID = 1

	data := o.ResolveFinalOutput(ctx)
	if data["query"] != "test query" {
		t.Errorf("query = %v, want test query", data["query"])
	}
	if data["pipeline_id"] != uint(1) {
		t.Errorf("pipeline_id = %v, want 1", data["pipeline_id"])
	}
	if data["specification"] != "" {
		t.Errorf("specification should be empty")
	}
	if data["code_files"] != nil {
		t.Error("code_files should be nil when Code is empty")
	}
}

func TestResolveFinalOutputFull(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 0)
	ctx := agent.NewDevOpsContext("full test")
	ctx.PipelineID = 2
	ctx.Specification = "spec data"
	ctx.Architecture = "arch data"
	ctx.Code = map[string]string{"main.go": "package main"}
	ctx.TestResults = "test data"
	ctx.DeploymentPlan = "deploy data"
	ctx.AddOutput(agent.AgentOutput{AgentName: "analyst", Signal: "approve", Confidence: 0.9})
	ctx.AddRiskFlag("security", "risk desc", "high", "tester")

	data := o.ResolveFinalOutput(ctx)
	if data["specification"] != "spec data" {
		t.Errorf("specification = %v, want spec data", data["specification"])
	}
	if data["code_files"] == nil {
		t.Error("code_files should be present when Code is non-empty")
	}
	outputs := data["agent_outputs"].([]map[string]interface{})
	if len(outputs) != 1 {
		t.Errorf("agent_outputs length = %d, want 1", len(outputs))
	}
	riskFlags := data["risk_flags"].([]map[string]interface{})
	if len(riskFlags) != 1 {
		t.Errorf("risk_flags length = %d, want 1", len(riskFlags))
	}
}

func TestBuildContentSummaryEmpty(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 0)
	ctx := agent.NewDevOpsContext("test")
	result := o.BuildContentSummary(ctx)
	if result != "" {
		t.Errorf("empty context should produce empty summary, got: %q", result)
	}
}

func TestBuildContentSummaryPartial(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 0)
	ctx := agent.NewDevOpsContext("add SMS verification")
	ctx.Specification = "structured requirement specification"

	result := o.BuildContentSummary(ctx)
	if len(result) == 0 {
		t.Error("partial context should produce non-empty summary")
	}
}

func TestBuildContentSummaryFull(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 0)
	ctx := agent.NewDevOpsContext("payment refactoring")
	ctx.Specification = "spec"
	ctx.Architecture = "arch"
	ctx.Code = map[string]string{"handler.go": "code", "service.go": "code2"}
	ctx.TestResults = "tests"
	ctx.DeploymentPlan = "deploy"

	result := o.BuildContentSummary(ctx)
	if len(result) == 0 {
		t.Error("full context should produce non-empty summary")
	}
	if len(result) < 50 {
		t.Errorf("summary seems too short: %d chars", len(result))
	}
}

func TestNewDevOpsOrchestrator(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, agent.ModeFull, 60000)
	if o.Mode != agent.ModeFull {
		t.Errorf("Mode = %q, want full", o.Mode)
	}
	if o.TimeoutMs != 60000 {
		t.Errorf("TimeoutMs = %d, want 60000", o.TimeoutMs)
	}
}

func TestNewDevOpsOrchestratorDefaultMode(t *testing.T) {
	o := NewDevOpsOrchestrator(nil, "", 0)
	if o.Mode != agent.ModeFull {
		t.Errorf("empty mode should default to full, got %q", o.Mode)
	}
}
