import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { chatApi } from '@/services/api'

export const useChatStore = defineStore('chat', () => {
  // State
  const messages = ref([])
  const isLoading = ref(false)
  const error = ref(null)
  const currentProjectId = ref(null)
  const suggestions = ref([])

  // Getters
  const hasMessages = computed(() => messages.value.length > 0)
  const lastMessage = computed(() => messages.value[messages.value.length - 1] || null)
  const userMessages = computed(() => messages.value.filter(m => m.role === 'user'))
  const assistantMessages = computed(() => messages.value.filter(m => m.role === 'assistant'))

  // Actions
  async function sendMessage(content, projectId = null) {
    try {
      isLoading.value = true
      error.value = null

      // Add user message to local state immediately
      const userMessage = {
        id: Date.now().toString(),
        role: 'user',
        content,
        project_id: projectId,
        created_at: new Date().toISOString()
      }
      messages.value.push(userMessage)

      // Send to API
      const response = await chatApi.sendMessage({
        message: content,
        project_id: projectId,
        context: formatContextForAPI()
      })

      if (response.data.success) {
        // Add assistant response
        messages.value.push(response.data.message)

        // Update suggestions if any
        if (response.data.suggestions && response.data.suggestions.length > 0) {
          suggestions.value = response.data.suggestions
        }

        return response.data
      } else {
        throw new Error(response.data.error || 'Failed to send message')
      }
    } catch (err) {
      error.value = err.message
      console.error('Failed to send chat message:', err)

      // Remove the user message if API call failed
      messages.value.pop()
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadChatHistory(projectId = null) {
    try {
      isLoading.value = true
      error.value = null

      const response = await chatApi.getChatHistory(projectId)
      if (response.data.success) {
        messages.value = response.data.history?.messages || []
        currentProjectId.value = projectId
      } else {
        throw new Error(response.data.error || 'Failed to load chat history')
      }
    } catch (err) {
      error.value = err.message
      console.error('Failed to load chat history:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  function addSystemMessage(content) {
    const systemMessage = {
      id: Date.now().toString(),
      role: 'system',
      content,
      created_at: new Date().toISOString()
    }
    messages.value.push(systemMessage)
  }

  function clearMessages() {
    messages.value = []
    suggestions.value = []
  }

  function clearError() {
    error.value = null
  }

  function applySuggestion(suggestion) {
    // Mark suggestion as applied
    const suggestionIndex = suggestions.value.findIndex(s => s.type === suggestion.type && s.value === suggestion.value)
    if (suggestionIndex !== -1) {
      suggestions.value[suggestionIndex].apply = true
    }

    // Add system message about applied suggestion
    addSystemMessage(`Applied suggestion: ${suggestion.type} = ${suggestion.value}. Reason: ${suggestion.reason}`)
  }

  function dismissSuggestion(suggestion) {
    suggestions.value = suggestions.value.filter(s => !(s.type === suggestion.type && s.value === suggestion.value))
  }

  function formatContextForAPI() {
    // Get last 5 messages for context
    const recentMessages = messages.value.slice(-5)
    return recentMessages.map(m => `${m.role}: ${m.content}`).join('\n')
  }

  function setCurrentProjectId(projectId) {
    currentProjectId.value = projectId
  }

  return {
    // State
    messages,
    isLoading,
    error,
    currentProjectId,
    suggestions,

    // Getters
    hasMessages,
    lastMessage,
    userMessages,
    assistantMessages,

    // Actions
    sendMessage,
    loadChatHistory,
    addSystemMessage,
    clearMessages,
    clearError,
    applySuggestion,
    dismissSuggestion,
    setCurrentProjectId
  }
})