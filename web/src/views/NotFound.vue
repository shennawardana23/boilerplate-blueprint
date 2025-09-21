<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center px-4 sm:px-6 lg:px-8">
    <div class="max-w-lg w-full text-center">
      <!-- 404 Illustration -->
      <div class="mb-8">
        <div class="relative">
          <!-- Large 404 Text -->
          <h1 class="text-9xl font-bold text-gray-200 select-none">
            404
          </h1>
          
          <!-- Floating Elements -->
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="relative">
              <!-- Code Brackets Animation -->
              <div class="absolute -top-4 -left-8 text-4xl text-primary-400 animate-bounce" style="animation-delay: 0.1s">
                {
              </div>
              <div class="absolute -top-4 -right-8 text-4xl text-primary-400 animate-bounce" style="animation-delay: 0.3s">
                }
              </div>
              
              <!-- Center Icon -->
              <div class="w-16 h-16 bg-primary-100 rounded-full flex items-center justify-center">
                <ExclamationTriangleIcon class="w-8 h-8 text-primary-600" />
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Content -->
      <div class="space-y-6">
        <div>
          <h2 class="text-3xl font-bold text-gray-900 mb-2">
            Oops! Page Not Found
          </h2>
          <p class="text-lg text-gray-600">
            The page you're looking for seems to have been moved, deleted, or doesn't exist.
          </p>
        </div>

        <!-- Suggestions -->
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">Here's what you can do:</h3>
          <div class="space-y-3 text-left">
            <div class="flex items-start">
              <CheckIcon class="w-5 h-5 text-green-500 mr-3 mt-0.5 flex-shrink-0" />
              <span class="text-gray-700">Double-check the URL for any typing errors</span>
            </div>
            <div class="flex items-start">
              <CheckIcon class="w-5 h-5 text-green-500 mr-3 mt-0.5 flex-shrink-0" />
              <span class="text-gray-700">Go back to the previous page using your browser's back button</span>
            </div>
            <div class="flex items-start">
              <CheckIcon class="w-5 h-5 text-green-500 mr-3 mt-0.5 flex-shrink-0" />
              <span class="text-gray-700">Visit our homepage to start fresh</span>
            </div>
            <div class="flex items-start">
              <CheckIcon class="w-5 h-5 text-green-500 mr-3 mt-0.5 flex-shrink-0" />
              <span class="text-gray-700">Use the navigation menu to find what you need</span>
            </div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <router-link
            to="/"
            class="btn btn-primary text-lg px-6 py-3"
          >
            <HomeIcon class="w-5 h-5 mr-2" />
            Go to Homepage
          </router-link>
          
          <router-link
            to="/generator"
            class="btn btn-outline text-lg px-6 py-3"
          >
            <RocketLaunchIcon class="w-5 h-5 mr-2" />
            Start Generating
          </router-link>
          
          <button
            @click="goBack"
            class="btn btn-outline text-lg px-6 py-3"
          >
            <ArrowLeftIcon class="w-5 h-5 mr-2" />
            Go Back
          </button>
        </div>

        <!-- Help Section -->
        <div class="pt-8 border-t border-gray-200">
          <p class="text-gray-600 mb-4">
            Still can't find what you're looking for?
          </p>
          <div class="flex flex-col sm:flex-row gap-3 justify-center">
            <a
              href="mailto:support@boilerplate-blueprint.com"
              class="text-primary-600 hover:text-primary-700 font-medium inline-flex items-center"
            >
              <EnvelopeIcon class="w-4 h-4 mr-1" />
              Contact Support
            </a>
            <span class="hidden sm:inline text-gray-300">•</span>
            <a
              href="https://github.com/your-username/boilerplate-blueprint/issues"
              target="_blank"
              rel="noopener noreferrer"
              class="text-primary-600 hover:text-primary-700 font-medium inline-flex items-center"
            >
              <ExclamationCircleIcon class="w-4 h-4 mr-1" />
              Report an Issue
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import {
  ExclamationTriangleIcon,
  CheckIcon,
  HomeIcon,
  RocketLaunchIcon,
  ArrowLeftIcon,
  EnvelopeIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/24/outline'

const router = useRouter()

function goBack() {
  // Go back in history if possible, otherwise go to homepage
  if (window.history.length > 1) {
    router.go(-1)
  } else {
    router.push('/')
  }
}
</script>

<style scoped>
/* Custom animations for floating elements */
@keyframes bounce {
  0%, 20%, 53%, 80%, 100% {
    transform: translate3d(0,0,0);
  }
  40%, 43% {
    transform: translate3d(0,-8px,0);
  }
  70% {
    transform: translate3d(0,-4px,0);
  }
  90% {
    transform: translate3d(0,-2px,0);
  }
}
</style>
