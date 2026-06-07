import api from './client'

// 1. Automation Configuration
export const getAutomationConfig = () => api.get('/automation/config')
export const updateAutomationConfig = (data) => api.put('/automation/config', data)
export const getReviewNodes = () => api.get('/automation/review-nodes')
export const updateReviewNode = (id, data) => api.put(`/automation/review-nodes/${id}`, data)

// 2. AI Agent Configuration
export const getAITeams = () => api.get('/ai-agent/teams')
export const getAITeam = (id) => api.get(`/ai-agent/teams/${id}`)
export const updateAITeam = (id, data) => api.put(`/ai-agent/teams/${id}`, data)
export const getAIAgents = (teamId) => api.get(`/ai-agent/agents?team_id=${teamId}`)
export const updateAIAgent = (id, data) => api.put(`/ai-agent/agents/${id}`, data)

// 3. Development Environment
export const getCodeRepos = () => api.get('/dev-environment/code-repos')
export const updateCodeRepo = (id, data) => api.put(`/dev-environment/code-repos/${id}`, data)
export const getCICDServices = () => api.get('/dev-environment/cicd-services')
export const updateCICDService = (id, data) => api.put(`/dev-environment/cicd-services/${id}`, data)
export const getCloudCredentials = () => api.get('/dev-environment/cloud-credentials')
export const updateCloudCredential = (id, data) => api.put(`/dev-environment/cloud-credentials/${id}`, data)
export const getSSHKeys = () => api.get('/dev-environment/ssh-keys')
export const createSSHKey = (data) => api.post('/dev-environment/ssh-keys', data)
export const deleteSSHKey = (id) => api.delete(`/dev-environment/ssh-keys/${id}`)
export const getEnvVars = () => api.get('/dev-environment/env-vars')
export const createEnvVar = (data) => api.post('/dev-environment/env-vars', data)
export const updateEnvVar = (id, data) => api.put(`/dev-environment/env-vars/${id}`, data)
export const deleteEnvVar = (id) => api.delete(`/dev-environment/env-vars/${id}`)

// 4. Integration Configuration
export const getConnectedServices = () => api.get('/integrations/services')
export const createConnectedService = (data) => api.post('/integrations/services', data)
export const updateConnectedService = (id, data) => api.put(`/integrations/services/${id}`, data)
export const deleteConnectedService = (id) => api.delete(`/integrations/services/${id}`)
export const getAPIKeys = () => api.get('/integrations/api-keys')
export const createAPIKey = (data) => api.post('/integrations/api-keys', data)
export const deleteAPIKey = (id) => api.delete(`/integrations/api-keys/${id}`)
export const refreshAPIKey = (id) => api.put(`/integrations/api-keys/${id}/refresh`)
export const getWebhookConfigs = () => api.get('/integrations/webhook')
export const updateWebhookConfig = (id, data) => api.put(`/integrations/webhook/${id}`, data)
export const getSyncConfig = () => api.get('/integrations/sync')
export const updateSyncConfig = (data) => api.put('/integrations/sync', data)

// 5. Model Management
export const getAIModels = () => api.get('/models')
export const createAIModel = (data) => api.post('/models', data)
export const updateAIModel = (id, data) => api.put(`/models/${id}`, data)
export const deleteAIModel = (id) => api.delete(`/models/${id}`)
export const testAIModelConnection = (id) => api.post(`/models/${id}/test`)

// 6. Notification Settings
export const getNotificationChannels = () => api.get('/notifications/channels')
export const updateNotificationChannel = (id, data) => api.put(`/notifications/channels/${id}`, data)
export const getNotificationEvents = () => api.get('/notifications/events')
export const updateNotificationEvent = (id, data) => api.put(`/notifications/events/${id}`, data)
export const getNotificationPreferences = () => api.get('/notifications/preferences')
export const updateNotificationPreferences = (data) => api.put('/notifications/preferences', data)

// 7. Review Rules
export const getReviewRules = () => api.get('/review-rules')
export const createReviewRule = (data) => api.post('/review-rules', data)
export const updateReviewRule = (id, data) => api.put(`/review-rules/${id}`, data)
export const deleteReviewRule = (id) => api.delete(`/review-rules/${id}`)
export const getReviewTemplates = () => api.get('/review-rules/templates')
export const getReviewTimeoutConfig = () => api.get('/review-rules/timeout')
export const updateReviewTimeoutConfig = (data) => api.put('/review-rules/timeout', data)

// 8. Security & Permissions
export const getRoles = () => api.get('/security/roles')
export const getPermissions = () => api.get('/security/permissions')
export const getRolePermissions = () => api.get('/security/role-permissions')
export const updateRolePermission = (id, data) => api.put(`/security/role-permissions/${id}`, data)
export const getSecurityPolicies = () => api.get('/security/policies')
export const updateSecurityPolicy = (id, data) => api.put(`/security/policies/${id}`, data)
export const getPasswordPolicy = () => api.get('/security/password-policy')
export const updatePasswordPolicy = (data) => api.put('/security/password-policy', data)
export const getAuditLogs = (params) => api.get('/security/audit-logs', { params })
export const createAuditLog = (data) => api.post('/security/audit-logs', data)

// 9. Team Members
export const getTeamMembers = () => api.get('/team-members')
export const createTeamMember = (data) => api.post('/team-members', data)
export const updateTeamMember = (id, data) => api.put(`/team-members/${id}`, data)
export const deleteTeamMember = (id) => api.delete(`/team-members/${id}`)
export const getPendingInvites = () => api.get('/team-members/invites')
export const createPendingInvite = (data) => api.post('/team-members/invites', data)
export const deletePendingInvite = (id) => api.delete(`/team-members/invites/${id}`)
export const getMemberAITeamBindings = () => api.get('/team-members/bindings')
export const createMemberAITeamBinding = (data) => api.post('/team-members/bindings', data)
export const deleteMemberAITeamBinding = (id) => api.delete(`/team-members/bindings/${id}`)