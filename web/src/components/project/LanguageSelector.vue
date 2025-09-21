<template>
  <div class="space-y-4">
    <div>
      <label class="form-label">
        Project Language
        <span class="text-red-500">*</span>
      </label>
      <p class="text-sm text-gray-500 mb-3">
        Choose your preferred programming language and framework
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <!-- Go Option -->
      <label
        class="language-card"
        :class="{ 'language-card-selected': modelValue === 'go' }"
      >
        <input
          type="radio"
          :value="'go'"
          :checked="modelValue === 'go'"
          @change="$emit('update:modelValue', 'go')"
          class="sr-only"
        />

        <div class="flex items-start space-x-3">
          <!-- Go Logo -->
          <div class="flex-shrink-0 w-10 h-10 rounded-lg bg-blue-100 flex items-center justify-center">
            <svg class="w-6 h-6 text-blue-600" viewBox="0 0 24 24" fill="currentColor">
              <path d="M1.811 10.231c-.047 0-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046 0 .058.035.035.07l-.199.303c-.023.036-.082.07-.117.07H1.811zM.047 11.306c-.047 0-.059-.023-.035-.058l.245-.316c.023-.035.082-.058.129-.058h5.328c.047 0 .07.035.058.07l-.093.28c-.012.047-.058.082-.105.082H.047zM2.828 12.381c-.047 0-.059-.035-.035-.07l.163-.292c.023-.035.070-.070.117-.070h2.337c.047 0 .070.035.070.082l-.023.245c0 .047-.035.093-.082.093H2.828zM21.293 12.032c-1.32 0-2.234-.5-2.234-1.18 0-.656.914-1.156 2.234-1.156 1.32 0 2.234.5 2.234 1.156 0 .68-.914 1.18-2.234 1.18zm0-1.77c-.902 0-1.40.304-1.40.59 0 .285.498.589 1.40.589.903 0 1.401-.304 1.401-.589 0-.286-.498-.59-1.401-.59zM7.23 15.734c-1.168 0-1.961-.445-1.961-1.109 0-.688.793-1.133 1.961-1.133s1.961.445 1.961 1.133c0 .664-.793 1.109-1.961 1.109zm0-1.688c-.793 0-1.262.281-1.262.578 0 .297.469.578 1.262.578.793 0 1.261-.281 1.261-.578 0-.297-.468-.578-1.261-.578z"/>
            </svg>
          </div>

          <div class="flex-1">
            <h3 class="text-lg font-semibold text-gray-900 mb-1">Go</h3>
            <p class="text-sm text-gray-600 mb-3">
              Fast, reliable, and efficient. Perfect for APIs, microservices, and system programming.
            </p>

            <!-- Go Features -->
            <div class="space-y-2">
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Clean Architecture with Gin Framework
              </div>
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                17 Built-in Utility Packages
              </div>
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                JWT Authentication & Security
              </div>
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Database Support (PostgreSQL, MySQL, etc.)
              </div>
            </div>
          </div>
        </div>

        <!-- Selection Indicator -->
        <div class="absolute top-4 right-4">
          <div
            class="w-5 h-5 rounded-full border-2 flex items-center justify-center"
            :class="modelValue === 'go'
              ? 'border-primary-600 bg-primary-600'
              : 'border-gray-300'"
          >
            <CheckIcon
              v-if="modelValue === 'go'"
              class="w-3 h-3 text-white"
            />
          </div>
        </div>
      </label>

      <!-- PHP Option -->
      <label
        class="language-card"
        :class="{ 'language-card-selected': modelValue === 'php' }"
      >
        <input
          type="radio"
          :value="'php'"
          :checked="modelValue === 'php'"
          @change="$emit('update:modelValue', 'php')"
          class="sr-only"
        />

        <div class="flex items-start space-x-3">
          <!-- PHP Logo -->
          <div class="flex-shrink-0 w-10 h-10 rounded-lg bg-purple-100 flex items-center justify-center">
            <svg class="w-6 h-6 text-purple-600" viewBox="0 0 24 24" fill="currentColor">
              <path d="M7.01 10.207h-.944l-.515 2.648h.838c.556 0 .982-.122 1.292-.391.291-.291.464-.683.464-1.212 0-.33-.097-.581-.304-.753-.207-.17-.536-.292-.831-.292zm8.512-.207h-.943l-.515 2.648h.838c.556 0 .982-.122 1.292-.391.291-.291.464-.683.464-1.212 0-.33-.097-.581-.304-.753-.207-.17-.536-.292-.832-.292zM24 12c0 6.627-5.373 12-12 12S0 18.627 0 12 5.373 0 12 0s12 5.373 12 12zM7.99 16.15c-.664.664-1.567.982-2.704.982H2.72L4.142 8.47h2.567c1.137 0 2.04.318 2.704.982.664.664.982 1.567.982 2.704s-.318 2.04-.982 2.704-.982.982-2.704.982zm8.512 0c-.664.664-1.567.982-2.704.982h-2.567L12.654 8.47h2.567c1.137 0 2.04.318 2.704.982.664.664.982 1.567.982 2.704s-.318 2.04-.982 2.704z"/>
            </svg>
          </div>

          <div class="flex-1">
            <h3 class="text-lg font-semibold text-gray-900 mb-1">PHP CodeIgniter</h3>
            <p class="text-sm text-gray-600 mb-3">
              Rapid web development with MVC architecture. Great for web applications and APIs.
            </p>

            <!-- PHP Features -->
            <div class="space-y-2">
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Complete MVC Architecture
              </div>
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Built-in Security Features
              </div>
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Admin Panel & User Management
              </div>
              <div class="flex items-center text-sm text-gray-700">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Database Integration & Migrations
              </div>
            </div>
          </div>
        </div>

        <!-- Selection Indicator -->
        <div class="absolute top-4 right-4">
          <div
            class="w-5 h-5 rounded-full border-2 flex items-center justify-center"
            :class="modelValue === 'php'
              ? 'border-primary-600 bg-primary-600'
              : 'border-gray-300'"
          >
            <CheckIcon
              v-if="modelValue === 'php'"
              class="w-3 h-3 text-white"
            />
          </div>
        </div>
      </label>
    </div>

    <!-- Language-specific info -->
    <div v-if="modelValue" class="p-4 bg-blue-50 rounded-lg">
      <div class="flex items-start space-x-2">
        <InformationCircleIcon class="w-5 h-5 text-blue-500 flex-shrink-0 mt-0.5" />
        <div class="text-sm">
          <p class="font-medium text-blue-900 mb-1">
            {{ modelValue === 'go' ? 'Go Project' : 'PHP CodeIgniter Project' }} Selected
          </p>
          <p class="text-blue-700">
            {{ modelValue === 'go'
              ? 'You\'ll get a production-ready Go application with Clean Architecture, 17 utility packages, and modern best practices.'
              : 'You\'ll get a complete CodeIgniter application with MVC structure, security features, and admin functionality.'
            }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { CheckIcon, InformationCircleIcon } from '@heroicons/vue/24/outline'

defineProps({
  modelValue: {
    type: String,
    default: null
  }
})

defineEmits(['update:modelValue'])
</script>

<style scoped>
.language-card {
  @apply relative cursor-pointer border-2 border-gray-200 rounded-xl p-6 hover:border-primary-300 hover:shadow-md transition-all duration-200 bg-white;
}

.language-card-selected {
  @apply border-primary-600 shadow-md bg-primary-50;
}
</style>