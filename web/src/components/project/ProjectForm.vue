<template>
  <form @submit.prevent="handleSubmit" class="space-y-6">
    <!-- Project Basic Info -->
    <div class="card">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">Project Information</h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label for="project-name" class="form-label">
            Project Name
            <span class="text-red-500">*</span>
          </label>
          <input
            id="project-name"
            v-model="formData.name"
            type="text"
            required
            class="form-input"
            placeholder="my-awesome-project"
            @blur="validateProjectName"
          />
          <p v-if="errors.name" class="mt-1 text-sm text-red-600">{{ errors.name }}</p>
          <p v-else class="mt-1 text-sm text-gray-500">
            Use lowercase letters, numbers, and hyphens only
          </p>
        </div>

        <div>
          <label for="description" class="form-label">Description</label>
          <input
            id="description"
            v-model="formData.description"
            type="text"
            class="form-input"
            placeholder="A brief description of your project"
          />
        </div>
      </div>
    </div>

    <!-- Language Selection -->
    <div class="card">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">Language & Framework</h3>
      <LanguageSelector v-model="formData.language" />
    </div>

    <!-- Language-specific Options -->
    <div v-if="formData.language" class="card">
      <h3 class="text-lg font-semibold text-gray-900 mb-4">Configuration Options</h3>

      <!-- Go Options -->
      <div v-if="formData.language === 'go'" class="space-y-6">
        <!-- Framework Selection -->
        <div>
          <label for="go-framework" class="form-label">HTTP Framework</label>
          <select
            id="go-framework"
            v-model="formData.options.framework"
            class="form-select"
          >
            <option value="gin">Gin (Recommended)</option>
            <option value="chi">Chi</option>
            <option value="echo">Echo</option>
            <option value="standard">Standard Library</option>
          </select>
          <p class="mt-1 text-sm text-gray-500">
            Choose your preferred HTTP framework for the web server
          </p>
        </div>

        <!-- Database Selection -->
        <div>
          <label for="go-database" class="form-label">Database</label>
          <select
            id="go-database"
            v-model="formData.options.database"
            class="form-select"
          >
            <option value="postgresql">PostgreSQL (Recommended)</option>
            <option value="mysql">MySQL</option>
            <option value="sqlite">SQLite</option>
            <option value="mongodb">MongoDB</option>
          </select>
        </div>

        <!-- Authentication -->
        <div>
          <label for="go-auth" class="form-label">Authentication</label>
          <select
            id="go-auth"
            v-model="formData.options.authentication"
            class="form-select"
          >
            <option value="jwt">JWT (Recommended)</option>
            <option value="oauth">OAuth 2.0</option>
            <option value="basic">Basic Auth</option>
          </select>
        </div>

        <!-- Utility Packages -->
        <div>
          <label class="form-label">Utility Packages</label>
          <p class="text-sm text-gray-500 mb-3">
            Select the utility packages to include (all recommended for enterprise applications)
          </p>
          <div class="grid grid-cols-2 md:grid-cols-3 gap-3">
            <label
              v-for="utility in goUtilities"
              :key="utility.key"
              class="flex items-center space-x-2 cursor-pointer"
            >
              <input
                v-model="formData.options.utilities"
                type="checkbox"
                :value="utility.key"
                class="rounded border-gray-300 text-primary-600 focus:ring-primary-500"
              />
              <span class="text-sm text-gray-700">{{ utility.label }}</span>
            </label>
          </div>
        </div>
      </div>

      <!-- PHP Options -->
      <div v-if="formData.language === 'php'" class="space-y-6">
        <!-- CodeIgniter Version -->
        <div>
          <label for="php-version" class="form-label">CodeIgniter Version</label>
          <select
            id="php-version"
            v-model="formData.options.ci_version"
            class="form-select"
          >
            <option value="3">CodeIgniter 3 (Stable)</option>
            <option value="4">CodeIgniter 4 (Latest)</option>
          </select>
        </div>

        <!-- Database Selection -->
        <div>
          <label for="php-database" class="form-label">Database</label>
          <select
            id="php-database"
            v-model="formData.options.database"
            class="form-select"
          >
            <option value="postgresql">PostgreSQL</option>
            <option value="mysql">MySQL (Recommended)</option>
            <option value="sqlite">SQLite</option>
          </select>
        </div>

        <!-- Frontend Framework -->
        <div>
          <label for="php-frontend" class="form-label">Frontend Framework</label>
          <select
            id="php-frontend"
            v-model="formData.options.frontend"
            class="form-select"
          >
            <option value="bootstrap">Bootstrap (Recommended)</option>
            <option value="tailwind">Tailwind CSS</option>
            <option value="custom">Custom CSS</option>
          </select>
        </div>

        <!-- Features -->
        <div>
          <label class="form-label">Features</label>
          <p class="text-sm text-gray-500 mb-3">
            Select the features to include in your CodeIgniter application
          </p>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <label
              v-for="feature in phpFeatures"
              :key="feature.key"
              class="flex items-center space-x-2 cursor-pointer"
            >
              <input
                v-model="formData.options.features"
                type="checkbox"
                :value="feature.key"
                class="rounded border-gray-300 text-primary-600 focus:ring-primary-500"
              />
              <span class="text-sm text-gray-700">{{ feature.label }}</span>
            </label>
          </div>
        </div>
      </div>
    </div>

    <!-- Actions -->
    <div class="flex justify-between items-center pt-6 border-t border-gray-200">
      <button
        type="button"
        @click="resetForm"
        class="btn btn-secondary"
        :disabled="isLoading"
      >
        Reset Form
      </button>

      <div class="flex space-x-3">
        <button
          type="button"
          @click="saveAsDraft"
          class="btn btn-outline"
          :disabled="isLoading || !formData.name || !formData.language"
        >
          <DocumentIcon class="w-4 h-4 mr-2" />
          Save as Draft
        </button>

        <button
          type="submit"
          class="btn btn-primary"
          :disabled="isLoading || !isFormValid"
        >
          <LoadingSpinner v-if="isLoading" size="sm" color="white" class="mr-2" />
          <RocketLaunchIcon v-else class="w-4 h-4 mr-2" />
          {{ isLoading ? 'Creating...' : 'Create Project' }}
        </button>
      </div>
    </div>
  </form>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { DocumentIcon, RocketLaunchIcon } from '@heroicons/vue/24/outline'
import LanguageSelector from './LanguageSelector.vue'
import LoadingSpinner from '@/components/ui/LoadingSpinner.vue'

const props = defineProps({
  isLoading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit', 'draft'])

// Form data
const formData = ref({
  name: '',
  description: '',
  language: null,
  options: {
    // Go options
    framework: 'gin',
    database: 'postgresql',
    authentication: 'jwt',
    utilities: [
      'authentication', 'cache', 'common', 'constants', 'converter',
      'date', 'datatype', 'encryption', 'exception', 'exceptioncode',
      'helper', 'httphelper', 'json', 'logger', 'password',
      'queryhelper', 'sort', 'template', 'validator', 'alert'
    ],
    // PHP options
    ci_version: '3',
    frontend: 'bootstrap',
    features: ['authentication', 'user_management', 'dashboard']
  }
})

// Form validation
const errors = ref({})

const goUtilities = [
  { key: 'authentication', label: 'Authentication' },
  { key: 'cache', label: 'Cache' },
  { key: 'common', label: 'Common' },
  { key: 'constants', label: 'Constants' },
  { key: 'converter', label: 'Converter' },
  { key: 'date', label: 'Date' },
  { key: 'datatype', label: 'Data Type' },
  { key: 'encryption', label: 'Encryption' },
  { key: 'exception', label: 'Exception' },
  { key: 'exceptioncode', label: 'Exception Code' },
  { key: 'helper', label: 'Helper' },
  { key: 'httphelper', label: 'HTTP Helper' },
  { key: 'json', label: 'JSON' },
  { key: 'logger', label: 'Logger' },
  { key: 'password', label: 'Password' },
  { key: 'queryhelper', label: 'Query Helper' },
  { key: 'sort', label: 'Sort' },
  { key: 'template', label: 'Template' },
  { key: 'validator', label: 'Validator' },
  { key: 'alert', label: 'Alert' }
]

const phpFeatures = [
  { key: 'authentication', label: 'User Authentication' },
  { key: 'user_management', label: 'User Management' },
  { key: 'dashboard', label: 'Admin Dashboard' },
  { key: 'api', label: 'REST API' },
  { key: 'file_upload', label: 'File Upload' },
  { key: 'email', label: 'Email System' },
  { key: 'notifications', label: 'Notifications' },
  { key: 'logging', label: 'Activity Logging' }
]

const isFormValid = computed(() => {
  return formData.value.name && formData.value.language && !errors.value.name
})

// Validation
function validateProjectName() {
  const name = formData.value.name
  if (!name) {
    errors.value.name = 'Project name is required'
    return false
  }

  if (!/^[a-z0-9-]+$/.test(name)) {
    errors.value.name = 'Project name must contain only lowercase letters, numbers, and hyphens'
    return false
  }

  if (name.startsWith('-') || name.endsWith('-')) {
    errors.value.name = 'Project name cannot start or end with a hyphen'
    return false
  }

  errors.value.name = null
  return true
}

// Reset language-specific options when language changes
watch(() => formData.value.language, (newLang, oldLang) => {
  if (newLang !== oldLang) {
    if (newLang === 'go') {
      formData.value.options = {
        ...formData.value.options,
        framework: 'gin',
        database: 'postgresql',
        authentication: 'jwt',
        utilities: [...goUtilities.map(u => u.key)]
      }
    } else if (newLang === 'php') {
      formData.value.options = {
        ...formData.value.options,
        ci_version: '3',
        database: 'mysql',
        frontend: 'bootstrap',
        features: ['authentication', 'user_management', 'dashboard']
      }
    }
  }
})

// Form actions
function handleSubmit() {
  if (!validateProjectName()) return
  if (!isFormValid.value) return

  emit('submit', { ...formData.value })
}

function saveAsDraft() {
  if (!formData.value.name || !formData.value.language) return

  emit('draft', { ...formData.value })
}

function resetForm() {
  formData.value = {
    name: '',
    description: '',
    language: null,
    options: {
      framework: 'gin',
      database: 'postgresql',
      authentication: 'jwt',
      utilities: [...goUtilities.map(u => u.key)],
      ci_version: '3',
      frontend: 'bootstrap',
      features: ['authentication', 'user_management', 'dashboard']
    }
  }
  errors.value = {}
}
</script>