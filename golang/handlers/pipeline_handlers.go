package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"devops/agent"
	"devops/agent/agents"
	"devops/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ============================================================
// Pipeline Handlers — Create, Run, Monitor
// ============================================================

// PipelineHandler manages the DevOps pipeline lifecycle.
type PipelineHandler struct {
	DB         *gorm.DB
	Orch       *agents.DevOpsOrchestrator
	LLMAdapter *agent.LLMAdapter
	EventBus   *PipelineEventBus
}

// NewPipelineHandler creates a handler with DB access and event bus.
func NewPipelineHandler(db *gorm.DB) *PipelineHandler {
	return &PipelineHandler{
		DB:       db,
		EventBus: NewPipelineEventBus(),
	}
}

// CreateRequirement handles POST /api/requirements — creates a requirement from
// a single sentence and optionally triggers a pipeline.
func (h *PipelineHandler) CreateRequirement(c *gin.Context) {
	var input struct {
		Description string `json:"description"`
		Priority    string `json:"priority"`
		Author      string `json:"author"`
		RunPipeline bool   `json:"run_pipeline"`
		Mode        string `json:"mode"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description is required"})
		return
	}

	if input.Mode == "" {
		input.Mode = "full"
	}

	// Create requirement
	req := models.Requirement{
		Title:       truncate(input.Description, 50),
		Description: input.Description,
		Status:      "需求分析",
		Priority:    input.Priority,
		Author:      input.Author,
	}

	if err := h.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Optionally create and start pipeline
	if input.RunPipeline {
		pipeline := models.Pipeline{
			RequirementID: req.ID,
			Mode:          input.Mode,
			Status:        "pending",
			Query:         input.Description,
		}

		if err := h.DB.Create(&pipeline).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		req.PipelineID = &pipeline.ID
		req.Status = "架构设计"
		h.DB.Save(&req)

		// Start pipeline execution asynchronously
		go h.RunPipelineAsync(pipeline.ID, input.Description, input.Mode)

		c.JSON(http.StatusOK, gin.H{
			"data":        req,
			"pipeline_id": pipeline.ID,
			"status":      "pipeline_started",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": req})
}

// RunPipeline handles POST /api/pipeline/:id/run — manually trigger a pipeline.
func (h *PipelineHandler) RunPipeline(c *gin.Context) {
	id := c.Param("id")
	var pipeline models.Pipeline
	if err := h.DB.First(&pipeline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pipeline not found"})
		return
	}

	if pipeline.Status == "running" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pipeline already running"})
		return
	}

	// Update status
	pipeline.Status = "running"
	h.DB.Save(&pipeline)

	// Start execution asynchronously
	go h.RunPipelineAsync(pipeline.ID, pipeline.Query, pipeline.Mode)

	c.JSON(http.StatusOK, gin.H{"data": pipeline, "status": "pipeline_started"})
}

// RunPipelineAsync executes the multi-agent pipeline and persists results.
func (h *PipelineHandler) RunPipelineAsync(pipelineID uint, query string, mode string) {
	// Create LLM adapter (from env or DB config)
	llmAdapter := agent.NewLLMAdapterFromEnv()

	orch := agents.NewDevOpsOrchestrator(llmAdapter, agent.PipelineMode(mode), 0)
	orch.DB = h.DB

	h.EventBus.Publish(PipelineEvent{
		PipelineID: pipelineID,
		Type:       "pipeline_started",
		Message:    "Pipeline execution started",
		Status:     "running",
	})

	// Create stage records in DB
	var stages []models.PipelineStage
	agentChain := orch.BuildAgentChain()
	for _, ag := range agentChain {
		stage := models.PipelineStage{
			PipelineID: pipelineID,
			StageName:  ag.AgentName(),
			Status:     "pending",
		}
		h.DB.Create(&stage)
		stages = append(stages, stage)
		h.EventBus.Publish(PipelineEvent{
			PipelineID: pipelineID,
			Stage:      ag.AgentName(),
			Type:       "stage_pending",
			Message:    "Stage " + ag.AgentName() + " registered",
			Status:     "pending",
		})
	}

	// Execute the pipeline
	result := orch.Run(query, pipelineID)

	// Persist results
	var pipeline models.Pipeline
	h.DB.First(&pipeline, pipelineID)

	pipeline.Status = "completed"
	pipelineCompleteType := "pipeline_completed"
	pipelineCompleteMsg := "Pipeline execution completed successfully"
	if !result.Success {
		pipeline.Status = "failed"
		pipeline.Error = result.Error
		pipelineCompleteType = "pipeline_failed"
		pipelineCompleteMsg = "Pipeline execution failed: " + result.Error
	}
	h.EventBus.Publish(PipelineEvent{
		PipelineID: pipelineID,
		Type:       pipelineCompleteType,
		Message:    pipelineCompleteMsg,
		DurationMs: result.Stats.TotalDurationMs,
		Status:     pipeline.Status,
	})

	if v, ok := result.PipelineData["specification"].(string); ok {
		pipeline.Specification = v
	}
	if v, ok := result.PipelineData["architecture"].(string); ok {
		pipeline.Architecture = v
	}
	if v, ok := result.PipelineData["test_results"].(string); ok {
		pipeline.TestResults = v
	}
	if v, ok := result.PipelineData["deploy_plan"].(string); ok {
		pipeline.DeployPlan = v
	}
	pipeline.TotalTokens = result.TotalTokens
	pipeline.TotalDurationMs = result.Stats.TotalDurationMs
	pipeline.ModelsUsed = result.Model

	// Persist code artifacts
	if codeFiles, ok := result.PipelineData["code_files"].(map[string]any); ok {
		codeJSON, _ := json.Marshal(codeFiles)
		pipeline.CodeArtifacts = string(codeJSON)

		for filename, content := range codeFiles {
			if contentStr, ok := content.(string); ok {
				artifact := models.PipelineArtifact{
					PipelineID: pipelineID,
					StageName:  "developer",
					Type:       "code",
					Name:       filename,
					Content:    contentStr,
				}
				h.DB.Create(&artifact)
			}
		}
	}

	pipeline.Result = result.Content
	h.DB.Save(&pipeline)

	// Update stage records from the orchestrator stats
	if result.Stats != nil {
		for i, stageResult := range result.Stats.StageResults {
			if i < len(stages) {
				stages[i].Status = string(stageResult.Status)
				if rt, ok := stageResult.Meta["raw_text"].(string); ok {
					stages[i].Output = rt
				}
				stages[i].Signal = ""
				stages[i].Confidence = 0
				stages[i].Reasoning = ""
				stages[i].DurationMs = stageResult.DurationMs
				stages[i].TokensUsed = stageResult.TokensUsed
				if stageResult.Output != nil {
					stages[i].Signal = stageResult.Output.Signal
					stages[i].Confidence = stageResult.Output.Confidence
					stages[i].Reasoning = stageResult.Output.Reasoning
					stages[i].Output = stageResult.Output.Content
				}
				if stageResult.Status == agent.StageFailed {
					stages[i].Error = stageResult.Error
					h.EventBus.Publish(PipelineEvent{
						PipelineID: pipelineID,
						Stage:      stages[i].StageName,
						Type:       "stage_failed",
						Message:    "Stage " + stages[i].StageName + " failed: " + stageResult.Error,
						DurationMs: stageResult.DurationMs,
						Status:     "failed",
						Error:      stageResult.Error,
					})
				} else {
					h.EventBus.Publish(PipelineEvent{
						PipelineID: pipelineID,
						Stage:      stages[i].StageName,
						Type:       "stage_done",
						Message:    "Stage " + stages[i].StageName + " completed",
						DurationMs: stageResult.DurationMs,
						Status:     "completed",
					})
				}
				h.DB.Save(&stages[i])
			}
		}
	}

	// Update requirement status
	var requirement models.Requirement
	h.DB.Where("pipeline_id = ?", pipelineID).First(&requirement)
	requirement.Status = "已完成"
	if !result.Success {
		requirement.Status = "失败"
	}
	requirement.TaskCount = len(stages)
	h.DB.Save(&requirement)

	// Create notification
	notification := models.Notification{
		Title:       "管道完成通知",
		Type:        "progress",
		Description: "需求 " + truncate(query, 30) + " 的开发管道已完成",
		Route:       "/workflow/dashboard",
	}
	if !result.Success {
		notification.Type = "agent_exception"
		notification.Title = "管道执行失败"
		notification.Description = "需求 " + truncate(query, 30) + " 的开发管道失败: " + result.Error
	}
	h.DB.Create(&notification)
}

// GetPipeline handles GET /api/pipeline/:id — returns pipeline status.
func (h *PipelineHandler) GetPipeline(c *gin.Context) {
	id := c.Param("id")
	var pipeline models.Pipeline
	if err := h.DB.Preload("Stages").Preload("Requirement").First(&pipeline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pipeline not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pipeline})
}

// GetPipelines handles GET /api/pipeline — returns all pipelines.
func (h *PipelineHandler) GetPipelines(c *gin.Context) {
	var pipelines []models.Pipeline
	h.DB.Preload("Stages").Preload("Requirement").Find(&pipelines)
	c.JSON(http.StatusOK, gin.H{"data": pipelines})
}

// GetPipelineArtifacts handles GET /api/pipeline/:id/artifacts.
func (h *PipelineHandler) GetPipelineArtifacts(c *gin.Context) {
	id := c.Param("id")
	var artifacts []models.PipelineArtifact
	h.DB.Where("pipeline_id = ?", id).Find(&artifacts)
	c.JSON(http.StatusOK, gin.H{"data": artifacts})
}

// ============================================================
// Requirement Handlers
// ============================================================

// GetRequirements handles GET /api/requirements.
func (h *PipelineHandler) GetRequirements(c *gin.Context) {
	var requirements []models.Requirement
	query := h.DB
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if priority := c.Query("priority"); priority != "" {
		query = query.Where("priority = ?", priority)
	}
	query.Find(&requirements)
	c.JSON(http.StatusOK, gin.H{"data": requirements})
}

// GetRequirement handles GET /api/requirements/:id.
func (h *PipelineHandler) GetRequirement(c *gin.Context) {
	id := c.Param("id")
	var req models.Requirement
	if err := h.DB.Preload("Pipeline.Stages").First(&req, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

// ============================================================
// Notification Handlers
// ============================================================

// GetNotifications handles GET /api/notifications.
func (h *PipelineHandler) GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	h.DB.Order("created_at desc").Find(&notifications)
	c.JSON(http.StatusOK, gin.H{"data": notifications})
}

// MarkNotificationRead handles PUT /api/notifications/:id/read.
func (h *PipelineHandler) MarkNotificationRead(c *gin.Context) {
	id := c.Param("id")
	var notification models.Notification
	if err := h.DB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	notification.IsRead = true
	h.DB.Save(&notification)
	c.JSON(http.StatusOK, gin.H{"data": notification})
}

// StreamPipelineProgress handles GET /api/pipeline/:id/progress — SSE endpoint.
func (h *PipelineHandler) StreamPipelineProgress(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pipeline id"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	ch := h.EventBus.Subscribe(uint(id))
	defer h.EventBus.Unsubscribe(uint(id), ch)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case event := <-ch:
			fmt.Fprintf(c.Writer, "data: %s\n\n", event.ToJSON())
			c.Writer.Flush()
		case <-ticker.C:
			fmt.Fprintf(c.Writer, "data: {\"type\":\"heartbeat\",\"pipeline_id\":%d}\n\n", id)
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

// ResumePipelineReview handles POST /api/pipeline/:id/resume — approve/reject a paused review node.
func (h *PipelineHandler) ResumePipelineReview(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		StageName  string `json:"stage_name"`
		Decision   string `json:"decision"` // approve, reject, needs_revision
		Comment    string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pipeline models.Pipeline
	if err := h.DB.Preload("Stages").First(&pipeline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pipeline not found"})
		return
	}

	if pipeline.Status != "paused_for_review" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pipeline is not paused for review"})
		return
	}

	// Find the paused stage
	var pausedStage *models.PipelineStage
	for i := range pipeline.Stages {
		if pipeline.Stages[i].Status == "paused_for_review" && pipeline.Stages[i].StageName == input.StageName {
			pausedStage = &pipeline.Stages[i]
			break
		}
	}

	if pausedStage == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no paused review stage found for " + input.StageName})
		return
	}

	switch input.Decision {
	case "approve":
		pausedStage.Status = "completed"
		pausedStage.Signal = "approve"
		pausedStage.Reasoning = "Human review approved: " + input.Comment
		h.DB.Save(&pausedStage)
		h.EventBus.Publish(PipelineEvent{
			PipelineID: pipeline.ID,
			Stage:      pausedStage.StageName,
			Type:       "review_approved",
			Message:    "Review approved: " + input.Comment,
			Status:     "completed",
		})
		// Resume pipeline execution
		pipeline.Status = "running"
		h.DB.Save(&pipeline)
		go h.RunPipelineAsync(pipeline.ID, pipeline.Query, pipeline.Mode)

	case "reject":
		pausedStage.Status = "failed"
		pausedStage.Signal = "block"
		pausedStage.Reasoning = "Human review rejected: " + input.Comment
		h.DB.Save(&pausedStage)
		pipeline.Status = "failed"
		pipeline.Error = "Human review rejected at stage " + input.StageName
		h.DB.Save(&pipeline)
		h.EventBus.Publish(PipelineEvent{
			PipelineID: pipeline.ID,
			Stage:      pausedStage.StageName,
			Type:       "review_rejected",
			Message:    "Review rejected: " + input.Comment,
			Status:     "failed",
		})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "decision must be approve or reject"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pipeline, "decision": input.Decision})
}

// ============================================================
// Workflow Task Handlers
// ============================================================

// GetWorkflowTasks handles GET /api/workflow/tasks — kanban board data.
func (h *PipelineHandler) GetWorkflowTasks(c *gin.Context) {
	var tasks []models.WorkflowTask
	h.DB.Preload("Pipeline").Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
