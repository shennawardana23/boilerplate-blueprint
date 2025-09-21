<template>
  <div class="space-y-4">
    <!-- Project Info -->
    <div v-if="project" class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
      <div class="flex items-center space-x-3">
        <div class="w-10 h-10 bg-primary-100 rounded-lg flex items-center justify-center">
          <FolderIcon class="w-6 h-6 text-primary-600" />
        </div>
        <div>
          <h4 class="font-semibold text-gray-900">{{ project.name }}</h4>
          <p class="text-sm text-gray-500">
            {{ project.language.toUpperCase() }} •
            {{ fileCount }} files •
            {{ directoryCount }} directories
          </p>
        </div>
      </div>
      <span class="badge" :class="languageBadgeClass">
        {{ project.language.toUpperCase() }}
      </span>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex items-center justify-center py-12">
      <div class="text-center">
        <LoadingSpinner size="lg" class="mb-4" />
        <p class="text-gray-600">Generating project structure...</p>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="files.length === 0" class="text-center py-12">
      <FolderOpenIcon class="w-12 h-12 text-gray-400 mx-auto mb-4" />
      <h4 class="text-lg font-medium text-gray-900 mb-2">No Files Generated</h4>
      <p class="text-gray-500 mb-4">
        Click "Generate Files" to create your project structure
      </p>
    </div>

    <!-- File Tree -->
    <div v-else class="border border-gray-200 rounded-lg overflow-hidden">
      <div class="bg-gray-50 px-4 py-2 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h4 class="font-medium text-gray-900">Project Structure</h4>
          <div class="flex items-center space-x-2">
            <button
              @click="expandAll"
              class="text-xs text-gray-600 hover:text-gray-800"
            >
              Expand All
            </button>
            <span class="text-gray-300">|</span>
            <button
              @click="collapseAll"
              class="text-xs text-gray-600 hover:text-gray-800"
            >
              Collapse All
            </button>
          </div>
        </div>
      </div>

      <div class="max-h-96 overflow-y-auto font-mono text-sm">
        <FileTreeNode
          v-for="node in fileTree"
          :key="node.path"
          :node="node"
          :level="0"
          :expanded-nodes="expandedNodes"
          @toggle="toggleNode"
          @select="selectFile"
        />
      </div>
    </div>

    <!-- File Content Preview -->
    <div v-if="selectedFile && selectedFile.content" class="border border-gray-200 rounded-lg overflow-hidden">
      <div class="bg-gray-50 px-4 py-2 border-b border-gray-200">
        <h4 class="font-medium text-gray-900">{{ selectedFile.path }}</h4>
      </div>
      <div class="max-h-64 overflow-auto">
        <pre class="p-4 text-sm text-gray-800 whitespace-pre-wrap">{{ selectedFile.content }}</pre>
      </div>
    </div>

    <!-- Project Statistics -->
    <div v-if="files.length > 0" class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <div class="text-center p-3 bg-gray-50 rounded-lg">
        <div class="text-lg font-semibold text-gray-900">{{ fileCount }}</div>
        <div class="text-xs text-gray-600">Files</div>
      </div>
      <div class="text-center p-3 bg-gray-50 rounded-lg">
        <div class="text-lg font-semibold text-gray-900">{{ directoryCount }}</div>
        <div class="text-xs text-gray-600">Directories</div>
      </div>
      <div class="text-center p-3 bg-gray-50 rounded-lg">
        <div class="text-lg font-semibold text-gray-900">{{ codeFileCount }}</div>
        <div class="text-xs text-gray-600">Code Files</div>
      </div>
      <div class="text-center p-3 bg-gray-50 rounded-lg">
        <div class="text-lg font-semibold text-gray-900">{{ configFileCount }}</div>
        <div class="text-xs text-gray-600">Config Files</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import {
  FolderIcon,
  FolderOpenIcon
} from '@heroicons/vue/24/outline'
import LoadingSpinner from '@/components/ui/LoadingSpinner.vue'
import FileTreeNode from './FileTreeNode.vue'

const props = defineProps({
  project: {
    type: Object,
    default: null
  },
  files: {
    type: Array,
    default: () => []
  },
  isLoading: {
    type: Boolean,
    default: false
  }
})

const expandedNodes = ref(new Set())
const selectedFile = ref(null)

// Computed properties
const languageBadgeClass = computed(() => {
  return props.project?.language === 'go' ? 'badge-primary' : 'badge-secondary'
})

const fileCount = computed(() => props.files.filter(f => !f.is_directory).length)
const directoryCount = computed(() => props.files.filter(f => f.is_directory).length)

const codeFileCount = computed(() => {
  const codeExtensions = ['.go', '.php', '.js', '.vue', '.ts', '.jsx', '.tsx']
  return props.files.filter(f =>
    !f.is_directory && codeExtensions.some(ext => f.path.endsWith(ext))
  ).length
})

const configFileCount = computed(() => {
  const configFiles = ['go.mod', 'go.sum', 'package.json', 'composer.json', 'Makefile', 'Dockerfile', '.env']
  return props.files.filter(f =>
    !f.is_directory && configFiles.some(name => f.path.endsWith(name))
  ).length
})

const fileTree = computed(() => {
  if (!props.files.length) return []

  const tree = {}

  // Build tree structure
  props.files.forEach(file => {
    const parts = file.path.split('/')
    let current = tree

    parts.forEach((part, index) => {
      if (!current[part]) {
        current[part] = {
          name: part,
          path: parts.slice(0, index + 1).join('/'),
          is_directory: index < parts.length - 1 || file.is_directory,
          content: file.content || '',
          children: {}
        }
      }
      current = current[part].children
    })
  })

  // Convert to array and sort
  const convertToArray = (obj) => {
    return Object.values(obj)
      .map(node => ({
        ...node,
        children: Object.keys(node.children).length > 0 ? convertToArray(node.children) : []
      }))
      .sort((a, b) => {
        // Directories first, then files
        if (a.is_directory && !b.is_directory) return -1
        if (!a.is_directory && b.is_directory) return 1
        return a.name.localeCompare(b.name)
      })
  }

  return convertToArray(tree)
})

// Methods
function toggleNode(path) {
  if (expandedNodes.value.has(path)) {
    expandedNodes.value.delete(path)
  } else {
    expandedNodes.value.add(path)
  }
}

function expandAll() {
  const allPaths = new Set()
  const collectPaths = (nodes) => {
    nodes.forEach(node => {
      if (node.is_directory) {
        allPaths.add(node.path)
        if (node.children) collectPaths(node.children)
      }
    })
  }
  collectPaths(fileTree.value)
  expandedNodes.value = allPaths
}

function collapseAll() {
  expandedNodes.value.clear()
}

function selectFile(file) {
  if (!file.is_directory) {
    selectedFile.value = file
  }
}

// Auto-expand first level on file change
watch(() => props.files, (newFiles) => {
  if (newFiles.length > 0 && expandedNodes.value.size === 0) {
    // Auto-expand first level directories
    const firstLevelDirs = new Set()
    newFiles.forEach(file => {
      const parts = file.path.split('/')
      if (parts.length > 1) {
        firstLevelDirs.add(parts[0])
      }
    })
    expandedNodes.value = firstLevelDirs
  }
}, { immediate: true })
</script>