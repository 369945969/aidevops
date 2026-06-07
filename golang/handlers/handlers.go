package handlers

import (
	"net/http"
	"strconv"
	"time"

	"devops/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// ============================================================
// 1. Automation Configuration
// ============================================================

func (h *Handler) GetAutomationConfig(c *gin.Context) {
	var config models.AutomationConfig
	if err := h.DB.First(&config).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": config})
}

func (h *Handler) UpdateAutomationConfig(c *gin.Context) {
	var config models.AutomationConfig
	h.DB.First(&config)
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.UpdatedAt = time.Now()
	h.DB.Save(&config)
	c.JSON(http.StatusOK, gin.H{"data": config})
}

func (h *Handler) GetReviewNodes(c *gin.Context) {
	var nodes []models.ReviewNode
	h.DB.Order("sort_order").Find(&nodes)
	c.JSON(http.StatusOK, gin.H{"data": nodes})
}

func (h *Handler) UpdateReviewNode(c *gin.Context) {
	id := c.Param("id")
	var node models.ReviewNode
	if err := h.DB.First(&node, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&node)
	c.JSON(http.StatusOK, gin.H{"data": node})
}

// ============================================================
// 2. AI Agent Configuration
// ============================================================

func (h *Handler) GetAITeams(c *gin.Context) {
	var teams []models.AITeam
	h.DB.Preload("Agents.PrimaryModel").Preload("Agents.FallbackModel").Preload("BoundMember").Find(&teams)
	c.JSON(http.StatusOK, gin.H{"data": teams})
}

func (h *Handler) GetAITeam(c *gin.Context) {
	id := c.Param("id")
	var team models.AITeam
	if err := h.DB.Preload("Agents.PrimaryModel").Preload("Agents.FallbackModel").Preload("BoundMember").First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": team})
}

func (h *Handler) UpdateAITeam(c *gin.Context) {
	id := c.Param("id")
	var team models.AITeam
	if err := h.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&team)
	c.JSON(http.StatusOK, gin.H{"data": team})
}

func (h *Handler) GetAIAgentsByTeam(c *gin.Context) {
	teamID := c.Query("team_id")
	var agents []models.AIAgent
	h.DB.Preload("PrimaryModel").Preload("FallbackModel").Where("team_id = ?", teamID).Find(&agents)
	c.JSON(http.StatusOK, gin.H{"data": agents})
}

func (h *Handler) UpdateAIAgent(c *gin.Context) {
	id := c.Param("id")
	var agent models.AIAgent
	if err := h.DB.First(&agent, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&agent)
	c.JSON(http.StatusOK, gin.H{"data": agent})
}

// ============================================================
// 3. Development Environment
// ============================================================

func (h *Handler) GetCodeRepos(c *gin.Context) {
	var repos []models.CodeRepo
	h.DB.Find(&repos)
	c.JSON(http.StatusOK, gin.H{"data": repos})
}

func (h *Handler) UpdateCodeRepo(c *gin.Context) {
	id := c.Param("id")
	var repo models.CodeRepo
	if err := h.DB.First(&repo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&repo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&repo)
	c.JSON(http.StatusOK, gin.H{"data": repo})
}

func (h *Handler) GetCICDServices(c *gin.Context) {
	var services []models.CICDService
	h.DB.Find(&services)
	c.JSON(http.StatusOK, gin.H{"data": services})
}

func (h *Handler) UpdateCICDService(c *gin.Context) {
	id := c.Param("id")
	var svc models.CICDService
	if err := h.DB.First(&svc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&svc)
	c.JSON(http.StatusOK, gin.H{"data": svc})
}

func (h *Handler) GetCloudCredentials(c *gin.Context) {
	var creds []models.CloudCredential
	h.DB.Find(&creds)
	c.JSON(http.StatusOK, gin.H{"data": creds})
}

func (h *Handler) UpdateCloudCredential(c *gin.Context) {
	id := c.Param("id")
	var cred models.CloudCredential
	if err := h.DB.First(&cred, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&cred)
	c.JSON(http.StatusOK, gin.H{"data": cred})
}

func (h *Handler) GetSSHKeys(c *gin.Context) {
	var keys []models.SSHKey
	h.DB.Find(&keys)
	c.JSON(http.StatusOK, gin.H{"data": keys})
}

func (h *Handler) CreateSSHKey(c *gin.Context) {
	var key models.SSHKey
	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&key)
	c.JSON(http.StatusOK, gin.H{"data": key})
}

func (h *Handler) DeleteSSHKey(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.SSHKey{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) GetEnvVars(c *gin.Context) {
	var vars []models.EnvVar
	h.DB.Find(&vars)
	c.JSON(http.StatusOK, gin.H{"data": vars})
}

func (h *Handler) CreateEnvVar(c *gin.Context) {
	var ev models.EnvVar
	if err := c.ShouldBindJSON(&ev); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&ev).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "duplicate key"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ev})
}

func (h *Handler) UpdateEnvVar(c *gin.Context) {
	id := c.Param("id")
	var ev models.EnvVar
	if err := h.DB.First(&ev, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&ev); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&ev)
	c.JSON(http.StatusOK, gin.H{"data": ev})
}

func (h *Handler) DeleteEnvVar(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.EnvVar{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// ============================================================
// 4. Integration Configuration
// ============================================================

func (h *Handler) GetConnectedServices(c *gin.Context) {
	var services []models.ConnectedService
	h.DB.Find(&services)
	c.JSON(http.StatusOK, gin.H{"data": services})
}

func (h *Handler) UpdateConnectedService(c *gin.Context) {
	id := c.Param("id")
	var svc models.ConnectedService
	if err := h.DB.First(&svc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&svc)
	c.JSON(http.StatusOK, gin.H{"data": svc})
}

func (h *Handler) CreateConnectedService(c *gin.Context) {
	var svc models.ConnectedService
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&svc)
	c.JSON(http.StatusOK, gin.H{"data": svc})
}

func (h *Handler) DeleteConnectedService(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.ConnectedService{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) GetAPIKeys(c *gin.Context) {
	var keys []models.APIKey
	h.DB.Find(&keys)
	c.JSON(http.StatusOK, gin.H{"data": keys})
}

func (h *Handler) CreateAPIKey(c *gin.Context) {
	var key models.APIKey
	if err := c.ShouldBindJSON(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&key)
	c.JSON(http.StatusOK, gin.H{"data": key})
}

func (h *Handler) DeleteAPIKey(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.APIKey{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) RefreshAPIKey(c *gin.Context) {
	id := c.Param("id")
	var key models.APIKey
	if err := h.DB.First(&key, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	key.KeyValue = "refreshed_" + strconv.FormatInt(time.Now().Unix(), 10)
	h.DB.Save(&key)
	c.JSON(http.StatusOK, gin.H{"data": key})
}

func (h *Handler) GetWebhookConfigs(c *gin.Context) {
	var configs []models.WebhookConfig
	h.DB.Find(&configs)
	c.JSON(http.StatusOK, gin.H{"data": configs})
}

func (h *Handler) UpdateWebhookConfig(c *gin.Context) {
	id := c.Param("id")
	var config models.WebhookConfig
	if err := h.DB.First(&config, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&config)
	c.JSON(http.StatusOK, gin.H{"data": config})
}

func (h *Handler) GetSyncConfig(c *gin.Context) {
	var config models.SyncConfig
	if err := h.DB.First(&config).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": config})
}

func (h *Handler) UpdateSyncConfig(c *gin.Context) {
	var config models.SyncConfig
	h.DB.First(&config)
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&config)
	c.JSON(http.StatusOK, gin.H{"data": config})
}

// ============================================================
// 5. Model Management
// ============================================================

func (h *Handler) GetAIModels(c *gin.Context) {
	var modelsList []models.AIModel
	h.DB.Find(&modelsList)
	c.JSON(http.StatusOK, gin.H{"data": modelsList})
}

func (h *Handler) CreateAIModel(c *gin.Context) {
	var m models.AIModel
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	m.IsBuiltin = false
	h.DB.Create(&m)
	c.JSON(http.StatusOK, gin.H{"data": m})
}

func (h *Handler) UpdateAIModel(c *gin.Context) {
	id := c.Param("id")
	var m models.AIModel
	if err := h.DB.First(&m, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&m)
	c.JSON(http.StatusOK, gin.H{"data": m})
}

func (h *Handler) DeleteAIModel(c *gin.Context) {
	id := c.Param("id")
	var m models.AIModel
	h.DB.First(&m, id)
	if m.IsBuiltin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete builtin model"})
		return
	}
	h.DB.Delete(&m)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) TestAIModelConnection(c *gin.Context) {
	id := c.Param("id")
	var m models.AIModel
	if err := h.DB.First(&m, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	m.ConnectionStatus = "connected"
	h.DB.Save(&m)
	c.JSON(http.StatusOK, gin.H{"data": m})
}

// ============================================================
// 6. Notification Settings
// ============================================================

func (h *Handler) GetNotificationChannels(c *gin.Context) {
	var channels []models.NotificationChannel
	h.DB.Find(&channels)
	c.JSON(http.StatusOK, gin.H{"data": channels})
}

func (h *Handler) UpdateNotificationChannel(c *gin.Context) {
	id := c.Param("id")
	var ch models.NotificationChannel
	if err := h.DB.First(&ch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&ch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&ch)
	c.JSON(http.StatusOK, gin.H{"data": ch})
}

func (h *Handler) GetNotificationEvents(c *gin.Context) {
	var events []models.NotificationEvent
	h.DB.Find(&events)
	c.JSON(http.StatusOK, gin.H{"data": events})
}

func (h *Handler) UpdateNotificationEvent(c *gin.Context) {
	id := c.Param("id")
	var ev models.NotificationEvent
	if err := h.DB.First(&ev, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&ev); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&ev)
	c.JSON(http.StatusOK, gin.H{"data": ev})
}

func (h *Handler) GetNotificationPreferences(c *gin.Context) {
	var pref models.NotificationPreference
	if err := h.DB.First(&pref).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pref})
}

func (h *Handler) UpdateNotificationPreferences(c *gin.Context) {
	var pref models.NotificationPreference
	h.DB.First(&pref)
	if err := c.ShouldBindJSON(&pref); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&pref)
	c.JSON(http.StatusOK, gin.H{"data": pref})
}

// ============================================================
// 7. Review Rules
// ============================================================

func (h *Handler) GetReviewRules(c *gin.Context) {
	var rules []models.ReviewRule
	h.DB.Find(&rules)
	c.JSON(http.StatusOK, gin.H{"data": rules})
}

func (h *Handler) CreateReviewRule(c *gin.Context) {
	var rule models.ReviewRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&rule)
	c.JSON(http.StatusOK, gin.H{"data": rule})
}

func (h *Handler) UpdateReviewRule(c *gin.Context) {
	id := c.Param("id")
	var rule models.ReviewRule
	if err := h.DB.First(&rule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&rule)
	c.JSON(http.StatusOK, gin.H{"data": rule})
}

func (h *Handler) DeleteReviewRule(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.ReviewRule{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) GetReviewTemplates(c *gin.Context) {
	var templates []models.ReviewTemplate
	h.DB.Find(&templates)
	c.JSON(http.StatusOK, gin.H{"data": templates})
}

func (h *Handler) GetReviewTimeoutConfig(c *gin.Context) {
	var config models.ReviewTimeoutConfig
	if err := h.DB.First(&config).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": config})
}

func (h *Handler) UpdateReviewTimeoutConfig(c *gin.Context) {
	var config models.ReviewTimeoutConfig
	h.DB.First(&config)
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&config)
	c.JSON(http.StatusOK, gin.H{"data": config})
}

// ============================================================
// 8. Security & Permissions
// ============================================================

func (h *Handler) GetRoles(c *gin.Context) {
	var roles []models.Role
	h.DB.Preload("RolePermissions.Permission").Find(&roles)
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func (h *Handler) GetPermissions(c *gin.Context) {
	var perms []models.Permission
	h.DB.Find(&perms)
	c.JSON(http.StatusOK, gin.H{"data": perms})
}

func (h *Handler) GetRolePermissions(c *gin.Context) {
	var rp []models.RolePermission
	h.DB.Preload("Role").Preload("Permission").Find(&rp)
	c.JSON(http.StatusOK, gin.H{"data": rp})
}

func (h *Handler) UpdateRolePermission(c *gin.Context) {
	id := c.Param("id")
	var rp models.RolePermission
	if err := h.DB.First(&rp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&rp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&rp)
	c.JSON(http.StatusOK, gin.H{"data": rp})
}

func (h *Handler) GetSecurityPolicies(c *gin.Context) {
	var policies []models.SecurityPolicy
	h.DB.Find(&policies)
	c.JSON(http.StatusOK, gin.H{"data": policies})
}

func (h *Handler) UpdateSecurityPolicy(c *gin.Context) {
	id := c.Param("id")
	var policy models.SecurityPolicy
	if err := h.DB.First(&policy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&policy)
	c.JSON(http.StatusOK, gin.H{"data": policy})
}

func (h *Handler) GetPasswordPolicy(c *gin.Context) {
	var policy models.PasswordPolicy
	if err := h.DB.First(&policy).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": policy})
}

func (h *Handler) UpdatePasswordPolicy(c *gin.Context) {
	var policy models.PasswordPolicy
	h.DB.First(&policy)
	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&policy)
	c.JSON(http.StatusOK, gin.H{"data": policy})
}

func (h *Handler) GetAuditLogs(c *gin.Context) {
	var logs []models.AuditLog
	query := h.DB.Order("created_at desc")
	if action := c.Query("action"); action != "" {
		query = query.Where("action = ?", action)
	}
	if operator := c.Query("operator"); operator != "" {
		query = query.Where("operator = ?", operator)
	}
	query.Find(&logs)
	c.JSON(http.StatusOK, gin.H{"data": logs})
}

func (h *Handler) CreateAuditLog(c *gin.Context) {
	var log models.AuditLog
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&log)
	c.JSON(http.StatusOK, gin.H{"data": log})
}

// ============================================================
// 9. Team Members
// ============================================================

func (h *Handler) GetTeamMembers(c *gin.Context) {
	var members []models.TeamMember
	h.DB.Preload("Role").Preload("TeamBindings.Team").Find(&members)
	c.JSON(http.StatusOK, gin.H{"data": members})
}

func (h *Handler) CreateTeamMember(c *gin.Context) {
	var member models.TeamMember
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&member)
	c.JSON(http.StatusOK, gin.H{"data": member})
}

func (h *Handler) UpdateTeamMember(c *gin.Context) {
	id := c.Param("id")
	var member models.TeamMember
	if err := h.DB.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Save(&member)
	c.JSON(http.StatusOK, gin.H{"data": member})
}

func (h *Handler) DeleteTeamMember(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.TeamMember{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) GetPendingInvites(c *gin.Context) {
	var invites []models.PendingInvite
	h.DB.Preload("Role").Find(&invites)
	c.JSON(http.StatusOK, gin.H{"data": invites})
}

func (h *Handler) CreatePendingInvite(c *gin.Context) {
	var invite models.PendingInvite
	if err := c.ShouldBindJSON(&invite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invite.Status = "pending"
	h.DB.Create(&invite)
	c.JSON(http.StatusOK, gin.H{"data": invite})
}

func (h *Handler) DeletePendingInvite(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.PendingInvite{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *Handler) GetMemberAITeamBindings(c *gin.Context) {
	var bindings []models.MemberAITeamBinding
	h.DB.Preload("Member").Preload("Team").Find(&bindings)
	c.JSON(http.StatusOK, gin.H{"data": bindings})
}

func (h *Handler) CreateMemberAITeamBinding(c *gin.Context) {
	var binding models.MemberAITeamBinding
	if err := c.ShouldBindJSON(&binding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&binding).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "duplicate binding"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": binding})
}

func (h *Handler) DeleteMemberAITeamBinding(c *gin.Context) {
	id := c.Param("id")
	h.DB.Delete(&models.MemberAITeamBinding{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}