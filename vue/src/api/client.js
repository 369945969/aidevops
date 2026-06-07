import axios from 'axios'

const baseURL = import.meta.env.VITE_API_BASE_URL || '/api'

const api = axios.create({
  baseURL,
  timeout: 10000,
  headers: { 'Content-Type': 'application/json' },
})

api.interceptors.response.use(
  (res) => res.data?.data ?? res.data,
  (err) => {
    console.error('API error:', err.response?.data?.error || err.message)
    return Promise.reject(err)
  }
)

export default api