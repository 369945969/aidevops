package agent

import "encoding/json"

// ToolDefinition describes a tool that agents can invoke during execution.
type ToolDefinition struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]ParamSpec   `json:"parameters"`
}

type ParamSpec struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Required    bool     `json:"required"`
	Enum        []string `json:"enum,omitempty"`
}

// ToolResult is the output from executing a tool call.
type ToolResult struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
	Error   string `json:"error,omitempty"`
}

// ToolExecutor is the function signature for tool implementations.
type ToolExecutor func(params map[string]interface{}) ToolResult

// ToolRegistry holds available tools and their executors.
type ToolRegistry struct {
	tools map[string]ToolEntry
}

type ToolEntry struct {
	Definition ToolDefinition
	Executor   ToolExecutor
}

func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{tools: make(map[string]ToolEntry)}
}

func (r *ToolRegistry) Register(def ToolDefinition, executor ToolExecutor) {
	r.tools[def.Name] = ToolEntry{Definition: def, Executor: executor}
}

func (r *ToolRegistry) Get(name string) (ToolEntry, bool) {
	entry, ok := r.tools[name]
	return entry, ok
}

func (r *ToolRegistry) AllDefinitions() []ToolDefinition {
	defs := make([]ToolDefinition, 0, len(r.tools))
	for _, entry := range r.tools {
		defs = append(defs, entry.Definition)
	}
	return defs
}

func (r *ToolRegistry) Execute(name string, paramsJSON string) ToolResult {
	entry, ok := r.tools[name]
	if !ok {
		return ToolResult{Success: false, Error: "tool not found: " + name}
	}
	var params map[string]interface{}
	if err := json.Unmarshal([]byte(paramsJSON), &params); err != nil {
		return ToolResult{Success: false, Error: "invalid params JSON: " + err.Error()}
	}
	return entry.Executor(params)
}
