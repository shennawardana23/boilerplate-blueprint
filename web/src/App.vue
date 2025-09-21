<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <AppHeader />
    <main class="flex-1">
      <router-view />
    </main>
    <AppFooter />

    <!-- Global Loading Overlay -->
    <Teleport to="body">
      <div
        v-if="isGlobalLoading"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      >
        <div class="bg-white rounded-lg p-6 flex items-center space-x-3">
          <LoadingSpinner class="w-6 h-6" />
          <span class="text-gray-700">{{ loadingMessage }}</span>
        </div>
      </div>
    </Teleport>

    <!-- Global Error Toast -->
    <Teleport to="body">
      <Transition
        name="toast"
        enter-active-class="transition-all duration-300"
        enter-from-class="transform translate-x-full opacity-0"
        enter-to-class="transform translate-x-0 opacity-100"
        leave-active-class="transition-all duration-300"
        leave-from-class="transform translate-x-0 opacity-100"
        leave-to-class="transform translate-x-full opacity-0"
      >
        <div
          v-if="globalError"
          class="fixed top-4 right-4 max-w-sm bg-red-50 border border-red-200 rounded-lg p-4 shadow-lg z-50"
        >
          <div class="flex items-start">
            <div class="flex-shrink-0">
              <ExclamationTriangleIcon class="h-5 w-5 text-red-400" />
            </div>
            <div class="ml-3 flex-1">
              <h3 class="text-sm font-medium text-red-800">Error</h3>
              <p class="mt-1 text-sm text-red-700">{{ globalError }}</p>
            </div>
            <div class="ml-4 flex-shrink-0">
              <button
                type="button"
                class="inline-flex text-red-400 hover:text-red-500"
                @click="clearGlobalError"
              >
                <span class="sr-only">Close</span>
                <XMarkIcon class="h-5 w-5" />
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup>
import { computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ExclamationTriangleIcon, XMarkIcon } from '@heroicons/vue/24/outline'
import { useProjectStore } from '@/stores/project'
import { useChatStore } from '@/stores/chat'

import AppHeader from '@/components/layout/AppHeader.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import LoadingSpinner from '@/components/ui/LoadingSpinner.vue'

const route = useRoute()
const projectStore = useProjectStore()
const chatStore = useChatStore()

// Global loading state
const isGlobalLoading = computed(() =>
  projectStore.isLoading || chatStore.isLoading
)

const loadingMessage = computed(() => {
  if (projectStore.isLoading) {
    return 'Processing project...'
  }
  if (chatStore.isLoading) {
    return 'AI is thinking...'
  }
  return 'Loading...'
})

// Global error state
const globalError = computed(() =>
  projectStore.error || chatStore.error
)

function clearGlobalError() {
  projectStore.clearError()
  chatStore.clearError()
}

// Load initial data
onMounted(async () => {
  try {
    await projectStore.loadTemplates()
  } catch (error) {
    console.error('Failed to load initial data:', error)
  }
})

// Clear errors when route changes
watch(route, () => {
  clearGlobalError()
})
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}

.toast-enter-from {
  transform: translateX(100%);
  opacity: 0;
}

.toast-leave-to {
  transform: translateX(100%);
  opacity: 0;
}
</style>