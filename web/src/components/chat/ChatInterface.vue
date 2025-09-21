<template>
  <div class="flex flex-col h-full bg-white rounded-lg border border-gray-200">
    <!-- Chat Header -->
    <div class="flex items-center justify-between p-4 border-b border-gray-200">
      <div class="flex items-center space-x-3">
        <div class="w-8 h-8 bg-gradient-to-br from-primary-500 to-primary-600 rounded-full flex items-center justify-center">
          <span class="text-white text-sm font-bold">AI</span>
        </div>
        <div>
          <h3 class="text-sm font-medium text-gray-900">Project Assistant</h3>
          <p class="text-xs text-gray-500">
            {{ isOnline ? 'Online' : 'Offline' }}
          </p>
        </div>
      </div>

      <div class="flex items-center space-x-2">
        <!-- Status indicator -->
        <div class="flex items-center space-x-1">
          <div
            class="w-2 h-2 rounded-full"
            :class="isOnline ? 'bg-green-400' : 'bg-gray-400'"
          ></div>
          <span class="text-xs text-gray-500">
            {{ chatStore.isLoading ? 'Typing...' : 'Ready' }}
          </span>
        </div>

        <!-- Clear chat button -->
        <button
          type="button"
          @click="clearChat"
          class="p-1 text-gray-400 hover:text-gray-600 transition-colors"
          title="Clear chat"
        >
          <TrashIcon class="w-4 h-4" />
        </button>
      </div>
    </div>

    <!-- Messages Container -->
    <div
      ref="messagesContainer"
      class="flex-1 overflow-y-auto p-4 space-y-4 min-h-0"
      style="max-height: 400px"
    >
      <!-- Welcome message -->
      <div v-if="!chatStore.hasMessages" class="text-center py-8">
        <div class="w-12 h-12 bg-primary-100 rounded-full flex items-center justify-center mx-auto mb-3">
          <ChatBubbleLeftRightIcon class="w-6 h-6 text-primary-600" />
        </div>
        <h4 class="text-sm font-medium text-gray-900 mb-1">Welcome to Project Assistant!</h4>
        <p class="text-xs text-gray-500 max-w-xs mx-auto">
          I can help you configure your project and suggest the best options for your needs.
        </p>
      </div>

      <!-- Messages -->
      <div
        v-for="message in chatStore.messages"
        :key="message.id"
        class="flex"
        :class="message.role === 'user' ? 'justify-end' : 'justify-start'"
      >
        <div
          class="max-w-xs lg:max-w-md px-3 py-2 rounded-lg"
          :class="message.role === 'user'
            ? 'bg-primary-600 text-white'
            : 'bg-gray-100 text-gray-900'"
        >
          <p class="text-sm">{{ message.content }}</p>
          <p
            class="text-xs mt-1 opacity-75"
            :class="message.role === 'user' ? 'text-primary-100' : 'text-gray-500'"
          >
            {{ formatTime(message.created_at) }}
          </p>
        </div>
      </div>

      <!-- Typing indicator -->
      <div v-if="chatStore.isLoading" class="flex justify-start">
        <div class="bg-gray-100 rounded-lg px-3 py-2">
          <div class="flex items-center space-x-1">
            <div class="w-2 h-2 bg-gray-400 rounded-full animate-pulse"></div>
            <div class="w-2 h-2 bg-gray-400 rounded-full animate-pulse" style="animation-delay: 0.2s"></div>
            <div class="w-2 h-2 bg-gray-400 rounded-full animate-pulse" style="animation-delay: 0.4s"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Suggestions -->
    <div v-if="chatStore.suggestions.length > 0" class="px-4 pb-2">
      <div class="border-t border-gray-200 pt-3">
        <h4 class="text-xs font-medium text-gray-700 mb-2">Suggestions:</h4>
        <div class="flex flex-wrap gap-2">
          <button
            v-for="suggestion in chatStore.suggestions"
            :key="`${suggestion.type}-${suggestion.value}`"
            @click="applySuggestion(suggestion)"
            class="inline-flex items-center px-2 py-1 text-xs bg-blue-50 text-blue-700 rounded border border-blue-200 hover:bg-blue-100 transition-colors"
          >
            <CheckIcon class="w-3 h-3 mr-1" />
            {{ suggestion.type }}: {{ suggestion.value }}
          </button>
        </div>
      </div>
    </div>

    <!-- Input Area -->
    <div class="p-4 border-t border-gray-200">
      <form @submit.prevent="sendMessage" class="flex space-x-2">
        <input
          v-model="messageText"
          type="text"
          placeholder="Ask about your project configuration..."
          class="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500 text-sm"
          :disabled="chatStore.isLoading"
          maxlength="500"
        />
        <button
          type="submit"
          :disabled="!messageText.trim() || chatStore.isLoading"
          class="px-3 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
        >
          <PaperAirplaneIcon class="w-4 h-4" />
        </button>
      </form>

      <!-- Quick suggestions -->
      <div class="flex flex-wrap gap-2 mt-2">
        <button
          v-for="quickSuggestion in quickSuggestions"
          :key="quickSuggestion"
          @click="sendQuickMessage(quickSuggestion)"
          class="text-xs px-2 py-1 text-gray-600 bg-gray-100 rounded hover:bg-gray-200 transition-colors"
          :disabled="chatStore.isLoading"
        >
          {{ quickSuggestion }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick, onMounted } from 'vue'
import {
  ChatBubbleLeftRightIcon,
  PaperAirplaneIcon,
  CheckIcon,
  TrashIcon
} from '@heroicons/vue/24/outline'
import { useChatStore } from '@/stores/chat'

const props = defineProps({
  projectId: {
    type: String,
    default: null
  }
})

const emit = defineEmits(['suggestion-applied'])

const chatStore = useChatStore()
const messageText = ref('')
const messagesContainer = ref(null)
const isOnline = ref(true)

const quickSuggestions = [
  'What framework should I use?',
  'Best database for my project?',
  'Security recommendations?',
  'Deployment options?'
]

async function sendMessage() {
  const text = messageText.value.trim()
  if (!text || chatStore.isLoading) return

  messageText.value = ''

  try {
    await chatStore.sendMessage(text, props.projectId)
    await scrollToBottom()
  } catch (error) {
    console.error('Failed to send message:', error)
  }
}

async function sendQuickMessage(text) {
  if (chatStore.isLoading) return

  try {
    await chatStore.sendMessage(text, props.projectId)
    await scrollToBottom()
  } catch (error) {
    console.error('Failed to send quick message:', error)
  }
}

async function applySuggestion(suggestion) {
  chatStore.applySuggestion(suggestion)
  emit('suggestion-applied', suggestion)

  // Send confirmation message
  await chatStore.addSystemMessage(`Applied ${suggestion.type}: ${suggestion.value}`)
  await scrollToBottom()
}

function clearChat() {
  if (confirm('Are you sure you want to clear the chat history?')) {
    chatStore.clearMessages()
  }
}

async function scrollToBottom() {
  await nextTick()
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

function formatTime(timestamp) {
  const date = new Date(timestamp)
  return date.toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(async () => {
  // Load chat history if project ID is provided
  if (props.projectId) {
    try {
      await chatStore.loadChatHistory(props.projectId)
      await scrollToBottom()
    } catch (error) {
      console.error('Failed to load chat history:', error)
    }
  }
})
</script>