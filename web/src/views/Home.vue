<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 via-white to-gray-50">
    <!-- Hero Section -->
    <section class="relative overflow-hidden pt-16 pb-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <!-- Main heading -->
          <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold text-gray-900 mb-6">
            <span class="block">Generate Professional</span>
            <span class="block text-transparent bg-clip-text bg-gradient-to-r from-primary-600 to-purple-600">
              Boilerplate Projects
            </span>
            <span class="block">with AI Assistance</span>
          </h1>

          <!-- Subtitle -->
          <p class="text-xl text-gray-600 mb-8 max-w-3xl mx-auto leading-relaxed">
            Create production-ready Go and PHP CodeIgniter projects in seconds.
            AI-powered project generation with clean architecture, security features,
            and enterprise-grade code patterns.
          </p>

          <!-- CTA Buttons -->
          <div class="flex flex-col sm:flex-row gap-4 justify-center mb-12">
            <router-link
              to="/generator"
              class="btn btn-primary text-lg px-8 py-3 shadow-lg hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-200"
            >
              <RocketLaunchIcon class="w-5 h-5 mr-2" />
              Start Generating
            </router-link>
            <a
              href="https://github.com/your-username/boilerplate-blueprint"
              target="_blank"
              rel="noopener noreferrer"
              class="btn btn-outline text-lg px-8 py-3 shadow hover:shadow-md transition-all duration-200"
            >
              <CodeBracketIcon class="w-5 h-5 mr-2" />
              View on GitHub
            </a>
          </div>

          <!-- Quick stats -->
          <div class="grid grid-cols-2 md:grid-cols-4 gap-6 max-w-2xl mx-auto">
            <div class="text-center">
              <div class="text-2xl font-bold text-primary-600">2</div>
              <div class="text-sm text-gray-600">Languages</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-primary-600">17+</div>
              <div class="text-sm text-gray-600">Utility Packages</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-primary-600">100%</div>
              <div class="text-sm text-gray-600">AI Powered</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-primary-600">0s</div>
              <div class="text-sm text-gray-600">Setup Time</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Background decoration -->
      <div class="absolute inset-0 -z-10">
        <div class="absolute top-0 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <div class="w-[800px] h-[800px] bg-gradient-to-r from-primary-100 to-purple-100 rounded-full opacity-20 blur-3xl"></div>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
            Why Choose Boilerplate Blueprint?
          </h2>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto">
            Generate enterprise-grade projects with best practices, security features, and clean architecture patterns built-in.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <!-- Go Projects -->
          <div class="feature-card">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-4">
              <svg class="w-8 h-8 text-blue-600" viewBox="0 0 24 24" fill="currentColor">
                <path d="M1.811 10.231c-.047 0-.058-.023-.035-.059l.246-.315c.023-.035.081-.058.128-.058h4.172c.046 0 .058.035.035.07l-.199.303c-.023.036-.082.07-.117.07H1.811zM.047 11.306c-.047 0-.059-.023-.035-.058l.245-.316c.023-.035.082-.058.129-.058h5.328c.047 0 .07.035.058.07l-.093.28c-.012.047-.058.082-.105.082H.047zM2.828 12.381c-.047 0-.059-.035-.035-.07l.163-.292c.023-.035.070-.070.117-.070h2.337c.047 0 .070.035.070.082l-.023.245c0 .047-.035.093-.082.093H2.828z"/>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">Go Projects</h3>
            <p class="text-gray-600 mb-4">
              Complete Go applications with Clean Architecture, Gin framework, and 17 utility packages for enterprise development.
            </p>
            <ul class="space-y-2 text-sm text-gray-600">
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Clean Architecture Pattern
              </li>
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                JWT Authentication
              </li>
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Database Integration
              </li>
            </ul>
          </div>

          <!-- PHP Projects -->
          <div class="feature-card">
            <div class="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mb-4">
              <svg class="w-8 h-8 text-purple-600" viewBox="0 0 24 24" fill="currentColor">
                <path d="M7.01 10.207h-.944l-.515 2.648h.838c.556 0 .982-.122 1.292-.391.291-.291.464-.683.464-1.212 0-.33-.097-.581-.304-.753-.207-.17-.536-.292-.831-.292z"/>
              </svg>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">PHP CodeIgniter</h3>
            <p class="text-gray-600 mb-4">
              Full-featured CodeIgniter applications with MVC architecture, security features, and admin panels.
            </p>
            <ul class="space-y-2 text-sm text-gray-600">
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                MVC Architecture
              </li>
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Security Headers
              </li>
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Admin Dashboard
              </li>
            </ul>
          </div>

          <!-- AI Assistant -->
          <div class="feature-card">
            <div class="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-4">
              <ChatBubbleLeftRightIcon class="w-8 h-8 text-green-600" />
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">AI Assistant</h3>
            <p class="text-gray-600 mb-4">
              Get intelligent recommendations for frameworks, databases, and project configurations through our AI chat.
            </p>
            <ul class="space-y-2 text-sm text-gray-600">
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Smart Recommendations
              </li>
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Best Practice Guidance
              </li>
              <li class="flex items-center">
                <CheckIcon class="w-4 h-4 text-green-500 mr-2" />
                Real-time Help
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- How It Works -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
            How It Works
          </h2>
          <p class="text-lg text-gray-600">
            Generate your project in three simple steps
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="text-center">
            <div class="w-16 h-16 bg-primary-100 rounded-full flex items-center justify-center mx-auto mb-6">
              <span class="text-2xl font-bold text-primary-600">1</span>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">Choose Language</h3>
            <p class="text-gray-600">
              Select between Go or PHP CodeIgniter based on your project requirements.
            </p>
          </div>

          <div class="text-center">
            <div class="w-16 h-16 bg-primary-100 rounded-full flex items-center justify-center mx-auto mb-6">
              <span class="text-2xl font-bold text-primary-600">2</span>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">Configure with AI</h3>
            <p class="text-gray-600">
              Use our AI assistant to get personalized recommendations for your project configuration.
            </p>
          </div>

          <div class="text-center">
            <div class="w-16 h-16 bg-primary-100 rounded-full flex items-center justify-center mx-auto mb-6">
              <span class="text-2xl font-bold text-primary-600">3</span>
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">Download & Deploy</h3>
            <p class="text-gray-600">
              Get your complete project as a ZIP file, ready to run and deploy.
            </p>
          </div>
        </div>

        <div class="text-center mt-12">
          <router-link
            to="/generator"
            class="btn btn-primary text-lg px-8 py-3"
          >
            Try It Now
          </router-link>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="py-20 bg-gradient-to-r from-primary-600 to-purple-600">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">
          Ready to Build Something Amazing?
        </h2>
        <p class="text-xl text-primary-100 mb-8">
          Start generating professional boilerplate projects with AI assistance today.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <router-link
            to="/generator"
            class="btn bg-white text-primary-600 hover:bg-gray-100 text-lg px-8 py-3"
          >
            <RocketLaunchIcon class="w-5 h-5 mr-2" />
            Generate Now
          </router-link>
          <router-link
            to="/about"
            class="btn btn-outline border-white text-white hover:bg-white hover:text-primary-600 text-lg px-8 py-3"
          >
            Learn More
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import {
  RocketLaunchIcon,
  CodeBracketIcon,
  CheckIcon,
  ChatBubbleLeftRightIcon
} from '@heroicons/vue/24/outline'
</script>

<style scoped>
.feature-card {
  @apply bg-white p-8 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow duration-200;
}

.feature-card:hover {
  transform: translateY(-2px);
  transition: transform 0.2s ease-out;
}
</style>