import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', redirect: '/workflow/dashboard' },
  { path: '/workflow/dashboard', name: 'Dashboard', component: () => import('../views/workflow/Dashboard.vue') },
  { path: '/workflow/task/:id', name: 'TaskDetail', component: () => import('../views/workflow/TaskDetail.vue') },
  { path: '/requirements', name: 'Requirements', component: () => import('../views/requirements/Requirements.vue') },
  { path: '/requirements/:id', name: 'RequirementDetail', component: () => import('../views/requirements/RequirementDetail.vue') },
  { path: '/requirements/new', name: 'NewRequirement', component: () => import('../views/requirements/NewRequirement.vue') },
  { path: '/knowledge', name: 'KnowledgeBase', component: () => import('../views/knowledge/KnowledgeBase.vue') },
  { path: '/knowledge/detail/:id', name: 'KnowledgeDetail', component: () => import('../views/knowledge/KnowledgeDetail.vue') },
  { path: '/knowledge/requirement-docs', name: 'KnowledgeRequirementDocs', component: () => import('../views/knowledge/KnowledgeRequirementDocs.vue') },
  { path: '/knowledge/architecture', name: 'KnowledgeArchitecture', component: () => import('../views/knowledge/KnowledgeArchitecture.vue') },
  { path: '/knowledge/code-patterns', name: 'KnowledgeCodePatterns', component: () => import('../views/knowledge/KnowledgeCodePatterns.vue') },
  { path: '/knowledge/solutions', name: 'KnowledgeSolutions', component: () => import('../views/knowledge/KnowledgeSolutions.vue') },
  { path: '/knowledge/deploy-config', name: 'KnowledgeDeployConfig', component: () => import('../views/knowledge/KnowledgeDeployConfig.vue') },
  { path: '/knowledge/favorites', name: 'KnowledgeFavorites', component: () => import('../views/knowledge/KnowledgeFavorites.vue') },
  { path: '/settings', name: 'Settings', component: () => import('../views/settings/Settings.vue') },
  { path: '/settings/ai-agent', name: 'SettingsAiAgent', component: () => import('../views/settings/SettingsAiAgent.vue') },
  { path: '/settings/dev-environment', name: 'SettingsDevEnvironment', component: () => import('../views/settings/SettingsDevEnvironment.vue') },
  { path: '/settings/integrations', name: 'SettingsIntegrations', component: () => import('../views/settings/SettingsIntegrations.vue') },
  { path: '/settings/notifications', name: 'SettingsNotifications', component: () => import('../views/settings/SettingsNotifications.vue') },
  { path: '/settings/review-rules', name: 'SettingsReviewRules', component: () => import('../views/settings/SettingsReviewRules.vue') },
  { path: '/settings/security', name: 'SettingsSecurity', component: () => import('../views/settings/SettingsSecurity.vue') },
  { path: '/settings/team-members', name: 'SettingsTeamMembers', component: () => import('../views/settings/SettingsTeamMembers.vue') },
  { path: '/review/code', name: 'CodeReview', component: () => import('../views/review/CodeReview.vue') },
  { path: '/review/architecture', name: 'ArchitectureReview', component: () => import('../views/review/ArchitectureReview.vue') },
  { path: '/review/deploy', name: 'DeployReview', component: () => import('../views/review/DeployReview.vue') },
  { path: '/review/deploy-monitoring', name: 'DeployMonitoring', component: () => import('../views/review/DeployMonitoring.vue') },
  { path: '/notifications', name: 'NotificationPanel', component: () => import('../views/NotificationPanel.vue') },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router