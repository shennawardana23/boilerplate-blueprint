import axios from 'axios'

// Create axios instance
const api = axios.create({
  baseURL: '/api',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor
api.interceptors.request.use(
  (config) => {
    // Add any auth tokens or other headers here
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
api.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    // Handle global errors
    if (error.response) {
      // Server responded with error status
      const message = error.response.data?.error || error.response.data?.message || 'Server error'
      error.message = message
    } else if (error.request) {
      // Network error
      error.message = 'Network error - please check your connection'
    } else {
      // Something else happened
      error.message = 'An unexpected error occurred'
    }

    return Promise.reject(error)
  }
)

// Project API endpoints
export const projectApi = {
  // Get available templates
  getTemplates() {
    return api.get('/templates')
  },

  // Create a new project
  createProject(projectData) {
    return api.post('/projects', projectData)
  },

  // Get project by ID
  getProject(projectId) {
    return api.get(`/projects/${projectId}`)
  },

  // Generate project files
  generateProject(projectId) {
    return api.post(`/projects/${projectId}/generate`)
  },

  // Download project as ZIP
  downloadProject(projectId) {
    return api.get(`/projects/${projectId}/download`, {
      responseType: 'blob'
    })
  }
}

// Chat API endpoints
export const chatApi = {
  // Send a chat message
  sendMessage(messageData) {
    return api.post('/chat/message', messageData)
  },

  // Get chat history
  getChatHistory(projectId = null) {
    const params = projectId ? { project_id: projectId } : {}
    return api.get('/chat/history', { params })
  }
}

// Health check
export const healthApi = {
  check() {
    return api.get('/health')
  }
}

export default api