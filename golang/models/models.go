package models

import "time"

// ============================================================
// 1. Automation Configuration
// ============================================================

type AutomationConfig struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Level               string    `gorm:"default:balanced" json:"level"`
	TechPreference      string    `gorm:"default:neutral" json:"tech_preference"`
	CodeStyle           string    `gorm:"default:concise" json:"code_style"`
	TestCoverage        int       `gorm:"default:80" json:"test_coverage"`
	AnalysisTimeout     int       `gorm:"default:30" json:"analysis_timeout"`
	ArchitectureTimeout int       `gorm:"default:60" json:"architecture_timeout"`
	DevelopmentTimeout  int       `gorm:"default:120" json:"development_timeout"`
	TestExecutionTimeout int      `gorm:"default:60" json:"test_execution_timeout"`
	AutoRetryCount      int       `gorm:"default:2" json:"auto_retry_count"`
	RetryInterval       int       `gorm:"default:5" json:"retry_interval"`
	FailureHandling     string    `gorm:"default:notify_human" json:"failure_handling"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type ReviewNode struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Description string   `gorm:"default:''" json:"description"`
	Enabled    bool      `gorm:"default:true" json:"enabled"`
	SubOption  string    `gorm:"default:''" json:"sub_option"`
	SubInput   string    `gorm:"default:''" json:"sub_input"`
	SortOrder  int       `gorm:"default:0" json:"sort_order"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ============================================================
// 2. AI Agent Configuration
// ============================================================

type AITeam struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name"`
	Icon             string    `gorm:"default:''" json:"icon"`
	Status           string    `gorm:"default:running" json:"status"`
	Enabled          bool      `gorm:"default:true" json:"enabled"`
	BoundMemberID    *uint     `json:"bound_member_id"`
	BoundMember      *TeamMember `gorm:"foreignKey:BoundMemberID" json:"bound_member,omitempty"`
	ConsensusPassRate float64   `gorm:"default:0" json:"consensus_pass_rate"`
	TotalTasks       int       `gorm:"default:0" json:"total_tasks"`
	AvgDuration      int       `gorm:"default:0" json:"avg_duration"`
	SuccessRate      float64   `gorm:"default:0" json:"success_rate"`
	Agents           []AIAgent `gorm:"foreignKey:TeamID" json:"agents,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type AIAgent struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TeamID          uint      `json:"team_id"`
	Team            AITeam    `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Role            string    `json:"role"`
	PrimaryModelID  *uint     `json:"primary_model_id"`
	PrimaryModel    *AIModel  `gorm:"foreignKey:PrimaryModelID" json:"primary_model,omitempty"`
	FallbackModelID *uint     `json:"fallback_model_id"`
	FallbackModel   *AIModel  `gorm:"foreignKey:FallbackModelID" json:"fallback_model,omitempty"`
	TechPreference  string    `gorm:"default:structured_analysis" json:"tech_preference"`
	BehaviorParam   float64   `gorm:"default:0.7" json:"behavior_param"`
	TrustLevel      float64   `gorm:"default:0.85" json:"trust_level"`
	RigorLevel      float64   `gorm:"default:0.9" json:"rigor_level"`
	TotalTasks      int       `gorm:"default:0" json:"total_tasks"`
	AvgDuration     int       `gorm:"default:0" json:"avg_duration"`
	SuccessRate     float64   `gorm:"default:0" json:"success_rate"`
	ConsensusStatus string    `gorm:"default:pending" json:"consensus_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ============================================================
// 3. Development Environment
// ============================================================

type CodeRepo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Provider  string    `json:"provider"`
	Connected bool      `gorm:"default:false" json:"connected"`
	RepoURL   string    `gorm:"default:''" json:"repo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CICDService struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Provider     string    `json:"provider"`
	Configured   bool      `gorm:"default:false" json:"configured"`
	ServiceURL   string    `gorm:"default:''" json:"service_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CloudCredential struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Provider        string    `json:"provider"`
	AccessKeyID     string    `gorm:"default:''" json:"access_key_id"`
	SecretAccessKey string    `gorm:"default:''" json:"secret_access_key"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SSHKey struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	KeyType     string    `gorm:"default:'RSA 2048'" json:"key_type"`
	Description string    `gorm:"default:''" json:"description"`
	KeyContent  string    `gorm:"default:''" json:"key_content"`
	CreatedAt   time.Time `json:"created_at"`
	LastUsedAt  *time.Time `json:"last_used_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EnvVar struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Key       string    `gorm:"unique" json:"key"`
	Value     string    `gorm:"default:''" json:"value"`
	IsSecret  bool      `gorm:"default:false" json:"is_secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ============================================================
// 4. Integration Configuration
// ============================================================

type ConnectedService struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Icon       string    `gorm:"default:''" json:"icon"`
	Status     string    `gorm:"default:pending" json:"status"`
	ConfigData string    `gorm:"default:''" json:"config_data"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type APIKey struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Service     string    `json:"service"`
	KeyValue    string    `gorm:"default:''" json:"key_value"`
	Permissions string    `gorm:"default:''" json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WebhookConfig struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	URL          string    `gorm:"default:''" json:"url"`
	EventTypes   string    `gorm:"default:''" json:"event_types"`
	SignatureKey string    `gorm:"default:''" json:"signature_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SyncConfig struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	LastSyncAt  *time.Time `json:"last_sync_at"`
	NextSyncAt  *time.Time `json:"next_sync_at"`
	Frequency   string     `gorm:"default:hourly" json:"frequency"`
	Status      string     `gorm:"default:running" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ============================================================
// 5. Model Management
// ============================================================

type AIModel struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name"`
	Provider         string    `gorm:"default:''" json:"provider"`
	IsBuiltin        bool      `gorm:"default:false" json:"is_builtin"`
	APIEndpoint      string    `gorm:"default:''" json:"api_endpoint"`
	APIKey           string    `gorm:"default:''" json:"api_key"`
	MaxContextLength int       `gorm:"default:8192" json:"max_context_length"`
	ConnectionStatus string    `gorm:"default:not_tested" json:"connection_status"`
	CapabilityTags   string    `gorm:"default:''" json:"capability_tags"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ============================================================
// 6. Notification Settings
// ============================================================

type NotificationChannel struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ChannelType string   `gorm:"unique" json:"channel_type"`
	Enabled    bool      `gorm:"default:true" json:"enabled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type NotificationEvent struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventType string    `gorm:"unique" json:"event_type"`
	Icon      string    `gorm:"default:''" json:"icon"`
	Enabled   bool      `gorm:"default:true" json:"enabled"`
	Severity  string    `gorm:"default:info" json:"severity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NotificationPreference struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Frequency        string    `gorm:"default:instant" json:"frequency"`
	QuietHoursEnabled bool     `gorm:"default:false" json:"quiet_hours_enabled"`
	QuietHoursStart  string    `gorm:"default:'22:00'" json:"quiet_hours_start"`
	QuietHoursEnd    string    `gorm:"default:'08:00'" json:"quiet_hours_end"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ============================================================
// 7. Review Rules
// ============================================================

type ReviewRule struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	Description     string    `gorm:"default:''" json:"description"`
	TriggerCondition string   `gorm:"default:''" json:"trigger_condition"`
	HandlingMethod  string    `gorm:"default:''" json:"handling_method"`
	Enabled         bool      `gorm:"default:true" json:"enabled"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ReviewTemplate struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `gorm:"default:''" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ReviewTimeoutConfig struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Enabled        bool      `gorm:"default:true" json:"enabled"`
	ThresholdHours int       `gorm:"default:24" json:"threshold_hours"`
	Strategy       string    `gorm:"default:auto_approve_mark" json:"strategy"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ============================================================
// 8. Security & Permissions
// ============================================================

type Role struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique" json:"name"`
	DisplayName string    `gorm:"default:''" json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Permission struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Module    string    `json:"module"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RolePermission struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RoleID       uint      `json:"role_id"`
	Role         Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	PermissionID uint      `json:"permission_id"`
	Permission   Permission `gorm:"foreignKey:PermissionID" json:"permission,omitempty"`
	AccessLevel  string    `gorm:"default:granted" json:"access_level"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SecurityPolicy struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PolicyType  string    `gorm:"unique" json:"policy_type"`
	Enabled     bool      `gorm:"default:false" json:"enabled"`
	Description string    `gorm:"default:''" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PasswordPolicy struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	MinLength   int       `gorm:"default:12" json:"min_length"`
	ExpiryDays  int       `gorm:"default:90" json:"expiry_days"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AuditLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Operator  string    `json:"operator"`
	Action    string    `json:"action"`
	Target    string    `gorm:"default:''" json:"target"`
	Result    string    `gorm:"default:success" json:"result"`
	Detail    string    `gorm:"default:''" json:"detail"`
	CreatedAt time.Time `json:"created_at"`
}

// ============================================================
// 9. Team Members
// ============================================================

type TeamMember struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Title        string    `gorm:"default:''" json:"title"`
	Email        string    `gorm:"default:''" json:"email"`
	Avatar       string    `gorm:"default:''" json:"avatar"`
	Status       string    `gorm:"default:offline" json:"status"`
	RoleID       *uint     `json:"role_id"`
	Role         *Role     `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	ActivityRate float64   `gorm:"default:0" json:"activity_rate"`
	TeamBindings []MemberAITeamBinding `gorm:"foreignKey:MemberID" json:"team_bindings,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PendingInvite struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Email       string    `json:"email"`
	RoleID      *uint     `json:"role_id"`
	InviteToken string    `gorm:"default:''" json:"invite_token"`
	Status      string    `gorm:"default:pending" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MemberAITeamBinding struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	MemberID  uint      `json:"member_id"`
	Member    TeamMember `gorm:"foreignKey:MemberID" json:"member,omitempty"`
	TeamID    uint      `json:"team_id"`
	Team      AITeam    `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}