# Codebase Structure

## Project Root
```
boilerplate-blueprint/
├── .claude/                    # Claude Code configuration
├── .cursor/                    # Cursor IDE configuration  
├── .github/                    # GitHub workflows
├── .serena/                    # Serena MCP memory storage
├── .specify/                   # Specification files
├── cmd/                        # Go application entrypoints
├── internal/                   # Private Go packages
├── web/                        # Vue.js frontend application
├── go.mod, go.sum             # Go module files
├── Makefile                   # Development commands
├── README.md                  # Basic project info
└── project-description.md     # Detailed requirements
```

## Backend Structure (Go)
```
cmd/
└── main.go                    # Application entry point

internal/
├── api/
│   ├── handlers.go           # HTTP request handlers
│   └── routes.go             # Route definitions
├── models/
│   ├── chat.go              # Chat-related data structures
│   └── project.go           # Project data structures  
└── services/
    ├── chat.go              # AI chat business logic
    ├── project.go           # Project generation logic
    └── template.go          # Template management logic
```

## Frontend Structure (Vue.js)
```
web/
├── src/                      # Vue.js source code
├── dist/                     # Built files (generated)
├── index.html               # Main HTML template
├── package.json             # NPM dependencies
├── vite.config.js           # Vite build configuration
├── tailwind.config.js       # TailwindCSS configuration
└── postcss.config.js        # PostCSS configuration
```

## Key Architecture Patterns

### Backend (Go)
- **Clean Architecture**: Separation of concerns with API → Services → Models
- **Dependency Injection**: Services injected into handlers
- **Interface-Based Design**: For testability and modularity
- **Static File Serving**: Serves Vue.js build from `/web/dist`

### Frontend (Vue.js)  
- **Single Page Application**: Vue Router for client-side routing
- **Component-Based**: Reusable Vue components
- **State Management**: Pinia for centralized state
- **Build Tool**: Vite for fast development and optimized builds

## Data Flow
1. **User Interaction**: Vue.js frontend receives user input
2. **API Calls**: Axios sends HTTP requests to Go backend
3. **Request Handling**: Gin routes to appropriate handlers
4. **Business Logic**: Services process the request
5. **Response**: JSON data returned to frontend
6. **UI Update**: Vue reactivity updates the interface