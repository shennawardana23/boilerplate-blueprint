# Technology Stack

## Backend (Go)
- **Framework**: Gin (github.com/gin-gonic/gin v1.9.1)
- **Language**: Go 1.21+
- **Key Dependencies**:
  - CORS: github.com/gin-contrib/cors v1.4.0
  - UUID generation: github.com/google/uuid v1.6.0
  - Environment: github.com/joho/godotenv v1.5.1

## Frontend (Vue.js)
- **Framework**: Vue 3.3.11 with Composition API
- **Build Tool**: Vite 5.0.8
- **Styling**: TailwindCSS 3.3.6
- **State Management**: Pinia 2.1.7
- **Routing**: Vue Router 4.2.5
- **HTTP Client**: Axios 1.6.2
- **UI Components**: 
  - Headless UI (@headlessui/vue 1.7.16)
  - Heroicons (@heroicons/vue 2.0.18)

## Development Tools
- **Linting**: ESLint with Vue plugin
- **CSS Processing**: PostCSS with Autoprefixer
- **Type System**: JavaScript (no TypeScript configured)

## Project Structure
```
boilerplate-blueprint/
├── cmd/main.go                 # Go application entry point
├── internal/
│   ├── api/                   # REST API handlers and routes
│   ├── services/              # Business logic (project, template, chat)
│   └── models/                # Data structures
├── web/                       # Vue.js frontend
│   ├── src/                   # Vue components and logic
│   ├── package.json           # NPM dependencies
│   └── vite.config.js         # Build configuration
├── go.mod                     # Go module definition
└── Makefile                   # Development commands
```