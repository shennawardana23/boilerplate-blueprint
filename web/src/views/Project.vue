<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Loading State -->
      <div v-if="isLoading" class="flex items-center justify-center py-20">
        <LoadingSpinner size="lg" />
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="text-center py-20">
        <div class="text-red-500 mb-4">
          <ExclamationTriangleIcon class="w-16 h-16 mx-auto mb-4" />
          <h2 class="text-2xl font-bold">Project Not Found</h2>
          <p class="text-gray-600 mt-2">{{ error }}</p>
        </div>
        <router-link to="/generator" class="btn btn-primary">
          Back to Generator
        </router-link>
      </div>

      <!-- Project Details -->
      <div v-else-if="project" class="space-y-8">
        <!-- Header -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div class="flex items-start justify-between">
            <div class="flex items-center space-x-4">
              <div class="w-16 h-16 bg-primary-100 rounded-xl flex items-center justify-center">
                <FolderIcon class="w-8 h-8 text-primary-600" />
              </div>
              <div>
                <h1 class="text-3xl font-bold text-gray-900">{{ project.name }}</h1>
                <p class="text-lg text-gray-600 mt-1">
                  {{ project.description || 'No description provided' }}
                </p>
                <div class="flex items-center space-x-4 mt-2">
                  <span class="badge badge-primary">{{ project.language.toUpperCase() }}</span>
                  <span class="text-sm text-gray-500">
                    Created {{ formatDate(project.created_at) }}
                  </span>
                </div>
              </div>
            </div>
            <div class="flex items-center space-x-3">
              <button
                @click="generateFiles"
                :disabled="projectStore.isLoading"
                class="btn btn-outline"
              >
                <DocumentIcon class="w-4 h-4 mr-2" />
                Generate Files
              </button>
              <button
                @click="downloadProject"
                :disabled="projectStore.isLoading || projectStore.currentProjectFiles.length === 0"
                class="btn btn-primary"
              >
                <ArrowDownTrayIcon class="w-4 h-4 mr-2" />
                Download ZIP
              </button>
            </div>
          </div>
        </div>

        <!-- Project Configuration -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <!-- Configuration Details -->
          <div class="lg:col-span-2 space-y-6">
            <!-- Project Options -->
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">Project Configuration</h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">Language</label>
                  <p class="text-sm text-gray-900">{{ project.language.toUpperCase() }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">Framework</label>
                  <p class="text-sm text-gray-900">{{ project.options.framework || 'Default' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">Database</label>
                  <p class="text-sm text-gray-900">{{ project.options.database || 'None' }}</p>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1">Authentication</label>
                  <p class="text-sm text-gray-900">{{ project.options.auth_type || 'Basic' }}</p>
                </div>
              </div>
            </div>

            <!-- Project Structure -->
            <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">Project Structure</h3>
              <ProjectStructurePreview
                :project="project"
                :files="projectStore.currentProjectFiles"
                :is-loading="projectStore.isLoading"
              />
            </div>
          </div>

          <!-- AI Assistant -->
          <div class="lg:col-span-1">
            <div class="sticky top-8">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">Project Assistant</h3>
              <div class="h-[600px]">
                <ChatInterface
                  :project-id="project.id"
                  @suggestion-applied="handleSuggestionApplied"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  FolderIcon,
  DocumentIcon,
  ArrowDownTrayIcon,
  ExclamationTriangleIcon
} from '@heroicons/vue/24/outline'

import { useProjectStore } from '@/stores/project'
import { useChatStore } from '@/stores/chat'

import LoadingSpinner from '@/components/ui/LoadingSpinner.vue'
import ProjectStructurePreview from '@/components/project/ProjectStructurePreview.vue'
import ChatInterface from '@/components/chat/ChatInterface.vue'

const route = useRoute()
const router = useRouter()
const projectStore = useProjectStore()
const chatStore = useChatStore()

const isLoading = ref(true)
const error = ref(null)
const project = computed(() => projectStore.currentProject)

// Load project on mount
onMounted(async () => {
  const projectId = route.params.id
  
  if (!projectId) {
    error.value = 'No project ID provided'
    isLoading.value = false
    return
  }

  try {
    await projectStore.loadProject(projectId)
    await chatStore.loadChatHistory(projectId)
    
    if (!project.value) {
      error.value = 'Project not found'
    }
  } catch (err) {
    console.error('Failed to load project:', err)
    error.value = 'Failed to load project'
  } finally {
    isLoading.value = false
  }
})

// Generate project files
async function generateFiles() {
  if (!project.value) return

  try {
    await projectStore.generateProjectFiles(project.value.id)
    
    // Add system message
    await chatStore.addSystemMessage(
      `Files generated successfully! Your ${project.value.language.toUpperCase()} project is ready for download.`
    )
  } catch (error) {
    console.error('Failed to generate files:', error)
  }
}

// Download project
async function downloadProject() {
  if (!project.value) return

  try {
    await projectStore.downloadProject(project.value.id)
  } catch (error) {
    console.error('Failed to download project:', error)
  }
}

// Handle AI suggestion applied
function handleSuggestionApplied(suggestion) {
  const updates = {}
  updates[suggestion.type] = suggestion.value
  projectStore.updateProjectOptions(updates)
}

// Format date
function formatDate(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>
