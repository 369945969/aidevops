package agents

import (
	"strings"
	"testing"

	"devops/agent"
)

func TestRequirementAnalystAgent_AgentName(t *testing.T) {
	a := &RequirementAnalystAgent{}
	if name := a.AgentName(); name != "requirement_analyst" {
		t.Errorf("AgentName = %q, want requirement_analyst", name)
	}
}

func TestRequirementAnalystAgent_MaxSteps(t *testing.T) {
	a := &RequirementAnalystAgent{}
	if s := a.MaxSteps(); s != 1 {
		t.Errorf("MaxSteps = %d, want 1", s)
	}
}

func TestRequirementAnalystAgent_SystemPrompt(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("test")
	prompt := a.SystemPrompt(ctx)
	if prompt == "" {
		t.Error("SystemPrompt should not be empty")
	}
	if !strings.Contains(prompt, "需求分析师") {
		t.Error("SystemPrompt should contain Chinese identifier")
	}
	if !strings.Contains(prompt, "JSON") {
		t.Error("SystemPrompt should mention JSON output")
	}
}

func TestRequirementAnalystAgent_BuildUserMessage(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("add SMS verification")
	msg := a.BuildUserMessage(ctx)
	if !strings.Contains(msg, "add SMS verification") {
		t.Error("BuildUserMessage should include ctx.Query")
	}
}

func TestRequirementAnalystAgent_PostProcess_ValidJSON(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{
		"title": "SMS Verification",
		"priority": "高",
		"description": "add SMS code verification",
		"functional_requirements": [{"id":"FR-001","title":"send code","description":"send","acceptance_criteria":["works"]}],
		"non_functional_requirements": [],
		"assumptions": [],
		"risks": [],
		"dependencies": [],
		"signal": "approve",
		"confidence": 0.9,
		"reasoning": "clear requirement"
	}`
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should return non-nil output")
	}
	if output.Signal != "approve" {
		t.Errorf("Signal = %q, want approve", output.Signal)
	}
	if output.Confidence != 0.9 {
		t.Errorf("Confidence = %v, want 0.9", output.Confidence)
	}
	if output.Reasoning != "clear requirement" {
		t.Errorf("Reasoning = %q, want clear requirement", output.Reasoning)
	}
	if !strings.Contains(output.Content, "SMS Verification") {
		t.Errorf("Content should contain title")
	}
	if output.RawData["title"] != "SMS Verification" {
		t.Errorf("RawData title = %v, want SMS Verification", output.RawData["title"])
	}
}

func TestRequirementAnalystAgent_PostProcess_MarkdownJSON(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "Here is the spec:\n```json\n{\"title\": \"Markdown Spec\", \"signal\": \"approve\", \"confidence\": 0.8}\n```\nEnd."
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should extract JSON from markdown")
	}
	if output.RawData["title"] != "Markdown Spec" {
		t.Errorf("title = %v, want Markdown Spec", output.RawData["title"])
	}
}

func TestRequirementAnalystAgent_PostProcess_RawTextFallback(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := "This is plain text with no JSON at all. Just a description."
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should fallback for non-JSON")
	}
	if output.Content != raw {
		t.Errorf("Content should be raw text on fallback")
	}
	if output.Confidence != 0.6 {
		t.Errorf("fallback Confidence = %v, want 0.6", output.Confidence)
	}
}

func TestRequirementAnalystAgent_PostProcess_BlockSignal(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{"title":"Bad","signal":"block","confidence":0.2,"reasoning":"requirement is impossible"}`
	output := a.PostProcess(ctx, raw)
	if output == nil {
		t.Fatal("PostProcess should not return nil")
	}
	if output.Signal != "block" {
		t.Errorf("Signal = %q, want block", output.Signal)
	}
}

func TestRequirementAnalystAgent_PostProcess_RiskFlags(t *testing.T) {
	a := &RequirementAnalystAgent{}
	ctx := agent.NewDevOpsContext("test")
	raw := `{
		"title":"Risk Test","signal":"approve","confidence":0.7,
		"risks":[{"description":"security risk","severity":"high","mitigation":"encrypt"}]
	}`
	_ = a.PostProcess(ctx, raw)
	if !ctx.HasRiskFlags() {
		t.Error("risk flags should be added to context")
	}
}

func TestExtractJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantNil bool
	}{
		{"json_block", "```json\n{\"a\":1}\n```", false},
		{"raw_json", "some text {\"b\":2} more", false},
		{"no_json", "just text", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSON(tt.input)
			if tt.wantNil && result != nil {
				t.Errorf("expected nil, got %s", string(result))
			}
			if !tt.wantNil && result == nil {
				t.Error("expected non-nil result")
			}
		})
	}
}

func TestGetString(t *testing.T) {
	m := map[string]interface{}{"key": "value", "num": 42}
	if getString(m, "key", "default") != "value" {
		t.Error("getString should return value")
	}
	if getString(m, "missing", "default") != "default" {
		t.Error("getString should return default for missing key")
	}
	if getString(m, "num", "default") != "default" {
		t.Error("getString should return default for non-string type")
	}
	if getString(nil, "key", "default") != "default" {
		t.Error("getString should return default for nil map")
	}
}

func TestGetFloat(t *testing.T) {
	m := map[string]interface{}{"f": 3.14, "i": 5}
	if getFloat(m, "f", 0) != 3.14 {
		t.Error("getFloat should return float64")
	}
	if getFloat(m, "i", 0) != 5 {
		t.Error("getFloat should convert int to float64")
	}
	if getFloat(m, "missing", 1.0) != 1.0 {
		t.Error("getFloat should return default for missing key")
	}
}

func TestGetInt(t *testing.T) {
	m := map[string]interface{}{"i": 10, "f": 3.5}
	if getInt(m, "i", 0) != 10 {
		t.Error("getInt should return int")
	}
	if getInt(m, "f", 0) != 3 {
		t.Error("getInt should convert float64 to int")
	}
	if getInt(m, "missing", 99) != 99 {
		t.Error("getInt should return default for missing key")
	}
}
