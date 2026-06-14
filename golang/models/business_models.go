package models

import "time"

// ============================================================
// Core Business Models — Requirements, Pipelines, Notifications
// ============================================================

// Requirement represents a user-submitted requirement that triggers
// a DevOps pipeline. This is the "single sentence" entry point.
type Requirement struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `gorm:"default:需求分析" json:"status"`
	Priority    string    `gorm:"default:中" json:"priority"`
	Author      string    `gorm:"default:''" json:"author"`
	AuthorAvatar string   `gorm:"default:''" json:"author_avatar"`
	TaskCount   int       `gorm:"default:0" json:"task_count"`
	PipelineID  *uint     `json:"pipeline_id"`
	Pipeline    *Pipeline `gorm:"foreignKey:PipelineID" json:"pipeline,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Pipeline represents a multi-agent execution pipeline triggered by a requirement.
// The pipeline tracks the full chain: RequirementAnalyst → Architect → Developer → Tester → Deployer.
type Pipeline struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	RequirementID uint      `json:"requirement_id"`
	Requirement   Requirement `gorm:"foreignKey:RequirementID" json:"requirement,omitempty"`
	Mode          string    `gorm:"default:full" json:"mode"` // quick, standard, full, review
	Status        string    `gorm:"default:pending" json:"status"` // pending, running, completed, failed, paused_for_review
	Query         string    `json:"query"` // original requirement text
	Specification string    `gorm:"type:text" json:"specification"`
	Architecture  string    `gorm:"type:text" json:"architecture"`
	CodeArtifacts string    `gorm:"type:text" json:"code_artifacts"` // JSON-encoded map of filename→code
	TestResults   string    `gorm:"type:text" json:"test_results"`
	DeployPlan    string    `gorm:"type:text" json:"deploy_plan"`
	Result        string    `gorm:"type:text" json:"result"` // full pipeline result summary
	Error         string    `gorm:"default:''" json:"error"`
	TotalTokens   int       `gorm:"default:0" json:"total_tokens"`
	TotalDurationMs int64   `gorm:"default:0" json:"total_duration_ms"`
	ModelsUsed    string    `gorm:"default:''" json:"models_used"` // comma-separated
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Stages        []PipelineStage `gorm:"foreignKey:PipelineID" json:"stages,omitempty"`
}

// PipelineStage represents a single agent execution within a pipeline.
// Each stage corresponds to one DevOpsAgent in the chain.
type PipelineStage struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PipelineID  uint      `json:"pipeline_id"`
	Pipeline    Pipeline  `gorm:"foreignKey:PipelineID" json:"pipeline,omitempty"`
	StageName   string    `json:"stage_name"` // requirement_analyst, architect, developer, tester, deployer
	Status      string    `gorm:"default:pending" json:"status"` // pending, running, completed, failed, skipped
	Output      string    `gorm:"type:text" json:"output"` // AgentOutput.Content
	Signal      string    `gorm:"default:''" json:"signal"` // approve, needs_revision, block
	Confidence  float64   `gorm:"default:0" json:"confidence"`
	Reasoning   string    `gorm:"type:text" json:"reasoning"`
	Error       string    `gorm:"default:''" json:"error"`
	DurationMs  int64     `gorm:"default:0" json:"duration_ms"`
	TokensUsed  int       `gorm:"default:0" json:"tokens_used"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PipelineArtifact stores individual code files or other artifacts produced by agents.
type PipelineArtifact struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PipelineID uint      `json:"pipeline_id"`
	Pipeline   Pipeline  `gorm:"foreignKey:PipelineID" json:"pipeline,omitempty"`
	StageName  string    `json:"stage_name"` // which agent produced this
	Type       string    `json:"type"` // code, test, deploy_config, etc.
	Name       string    `json:"name"` // filename or artifact name
	Content    string    `gorm:"type:text" json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

// Notification represents an in-app notification for the DevOps platform.
type Notification struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"` // review, progress, system, timeout, agent_exception
	TypeColor   string    `gorm:"default:''" json:"type_color"`
	TypeBg      string    `gorm:"default:''" json:"type_bg"`
	Icon        string    `gorm:"default:''" json:"icon"`
	BorderColor string    `gorm:"default:''" json:"border_color"`
	IconBg      string    `gorm:"default:''" json:"icon_bg"`
	Description string    `json:"description"`
	Time        string    `json:"time"`
	Route       string    `gorm:"default:''" json:"route"`
	IsRead      bool      `gorm:"default:false" json:"is_read"`
	CreatedAt   time.Time `json:"created_at"`
}

// WorkflowTask represents a task on the kanban dashboard.
// This is derived from the pipeline stages and task decomposition.
type WorkflowTask struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PipelineID   uint      `json:"pipeline_id"`
	Pipeline     Pipeline  `gorm:"foreignKey:PipelineID" json:"pipeline,omitempty"`
	Title        string    `json:"title"`
	Agent        string    `gorm:"default:''" json:"agent"` // e.g. "后端开发"
	Priority     string    `gorm:"default:中" json:"priority"`
	Progress     int       `gorm:"default:0" json:"progress"` // 0-100
	Estimate     string    `gorm:"default:''" json:"estimate"`
	Status       string    `gorm:"default:需求分析" json:"status"`
	ReviewRoute  string    `gorm:"default:''" json:"review_route"`
	Extra        string    `gorm:"default:''" json:"extra"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
