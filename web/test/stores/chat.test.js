import { expect } from 'chai'
import { setActivePinia, createPinia } from 'pinia'
import { useChatStore } from '../../src/stores/chat.js'

// Mock the API
const mockChatApi = {
  sendMessage: () => Promise.resolve({ data: { success: true, message: {} } }),
  getChatHistory: () => Promise.resolve({ data: { success: true, history: {} } })
}

describe('Chat Store', () => {
  let store

  beforeEach(() => {
    setActivePinia(createPinia())
    store = useChatStore()
    
    // Mock the API module
    global.sinon.stub(require('../../src/services/api'), 'chatApi').value(mockChatApi)
  })

  afterEach(() => {
    global.sinon.restore()
  })

  describe('Initial State', () => {
    it('should have initial state values', () => {
      expect(store.messages).to.be.an('array').that.is.empty
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
      expect(store.currentProjectId).to.be.null
      expect(store.suggestions).to.be.an('array').that.is.empty
    })

    it('should have computed getters', () => {
      expect(store.hasMessages).to.be.false
      expect(store.lastMessage).to.be.null
      expect(store.userMessages).to.be.an('array').that.is.empty
      expect(store.assistantMessages).to.be.an('array').that.is.empty
    })
  })

  describe('sendMessage', () => {
    it('should send message successfully', async () => {
      const messageContent = 'Hello, I need help with my project'
      const projectId = 'test-project'
      
      const mockAssistantMessage = {
        id: 'assistant-msg-id',
        role: 'assistant',
        content: 'I can help you with that!',
        project_id: projectId,
        created_at: new Date().toISOString()
      }

      const mockSuggestions = [
        {
          type: 'language',
          value: 'go',
          reason: 'You mentioned Go in your message',
          confidence: 0.9,
          apply: true
        }
      ]

      mockChatApi.sendMessage = () => Promise.resolve({ 
        data: { 
          success: true, 
          message: mockAssistantMessage,
          suggestions: mockSuggestions
        } 
      })

      const result = await store.sendMessage(messageContent, projectId)

      expect(store.messages).to.have.length(2)
      expect(store.messages[0].role).to.equal('user')
      expect(store.messages[0].content).to.equal(messageContent)
      expect(store.messages[1]).to.deep.equal(mockAssistantMessage)
      expect(store.suggestions).to.deep.equal(mockSuggestions)
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
      expect(result.message).to.deep.equal(mockAssistantMessage)
    })

    it('should handle send message error', async () => {
      const messageContent = 'Hello'
      const errorMessage = 'Failed to send message'
      
      mockChatApi.sendMessage = () => Promise.resolve({ 
        data: { success: false, error: errorMessage } 
      })

      try {
        await store.sendMessage(messageContent)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.messages).to.be.empty // User message should be removed on error
      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })

    it('should handle API exception', async () => {
      const messageContent = 'Hello'
      const errorMessage = 'Network error'
      
      mockChatApi.sendMessage = () => Promise.reject(new Error(errorMessage))

      try {
        await store.sendMessage(messageContent)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.messages).to.be.empty // User message should be removed on error
      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('loadChatHistory', () => {
    it('should load chat history successfully', async () => {
      const projectId = 'test-project'
      const mockHistory = {
        project_id: projectId,
        messages: [
          {
            id: 'msg1',
            role: 'user',
            content: 'Hello',
            created_at: new Date().toISOString()
          },
          {
            id: 'msg2',
            role: 'assistant',
            content: 'Hi there!',
            created_at: new Date().toISOString()
          }
        ],
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }

      mockChatApi.getChatHistory = () => Promise.resolve({ 
        data: { success: true, history: mockHistory } 
      })

      await store.loadChatHistory(projectId)

      expect(store.messages).to.deep.equal(mockHistory.messages)
      expect(store.currentProjectId).to.equal(projectId)
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should load chat history without projectId', async () => {
      const mockHistory = {
        project_id: 'general',
        messages: [],
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }

      mockChatApi.getChatHistory = () => Promise.resolve({ 
        data: { success: true, history: mockHistory } 
      })

      await store.loadChatHistory()

      expect(store.messages).to.deep.equal(mockHistory.messages)
      expect(store.currentProjectId).to.be.null
      expect(store.isLoading).to.be.false
      expect(store.error).to.be.null
    })

    it('should handle load chat history error', async () => {
      const projectId = 'test-project'
      const errorMessage = 'Failed to load chat history'
      
      mockChatApi.getChatHistory = () => Promise.resolve({ 
        data: { success: false, error: errorMessage } 
      })

      try {
        await store.loadChatHistory(projectId)
      } catch (err) {
        expect(err.message).to.equal(errorMessage)
      }

      expect(store.error).to.equal(errorMessage)
      expect(store.isLoading).to.be.false
    })
  })

  describe('addSystemMessage', () => {
    it('should add system message to messages array', () => {
      const systemContent = 'System message'

      store.addSystemMessage(systemContent)

      expect(store.messages).to.have.length(1)
      expect(store.messages[0].role).to.equal('system')
      expect(store.messages[0].content).to.equal(systemContent)
      expect(store.messages[0].id).to.be.a('string')
      expect(store.messages[0].created_at).to.be.a('string')
    })
  })

  describe('clearMessages', () => {
    it('should clear all messages and suggestions', () => {
      store.messages = [
        { id: 'msg1', role: 'user', content: 'Hello' },
        { id: 'msg2', role: 'assistant', content: 'Hi' }
      ]
      store.suggestions = [
        { type: 'language', value: 'go' }
      ]

      store.clearMessages()

      expect(store.messages).to.be.empty
      expect(store.suggestions).to.be.empty
    })
  })

  describe('clearError', () => {
    it('should clear error', () => {
      store.error = 'Some error message'

      store.clearError()

      expect(store.error).to.be.null
    })
  })

  describe('applySuggestion', () => {
    it('should apply suggestion and add system message', () => {
      const suggestion = {
        type: 'framework',
        value: 'gin',
        reason: 'Gin is a popular framework'
      }

      store.suggestions = [suggestion]

      store.applySuggestion(suggestion)

      expect(store.suggestions[0].apply).to.be.true
      expect(store.messages).to.have.length(1)
      expect(store.messages[0].role).to.equal('system')
      expect(store.messages[0].content).to.include('Applied suggestion')
      expect(store.messages[0].content).to.include('framework = gin')
    })

    it('should not apply suggestion if not found', () => {
      const suggestion = {
        type: 'framework',
        value: 'gin',
        reason: 'Gin is a popular framework'
      }

      store.suggestions = []

      store.applySuggestion(suggestion)

      expect(store.messages).to.be.empty
    })
  })

  describe('dismissSuggestion', () => {
    it('should remove suggestion from suggestions array', () => {
      const suggestion1 = { type: 'framework', value: 'gin' }
      const suggestion2 = { type: 'database', value: 'postgresql' }

      store.suggestions = [suggestion1, suggestion2]

      store.dismissSuggestion(suggestion1)

      expect(store.suggestions).to.have.length(1)
      expect(store.suggestions[0]).to.deep.equal(suggestion2)
    })

    it('should not remove suggestion if not found', () => {
      const suggestion1 = { type: 'framework', value: 'gin' }
      const suggestion2 = { type: 'database', value: 'postgresql' }

      store.suggestions = [suggestion1]

      store.dismissSuggestion(suggestion2)

      expect(store.suggestions).to.have.length(1)
      expect(store.suggestions[0]).to.deep.equal(suggestion1)
    })
  })

  describe('formatContextForAPI', () => {
    it('should format recent messages for API context', () => {
      store.messages = [
        { role: 'user', content: 'Hello' },
        { role: 'assistant', content: 'Hi there!' },
        { role: 'user', content: 'How are you?' },
        { role: 'assistant', content: 'I am doing well' },
        { role: 'user', content: 'Great!' },
        { role: 'assistant', content: 'Thanks!' }
      ]

      const context = store.formatContextForAPI()

      expect(context).to.equal('assistant: Thanks!\nuser: Great!\nassistant: I am doing well\nuser: How are you?\nassistant: Hi there!')
    })

    it('should handle empty messages array', () => {
      store.messages = []

      const context = store.formatContextForAPI()

      expect(context).to.equal('')
    })
  })

  describe('setCurrentProjectId', () => {
    it('should set current project ID', () => {
      const projectId = 'test-project'

      store.setCurrentProjectId(projectId)

      expect(store.currentProjectId).to.equal(projectId)
    })
  })

  describe('Computed Properties', () => {
    it('should compute hasMessages correctly', () => {
      expect(store.hasMessages).to.be.false

      store.messages = [{ id: 'msg1', role: 'user', content: 'Hello' }]
      expect(store.hasMessages).to.be.true
    })

    it('should compute lastMessage correctly', () => {
      expect(store.lastMessage).to.be.null

      const messages = [
        { id: 'msg1', role: 'user', content: 'Hello' },
        { id: 'msg2', role: 'assistant', content: 'Hi' }
      ]
      store.messages = messages

      expect(store.lastMessage).to.deep.equal(messages[1])
    })

    it('should compute userMessages correctly', () => {
      const messages = [
        { id: 'msg1', role: 'user', content: 'Hello' },
        { id: 'msg2', role: 'assistant', content: 'Hi' },
        { id: 'msg3', role: 'user', content: 'How are you?' }
      ]
      store.messages = messages

      expect(store.userMessages).to.have.length(2)
      expect(store.userMessages[0].content).to.equal('Hello')
      expect(store.userMessages[1].content).to.equal('How are you?')
    })

    it('should compute assistantMessages correctly', () => {
      const messages = [
        { id: 'msg1', role: 'user', content: 'Hello' },
        { id: 'msg2', role: 'assistant', content: 'Hi' },
        { id: 'msg3', role: 'user', content: 'How are you?' },
        { id: 'msg4', role: 'assistant', content: 'I am doing well' }
      ]
      store.messages = messages

      expect(store.assistantMessages).to.have.length(2)
      expect(store.assistantMessages[0].content).to.equal('Hi')
      expect(store.assistantMessages[1].content).to.equal('I am doing well')
    })
  })
})