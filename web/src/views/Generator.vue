<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="text-center mb-8">
        <h1 class="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
          Project Generator
        </h1>
        <p class="text-lg text-gray-600 max-w-2xl mx-auto">
          Create professional boilerplate projects with AI assistance. Configure your project settings and get a complete, production-ready codebase.
        </p>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Project Form - Left Column -->
        <div class="lg:col-span-2">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200">
            <div class="p-6">
              <h2 class="text-xl font-semibold text-gray-900 mb-6">
                Configure Your Project
              </h2>

              <ProjectForm
                :is-loading="projectStore.isLoading"
                @submit="handleCreateProject"
                @draft="handleSaveDraft"
              />
            </div>
          </div>

          <!-- Project Preview -->
          <div v-if="projectStore.hasCurrentProject" class="mt-8 bg-white rounded-lg shadow-sm border border-gray-200">
            <div class="p-6">
              <div class="flex items-center justify-between mb-6">
                <h3 class="text-lg font-semibold text-gray-900">
                  Project Structure Preview
                </h3>
                <div class="flex space-x-2">
                  <button
                    @click="generateFiles"
                    :disabled="projectStore.isLoading"
                    class="btn btn-outline btn-sm"
                  >
                    <DocumentIcon class="w-4 h-4 mr-1" />
                    Generate Files
                  </button>
                  <button
                    @click="downloadProject"
                    :disabled="projectStore.isLoading || projectStore.currentProjectFiles.length === 0"
                    class="btn btn-primary btn-sm"
                  >
                    <ArrowDownTrayIcon class="w-4 h-4 mr-1" />
                    Download ZIP
                  </button>
                </div>
              </div>

              <ProjectStructurePreview
                :project="projectStore.currentProject"
                :files="projectStore.currentProjectFiles"
                :is-loading="projectStore.isLoading"
              />
            </div>
          </div>
        </div>

        <!-- AI Chat - Right Column -->
        <div class="lg:col-span-1">
          <div class="sticky top-8">
            <h2 class="text-xl font-semibold text-gray-900 mb-4">
              AI Assistant
            </h2>
            <div class="h-[600px]">
              <ChatInterface
                :project-id="projectStore.currentProject?.id"
                @suggestion-applied="handleSuggestionApplied"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Projects -->
      <div v-if="recentProjects.length > 0" class="mt-12">
        <h2 class="text-xl font-semibold text-gray-900 mb-6">Recent Projects</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="project in recentProjects"
            :key="project.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow cursor-pointer"
            @click="loadProject(project.id)"
          >
            <div class="flex items-start justify-between mb-3">
              <div>
                <h3 class="font-semibold text-gray-900">{{ project.name }}</h3>
                <p class="text-sm text-gray-500">{{ project.description || 'No description' }}</p>
              </div>
              <span class="badge badge-primary">{{ project.language.toUpperCase() }}</span>
            </div>
            <div class="text-xs text-gray-500">
              Created {{ formatDate(project.created_at) }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Success Modal -->
    <Teleport to="body">
      <div
        v-if="showSuccessModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      >
        <div class="bg-white rounded-lg max-w-md w-full p-6">
          <div class="text-center">
            <div class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <CheckIcon class="w-6 h-6 text-green-600" />
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mb-2">
              Project Created Successfully!
            </h3>
            <p class="text-gray-600 mb-6">
              Your {{ projectStore.currentProject?.language.toUpperCase() }} project "{{ projectStore.currentProject?.name }}" has been created and is ready for development.
            </p>
            <div class="flex space-x-3 justify-center">
              <button
                @click="showSuccessModal = false"
                class="btn btn-secondary"
              >
                Continue Editing
              </button>
              <button
                @click="downloadProject"
                class="btn btn-primary"
              >
                <ArrowDownTrayIcon class="w-4 h-4 mr-1" />
                Download Now
              </button>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  DocumentIcon,
  ArrowDownTrayIcon,
  CheckIcon
} from '@heroicons/vue/24/outline'

import { useProjectStore } from '@/stores/project'
import { useChatStore } from '@/stores/chat'

import ProjectForm from '@/components/project/ProjectForm.vue'
import ChatInterface from '@/components/chat/ChatInterface.vue'
import ProjectStructurePreview from '@/components/project/ProjectStructurePreview.vue'

const router = useRouter()
const projectStore = useProjectStore()
const chatStore = useChatStore()

const showSuccessModal = ref(false)
const recentProjects = ref([])

// Handle project creation
async function handleCreateProject(projectData) {
  try {
    const project = await projectStore.createProject(projectData)

    // Set project context for chat
    chatStore.setCurrentProjectId(project.id)

    // Show success modal
    showSuccessModal.value = true

    // Add welcome message from AI
    await chatStore.addSystemMessage(
      `Great! I've helped create your ${project.language.toUpperCase()} project "${project.name}". You can now generate the files and download your complete project structure. Is there anything specific you'd like to configure?`
    )

    // Update recent projects
    updateRecentProjects(project)

  } catch (error) {
    console.error('Failed to create project:', error)
  }
}

// Handle draft saving
async function handleSaveDraft(projectData) {
  try {
    const project = await projectStore.createProject({
      ...projectData,
      name: `${projectData.name}-draft-${Date.now()}`
    })

    // Add system message
    await chatStore.addSystemMessage(
      `Draft saved! You can continue configuring your ${project.language.toUpperCase()} project.`
    )

    updateRecentProjects(project)

  } catch (error) {
    console.error('Failed to save draft:', error)
  }
}

// Handle AI suggestion applied
function handleSuggestionApplied(suggestion) {
  // Update project options based on suggestion
  const updates = {}
  updates[suggestion.type] = suggestion.value
  projectStore.updateProjectOptions(updates)
}

// Generate project files
async function generateFiles() {
  if (!projectStore.currentProject) return

  try {
    await projectStore.generateProjectFiles(projectStore.currentProject.id)

    // Add system message
    await chatStore.addSystemMessage(
      `Files generated successfully! Your ${projectStore.currentProject.language.toUpperCase()} project is ready for download.`
    )
  } catch (error) {
    console.error('Failed to generate files:', error)
  }
}

// Download project
async function downloadProject() {
  if (!projectStore.currentProject) return

  try {
    await projectStore.downloadProject(projectStore.currentProject.id)
    showSuccessModal.value = false
  } catch (error) {
    console.error('Failed to download project:', error)
  }
}

// Load existing project
async function loadProject(projectId) {
  try {
    await projectStore.loadProject(projectId)
    await chatStore.loadChatHistory(projectId)
  } catch (error) {
    console.error('Failed to load project:', error)
  }
}

// Utility functions
function updateRecentProjects(project) {
  // Add to recent projects (keep only last 6)
  recentProjects.value = [project, ...recentProjects.value.slice(0, 5)]

  // Save to localStorage
  localStorage.setItem('recentProjects', JSON.stringify(recentProjects.value))
}

function loadRecentProjects() {
  try {
    const saved = localStorage.getItem('recentProjects')
    if (saved) {
      recentProjects.value = JSON.parse(saved)
    }
  } catch (error) {
    console.error('Failed to load recent projects:', error)
  }
}

function formatDate(dateString) {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

// Lifecycle
onMounted(() => {
  loadRecentProjects()
})
</script>

<style scoped>
.btn-sm {
  @apply px-3 py-1.5 text-sm;
}
</style>