import { ref, onUnmounted } from 'vue'

const API_BASE = 'http://localhost:8080/api'

export function usePipelineProgress(pipelineId) {
  const events = ref([])
  const status = ref('connecting')
  const stages = ref({})
  const error = ref(null)
  let eventSource = null

  function connect() {
    if (!pipelineId) return
    status.value = 'connecting'

    eventSource = new EventSource(`${API_BASE}/pipeline/${pipelineId}/progress`)

    eventSource.onopen = () => {
      status.value = 'connected'
    }

    eventSource.onmessage = (e) => {
      try {
        const event = JSON.parse(e.data)
        if (event.type === 'heartbeat') return

        events.value.push(event)

        if (event.stage) {
          stages.value[event.stage] = {
            status: event.status,
            message: event.message,
            durationMs: event.duration_ms,
            error: event.error,
          }
        }

        if (event.type === 'pipeline_completed') {
          status.value = 'completed'
          disconnect()
        } else if (event.type === 'pipeline_failed') {
          status.value = 'failed'
          error.value = event.error || event.message
          disconnect()
        } else if (event.type === 'review_pause') {
          status.value = 'paused_for_review'
        }
      } catch (err) {
        console.warn('Failed to parse SSE event:', err)
      }
    }

    eventSource.onerror = () => {
      status.value = 'error'
      error.value = 'SSE connection lost'
      eventSource.close()
      eventSource = null
      // Auto reconnect after 5s
      setTimeout(connect, 5000)
    }
  }

  function disconnect() {
    if (eventSource) {
      eventSource.close()
      eventSource = null
      status.value = 'disconnected'
    }
  }

  connect()

  onUnmounted(() => {
    disconnect()
  })

  return {
    events,
    status,
    stages,
    error,
    disconnect,
    reconnect: connect,
  }
}
