<template>
  <div class="min-h-screen bg-gray-50 py-16">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="text-center mb-16">
        <h1 class="text-4xl sm:text-5xl font-bold text-gray-900 mb-6">
          About Boilerplate Blueprint
        </h1>
        <p class="text-xl text-gray-600 max-w-3xl mx-auto">
          Empowering developers to build faster with AI-powered project generation and enterprise-grade boilerplate code.
        </p>
      </div>

      <!-- Mission Section -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 mb-12">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
          <div>
            <h2 class="text-3xl font-bold text-gray-900 mb-4">Our Mission</h2>
            <p class="text-lg text-gray-600 mb-6">
              We believe that every developer deserves to start their projects with clean, well-structured, and production-ready code. Boilerplate Blueprint eliminates the repetitive setup work, letting you focus on what matters most - building features that delight your users.
            </p>
            <div class="space-y-3">
              <div class="flex items-center">
                <CheckIcon class="w-5 h-5 text-green-500 mr-3" />
                <span class="text-gray-700">Save hours of initial project setup</span>
              </div>
              <div class="flex items-center">
                <CheckIcon class="w-5 h-5 text-green-500 mr-3" />
                <span class="text-gray-700">Follow industry best practices</span>
              </div>
              <div class="flex items-center">
                <CheckIcon class="w-5 h-5 text-green-500 mr-3" />
                <span class="text-gray-700">Build with confidence from day one</span>
              </div>
            </div>
          </div>
          <div class="flex justify-center">
            <div class="w-64 h-64 bg-gradient-to-br from-primary-100 to-purple-100 rounded-full flex items-center justify-center">
              <RocketLaunchIcon class="w-24 h-24 text-primary-600" />
            </div>
          </div>
        </div>
      </div>

      <!-- Features Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-12">
        <!-- Go Projects -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-8">
          <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-6">
            <svg class="w-8 h-8 text-blue-600" viewBox="0 0 24 24" fill="currentColor">
              <path d="M1.811 10.231c-.047 0-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046 0 .058.035.035.07l-.199.303c-.023.036-.082.07-.117.07H1.811zM.047 11.306c-.047 0-.059-.023-.035-.058l.245-.316c.023-.035.082-.058.129-.058h5.328c.047 0 .07.035.058.07l-.093.28c-.012.047-.058.082-.105.082H.047zM2.828 12.381c-.047 0-.059-.035-.035-.07l.163-.292c.023-.035.070-.070.117-.070h2.337c.047 0 .070.035.070.082l-.023.245c0 .047-.035.093-.082.093H2.828z"/>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-4">Go Projects</h3>
          <p class="text-gray-600 mb-4">
            Generate production-ready Go applications with Clean Architecture, comprehensive utility packages, and enterprise-grade patterns.
          </p>
          <ul class="space-y-2 text-sm text-gray-600">
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              17 Essential Utility Packages
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Clean Architecture Pattern
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Gin Framework Integration
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              JWT Authentication & Security
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Database Integration & Migration
            </li>
          </ul>
        </div>

        <!-- PHP Projects -->
        <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-8">
          <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mb-6">
            <svg class="w-8 h-8 text-purple-600" viewBox="0 0 24 24" fill="currentColor">
              <path d="M7.01 10.207h-.944l-.515 2.648h.838c.556 0 .982-.122 1.292-.391.291-.291.464-.683.464-1.212 0-.33-.097-.581-.304-.753-.207-.17-.536-.292-.831-.292z"/>
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-4">PHP CodeIgniter</h3>
          <p class="text-gray-600 mb-4">
            Create comprehensive CodeIgniter applications with MVC architecture, security features, and admin panels ready for deployment.
          </p>
          <ul class="space-y-2 text-sm text-gray-600">
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Complete MVC Architecture
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Security Headers & CSRF Protection
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Admin Dashboard Interface
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              Database Integration & Seeding
            </li>
            <li class="flex items-center">
              <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
              User Authentication System
            </li>
          </ul>
        </div>
      </div>

      <!-- Technology Stack -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 mb-12">
        <h2 class="text-2xl font-bold text-gray-900 mb-6 text-center">Built With Modern Technology</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
          <div class="text-center">
            <div class="w-16 h-16 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-3">
              <svg class="w-8 h-8 text-blue-600" fill="currentColor" viewBox="0 0 24 24">
                <path d="M1.5 0h21l-1.91 21.563L11.977 24l-8.564-2.438L1.5 0zm7.031 9.75l-.232-2.718 10.059.003.23-2.622L5.412 4.41l.698 8.01h9.126l-.326 3.426-2.91.804-2.955-.81-.188-2.11H6.248l.33 4.171L12 19.351l5.379-1.443.744-8.157H8.531z"/>
              </svg>
            </div>
            <h4 class="font-semibold text-gray-900">Vue.js 3</h4>
            <p class="text-sm text-gray-500">Modern Frontend</p>
          </div>
          <div class="text-center">
            <div class="w-16 h-16 bg-cyan-100 rounded-lg flex items-center justify-center mx-auto mb-3">
              <svg class="w-8 h-8 text-cyan-600" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0C5.374 0 0 5.373 0 12 0 17.302 3.438 21.8 8.207 23.387c.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0112 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"/>
              </svg>
            </div>
            <h4 class="font-semibold text-gray-900">Go Lang</h4>
            <p class="text-sm text-gray-500">Backend API</p>
          </div>
          <div class="text-center">
            <div class="w-16 h-16 bg-green-100 rounded-lg flex items-center justify-center mx-auto mb-3">
              <svg class="w-8 h-8 text-green-600" fill="currentColor" viewBox="0 0 24 24">
                <path d="M11.572 0c-.176 0-.31.001-.358.007a19.76 19.76 0 01-.364.033C7.443.346 4.25 2.185 2.228 5.012a11.875 11.875 0 00-2.119 5.243c-.096.659-.108.854-.108 1.747s.012 1.089.108 1.748c.652 4.506 3.86 8.292 8.209 9.695.779.25 1.6.422 2.534.525.363.04 1.935.04 2.299 0 1.611-.178 2.977-.577 4.323-1.264.207-.106.247-.134.219-.158-.02-.013-.9-1.193-1.955-2.62l-1.919-2.592-2.404-3.558a338.739 338.739 0 00-2.422-3.556c-.009-.002-.018 1.579-.023 3.51-.007 3.38-.01 3.515-.052 3.595a.426.426 0 01-.206.214c-.075.037-.14.044-.495.044H7.81l-.108-.068a.438.438 0 01-.157-.171l-.05-.106.006-4.703.007-4.705.072-.092a.645.645 0 01.174-.143c.096-.047.134-.051.54-.051.478 0 .558.018.682.154.035.038 1.337 1.999 2.895 4.361a10760.433 10760.433 0 004.735 7.17l1.9 2.879.096-.063a12.317 12.317 0 002.466-2.163 11.944 11.944 0 002.824-6.134c.096-.66.108-.854.108-1.748 0-.893-.012-1.088-.108-1.747C23.708 4.367 20.049.631 15.24.092a23.65 23.65 0 00-3.668.008z"/>
              </svg>
            </div>
            <h4 class="font-semibold text-gray-900">Tailwind CSS</h4>
            <p class="text-sm text-gray-500">Styling Framework</p>
          </div>
          <div class="text-center">
            <div class="w-16 h-16 bg-yellow-100 rounded-lg flex items-center justify-center mx-auto mb-3">
              <SparklesIcon class="w-8 h-8 text-yellow-600" />
            </div>
            <h4 class="font-semibold text-gray-900">AI Integration</h4>
            <p class="text-sm text-gray-500">Smart Assistance</p>
          </div>
        </div>
      </div>

      <!-- Team Section -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-8 mb-12">
        <h2 class="text-2xl font-bold text-gray-900 mb-6 text-center">Open Source & Community Driven</h2>
        <div class="text-center">
          <p class="text-lg text-gray-600 mb-6">
            Boilerplate Blueprint is open source and built by developers, for developers. We believe in transparency, collaboration, and continuous improvement through community feedback.
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <a
              href="https://github.com/your-username/boilerplate-blueprint"
              target="_blank"
              rel="noopener noreferrer"
              class="btn btn-outline text-lg px-6 py-3"
            >
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
              View on GitHub
            </a>
            <router-link
              to="/generator"
              class="btn btn-primary text-lg px-6 py-3"
            >
              <RocketLaunchIcon class="w-5 h-5 mr-2" />
              Start Building
            </router-link>
          </div>
        </div>
      </div>

      <!-- Contact Section -->
      <div class="text-center">
        <h2 class="text-2xl font-bold text-gray-900 mb-4">Have Questions or Feedback?</h2>
        <p class="text-lg text-gray-600 mb-6">
          We're always looking to improve and would love to hear from you.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <a
            href="mailto:support@boilerplate-blueprint.com"
            class="btn btn-outline"
          >
            <EnvelopeIcon class="w-4 h-4 mr-2" />
            Contact Us
          </a>
          <a
            href="https://github.com/your-username/boilerplate-blueprint/issues"
            target="_blank"
            rel="noopener noreferrer"
            class="btn btn-outline"
          >
            <ExclamationCircleIcon class="w-4 h-4 mr-2" />
            Report Issue
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  CheckIcon,
  RocketLaunchIcon,
  SparklesIcon,
  EnvelopeIcon,
  ExclamationCircleIcon
} from '@heroicons/vue/24/outline'
</script>
