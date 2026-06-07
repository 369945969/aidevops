package main

import (
	"devops/models"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	// Only seed if tables are empty

	// Roles
	if db.First(&models.Role{}).Error != nil {
		db.Create(&models.Role{Name: "admin", DisplayName: "超级管理员"})
		db.Create(&models.Role{Name: "tech_lead", DisplayName: "技术负责人"})
		db.Create(&models.Role{Name: "developer", DisplayName: "开发者"})
		db.Create(&models.Role{Name: "observer", DisplayName: "观察者"})
	}

	// Permissions
	if db.First(&models.Permission{}).Error != nil {
	 perms := []models.Permission{
			{Module: "system_config", Action: "view"},
			{Module: "system_config", Action: "edit"},
			{Module: "requirement", Action: "view"},
			{Module: "requirement", Action: "create"},
			{Module: "code_review", Action: "approve"},
			{Module: "deploy", Action: "execute"},
			{Module: "knowledge", Action: "manage"},
		}
		db.Create(&perms)
	}

	// Role-Permissions (admin=all granted)
	if db.First(&models.RolePermission{}).Error != nil {
		// admin (id=1): all granted
		for i := 1; i <= 7; i++ {
			db.Create(&models.RolePermission{RoleID: 1, PermissionID: uint(i), AccessLevel: "granted"})
		}
		// tech_lead (id=2)
		rp2 := []models.RolePermission{
			{RoleID: 2, PermissionID: 1, AccessLevel: "granted"},
			{RoleID: 2, PermissionID: 2, AccessLevel: "partial"},
			{RoleID: 2, PermissionID: 3, AccessLevel: "granted"},
			{RoleID: 2, PermissionID: 4, AccessLevel: "granted"},
			{RoleID: 2, PermissionID: 5, AccessLevel: "granted"},
			{RoleID: 2, PermissionID: 6, AccessLevel: "partial"},
			{RoleID: 2, PermissionID: 7, AccessLevel: "granted"},
		}
		db.Create(&rp2)
		// developer (id=3)
		rp3 := []models.RolePermission{
			{RoleID: 3, PermissionID: 1, AccessLevel: "partial"},
			{RoleID: 3, PermissionID: 2, AccessLevel: "denied"},
			{RoleID: 3, PermissionID: 3, AccessLevel: "granted"},
			{RoleID: 3, PermissionID: 4, AccessLevel: "granted"},
			{RoleID: 3, PermissionID: 5, AccessLevel: "partial"},
			{RoleID: 3, PermissionID: 6, AccessLevel: "denied"},
			{RoleID: 3, PermissionID: 7, AccessLevel: "partial"},
		}
		db.Create(&rp3)
		// observer (id=4): only view
		rp4 := []models.RolePermission{
			{RoleID: 4, PermissionID: 1, AccessLevel: "granted"},
			{RoleID: 4, PermissionID: 2, AccessLevel: "denied"},
			{RoleID: 4, PermissionID: 3, AccessLevel: "granted"},
			{RoleID: 4, PermissionID: 4, AccessLevel: "denied"},
			{RoleID: 4, PermissionID: 5, AccessLevel: "denied"},
			{RoleID: 4, PermissionID: 6, AccessLevel: "denied"},
			{RoleID: 4, PermissionID: 7, AccessLevel: "denied"},
		}
		db.Create(&rp4)
	}

	// Security policies
	if db.First(&models.SecurityPolicy{}).Error != nil {
		db.Create(&[]models.SecurityPolicy{
			{PolicyType: "db_encryption", Enabled: true, Description: "数据库加密 (AES-256)"},
			{PolicyType: "api_encryption", Enabled: true, Description: "API通信加密 (TLS 1.3)"},
			{PolicyType: "two_factor_auth", Enabled: true, Description: "双因素认证 (TOTP)"},
			{PolicyType: "ip_whitelist", Enabled: false, Description: "IP白名单"},
		})
	}

	// Password policy
	if db.First(&models.PasswordPolicy{}).Error != nil {
		db.Create(&models.PasswordPolicy{MinLength: 12, ExpiryDays: 90})
	}

	// Notification channels
	if db.First(&models.NotificationChannel{}).Error != nil {
		db.Create(&[]models.NotificationChannel{
			{ChannelType: "email", Enabled: true},
			{ChannelType: "slack", Enabled: true},
			{ChannelType: "in_app", Enabled: true},
			{ChannelType: "sms", Enabled: false},
		})
	}

	// Notification events
	if db.First(&models.NotificationEvent{}).Error != nil {
		db.Create(&[]models.NotificationEvent{
			{EventType: "review_passed", Icon: "check-circle", Enabled: true, Severity: "info"},
			{EventType: "merge_request", Icon: "git-merge", Enabled: true, Severity: "info"},
			{EventType: "deploy_completed", Icon: "rocket", Enabled: true, Severity: "info"},
			{EventType: "review_timeout", Icon: "clock", Enabled: false, Severity: "warning"},
			{EventType: "agent_exception", Icon: "alert-triangle", Enabled: true, Severity: "critical"},
		})
	}

	// Notification preferences
	if db.First(&models.NotificationPreference{}).Error != nil {
		db.Create(&models.NotificationPreference{Frequency: "instant", QuietHoursEnabled: true, QuietHoursStart: "22:00", QuietHoursEnd: "08:00"})
	}

	// Automation config
	if db.First(&models.AutomationConfig{}).Error != nil {
		db.Create(&models.AutomationConfig{
			Level: "balanced", TechPreference: "neutral", CodeStyle: "concise",
			TestCoverage: 80, AnalysisTimeout: 30, ArchitectureTimeout: 60,
			DevelopmentTimeout: 120, TestExecutionTimeout: 60,
			AutoRetryCount: 2, RetryInterval: 5, FailureHandling: "notify_human",
		})
	}

	// Review nodes
	if db.First(&models.ReviewNode{}).Error != nil {
		db.Create(&[]models.ReviewNode{
			{Name: "requirement_confirm", Description: "需求理解确认", Enabled: true, SortOrder: 1},
			{Name: "architecture_review", Description: "架构方案审核", Enabled: true, SortOrder: 2},
			{Name: "code_merge_review", Description: "代码合并审核", Enabled: true, SubOption: "core_modules_only", SubInput: "auth, payment, user", SortOrder: 3},
			{Name: "test_confirm", Description: "测试结果确认", Enabled: false, SortOrder: 4},
			{Name: "deploy_review", Description: "集群部署审核", Enabled: true, SubOption: "production_only", SortOrder: 5},
		})
	}

	// AI Models (built-in)
	if db.First(&models.AIModel{}).Error != nil {
		db.Create(&[]models.AIModel{
			{Name: "Claude Opus", Provider: "Anthropic", IsBuiltin: true, MaxContextLength: 200000, ConnectionStatus: "connected", CapabilityTags: "complex_reasoning,long_text,multi_step_reasoning"},
			{Name: "Claude Sonnet", Provider: "Anthropic", IsBuiltin: true, MaxContextLength: 128000, ConnectionStatus: "connected", CapabilityTags: "balanced_performance,fast_output,code_generation"},
			{Name: "GPT-4o", Provider: "OpenAI", IsBuiltin: true, MaxContextLength: 128000, ConnectionStatus: "connected", CapabilityTags: "multi_modal,fast_reasoning,general"},
			{Name: "Gemini Pro", Provider: "Google", IsBuiltin: true, MaxContextLength: 1000000, ConnectionStatus: "not_tested", CapabilityTags: "multi_modal,long_context"},
			{Name: "DeepSeek", Provider: "DeepSeek", IsBuiltin: true, MaxContextLength: 64000, ConnectionStatus: "connected", CapabilityTags: "code_generation,math_reasoning,low_cost"},
		})
	}

	// Review rules
	if db.First(&models.ReviewRule{}).Error != nil {
		db.Create(&[]models.ReviewRule{
			{Name: "架构变更审核", Description: "核心架构模块变更时触发审核", TriggerCondition: "核心模块变更", HandlingMethod: "技术负责人审核", Enabled: true},
			{Name: "代码合并审核", Description: "安全敏感文件的合并请求", TriggerCondition: "安全敏感文件", HandlingMethod: "双人审核", Enabled: true},
			{Name: "生产部署审核", Description: "生产环境部署操作", TriggerCondition: "生产环境部署", HandlingMethod: "运维负责人审核", Enabled: true},
			{Name: "依赖变更审核", Description: "依赖版本变更", TriggerCondition: "依赖版本变更", HandlingMethod: "自动安全扫描", Enabled: false},
		})
	}

	// Review template
	if db.First(&models.ReviewTemplate{}).Error != nil {
		db.Create(&models.ReviewTemplate{Name: "安全优先模板", Description: "包含6个预定义审核规则，涵盖架构变更、代码合并、生产部署等"})
	}

	// Review timeout config
	if db.First(&models.ReviewTimeoutConfig{}).Error != nil {
		db.Create(&models.ReviewTimeoutConfig{Enabled: true, ThresholdHours: 24, Strategy: "auto_approve_mark"})
	}

	// Code repos
	if db.First(&models.CodeRepo{}).Error != nil {
		db.Create(&[]models.CodeRepo{
			{Provider: "github", Connected: true, RepoURL: "https://github.com/ai-devops/main-project"},
			{Provider: "gitlab", Connected: false},
		})
	}

	// CI/CD services
	if db.First(&models.CICDService{}).Error != nil {
		db.Create(&[]models.CICDService{
			{Provider: "jenkins", Configured: true, ServiceURL: "https://jenkins.ai-devops.com"},
			{Provider: "gitlab_ci", Configured: false},
		})
	}

	// Cloud credentials
	if db.First(&models.CloudCredential{}).Error != nil {
		db.Create(&[]models.CloudCredential{
			{Provider: "aws", AccessKeyID: "AKIA****************", SecretAccessKey: "********************"},
			{Provider: "alibaba_cloud", AccessKeyID: "LTAI****************", SecretAccessKey: "********************"},
		})
	}

	// SSH keys
	if db.First(&models.SSHKey{}).Error != nil {
		db.Create(&[]models.SSHKey{
			{Name: "部署密钥", KeyType: "RSA 2048", Description: "用于生产服务器部署访问"},
			{Name: "CI密钥", KeyType: "Ed25519", Description: "用于CI/CD流水线自动拉取代码"},
		})
	}

	// Environment variables
	if db.First(&models.EnvVar{}).Error != nil {
		db.Create(&[]models.EnvVar{
			{Key: "NODE_ENV", Value: "production", IsSecret: false},
			{Key: "API_BASE_URL", Value: "https://api.ai-devops.com/v1", IsSecret: false},
			{Key: "DB_CONNECTION", Value: "mongodb+srv://***:***@cluster.ai-devops.mongodb.net", IsSecret: true},
		})
	}

	// Connected services
	if db.First(&models.ConnectedService{}).Error != nil {
		db.Create(&[]models.ConnectedService{
			{Name: "GitHub Enterprise", Icon: "lucide:github", Status: "connected", ConfigData: "代码仓库与CI/CD集成"},
			{Name: "Slack", Icon: "lucide:message-circle", Status: "connected", ConfigData: "即时通讯与通知集成"},
			{Name: "Jira", Icon: "lucide:trello", Status: "pending", ConfigData: "项目管理与需求追踪"},
			{Name: "AWS CloudWatch", Icon: "lucide:cloud", Status: "connected", ConfigData: "云监控与日志集成"},
		})
	}

	// API keys
	if db.First(&models.APIKey{}).Error != nil {
		db.Create(&[]models.APIKey{
			{Name: "GitHub API Token", Service: "github", KeyValue: "ghp_9xK4***", Permissions: "repo, workflow"},
			{Name: "Slack Bot Token", Service: "slack", KeyValue: "xoxb-2847***", Permissions: "channels:read, chat:write"},
			{Name: "AWS Access Key", Service: "aws", KeyValue: "AKIA***3QR", Permissions: "cloudwatch, ec2"},
		})
	}

	// Webhook config
	if db.First(&models.WebhookConfig{}).Error != nil {
		db.Create(&models.WebhookConfig{URL: "https://ai-devops.com/api/webhook/events", EventTypes: "push, pull_request", SignatureKey: "whsec_***"})
	}

	// Sync config
	if db.First(&models.SyncConfig{}).Error != nil {
		db.Create(&models.SyncConfig{Frequency: "hourly", Status: "running"})
	}

	// Audit logs (sample)
	if db.First(&models.AuditLog{}).Error != nil {
		db.Create(&[]models.AuditLog{
			{Operator: "Alex", Action: "修改自动化配置", Target: "automation_config", Result: "success"},
			{Operator: "Sarah", Action: "审核REQ-0008架构方案", Target: "requirement", Result: "success"},
			{Operator: "Alex", Action: "添加团队成员Bob", Target: "team_member", Result: "success"},
			{Operator: "Alex", Action: "更新安全策略配置", Target: "security_policy", Result: "success"},
		})
	}

	// Team members
	if db.First(&models.TeamMember{}).Error != nil {
		db.Create(&[]models.TeamMember{
			{Name: "Alex Chen", Title: "技术负责人", Email: "alex.chen@company.com", Avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=alex", Status: "online", RoleID: uintPtr(1), ActivityRate: 95},
			{Name: "Sarah Wang", Title: "产品经理", Email: "sarah.wang@company.com", Avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=sarah", Status: "online", RoleID: uintPtr(2), ActivityRate: 80},
			{Name: "Mike Liu", Title: "前端开发者", Email: "mike.liu@company.com", Avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=mike", Status: "offline", RoleID: uintPtr(3), ActivityRate: 60},
			{Name: "Emma Zhang", Title: "后端开发者", Email: "emma.zhang@company.com", Avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=emily", Status: "online", RoleID: uintPtr(3), ActivityRate: 85},
			{Name: "David Li", Title: "DevOps工程师", Email: "david.li@company.com", Avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=david", Status: "offline", RoleID: uintPtr(3), ActivityRate: 45},
		})
	}

	// Pending invites
	if db.First(&models.PendingInvite{}).Error != nil {
		db.Create(&[]models.PendingInvite{
			{Email: "david.liu@company.com", RoleID: uintPtr(3), Status: "pending"},
			{Email: "jenny.wu@company.com", RoleID: uintPtr(4), Status: "pending"},
		})
	}

	// AI teams
	if db.First(&models.AITeam{}).Error != nil {
		db.Create(&[]models.AITeam{
			{Name: "需求分析师 Team", Icon: "lucide:file-text", Status: "running", Enabled: true, BoundMemberID: uintPtr(2), ConsensusPassRate: 0.67, TotalTasks: 42, AvgDuration: 20, SuccessRate: 93},
			{Name: "架构师 Team", Icon: "lucide:layout", Status: "running", Enabled: true, BoundMemberID: uintPtr(1), ConsensusPassRate: 1.0, TotalTasks: 35, AvgDuration: 30, SuccessRate: 96},
			{Name: "前端开发 Team", Icon: "lucide:monitor", Status: "running", Enabled: true, BoundMemberID: uintPtr(3), ConsensusPassRate: 0.5, TotalTasks: 28, AvgDuration: 45, SuccessRate: 89},
			{Name: "后端开发 Team", Icon: "lucide:database", Status: "running", Enabled: true, BoundMemberID: uintPtr(4), ConsensusPassRate: 1.0, TotalTasks: 38, AvgDuration: 35, SuccessRate: 92},
			{Name: "测试 Team", Icon: "lucide:test-tubes", Status: "standby", Enabled: true, BoundMemberID: uintPtr(1), ConsensusPassRate: 0, TotalTasks: 0, AvgDuration: 0, SuccessRate: 0},
			{Name: "DevOps Team", Icon: "lucide:server", Status: "running", Enabled: true, BoundMemberID: uintPtr(5), ConsensusPassRate: 1.0, TotalTasks: 20, AvgDuration: 15, SuccessRate: 95},
		})
	}

	// AI agents (for requirement analyst team, id=1)
	if db.First(&models.AIAgent{}).Error != nil {
		db.Create(&[]models.AIAgent{
			{TeamID: 1, Role: "撰写", PrimaryModelID: uintPtr(1), FallbackModelID: uintPtr(3), TechPreference: "structured_analysis", BehaviorParam: 0.7, TrustLevel: 0.85, TotalTasks: 42, AvgDuration: 15, SuccessRate: 97, ConsensusStatus: "confirmed"},
			{TeamID: 1, Role: "审查", PrimaryModelID: uintPtr(2), FallbackModelID: uintPtr(3), TechPreference: "structured_analysis", BehaviorParam: 0.7, TrustLevel: 0.75, RigorLevel: 0.9, TotalTasks: 38, AvgDuration: 10, SuccessRate: 95, ConsensusStatus: "confirmed"},
		})
	}

	// Member-AI team bindings
	if db.First(&models.MemberAITeamBinding{}).Error != nil {
		db.Create(&[]models.MemberAITeamBinding{
			{MemberID: 1, TeamID: 2},
			{MemberID: 1, TeamID: 5},
			{MemberID: 2, TeamID: 1},
			{MemberID: 3, TeamID: 3},
			{MemberID: 4, TeamID: 4},
			{MemberID: 5, TeamID: 6},
		})
	}
}

func uintPtr(v uint) *uint {
	return &v
}