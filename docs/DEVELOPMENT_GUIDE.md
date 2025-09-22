# Boilerplate Blueprint - Development Guide

## 🚀 Quick Start

### Prerequisites
- **Go 1.21+**: [Download](https://golang.org/dl/)
- **Node.js 18+**: [Download](https://nodejs.org/)
- **Git**: [Download](https://git-scm.com/)

### Setup
```bash
# Clone repository
git clone <repository-url>
cd boilerplate-blueprint

# Backend setup
go mod tidy

# Frontend setup
cd web
npm install
cd ..

# Start development
make run          # Backend (port 8080)
cd web && npm run dev  # Frontend (port 5173)
```

## 🏗️ Architecture Deep Dive

### Backend Architecture

#### Clean Architecture Implementation
The backend follows Clean Architecture principles with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                        HTTP Layer                           │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐ │
│  │   Handlers      │  │     Routes      │  │ Middleware  │ │
│  │                 │  │                 │  │             │ │
│  └─────────────────┘  └─────────────────┘  └─────────────┘ │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                     Business Layer                          │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐ │
│  │ ProjectService  │  │  ChatService    │  │TemplateSvc  │ │
│  │                 │  │                 │  │             │ │
│  └─────────────────┘  └─────────────────┘  └─────────────┘ │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                      Data Layer                             │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐ │
│  │    Models       │  │   Templates     │  │   Storage   │ │
│  │                 │  │                 │  │             │ │
│  └─────────────────┘  └─────────────────┘  └─────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

#### Service Layer Details

**ProjectService** (`internal/services/project.go`)
- **Thread Safety**: Uses `sync.RWMutex` for concurrent access
- **Validation**: Validates language and project options
- **Default Options**: Automatically sets sensible defaults
- **File Generation**: Orchestrates template generation
- **ZIP Creation**: Handles archive creation and cleanup

**ChatService** (`internal/services/chat.go`)
- **Rule-Based AI**: Intelligent response generation
- **Context Awareness**: Maintains conversation context
- **Suggestion Engine**: Generates project recommendations
- **History Management**: Stores and retrieves chat history

**TemplateService** (`internal/services/template.go`)
- **Dynamic Generation**: Creates files based on project specs
- **Template System**: Uses Go's `text/template` package
- **Directory Structure**: Generates complete project hierarchies
- **ZIP Archives**: Creates downloadable project packages

### Frontend Architecture

#### Vue.js 3 Composition API
The frontend uses modern Vue.js patterns with Composition API:

```javascript
// Store Pattern (Pinia)
export const useProjectStore = defineStore('project', () => {
  // Reactive state
  const currentProject = ref(null)
  const isLoading = ref(false)
  
  // Computed properties
  const hasCurrentProject = computed(() => currentProject.value !== null)
  
  // Actions
  async function createProject(projectData) {
    // Implementation
  }
  
  return {
    currentProject,
    isLoading,
    hasCurrentProject,
    createProject
  }
})
```

#### Component Structure
```
web/src/components/
├── layout/           # Layout components
│   ├── AppHeader.vue
│   └── AppFooter.vue
├── project/          # Project-related components
│   ├── ProjectForm.vue
│   ├── ProjectPreview.vue
│   └── ProjectDownload.vue
├── chat/             # Chat components
│   ├── ChatInterface.vue
│   └── MessageList.vue
└── ui/               # Reusable UI components
    └── LoadingSpinner.vue
```

## 🔧 Development Workflow

### Backend Development

#### Adding New Endpoints
1. **Define Route** (`internal/api/routes.go`):
```go
api.POST("/projects/:id/custom", handlers.CustomAction)
```

2. **Implement Handler** (`internal/api/handlers.go`):
```go
func (h *Handlers) CustomAction(c *gin.Context) {
    // Implementation
}
```

3. **Add Service Method** (`internal/services/project.go`):
```go
func (s *ProjectService) CustomMethod(projectID string) error {
    // Implementation
}
```

4. **Write Tests** (`tests/api/handlers_test.go`):
```go
func TestHandlers_CustomAction(t *testing.T) {
    // Test implementation
}
```

#### Adding New Project Templates
1. **Update Template Service** (`internal/services/template.go`):
```go
func (s *TemplateService) GenerateCustomProject(project *models.Project) ([]models.ProjectFile, error) {
    // Template generation logic
}
```

2. **Add Template Info**:
```go
func (s *TemplateService) GetAvailableTemplates() []models.TemplateInfo {
    return []models.TemplateInfo{
        // Existing templates...
        {
            Language:    models.LanguageCustom,
            Name:        "Custom Template",
            Description: "Description of custom template",
            Options:     []models.TemplateOption{...},
        },
    }
}
```

3. **Update Models** (`internal/models/project.go`):
```go
const (
    LanguageGo     ProjectLanguage = "go"
    LanguagePHP   ProjectLanguage = "php"
    LanguageCustom ProjectLanguage = "custom"  // New language
)
```

### Frontend Development

#### Adding New Components
1. **Create Component** (`web/src/components/NewComponent.vue`):
```vue
<template>
  <div class="new-component">
    <!-- Template -->
  </div>
</template>

<script setup>
// Component logic
</script>

<style scoped>
/* Component styles */
</style>
```

2. **Add to Router** (`web/src/router/index.js`):
```javascript
{
  path: '/new-route',
  name: 'NewRoute',
  component: () => import('../views/NewView.vue')
}
```

3. **Create Store Action** (if needed):
```javascript
// In appropriate store
async function newAction(data) {
  try {
    isLoading.value = true
    // Implementation
  } catch (error) {
    error.value = error.message
  } finally {
    isLoading.value = false
  }
}
```

#### Adding New API Endpoints
1. **Update API Service** (`web/src/services/api.js`):
```javascript
export const customApi = {
  customEndpoint(data) {
    return api.post('/custom-endpoint', data)
  }
}
```

2. **Use in Store**:
```javascript
import { customApi } from '@/services/api'

async function useCustomEndpoint(data) {
  const response = await customApi.customEndpoint(data)
  // Handle response
}
```

## 🧪 Testing Strategy

### Backend Testing

#### Unit Tests
```go
func TestProjectService_CreateProject(t *testing.T) {
    // Arrange
    service := services.NewProjectService(templateService)
    req := &models.ProjectRequest{
        Name:     "test-project",
        Language: models.LanguageGo,
    }
    
    // Act
    project, err := service.CreateProject(req)
    
    // Assert
    require.NoError(t, err)
    assert.NotNil(t, project)
    assert.Equal(t, req.Name, project.Name)
}
```

#### Integration Tests
```go
func TestHandlers_CreateProject_Integration(t *testing.T) {
    // Setup
    handlers := setupTestHandlers()
    router := gin.New()
    router.POST("/projects", handlers.CreateProject)
    
    // Test
    req := httptest.NewRequest("POST", "/projects", jsonData)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
}
```

#### Concurrency Tests
```go
func TestProjectService_ConcurrentAccess(t *testing.T) {
    service := services.NewProjectService(templateService)
    
    // Test concurrent access
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            // Test concurrent operations
        }(i)
    }
    wg.Wait()
}
```

### Frontend Testing

#### Store Tests
```javascript
describe('Project Store', () => {
  it('should create project successfully', async () => {
    const projectData = {
      name: 'test-project',
      language: 'go'
    }
    
    const result = await store.createProject(projectData)
    
    expect(result.name).to.equal('test-project')
    expect(store.currentProject).to.deep.equal(result)
  })
})
```

#### Component Tests
```javascript
describe('ProjectForm Component', () => {
  it('should render form fields', () => {
    const wrapper = mount(ProjectForm)
    
    expect(wrapper.find('input[name="name"]')).to.exist
    expect(wrapper.find('select[name="language"]')).to.exist
  })
})
```

## 🚀 Build & Deployment

### Development Build
```bash
# Backend
go build -o boilerplate-blueprint cmd/main.go

# Frontend
cd web && npm run build
```

### Production Build
```bash
# Backend (optimized)
CGO_ENABLED=0 go build -ldflags="-s -w" -o boilerplate-blueprint cmd/main.go

# Frontend (optimized)
cd web && npm run build
```

### Docker Build (Future)
```dockerfile
# Multi-stage build
FROM golang:1.21-alpine AS backend
WORKDIR /app
COPY . .
RUN go build -o boilerplate-blueprint cmd/main.go

FROM node:18-alpine AS frontend
WORKDIR /app/web
COPY web/package*.json ./
RUN npm ci --only=production
COPY web/ .
RUN npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=backend /app/boilerplate-blueprint .
COPY --from=frontend /app/web/dist ./web/dist
CMD ["./boilerplate-blueprint"]
```

## 🔍 Debugging

### Backend Debugging
```bash
# Run with debug logging
GIN_MODE=debug go run cmd/main.go

# Use delve debugger
dlv debug cmd/main.go

# Add debug prints
log.Printf("Debug: %+v", data)
```

### Frontend Debugging
```bash
# Run with source maps
npm run dev

# Use Vue DevTools
# Install browser extension for Vue DevTools

# Add debug logs
console.log('Debug:', data)
```

## 📊 Performance Optimization

### Backend Optimization
- **Connection Pooling**: Use connection pools for external services
- **Caching**: Implement Redis caching for frequently accessed data
- **Compression**: Enable gzip compression for responses
- **Profiling**: Use `go tool pprof` for performance analysis

### Frontend Optimization
- **Code Splitting**: Implement route-based code splitting
- **Lazy Loading**: Load components on demand
- **Bundle Analysis**: Use `npm run build -- --analyze`
- **Image Optimization**: Optimize images and use WebP format

## 🔒 Security Best Practices

### Backend Security
- **Input Validation**: Validate all input data
- **SQL Injection**: Use parameterized queries
- **XSS Protection**: Sanitize user input
- **CSRF Protection**: Implement CSRF tokens
- **Rate Limiting**: Implement rate limiting for APIs

### Frontend Security
- **XSS Prevention**: Use Vue's built-in XSS protection
- **Content Security Policy**: Implement CSP headers
- **HTTPS**: Use HTTPS in production
- **Dependency Scanning**: Regularly scan for vulnerabilities

## 📈 Monitoring & Logging

### Structured Logging
```go
import "github.com/sirupsen/logrus"

log := logrus.WithFields(logrus.Fields{
    "service": "boilerplate-blueprint",
    "version": "1.0.0",
    "request_id": requestID,
})

log.Info("Project created successfully")
log.Error("Failed to create project", err)
```

### Metrics Collection
```go
// Add metrics endpoints
router.GET("/metrics", gin.WrapH(promhttp.Handler()))
```

## 🚀 Future Enhancements

### Planned Features
1. **Database Integration**: Add PostgreSQL for persistence
2. **User Authentication**: Implement JWT-based auth
3. **Real-time Updates**: WebSocket integration
4. **Advanced Templates**: More project types
5. **CI/CD Pipeline**: Automated testing and deployment
6. **Monitoring**: Application performance monitoring
7. **Caching**: Redis integration
8. **API Documentation**: Swagger/OpenAPI docs

### Technical Debt
- [ ] Add comprehensive error handling
- [ ] Implement proper logging
- [ ] Add database migrations
- [ ] Implement caching layer
- [ ] Add monitoring and metrics
- [ ] Improve test coverage
- [ ] Add API documentation
- [ ] Implement proper security measures

## 🤝 Contributing

### Development Process
1. **Fork Repository**: Create your fork
2. **Create Branch**: `git checkout -b feature/new-feature`
3. **Make Changes**: Implement your feature
4. **Write Tests**: Add comprehensive tests
5. **Run Tests**: Ensure all tests pass
6. **Submit PR**: Create pull request with description

### Code Standards
- **Go**: Follow `gofmt` and `golint` standards
- **JavaScript**: Follow ESLint and Prettier rules
- **Commits**: Use conventional commit messages
- **Documentation**: Update docs for new features

### Review Process
- All PRs require review
- Tests must pass
- Code must follow standards
- Documentation must be updated

---

**Happy Coding!** 🚀

For questions or help, please:
- Check existing documentation
- Create an issue in the repository
- Review the codebase structure
- Follow the development guidelines