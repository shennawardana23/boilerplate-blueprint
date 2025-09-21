<template>
  <header class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-40">
    <nav class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <!-- Logo and brand -->
        <div class="flex items-center">
          <router-link
            to="/"
            class="flex items-center space-x-3 text-xl font-bold text-gray-900 hover:text-primary-600 transition-colors"
          >
            <div class="w-8 h-8 bg-gradient-to-br from-primary-500 to-primary-600 rounded-lg flex items-center justify-center">
              <span class="text-white text-sm font-bold">BB</span>
            </div>
            <span class="hidden sm:block">Boilerplate Blueprint</span>
          </router-link>
        </div>

        <!-- Navigation Links -->
        <div class="hidden md:flex items-center space-x-8">
          <router-link
            to="/"
            class="nav-link"
            :class="{ 'nav-link-active': $route.name === 'Home' }"
          >
            Home
          </router-link>
          <router-link
            to="/generator"
            class="nav-link"
            :class="{ 'nav-link-active': $route.name === 'Generator' }"
          >
            Generator
          </router-link>
          <router-link
            to="/about"
            class="nav-link"
            :class="{ 'nav-link-active': $route.name === 'About' }"
          >
            About
          </router-link>
        </div>

        <!-- Actions -->
        <div class="flex items-center space-x-4">
          <!-- GitHub Link -->
          <a
            href="https://github.com/your-username/boilerplate-blueprint"
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 hover:text-gray-700 transition-colors"
            title="View on GitHub"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 0C5.374 0 0 5.373 0 12 0 17.302 3.438 21.8 8.207 23.387c.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0112 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/>
            </svg>
          </a>

          <!-- Get Started Button -->
          <router-link
            to="/generator"
            class="btn btn-primary hidden sm:inline-flex"
          >
            Get Started
          </router-link>

          <!-- Mobile menu button -->
          <button
            type="button"
            class="md:hidden p-2 rounded-lg text-gray-500 hover:text-gray-700 hover:bg-gray-100 transition-colors"
            @click="toggleMobileMenu"
          >
            <span class="sr-only">Open menu</span>
            <Bars3Icon v-if="!isMobileMenuOpen" class="w-6 h-6" />
            <XMarkIcon v-else class="w-6 h-6" />
          </button>
        </div>
      </div>

      <!-- Mobile menu -->
      <Transition
        enter-active-class="transition-all duration-200"
        enter-from-class="transform -translate-y-2 opacity-0"
        enter-to-class="transform translate-y-0 opacity-100"
        leave-active-class="transition-all duration-150"
        leave-from-class="transform translate-y-0 opacity-100"
        leave-to-class="transform -translate-y-2 opacity-0"
      >
        <div v-if="isMobileMenuOpen" class="md:hidden border-t border-gray-200 py-4">
          <div class="space-y-3">
            <router-link
              to="/"
              class="mobile-nav-link"
              :class="{ 'mobile-nav-link-active': $route.name === 'Home' }"
              @click="closeMobileMenu"
            >
              Home
            </router-link>
            <router-link
              to="/generator"
              class="mobile-nav-link"
              :class="{ 'mobile-nav-link-active': $route.name === 'Generator' }"
              @click="closeMobileMenu"
            >
              Generator
            </router-link>
            <router-link
              to="/about"
              class="mobile-nav-link"
              :class="{ 'mobile-nav-link-active': $route.name === 'About' }"
              @click="closeMobileMenu"
            >
              About
            </router-link>
            <div class="pt-2">
              <router-link
                to="/generator"
                class="btn btn-primary w-full justify-center"
                @click="closeMobileMenu"
              >
                Get Started
              </router-link>
            </div>
          </div>
        </div>
      </Transition>
    </nav>
  </header>
</template>

<script setup>
import { ref } from 'vue'
import { Bars3Icon, XMarkIcon } from '@heroicons/vue/24/outline'

const isMobileMenuOpen = ref(false)

function toggleMobileMenu() {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

function closeMobileMenu() {
  isMobileMenuOpen.value = false
}
</script>

<style scoped>
.nav-link {
  @apply text-gray-600 hover:text-gray-900 px-3 py-2 text-sm font-medium transition-colors rounded-lg hover:bg-gray-100;
}

.nav-link-active {
  @apply text-primary-600 bg-primary-50;
}

.mobile-nav-link {
  @apply block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-100 rounded-lg transition-colors;
}

.mobile-nav-link-active {
  @apply text-primary-600 bg-primary-50;
}
</style>