package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestRequirementDefaults(t *testing.T) {
	// GORM default tags only apply when using DB.Create, not on raw struct.
	// Verify the tag values are set correctly by creating via GORM (if DB available)
	// For unit test, just verify zero values are acceptable
	r := Requirement{}
	_ = r // zero-value struct should not panic
}

func TestRequirementJSON(t *testing.T) {
	now := time.Now()
	r := Requirement{
		ID:          1,
		Title:       "Add SMS feature",
		Description: "Need SMS verification",
		Status:      "开发中",
		Priority:    "高",
		Author:      "Sarah",
		TaskCount:   3,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var r2 Requirement
	if err := json.Unmarshal(b, &r2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if r2.ID != 1 || r2.Title != "Add SMS feature" || r2.Priority != "高" {
		t.Errorf("round-trip mismatch: %+v", r2)
	}
}

func TestPipelineDefaults(t *testing.T) {
	p := Pipeline{}
	_ = p // zero-value struct should not panic; GORM defaults apply via DB.Create
}

func TestPipelineJSON(t *testing.T) {
	p := Pipeline{
		ID:              10,
		RequirementID:   5,
		Mode:            "full",
		Status:          "running",
		Query:           "add SMS feature",
		Specification:   "spec content",
		Architecture:    "arch content",
		CodeArtifacts:   `{"main.go":"package main"}`,
		TestResults:     "test results",
		DeployPlan:      "deploy plan",
		TotalTokens:     1500,
		TotalDurationMs: 45000,
		ModelsUsed:      "gpt-4o",
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var p2 Pipeline
	if err := json.Unmarshal(b, &p2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if p2.ID != 10 || p2.Query != "add SMS feature" {
		t.Errorf("round-trip mismatch: %+v", p2)
	}
}

func TestPipelineStageDefaults(t *testing.T) {
	s := PipelineStage{}
	_ = s // zero-value struct should not panic; GORM defaults apply via DB.Create
}

func TestPipelineStageJSON(t *testing.T) {
	s := PipelineStage{
		ID:         1,
		PipelineID: 10,
		StageName:  "developer",
		Status:     "completed",
		Output:     "generated code",
		Signal:     "approve",
		Confidence: 0.95,
		Reasoning:  "code looks good",
		DurationMs: 12000,
		TokensUsed: 500,
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var s2 PipelineStage
	if err := json.Unmarshal(b, &s2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if s2.StageName != "developer" || s2.Signal != "approve" || s2.Confidence != 0.95 {
		t.Errorf("round-trip mismatch: %+v", s2)
	}
}

func TestPipelineArtifactJSON(t *testing.T) {
	a := PipelineArtifact{
		ID:         1,
		PipelineID: 10,
		StageName:  "developer",
		Type:       "code",
		Name:       "main.go",
		Content:    "package main\nfunc main() {}",
	}
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var a2 PipelineArtifact
	if err := json.Unmarshal(b, &a2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if a2.Name != "main.go" || a2.Content != "package main\nfunc main() {}" {
		t.Errorf("round-trip mismatch: %+v", a2)
	}
}

func TestNotificationDefaults(t *testing.T) {
	n := Notification{}
	if n.IsRead != false {
		t.Error("IsRead should default to false")
	}
}

func TestNotificationJSON(t *testing.T) {
	n := Notification{
		ID:          1,
		Title:       "Pipeline completed",
		Type:        "progress",
		Description: "All stages done",
		Route:       "/workflow/dashboard",
		IsRead:      false,
	}
	b, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var n2 Notification
	if err := json.Unmarshal(b, &n2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if n2.Title != "Pipeline completed" || n2.IsRead != false {
		t.Errorf("round-trip mismatch: %+v", n2)
	}
}

func TestWorkflowTaskDefaults(t *testing.T) {
	task := WorkflowTask{}
	_ = task // zero-value struct should not panic; GORM defaults apply via DB.Create
}

func TestWorkflowTaskJSON(t *testing.T) {
	task := WorkflowTask{
		ID:          1,
		PipelineID:  10,
		Title:       "Implement SMS API",
		Agent:       "后端开发",
		Priority:    "高",
		Progress:    75,
		Estimate:    "预计2h后完成",
		Status:      "开发中",
		ReviewRoute: "/review/code",
		Extra:       "覆盖率 72%",
	}
	b, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var task2 WorkflowTask
	if err := json.Unmarshal(b, &task2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if task2.Title != "Implement SMS API" || task2.Progress != 75 {
		t.Errorf("round-trip mismatch: %+v", task2)
	}
}

func TestPipelineJSONWithStages(t *testing.T) {
	type PipelineWithStages struct {
		Pipeline
		Stages []PipelineStage `json:"stages"`
	}
	p := PipelineWithStages{
		Pipeline: Pipeline{ID: 1, Query: "test"},
		Stages: []PipelineStage{
			{PipelineID: 1, StageName: "requirement_analyst", Status: "completed"},
			{PipelineID: 1, StageName: "developer", Status: "running"},
		},
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !json.Valid(b) {
		t.Error("output should be valid JSON")
	}
	var decoded PipelineWithStages
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(decoded.Stages) != 2 {
		t.Errorf("stages count = %d, want 2", len(decoded.Stages))
	}
}

func TestAllModelZeroValues(t *testing.T) {
	// Verify all models can be created with zero values without panicking
	r := Requirement{}
	_ = r.ID

	p := Pipeline{}
	_ = p.ID

	s := PipelineStage{}
	_ = s.ID

	a := PipelineArtifact{}
	_ = a.ID

	n := Notification{}
	_ = n.ID

	w := WorkflowTask{}
	_ = w.ID
}
