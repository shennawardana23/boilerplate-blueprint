<template>
  <div>
    <div
      class="flex items-center py-1 px-2 hover:bg-gray-50 cursor-pointer"
      :style="{ paddingLeft: `${level * 20 + 8}px` }"
      @click="handleClick"
    >
      <!-- Expand/Collapse Icon -->
      <button
        v-if="node.is_directory"
        class="w-4 h-4 mr-1 flex items-center justify-center text-gray-500 hover:text-gray-700"
        @click.stop="$emit('toggle', node.path)"
      >
        <ChevronRightIcon
          v-if="!isExpanded"
          class="w-3 h-3"
        />
        <ChevronDownIcon
          v-else
          class="w-3 h-3"
        />
      </button>
      <div v-else class="w-5 mr-1"></div>

      <!-- File/Folder Icon -->
      <div class="w-4 h-4 mr-2 flex items-center justify-center">
        <component
          :is="fileIcon"
          class="w-4 h-4"
          :class="fileIconColor"
        />
      </div>

      <!-- File/Folder Name -->
      <span
        class="text-sm truncate"
        :class="node.is_directory ? 'font-medium text-gray-900' : 'text-gray-700'"
      >
        {{ node.name }}
      </span>

      <!-- File Size/Type Badge -->
      <span
        v-if="!node.is_directory && fileBadge"
        class="ml-auto text-xs px-1.5 py-0.5 bg-gray-100 text-gray-600 rounded"
      >
        {{ fileBadge }}
      </span>
    </div>

    <!-- Children (when expanded) -->
    <div v-if="node.is_directory && isExpanded && node.children">
      <FileTreeNode
        v-for="child in node.children"
        :key="child.path"
        :node="child"
        :level="level + 1"
        :expanded-nodes="expandedNodes"
        @toggle="$emit('toggle', $event)"
        @select="$emit('select', $event)"
      />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import {
  ChevronRightIcon,
  ChevronDownIcon,
  FolderIcon,
  DocumentIcon,
  Cog6ToothIcon,
  CodeBracketIcon,
  DocumentTextIcon,
  PhotoIcon
} from '@heroicons/vue/24/outline'

const props = defineProps({
  node: {
    type: Object,
    required: true
  },
  level: {
    type: Number,
    default: 0
  },
  expandedNodes: {
    type: Set,
    required: true
  }
})

defineEmits(['toggle', 'select'])

const isExpanded = computed(() => props.expandedNodes.has(props.node.path))

const fileIcon = computed(() => {
  if (props.node.is_directory) {
    return FolderIcon
  }

  const extension = props.node.name.split('.').pop()?.toLowerCase()

  switch (extension) {
    case 'go':
    case 'php':
    case 'js':
    case 'ts':
    case 'jsx':
    case 'tsx':
    case 'vue':
      return CodeBracketIcon
    case 'json':
    case 'yml':
    case 'yaml':
    case 'toml':
    case 'env':
      return Cog6ToothIcon
    case 'md':
    case 'txt':
    case 'rst':
      return DocumentTextIcon
    case 'png':
    case 'jpg':
    case 'jpeg':
    case 'gif':
    case 'svg':
      return PhotoIcon
    default:
      return DocumentIcon
  }
})

const fileIconColor = computed(() => {
  if (props.node.is_directory) {
    return 'text-blue-500'
  }

  const extension = props.node.name.split('.').pop()?.toLowerCase()

  switch (extension) {
    case 'go':
      return 'text-blue-600'
    case 'php':
      return 'text-purple-600'
    case 'js':
    case 'jsx':
      return 'text-yellow-600'
    case 'ts':
    case 'tsx':
      return 'text-blue-500'
    case 'vue':
      return 'text-green-600'
    case 'json':
      return 'text-orange-500'
    case 'yml':
    case 'yaml':
      return 'text-red-500'
    case 'md':
      return 'text-gray-600'
    case 'env':
      return 'text-green-500'
    default:
      return 'text-gray-500'
  }
})

const fileBadge = computed(() => {
  if (props.node.is_directory) return null

  const extension = props.node.name.split('.').pop()?.toLowerCase()

  switch (extension) {
    case 'go':
      return 'GO'
    case 'php':
      return 'PHP'
    case 'js':
      return 'JS'
    case 'ts':
      return 'TS'
    case 'vue':
      return 'VUE'
    case 'json':
      return 'JSON'
    case 'yml':
    case 'yaml':
      return 'YAML'
    case 'md':
      return 'MD'
    case 'toml':
      return 'TOML'
    default:
      return null
  }
})

function handleClick() {
  if (props.node.is_directory) {
    // Toggle directory expansion
    // Note: We don't emit toggle here to avoid double-toggle with the button click
  } else {
    // Select file for preview
    emit('select', props.node)
  }
}
</script>