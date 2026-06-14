package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"devops/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupPipelineTestDB creates an in-memory SQLite DB and migrates models.
func setupPipelineTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(
		&models.Requirement{},
		&models.Pipeline{},
		&models.PipelineStage{},
		&models.PipelineArtifact{},
		&models.Notification{},
		&models.WorkflowTask{},
	); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

// setupPipelineTest creates a PipelineHandler with an in-memory DB and a Gin engine.
func setupPipelineTest(t *testing.T) (*PipelineHandler, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	db := setupPipelineTestDB(t)
	h := NewPipelineHandler(db)
	r := gin.New()
	return h, r
}

func createTestRequirement(t *testing.T, db *gorm.DB, title, desc, status, priority string) models.Requirement {
	req := models.Requirement{
		Title:       title,
		Description: desc,
		Status:      status,
		Priority:    priority,
	}
	if err := db.Create(&req).Error; err != nil {
		t.Fatalf("failed to create requirement: %v", err)
	}
	return req
}

func createTestPipeline(t *testing.T, db *gorm.DB, reqID uint, mode, status, query string) models.Pipeline {
	p := models.Pipeline{
		RequirementID: reqID,
		Mode:          mode,
		Status:        status,
		Query:         query,
	}
	if err := db.Create(&p).Error; err != nil {
		t.Fatalf("failed to create pipeline: %v", err)
	}
	return p
}

// ============================================================
// CreateRequirement Tests
// ============================================================

func TestCreateRequirement_Success(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/requirements", h.CreateRequirement)

	body := `{"description":"添加短信验证功能","priority":"高","author":"Sarah"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/requirements", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp struct {
		Data models.Requirement `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v. body=%s", err, w.Body.String())
	}
	if resp.Data.ID == 0 {
		t.Error("requirement should have ID > 0")
	}
	if resp.Data.Status != "需求分析" {
		t.Errorf("status = %q, want 需求分析", resp.Data.Status)
	}
}

func TestCreateRequirement_WithPipeline(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/requirements", h.CreateRequirement)

	body := `{"description":"实现支付系统","priority":"高","run_pipeline":true,"mode":"standard"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/requirements", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200. body=%s", w.Code, w.Body.String())
	}
	var resp struct {
		Data       models.Requirement `json:"data"`
		PipelineID uint               `json:"pipeline_id"`
		Status     string             `json:"status"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v. body=%s", err, w.Body.String())
	}
	if resp.PipelineID == 0 {
		t.Error("pipeline_id should be > 0")
	}
	if resp.Status != "pipeline_started" {
		t.Errorf("status = %q, want pipeline_started", resp.Status)
	}
	// Verify pipeline was created in DB
	var pipeline models.Pipeline
	if err := h.DB.First(&pipeline, resp.PipelineID).Error; err != nil {
		t.Fatalf("pipeline should exist in db: %v", err)
	}
	if pipeline.Mode != "standard" {
		t.Errorf("pipeline.Mode = %q, want standard", pipeline.Mode)
	}
}

func TestCreateRequirement_EmptyDescription(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/requirements", h.CreateRequirement)

	body := `{"description":"","priority":"低"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/requirements", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestCreateRequirement_InvalidJSON(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/requirements", h.CreateRequirement)

	body := `{invalid json}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/requirements", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

// ============================================================
// RunPipeline Tests
// ============================================================

func TestRunPipeline_Success(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/run", h.RunPipeline)

	req := createTestRequirement(t, h.DB, "test", "test desc", "需求分析", "中")
	p := createTestPipeline(t, h.DB, req.ID, "full", "pending", "test query")

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/run", nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200. body=%s", w.Code, w.Body.String())
	}
	var resp struct {
		Data   models.Pipeline `json:"data"`
		Status string          `json:"status"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.Status != "pipeline_started" {
		t.Errorf("status = %q, want pipeline_started", resp.Status)
	}
}

func TestRunPipeline_NotFound(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/run", h.RunPipeline)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/pipeline/99999/run", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

func TestRunPipeline_AlreadyRunning(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/run", h.RunPipeline)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "中")
	p := createTestPipeline(t, h.DB, req.ID, "full", "running", "test")

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/run", nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

// ============================================================
// GetPipeline Tests
// ============================================================

func TestGetPipeline_Found(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/pipeline/:id", h.GetPipeline)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "中")
	p := createTestPipeline(t, h.DB, req.ID, "full", "completed", "test query")

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("GET", "/api/pipeline/"+formatUint(p.ID), nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestGetPipeline_NotFound(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/pipeline/:id", h.GetPipeline)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/pipeline/99999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

// ============================================================
// GetPipelines Test
// ============================================================

func TestGetPipelines(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/pipeline", h.GetPipelines)

	req := createTestRequirement(t, h.DB, "test1", "desc1", "需求分析", "中")
	createTestPipeline(t, h.DB, req.ID, "full", "completed", "q1")
	createTestPipeline(t, h.DB, req.ID, "quick", "running", "q2")

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("GET", "/api/pipeline", nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp struct {
		Data []models.Pipeline `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(resp.Data) != 2 {
		t.Errorf("pipeline count = %d, want 2", len(resp.Data))
	}
}

// ============================================================
// GetPipelineArtifacts Test
// ============================================================

func TestGetPipelineArtifacts(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/pipeline/:id/artifacts", h.GetPipelineArtifacts)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "中")
	p := createTestPipeline(t, h.DB, req.ID, "full", "completed", "q")

	h.DB.Create(&models.PipelineArtifact{
		PipelineID: p.ID,
		StageName:  "developer",
		Type:       "code",
		Name:       "main.go",
		Content:    "package main",
	})

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("GET", "/api/pipeline/"+formatUint(p.ID)+"/artifacts", nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp struct {
		Data []models.PipelineArtifact `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(resp.Data) != 1 || resp.Data[0].Name != "main.go" {
		t.Errorf("unexpected artifacts: %+v", resp.Data)
	}
}

// ============================================================
// GetRequirements Tests
// ============================================================

func TestGetRequirements(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/requirements", h.GetRequirements)

	createTestRequirement(t, h.DB, "需求A", "descA", "需求分析", "高")
	createTestRequirement(t, h.DB, "需求B", "descB", "开发中", "中")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/requirements", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp struct {
		Data []models.Requirement `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(resp.Data) != 2 {
		t.Errorf("count = %d, want 2", len(resp.Data))
	}
}

func TestGetRequirements_FilterByStatus(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/requirements", h.GetRequirements)

	createTestRequirement(t, h.DB, "需求A", "descA", "需求分析", "高")
	createTestRequirement(t, h.DB, "需求B", "descB", "开发中", "中")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/requirements?status=需求分析", nil)
	r.ServeHTTP(w, req)

	var resp struct {
		Data []models.Requirement `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp.Data) != 1 {
		t.Errorf("filtered count = %d, want 1", len(resp.Data))
	}
}

func TestGetRequirements_FilterByPriority(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/requirements", h.GetRequirements)

	createTestRequirement(t, h.DB, "需求A", "descA", "需求分析", "高")
	createTestRequirement(t, h.DB, "需求B", "descB", "开发中", "中")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/requirements?priority=高", nil)
	r.ServeHTTP(w, req)

	var resp struct {
		Data []models.Requirement `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp.Data) != 1 {
		t.Errorf("filtered count = %d, want 1", len(resp.Data))
	}
}

// ============================================================
// GetRequirement Tests
// ============================================================

func TestGetRequirement_Found(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/requirements/:id", h.GetRequirement)

	req := createTestRequirement(t, h.DB, "测试需求", "desc", "需求分析", "高")

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("GET", "/api/requirements/"+formatUint(req.ID), nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestGetRequirement_NotFound(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/requirements/:id", h.GetRequirement)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/requirements/99999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

// ============================================================
// Notification Tests
// ============================================================

func TestGetNotifications(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/notifications", h.GetNotifications)

	h.DB.Create(&models.Notification{Title: "测试通知", Type: "progress", Description: "描述"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/notifications", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp struct {
		Data []models.Notification `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp.Data) != 1 {
		t.Errorf("count = %d, want 1", len(resp.Data))
	}
}

func TestMarkNotificationRead_Found(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.PUT("/api/notifications/:id/read", h.MarkNotificationRead)

	n := models.Notification{Title: "test", Type: "progress", IsRead: false}
	h.DB.Create(&n)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/notifications/"+formatUint(n.ID)+"/read", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var updated models.Notification
	h.DB.First(&updated, n.ID)
	if !updated.IsRead {
		t.Error("notification should be marked as read")
	}
}

func TestMarkNotificationRead_NotFound(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.PUT("/api/notifications/:id/read", h.MarkNotificationRead)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/notifications/99999/read", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

// ============================================================
// ResumePipelineReview Tests
// ============================================================

func TestResumePipelineReview_Approve(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/resume", h.ResumePipelineReview)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "高")
	p := createTestPipeline(t, h.DB, req.ID, "full", "paused_for_review", "query")
	stage := models.PipelineStage{
		PipelineID: p.ID,
		StageName:  "requirement_analyst",
		Status:     "paused_for_review",
	}
	h.DB.Create(&stage)

	body := `{"stage_name":"requirement_analyst","decision":"approve","comment":"looks good"}`
	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/resume",
		strings.NewReader(body))
	httpreq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200. body=%s", w.Code, w.Body.String())
	}
	var resp struct {
		Data     models.Pipeline `json:"data"`
		Decision string          `json:"decision"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.Decision != "approve" {
		t.Errorf("decision = %q, want approve", resp.Decision)
	}
}

func TestResumePipelineReview_Reject(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/resume", h.ResumePipelineReview)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "高")
	p := createTestPipeline(t, h.DB, req.ID, "full", "paused_for_review", "query")
	stage := models.PipelineStage{
		PipelineID: p.ID,
		StageName:  "developer",
		Status:     "paused_for_review",
	}
	h.DB.Create(&stage)

	body := `{"stage_name":"developer","decision":"reject","comment":"needs fix"}`
	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/resume",
		strings.NewReader(body))
	httpreq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	// Verify pipeline is now failed
	var pipeline models.Pipeline
	h.DB.First(&pipeline, p.ID)
	if pipeline.Status != "failed" {
		t.Errorf("pipeline status = %q, want failed", pipeline.Status)
	}
}

func TestResumePipelineReview_NotPaused(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/resume", h.ResumePipelineReview)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "高")
	p := createTestPipeline(t, h.DB, req.ID, "full", "running", "query")

	body := `{"stage_name":"developer","decision":"approve"}`
	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/resume",
		strings.NewReader(body))
	httpreq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400. body=%s", w.Code, w.Body.String())
	}
}

func TestResumePipelineReview_InvalidDecision(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/resume", h.ResumePipelineReview)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "高")
	p := createTestPipeline(t, h.DB, req.ID, "full", "paused_for_review", "query")
	stage := models.PipelineStage{
		PipelineID: p.ID,
		StageName:  "developer",
		Status:     "paused_for_review",
	}
	h.DB.Create(&stage)

	body := `{"stage_name":"developer","decision":"invalid_decision"}`
	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/resume",
		strings.NewReader(body))
	httpreq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestResumePipelineReview_PipelineNotFound(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/resume", h.ResumePipelineReview)

	body := `{"stage_name":"developer","decision":"approve"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/pipeline/99999/resume",
		strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

func TestResumePipelineReview_StageNotFound(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.POST("/api/pipeline/:id/resume", h.ResumePipelineReview)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "高")
	p := createTestPipeline(t, h.DB, req.ID, "full", "paused_for_review", "query")

	body := `{"stage_name":"nonexistent_stage","decision":"approve"}`
	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("POST", "/api/pipeline/"+formatUint(p.ID)+"/resume",
		strings.NewReader(body))
	httpreq.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400. body=%s", w.Code, w.Body.String())
	}
}

// ============================================================
// GetWorkflowTasks Test
// ============================================================

func TestGetWorkflowTasks(t *testing.T) {
	h, r := setupPipelineTest(t)
	r.GET("/api/workflow/tasks", h.GetWorkflowTasks)

	req := createTestRequirement(t, h.DB, "test", "desc", "需求分析", "高")
	p := createTestPipeline(t, h.DB, req.ID, "full", "running", "query")
	h.DB.Create(&models.WorkflowTask{
		PipelineID: p.ID,
		Title:      "实现需求A",
		Agent:      "后端开发",
		Status:     "开发中",
		Priority:   "高",
	})

	w := httptest.NewRecorder()
	httpreq, _ := http.NewRequest("GET", "/api/workflow/tasks", nil)
	r.ServeHTTP(w, httpreq)

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	var resp struct {
		Data []models.WorkflowTask `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if len(resp.Data) != 1 {
		t.Errorf("count = %d, want 1", len(resp.Data))
	}
}

// ============================================================
// truncate helper test
// ============================================================

func TestTruncate(t *testing.T) {
	tests := []struct {
		input  string
		maxLen int
		want   string
	}{
		{"hello", 10, "hello"},
		{"hello world", 5, "hello..."},
		{"short", 20, "short"},
		{"", 5, ""},
	}
	for _, tc := range tests {
		got := truncate(tc.input, tc.maxLen)
		if got != tc.want {
			t.Errorf("truncate(%q, %d) = %q, want %q", tc.input, tc.maxLen, got, tc.want)
		}
	}
}

// formatUint converts uint to string for URL construction.
func formatUint(n uint) string {
	return fmt.Sprintf("%d", n)
}
