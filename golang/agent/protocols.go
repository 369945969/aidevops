package agent

import (
	"sync"
	"time"
)

// ============================================================
// Enums — Stage lifecycle status
// ============================================================

type StageStatus string

const (
	StagePending   StageStatus = "pending"
	StageRunning   StageStatus = "running"
	StageCompleted StageStatus = "completed"
	StageFailed    StageStatus = "failed"
	StageSkipped   StageStatus = "skipped"
)

// PipelineMode defines the depth of the agent chain.
type PipelineMode string

const (
	ModeQuick     PipelineMode = "quick"     // RequirementAnalyst → Developer
	ModeStandard  PipelineMode = "standard"  // RequirementAnalyst → Architect → Developer
	ModeFull      PipelineMode = "full"      // RequirementAnalyst → Architect → Developer → Tester → Deployer
	ModeReview    PipelineMode = "review"    // Full + human review nodes at architecture, code, deploy
)

// ============================================================
// DevOpsContext — shared state bag for a single pipeline run
// ============================================================

// DevOpsContext carries all accumulated state across agents in a pipeline.
// Each agent reads from and writes to the same context object, mirroring
// the AgentContext pattern from daily_stock_analysis/src/agent/protocols.py.
type DevOpsContext struct {
	mu sync.RWMutex

	// --- identity ---
	Query      string // original user requirement (single sentence)
	PipelineID uint   // database ID for this pipeline run
	SessionID  string // chat/conversation session ID

	// --- accumulated artifacts (populated by agents in order) ---
	Specification   string            // RequirementAnalyst output: structured requirement doc
	Architecture    string            // Architect output: design decisions, module layout
	Code            map[string]string // Developer output: filename → code content
	TestPlan        string            // Tester output: test strategy and cases
	TestResults     string            // Tester output: execution results
	DeploymentPlan  string            // Deployer output: deployment configuration

	// --- agent outputs (ordered, mirroring ctx.opinions) ---
	Outputs []AgentOutput

	// --- risk/review flags ---
	RiskFlags []RiskFlag

	// --- arbitrary metadata ---
	Meta map[string]interface{}

	// --- timing ---
	CreatedAt time.Time
}

// RiskFlag represents a risk or concern raised by any agent.
type RiskFlag struct {
	Category   string
	Description string
	Severity   string // "low", "medium", "high", "critical"
	AgentName  string
	Timestamp  time.Time
}

// NewDevOpsContext seeds a context from the user's requirement text.
func NewDevOpsContext(query string) *DevOpsContext {
	return &DevOpsContext{
		Query:      query,
		Code:       make(map[string]string),
		Outputs:    []AgentOutput{},
		RiskFlags:  []RiskFlag{},
		Meta:       make(map[string]interface{}),
		CreatedAt:  time.Now(),
	}
}

// AddOutput appends an agent output to the context.
func (ctx *DevOpsContext) AddOutput(output AgentOutput) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if output.Timestamp.IsZero() {
		output.Timestamp = time.Now()
	}
	ctx.Outputs = append(ctx.Outputs, output)
}

// AddRiskFlag adds a risk flag to the context.
func (ctx *DevOpsContext) AddRiskFlag(category, description, severity, agentName string) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.RiskFlags = append(ctx.RiskFlags, RiskFlag{
		Category:    category,
		Description: description,
		Severity:    severity,
		AgentName:   agentName,
		Timestamp:   time.Now(),
	})
}

// SetMeta stores a metadata value.
func (ctx *DevOpsContext) SetMeta(key string, value interface{}) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.Meta[key] = value
}

// GetMeta retrieves a metadata value.
func (ctx *DevOpsContext) GetMeta(key string) interface{} {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	return ctx.Meta[key]
}

// SetData stores an artifact in the context (convenience method for Meta).
func (ctx *DevOpsContext) SetData(key string, value interface{}) {
	ctx.SetMeta(key, value)
}

// GetData retrieves an artifact from the context.
func (ctx *DevOpsContext) GetData(key string) interface{} {
	return ctx.GetMeta(key)
}

// HasRiskFlags returns whether any risk flags exist.
func (ctx *DevOpsContext) HasRiskFlags() bool {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	return len(ctx.RiskFlags) > 0
}

// LatestOutput returns the most recent output from an agent matching the given name.
func (ctx *DevOpsContext) LatestOutput(agentName string) *AgentOutput {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	for i := len(ctx.Outputs) - 1; i >= 0; i-- {
		if ctx.Outputs[i].AgentName == agentName {
			return &ctx.Outputs[i]
		}
	}
	return nil
}

// ============================================================
// AgentOutput — structured output from any single agent
// ============================================================

// AgentOutput is the DevOps equivalent of AgentOpinion. Each agent
// produces one AgentOutput appended to DevOpsContext.Outputs.
type AgentOutput struct {
	AgentName  string                 // unique identifier (e.g. "requirement_analyst")
	Signal     string                 // "approve", "needs_revision", "block"
	Confidence float64                // 0.0 – 1.0
	Reasoning  string                 // explanation of the agent's decision
	Content    string                 // the main output text (specification, architecture, etc.)
	RawData    map[string]interface{} // any extra structured payload
	KeyLevels  map[string]string      // analogous to price levels, but for DevOps (e.g. "complexity", "risk")
	Timestamp  time.Time
}

// ============================================================
// StageResult — return type from a single pipeline stage
// ============================================================

// StageResult captures the outcome of one pipeline stage (one agent execution).
type StageResult struct {
	StageName       string
	Status          StageStatus
	Output          *AgentOutput
	Error           string
	DurationMs      int64
	TokensUsed      int
	ToolCallsCount  int
	Meta            map[string]interface{}
}

// IsSuccess returns true if the stage completed successfully.
func (sr *StageResult) IsSuccess() bool {
	return sr.Status == StageCompleted
}

// ============================================================
// AgentRunStats — aggregate statistics for an entire pipeline run
// ============================================================

// AgentRunStats collects statistics across all agents in a pipeline.
type AgentRunStats struct {
	TotalStages      int
	CompletedStages  int
	FailedStages     int
	SkippedStages    int
	TotalTokens      int
	TotalToolCalls   int
	TotalDurationMs  int64
	ModelsUsed       []string
	StageResults     []StageResult
}

// RecordStage adds a stage result and updates counters.
func (stats *AgentRunStats) RecordStage(result StageResult) {
	stats.StageResults = append(stats.StageResults, result)
	stats.TotalStages++
	stats.TotalTokens += result.TokensUsed
	stats.TotalToolCalls += result.ToolCallsCount
	stats.TotalDurationMs += result.DurationMs

	switch result.Status {
	case StageCompleted:
		stats.CompletedStages++
	case StageFailed:
		stats.FailedStages++
	case StageSkipped:
		stats.SkippedStages++
	}
}

// ============================================================
// OrchestratorResult — unified result from a pipeline run
// ============================================================

// OrchestratorResult is the final output of the multi-agent pipeline.
type OrchestratorResult struct {
	Success      bool
	Content      string                 // final assembled text
	PipelineData map[string]interface{} // structured pipeline output
	ToolCallsLog []map[string]interface{}
	TotalSteps   int
	TotalTokens  int
	Provider     string
	Model        string
	Error        string
	Stats        *AgentRunStats
}
