import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { projectApi } from '@/services/api'

export const useProjectStore = defineStore('project', () => {
  // State
  const currentProject = ref(null)
  const projects = ref([])
  const templates = ref([])
  const isLoading = ref(false)
  const error = ref(null)

  // Getters
  const hasCurrentProject = computed(() => currentProject.value !== null)
  const currentProjectFiles = computed(() => currentProject.value?.files || [])
  const isGoProject = computed(() => currentProject.value?.language === 'go')
  const isPHPProject = computed(() => currentProject.value?.language === 'php')

  // Actions
  async function loadTemplates() {
    try {
      isLoading.value = true
      error.value = null
      const response = await projectApi.getTemplates()
      templates.value = response.data.templates
    } catch (err) {
      error.value = err.message
      console.error('Failed to load templates:', err)
    } finally {
      isLoading.value = false
    }
  }

  async function createProject(projectData) {
    try {
      isLoading.value = true
      error.value = null

      const response = await projectApi.createProject(projectData)
      if (response.data.success) {
        currentProject.value = response.data.project
        projects.value.push(response.data.project)
        return response.data.project
      } else {
        throw new Error(response.data.error || 'Failed to create project')
      }
    } catch (err) {
      error.value = err.message
      console.error('Failed to create project:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function loadProject(projectId) {
    try {
      isLoading.value = true
      error.value = null

      const response = await projectApi.getProject(projectId)
      if (response.data.success) {
        currentProject.value = response.data.project
        return response.data.project
      } else {
        throw new Error(response.data.error || 'Failed to load project')
      }
    } catch (err) {
      error.value = err.message
      console.error('Failed to load project:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function generateProjectFiles(projectId) {
    try {
      isLoading.value = true
      error.value = null

      const response = await projectApi.generateProject(projectId)
      if (response.data.success) {
        // Update current project with generated files
        if (currentProject.value && currentProject.value.id === projectId) {
          currentProject.value.files = response.data.files
        }
        return response.data.files
      } else {
        throw new Error(response.data.error || 'Failed to generate project files')
      }
    } catch (err) {
      error.value = err.message
      console.error('Failed to generate project files:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  async function downloadProject(projectId) {
    try {
      isLoading.value = true
      error.value = null

      const response = await projectApi.downloadProject(projectId)

      // Create blob and download
      const blob = new Blob([response.data], { type: 'application/zip' })
      const url = window.URL.createObjectURL(blob)
      const link = document.createElement('a')
      link.href = url

      // Get filename from response headers
      const contentDisposition = response.headers['content-disposition']
      let filename = `project-${projectId}.zip`
      if (contentDisposition) {
        const filenameMatch = contentDisposition.match(/filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/)
        if (filenameMatch && filenameMatch[1]) {
          filename = filenameMatch[1].replace(/['"]/g, '')
        }
      }

      link.download = filename
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
      window.URL.revokeObjectURL(url)

    } catch (err) {
      error.value = err.message
      console.error('Failed to download project:', err)
      throw err
    } finally {
      isLoading.value = false
    }
  }

  function updateProjectOptions(options) {
    if (currentProject.value) {
      currentProject.value.options = { ...currentProject.value.options, ...options }
    }
  }

  function clearCurrentProject() {
    currentProject.value = null
  }

  function clearError() {
    error.value = null
  }

  return {
    // State
    currentProject,
    projects,
    templates,
    isLoading,
    error,

    // Getters
    hasCurrentProject,
    currentProjectFiles,
    isGoProject,
    isPHPProject,

    // Actions
    loadTemplates,
    createProject,
    loadProject,
    generateProjectFiles,
    downloadProject,
    updateProjectOptions,
    clearCurrentProject,
    clearError
  }
})