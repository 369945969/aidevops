import api from './client'

export const createRequirement = (data) => api.post('/requirements', data)
export const getRequirements = (params) => api.get('/requirements', { params })
export const getRequirement = (id) => api.get(`/requirements/${id}`)

export const getPipelines = () => api.get('/pipeline')
export const getPipeline = (id) => api.get(`/pipeline/${id}`)
export const runPipeline = (id, data) => api.post(`/pipeline/${id}/run`, data)
export const getPipelineArtifacts = (id) => api.get(`/pipeline/${id}/artifacts`)
export const resumePipelineReview = (id, data) => api.post(`/pipeline/${id}/resume`, data)

export const getWorkflowTasks = () => api.get('/workflow/tasks')

export const getNotifications = () => api.get('/biz-notifications')
export const markNotificationRead = (id) => api.put(`/biz-notifications/${id}/read`)
