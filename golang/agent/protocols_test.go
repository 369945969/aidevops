package agent

import (
	"encoding/json"
	"sync"
	"testing"
	"time"
)

func TestNewDevOpsContext(t *testing.T) {
	ctx := NewDevOpsContext("add SMS verification")
	if ctx.Query != "add SMS verification" {
		t.Errorf("Query = %q, want %q", ctx.Query, "add SMS verification")
	}
	if ctx.Code == nil {
		t.Error("Code map should be initialized")
	}
	if ctx.Outputs == nil {
		t.Error("Outputs should be initialized")
	}
	if ctx.RiskFlags == nil {
		t.Error("RiskFlags should be initialized")
	}
	if ctx.Meta == nil {
		t.Error("Meta should be initialized")
	}
	if ctx.CreatedAt.IsZero() {
		t.Error("CreatedAt should be set")
	}
}

func TestAddOutput(t *testing.T) {
	ctx := NewDevOpsContext("test")
	out := AgentOutput{AgentName: "analyst", Signal: "approve", Confidence: 0.9}
	ctx.AddOutput(out)

	if len(ctx.Outputs) != 1 {
		t.Fatalf("Outputs length = %d, want 1", len(ctx.Outputs))
	}
	if ctx.Outputs[0].AgentName != "analyst" {
		t.Errorf("AgentName = %q, want %q", ctx.Outputs[0].AgentName, "analyst")
	}
	if ctx.Outputs[0].Timestamp.IsZero() {
		t.Error("Timestamp should be auto-filled when zero")
	}

	outWithTime := AgentOutput{AgentName: "arch", Timestamp: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	ctx.AddOutput(outWithTime)
	if !ctx.Outputs[1].Timestamp.Equal(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Error("Explicit timestamp should not be overwritten")
	}
}

func TestAddRiskFlag(t *testing.T) {
	ctx := NewDevOpsContext("test")
	ctx.AddRiskFlag("security", "plain text storage", "high", "tester")

	if len(ctx.RiskFlags) != 1 {
		t.Fatalf("RiskFlags length = %d, want 1", len(ctx.RiskFlags))
	}
	f := ctx.RiskFlags[0]
	if f.Category != "security" || f.Severity != "high" || f.AgentName != "tester" {
		t.Errorf("RiskFlag fields mismatch: got %+v", f)
	}
}

func TestSetGetMeta(t *testing.T) {
	ctx := NewDevOpsContext("test")
	ctx.SetMeta("key1", "value1")
	ctx.SetMeta("key2", 42)

	if ctx.GetMeta("key1") != "value1" {
		t.Errorf("GetMeta(key1) = %v, want value1", ctx.GetMeta("key1"))
	}
	if ctx.GetMeta("key2") != 42 {
		t.Errorf("GetMeta(key2) = %v, want 42", ctx.GetMeta("key2"))
	}
	if ctx.GetMeta("nonexistent") != nil {
		t.Errorf("GetMeta(nonexistent) = %v, want nil", ctx.GetMeta("nonexistent"))
	}
}

func TestSetGetData(t *testing.T) {
	ctx := NewDevOpsContext("test")
	ctx.SetData("pipeline_data", "some output")
	if ctx.GetData("pipeline_data") != "some output" {
		t.Errorf("GetData = %v, want some output", ctx.GetData("pipeline_data"))
	}
}

func TestHasRiskFlags(t *testing.T) {
	ctx := NewDevOpsContext("test")
	if ctx.HasRiskFlags() {
		t.Error("empty context should have no risk flags")
	}
	ctx.AddRiskFlag("perf", "slow query", "medium", "tester")
	if !ctx.HasRiskFlags() {
		t.Error("context with risk flag should return true")
	}
}

func TestLatestOutput(t *testing.T) {
	ctx := NewDevOpsContext("test")
	ctx.AddOutput(AgentOutput{AgentName: "analyst", Confidence: 0.8})
	ctx.AddOutput(AgentOutput{AgentName: "architect", Confidence: 0.9})
	ctx.AddOutput(AgentOutput{AgentName: "analyst", Confidence: 0.95})

	result := ctx.LatestOutput("analyst")
	if result == nil {
		t.Fatal("LatestOutput(analyst) should not be nil")
	}
	if result.Confidence != 0.95 {
		t.Errorf("LatestOutput(analyst).Confidence = %v, want 0.95", result.Confidence)
	}

	if ctx.LatestOutput("nonexistent") != nil {
		t.Error("LatestOutput for unknown agent should be nil")
	}
}

func TestStageResultIsSuccess(t *testing.T) {
	sr := StageResult{Status: StageCompleted}
	if !sr.IsSuccess() {
		t.Error("completed stage should be success")
	}
	sr.Status = StageFailed
	if sr.IsSuccess() {
		t.Error("failed stage should not be success")
	}
}

func TestAgentRunStatsRecordStage(t *testing.T) {
	stats := &AgentRunStats{}

	stats.RecordStage(StageResult{Status: StageCompleted, TokensUsed: 100, DurationMs: 5000})
	stats.RecordStage(StageResult{Status: StageFailed, TokensUsed: 50, DurationMs: 2000})
	stats.RecordStage(StageResult{Status: StageSkipped, TokensUsed: 0, DurationMs: 0})

	if stats.TotalStages != 3 {
		t.Errorf("TotalStages = %d, want 3", stats.TotalStages)
	}
	if stats.CompletedStages != 1 {
		t.Errorf("CompletedStages = %d, want 1", stats.CompletedStages)
	}
	if stats.FailedStages != 1 {
		t.Errorf("FailedStages = %d, want 1", stats.FailedStages)
	}
	if stats.SkippedStages != 1 {
		t.Errorf("SkippedStages = %d, want 1", stats.SkippedStages)
	}
	if stats.TotalTokens != 150 {
		t.Errorf("TotalTokens = %d, want 150", stats.TotalTokens)
	}
	if stats.TotalDurationMs != 7000 {
		t.Errorf("TotalDurationMs = %d, want 7000", stats.TotalDurationMs)
	}
	if len(stats.StageResults) != 3 {
		t.Errorf("StageResults length = %d, want 3", len(stats.StageResults))
	}
}

func TestOrchestratorResultJSON(t *testing.T) {
	r := OrchestratorResult{
		Success:     true,
		Content:     "test content",
		TotalSteps:  5,
		TotalTokens: 500,
		Model:       "gpt-4o",
		Stats:       &AgentRunStats{TotalStages: 5, CompletedStages: 5},
	}
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}
	var r2 OrchestratorResult
	if err := json.Unmarshal(b, &r2); err != nil {
		t.Fatalf("JSON unmarshal failed: %v", err)
	}
	if r2.Success != true || r2.TotalSteps != 5 || r2.Model != "gpt-4o" {
		t.Errorf("Round-trip mismatch: %+v", r2)
	}
}

func TestConcurrentContextAccess(t *testing.T) {
	ctx := NewDevOpsContext("concurrent test")
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ctx.AddOutput(AgentOutput{AgentName: "agent_" + string(rune(n)), Confidence: float64(n) / 100})
			ctx.AddRiskFlag("cat", "desc", "low", "agent_" + string(rune(n)))
			ctx.SetMeta("key_"+string(rune(n)), n)
		}(i)
	}

	wg.Wait()

	if len(ctx.Outputs) != 100 {
		t.Errorf("Outputs length = %d, want 100", len(ctx.Outputs))
	}
	if len(ctx.RiskFlags) != 100 {
		t.Errorf("RiskFlags length = %d, want 100", len(ctx.RiskFlags))
	}
}

func TestStageStatusConstants(t *testing.T) {
	if StagePending != "pending" {
		t.Errorf("StagePending = %q, want pending", StagePending)
	}
	if StageRunning != "running" {
		t.Errorf("StageRunning = %q, want running", StageRunning)
	}
	if StageCompleted != "completed" {
		t.Errorf("StageCompleted = %q, want completed", StageCompleted)
	}
	if StageFailed != "failed" {
		t.Errorf("StageFailed = %q, want failed", StageFailed)
	}
	if StageSkipped != "skipped" {
		t.Errorf("StageSkipped = %q, want skipped", StageSkipped)
	}
}

func TestPipelineModeConstants(t *testing.T) {
	if ModeQuick != "quick" {
		t.Errorf("ModeQuick = %q, want quick", ModeQuick)
	}
	if ModeStandard != "standard" {
		t.Errorf("ModeStandard = %q, want standard", ModeStandard)
	}
	if ModeFull != "full" {
		t.Errorf("ModeFull = %q, want full", ModeFull)
	}
	if ModeReview != "review" {
		t.Errorf("ModeReview = %q, want review", ModeReview)
	}
}
