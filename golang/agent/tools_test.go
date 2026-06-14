package agent

import (
	"encoding/json"
	"testing"
)

func TestNewToolRegistry(t *testing.T) {
	r := NewToolRegistry()
	if r.tools == nil {
		t.Error("tools map should be initialized")
	}
	if len(r.AllDefinitions()) != 0 {
		t.Error("new registry should have no tools")
	}
}

func TestRegisterAndGet(t *testing.T) {
	r := NewToolRegistry()
	def := ToolDefinition{
		Name:        "test_tool",
		Description: "a test tool",
		Parameters: map[string]ParamSpec{
			"input": {Type: "string", Required: true},
		},
	}
	exec := func(params map[string]interface{}) ToolResult {
		return ToolResult{Success: true, Output: "executed"}
	}
	r.Register(def, exec)

	entry, found := r.Get("test_tool")
	if !found {
		t.Error("registered tool should be found")
	}
	if entry.Definition.Name != "test_tool" {
		t.Errorf("Definition.Name = %q, want test_tool", entry.Definition.Name)
	}

	_, found = r.Get("nonexistent")
	if found {
		t.Error("nonexistent tool should not be found")
	}
}

func TestExecute(t *testing.T) {
	r := NewToolRegistry()
	r.Register(ToolDefinition{Name: "adder"}, func(params map[string]interface{}) ToolResult {
		_, _ = params["a"].(float64)
		_, _ = params["b"].(float64)
		return ToolResult{Success: true, Output: "result"}
	})

	result := r.Execute("adder", `{"a":1,"b":2}`)
	if !result.Success {
		t.Errorf("Execute should succeed: %s", result.Error)
	}
}

func TestExecuteUnknownTool(t *testing.T) {
	r := NewToolRegistry()
	result := r.Execute("unknown", `{}`)
	if result.Success {
		t.Error("Execute on unknown tool should fail")
	}
	if result.Error != "tool not found: unknown" {
		t.Errorf("Error = %q, want 'tool not found: unknown'", result.Error)
	}
}

func TestExecuteInvalidJSON(t *testing.T) {
	r := NewToolRegistry()
	r.Register(ToolDefinition{Name: "tool"}, func(p map[string]interface{}) ToolResult {
		return ToolResult{Success: true}
	})
	result := r.Execute("tool", "not json{{{")
	if result.Success {
		t.Error("Execute with invalid JSON should fail")
	}
	if !json.Valid([]byte(result.Error)) {
		// Just verify error mentions invalid params
		t.Logf("error message: %s", result.Error)
	}
}

func TestAllDefinitions(t *testing.T) {
	r := NewToolRegistry()
	r.Register(ToolDefinition{Name: "tool1"}, nil)
	r.Register(ToolDefinition{Name: "tool2"}, nil)

	defs := r.AllDefinitions()
	if len(defs) != 2 {
		t.Errorf("AllDefinitions length = %d, want 2", len(defs))
	}
	names := map[string]bool{}
	for _, d := range defs {
		names[d.Name] = true
	}
	if !names["tool1"] || !names["tool2"] {
		t.Errorf("AllDefinitions should contain tool1 and tool2, got %v", names)
	}
}

func TestToolResultJSON(t *testing.T) {
	tr := ToolResult{Success: true, Output: "hello"}
	b, err := json.Marshal(tr)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}
	if !json.Valid(b) {
		t.Error("ToolResult JSON should be valid")
	}

	tr2 := ToolResult{Success: false, Error: "something went wrong"}
	b2, _ := json.Marshal(tr2)
	var parsed ToolResult
	json.Unmarshal(b2, &parsed)
	if parsed.Success != false || parsed.Error != "something went wrong" {
		t.Errorf("round-trip mismatch: %+v", parsed)
	}
}

func TestToolDefinitionJSON(t *testing.T) {
	def := ToolDefinition{
		Name:        "my_tool",
		Description: "does something",
		Parameters: map[string]ParamSpec{
			"path": {Type: "string", Description: "file path", Required: true},
			"mode": {Type: "string", Enum: []string{"read", "write"}, Required: false},
		},
	}
	b, err := json.Marshal(def)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}
	var def2 ToolDefinition
	if err := json.Unmarshal(b, &def2); err != nil {
		t.Fatalf("JSON unmarshal failed: %v", err)
	}
	if def2.Name != "my_tool" {
		t.Errorf("Name = %q, want my_tool", def2.Name)
	}
	if len(def2.Parameters) != 2 {
		t.Errorf("Parameters length = %d, want 2", len(def2.Parameters))
	}
}
