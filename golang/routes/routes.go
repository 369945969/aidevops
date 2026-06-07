package routes

import (
	"devops/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handler) {
	api := r.Group("/api")

	// 1. Automation Configuration
	automation := api.Group("/automation")
	{
		automation.GET("/config", h.GetAutomationConfig)
		automation.PUT("/config", h.UpdateAutomationConfig)
		automation.GET("/review-nodes", h.GetReviewNodes)
		automation.PUT("/review-nodes/:id", h.UpdateReviewNode)
	}

	// 2. AI Agent Configuration
	aiAgent := api.Group("/ai-agent")
	{
		aiAgent.GET("/teams", h.GetAITeams)
		aiAgent.GET("/teams/:id", h.GetAITeam)
		aiAgent.PUT("/teams/:id", h.UpdateAITeam)
		aiAgent.GET("/agents", h.GetAIAgentsByTeam)
		aiAgent.PUT("/agents/:id", h.UpdateAIAgent)
	}

	// 3. Development Environment
	devEnv := api.Group("/dev-environment")
	{
		devEnv.GET("/code-repos", h.GetCodeRepos)
		devEnv.PUT("/code-repos/:id", h.UpdateCodeRepo)
		devEnv.GET("/cicd-services", h.GetCICDServices)
		devEnv.PUT("/cicd-services/:id", h.UpdateCICDService)
		devEnv.GET("/cloud-credentials", h.GetCloudCredentials)
		devEnv.PUT("/cloud-credentials/:id", h.UpdateCloudCredential)
		devEnv.GET("/ssh-keys", h.GetSSHKeys)
		devEnv.POST("/ssh-keys", h.CreateSSHKey)
		devEnv.DELETE("/ssh-keys/:id", h.DeleteSSHKey)
		devEnv.GET("/env-vars", h.GetEnvVars)
		devEnv.POST("/env-vars", h.CreateEnvVar)
		devEnv.PUT("/env-vars/:id", h.UpdateEnvVar)
		devEnv.DELETE("/env-vars/:id", h.DeleteEnvVar)
	}

	// 4. Integration Configuration
	integrations := api.Group("/integrations")
	{
		integrations.GET("/services", h.GetConnectedServices)
		integrations.POST("/services", h.CreateConnectedService)
		integrations.PUT("/services/:id", h.UpdateConnectedService)
		integrations.DELETE("/services/:id", h.DeleteConnectedService)
		integrations.GET("/api-keys", h.GetAPIKeys)
		integrations.POST("/api-keys", h.CreateAPIKey)
		integrations.DELETE("/api-keys/:id", h.DeleteAPIKey)
		integrations.PUT("/api-keys/:id/refresh", h.RefreshAPIKey)
		integrations.GET("/webhook", h.GetWebhookConfigs)
		integrations.PUT("/webhook/:id", h.UpdateWebhookConfig)
		integrations.GET("/sync", h.GetSyncConfig)
		integrations.PUT("/sync", h.UpdateSyncConfig)
	}

	// 5. Model Management
	models := api.Group("/models")
	{
		models.GET("", h.GetAIModels)
		models.POST("", h.CreateAIModel)
		models.PUT("/:id", h.UpdateAIModel)
		models.DELETE("/:id", h.DeleteAIModel)
		models.POST("/:id/test", h.TestAIModelConnection)
	}

	// 6. Notification Settings
	notifications := api.Group("/notifications")
	{
		notifications.GET("/channels", h.GetNotificationChannels)
		notifications.PUT("/channels/:id", h.UpdateNotificationChannel)
		notifications.GET("/events", h.GetNotificationEvents)
		notifications.PUT("/events/:id", h.UpdateNotificationEvent)
		notifications.GET("/preferences", h.GetNotificationPreferences)
		notifications.PUT("/preferences", h.UpdateNotificationPreferences)
	}

	// 7. Review Rules
	reviewRules := api.Group("/review-rules")
	{
		reviewRules.GET("", h.GetReviewRules)
		reviewRules.POST("", h.CreateReviewRule)
		reviewRules.PUT("/:id", h.UpdateReviewRule)
		reviewRules.DELETE("/:id", h.DeleteReviewRule)
		reviewRules.GET("/templates", h.GetReviewTemplates)
		reviewRules.GET("/timeout", h.GetReviewTimeoutConfig)
		reviewRules.PUT("/timeout", h.UpdateReviewTimeoutConfig)
	}

	// 8. Security & Permissions
	security := api.Group("/security")
	{
		security.GET("/roles", h.GetRoles)
		security.GET("/permissions", h.GetPermissions)
		security.GET("/role-permissions", h.GetRolePermissions)
		security.PUT("/role-permissions/:id", h.UpdateRolePermission)
		security.GET("/policies", h.GetSecurityPolicies)
		security.PUT("/policies/:id", h.UpdateSecurityPolicy)
		security.GET("/password-policy", h.GetPasswordPolicy)
		security.PUT("/password-policy", h.UpdatePasswordPolicy)
		security.GET("/audit-logs", h.GetAuditLogs)
		security.POST("/audit-logs", h.CreateAuditLog)
	}

	// 9. Team Members
	teamMembers := api.Group("/team-members")
	{
		teamMembers.GET("", h.GetTeamMembers)
		teamMembers.POST("", h.CreateTeamMember)
		teamMembers.PUT("/:id", h.UpdateTeamMember)
		teamMembers.DELETE("/:id", h.DeleteTeamMember)
		teamMembers.GET("/invites", h.GetPendingInvites)
		teamMembers.POST("/invites", h.CreatePendingInvite)
		teamMembers.DELETE("/invites/:id", h.DeletePendingInvite)
		teamMembers.GET("/bindings", h.GetMemberAITeamBindings)
		teamMembers.POST("/bindings", h.CreateMemberAITeamBinding)
		teamMembers.DELETE("/bindings/:id", h.DeleteMemberAITeamBinding)
	}
}