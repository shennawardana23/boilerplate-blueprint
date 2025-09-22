# Boilerplate Blueprint - Project Onboarding Guide

## 🎯 Project Overview

**Boilerplate Blueprint** is a web-based AI-powered boilerplate generator application that creates Go and PHP CodeIgniter projects. It provides an intuitive interface for developers to generate complete project structures with AI assistance through a chat interface.

### Key Features
- **Language Selection**: Support for Go and PHP CodeIgniter projects
- **AI Chat Integration**: Intelligent chat interface for project customization
- **Dynamic Project Generation**: Real-time folder and file creation
- **Project Preview**: Live preview of generated project structure
- **ZIP Download**: Complete project export functionality
- **Template System**: Pre-built project templates with customizable options

## 🏗️ Architecture Overview

### Technology Stack

#### Backend (Go)
- **Framework**: Gin (HTTP web framework)
- **Language**: Go 1.21+
- **Architecture**: Clean Architecture with dependency injection
- **Storage**: In-memory storage (no database required for prototype)
- **File Generation**: Go templates + archive/zip for ZIP creation
- **CORS**: Cross-origin resource sharing enabled

#### Frontend (Vue.js)
- **Framework**: Vue.js 3 with Composition API
- **Build Tool**: Vite
- **State Management**: Pinia
- **Routing**: Vue Router
- **Styling**: Tailwind CSS
- **UI Components**: Headless UI + Heroicons
- **HTTP Client**: Axios

### Project Structure

```
boilerplate-blueprint/
├── cmd/                    # Application entry point
│   └── main.go            # Main server file
├── internal/              # Internal application code
│   ├── api/               # HTTP handlers and routes
│   │   ├── handlers.go    # Request handlers
│   │   └── routes.go      # Route definitions
│   ├── models/            # Data models
│   │   ├── chat.go        # Chat-related models
│   │   └── project.go     # Project-related models
│   └── services/          # Business logic layer
│       ├── chat.go        # Chat service with AI integration
│       ├── project.go     # Project management service
│       └── template.go    # Template generation service
├── web/                   # Frontend application
│   ├── src/
│   │   ├── components/    # Vue components
│   │   ├── stores/        # Pinia stores
│   │   ├── services/      # API services
│   │   ├── router/        # Vue Router configuration
│   │   └── views/         # Page components
│   └── package.json       # Frontend dependencies
├── tests/                 # Test files
│   ├── api/               # API handler tests
│   ├── services/          # Service layer tests
│   └── models/            # Model tests
├── Makefile              # Build and development commands
├── go.mod                # Go dependencies
└── README.md             # Project documentation
```

## 🔧 Core Components

### Backend Services

#### 1. ProjectService (`internal/services/project.go`)
**Purpose**: Manages the complete project lifecycle

**Key Methods**:
- `CreateProject(req *ProjectRequest)`: Creates new projects with validation
- `GetProject(projectID string)`: Retrieves project by ID
- `GenerateProjectFiles(project *Project)`: Generates project files based on template
- `CreateProjectZIP(projectID string)`: Creates downloadable ZIP archives
- `ListProjects()`: Returns all created projects

**Features**:
- Thread-safe operations with mutex locks
- Automatic default option setting based on language
- Support for Go and PHP project generation
- ZIP file creation with proper directory structure

#### 2. ChatService (`internal/services/chat.go`)
**Purpose**: Handles AI chat interactions and suggestions

**Key Methods**:
- `ProcessMessage(req *ChatRequest)`: Processes user messages and generates AI responses
- `GetChatHistory(projectID string)`: Retrieves conversation history
- `generateRuleBasedResponse()`: Creates intelligent responses based on message content

**Features**:
- Rule-based AI responses (OpenAI integration ready)
- Language and framework detection
- Project suggestion generation
- Conversation history management
- Context-aware responses

#### 3. TemplateService (`internal/services/template.go`)
**Purpose**: Generates project files and structures

**Key Methods**:
- `GenerateGoProject(project *Project)`: Creates complete Go project structure
- `GeneratePHPProject(project *Project)`: Creates complete PHP CodeIgniter structure
- `CreateZIPArchive(project *Project)`: Generates downloadable ZIP files
- `GetAvailableTemplates()`: Returns available project templates

**Features**:
- Dynamic directory structure creation
- Template-based file generation
- Support for 17 Go utility packages
- Complete PHP MVC structure
- ZIP archive creation

### Frontend Components

#### 1. Project Store (`web/src/stores/project.js`)
**Purpose**: Manages project state and operations

**Key Actions**:
- `loadTemplates()`: Loads available project templates
- `createProject(projectData)`: Creates new projects
- `generateProjectFiles(projectId)`: Generates project files
- `downloadProject(projectId)`: Downloads project as ZIP

**State**:
- `currentProject`: Currently selected project
- `projects`: List of all projects
- `templates`: Available project templates
- `isLoading`: Loading state
- `error`: Error messages

#### 2. Chat Store (`web/src/stores/chat.js`)
**Purpose**: Manages chat interactions and AI responses

**Key Actions**:
- `sendMessage(content, projectId)`: Sends messages to AI
- `loadChatHistory(projectId)`: Loads conversation history
- `applySuggestion(suggestion)`: Applies AI suggestions
- `clearMessages()`: Clears conversation

**State**:
- `messages`: Chat message history
- `suggestions`: AI-generated suggestions
- `isLoading`: Loading state
- `error`: Error messages

#### 3. API Service (`web/src/services/api.js`)
**Purpose**: Handles HTTP communication with backend

**Endpoints**:
- `projectApi`: Project management endpoints
- `chatApi`: Chat and AI endpoints
- `healthApi`: Health check endpoint

**Features**:
- Axios-based HTTP client
- Request/response interceptors
- Error handling
- CORS support

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- Node.js 18+ and npm
- Git

### Backend Setup

1. **Clone and navigate to project**:
   ```bash
   git clone <repository-url>
   cd boilerplate-blueprint
   ```

2. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the server**:
   ```bash
   make run
   # or
   go run cmd/main.go
   ```

4. **Server will start on**: `http://localhost:8080`

### Frontend Setup

1. **Navigate to web directory**:
   ```bash
   cd web
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Start development server**:
   ```bash
   npm run dev
   ```

4. **Frontend will be available on**: `http://localhost:5173`

### Running Tests

#### Backend Tests
```bash
# Run all tests
make test

# Run specific test packages
go test ./internal/services/... -v
go test ./internal/api/... -v
```

#### Frontend Tests
```bash
cd web
npm test
```

## 📋 API Endpoints

### Project Management
- `GET /api/templates` - Get available project templates
- `POST /api/projects` - Create new project
- `GET /api/projects/:id` - Get project by ID
- `POST /api/projects/:id/generate` - Generate project files
- `GET /api/projects/:id/download` - Download project as ZIP

### Chat & AI
- `POST /api/chat/message` - Send chat message
- `GET /api/chat/history` - Get chat history

### Health
- `GET /api/health` - Health check endpoint

## 🎨 Project Templates

### Go Project Template
**Features**:
- Clean Architecture structure
- Gin HTTP framework
- PostgreSQL database support
- JWT authentication
- 17 utility packages:
  - authentication, cache, common, constants
  - converter, date, datatype, encryption
  - exception, exceptioncode, helper
  - httphelper, json, logger, password
  - queryhelper, sort, template, validator, alert

**Structure**:
```
project-name/
├── cmd/main.go
├── internal/
│   ├── app/database/
│   ├── app/middleware/
│   ├── controller/
│   ├── service/
│   ├── repository/
│   ├── entity/
│   ├── model/api/
│   ├── converter/
│   ├── routes/
│   └── util/
├── scripts/
├── tests/
└── api/
```

### PHP CodeIgniter Template
**Features**:
- MVC architecture
- CodeIgniter 3 framework
- MySQL/PostgreSQL support
- Bootstrap frontend
- Security features
- Authentication system

**Structure**:
```
project-name/
├── application/
│   ├── config/
│   ├── controllers/
│   ├── models/
│   ├── views/
│   ├── helpers/
│   ├── libraries/
│   └── core/
├── assets/
├── system/
└── vendor/
```

## 🔒 Security Considerations

### Backend Security
- CORS configuration for cross-origin requests
- Input validation on all endpoints
- Error handling without sensitive information exposure
- Thread-safe operations with mutex locks

### Frontend Security
- Input sanitization
- XSS protection through Vue.js
- Secure HTTP client configuration
- Error boundary handling

## 🧪 Testing Strategy

### Backend Testing
- **Unit Tests**: Comprehensive coverage for all services and handlers
- **Integration Tests**: API endpoint testing
- **Concurrency Tests**: Thread safety validation
- **Error Handling**: Edge case and error scenario testing

### Frontend Testing
- **Unit Tests**: Store and service testing with Mocha
- **Component Tests**: Vue component testing with Vue Test Utils
- **Integration Tests**: API integration testing
- **E2E Tests**: Full user workflow testing (Cypress ready)

## 🚀 Deployment

### Development
```bash
# Backend
make run

# Frontend
cd web && npm run dev
```

### Production
```bash
# Build frontend
cd web && npm run build

# Build backend
make build

# Run production server
./boilerplate-blueprint
```

### Docker (Future)
```bash
# Build Docker image
make docker-build

# Run container
make docker-run
```

## 🔧 Development Commands

### Backend (Makefile)
- `make run` - Start development server
- `make test` - Run all tests
- `make build` - Build binary
- `make clean` - Clean build artifacts
- `make fmt` - Format code
- `make lint` - Run linting
- `make vet` - Run go vet

### Frontend (npm scripts)
- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm test` - Run tests
- `npm run lint` - Run ESLint

## 📚 Key Dependencies

### Backend (go.mod)
- `github.com/gin-gonic/gin` - HTTP web framework
- `github.com/google/uuid` - UUID generation
- `github.com/joho/godotenv` - Environment variable loading
- `github.com/gin-contrib/cors` - CORS middleware
- `github.com/stretchr/testify` - Testing framework

### Frontend (package.json)
- `vue` - Frontend framework
- `vue-router` - Client-side routing
- `pinia` - State management
- `axios` - HTTP client
- `tailwindcss` - CSS framework
- `@headlessui/vue` - UI components
- `@heroicons/vue` - Icons

## 🎯 Future Enhancements

### Planned Features
1. **OpenAI Integration**: Replace rule-based AI with OpenAI API
2. **Database Persistence**: Add PostgreSQL for project storage
3. **User Authentication**: Implement user accounts and project ownership
4. **Project Sharing**: Allow users to share and collaborate on projects
5. **Advanced Templates**: More project types and configurations
6. **Real-time Updates**: WebSocket integration for live collaboration
7. **Docker Support**: Containerized deployment
8. **CI/CD Pipeline**: Automated testing and deployment

### Technical Improvements
1. **Performance Optimization**: Caching and optimization
2. **Monitoring**: Application performance monitoring
3. **Logging**: Structured logging with levels
4. **Security**: Enhanced security features
5. **Documentation**: API documentation with Swagger

## 🤝 Contributing

### Development Workflow
1. Fork the repository
2. Create a feature branch
3. Make changes with tests
4. Run all tests
5. Submit a pull request

### Code Standards
- **Go**: Follow Go conventions and use `gofmt`
- **JavaScript**: Follow ESLint rules and use Prettier
- **Testing**: Maintain high test coverage
- **Documentation**: Update documentation for new features

## 📞 Support

For questions, issues, or contributions:
- Create an issue in the repository
- Check existing documentation
- Review the codebase structure
- Follow the development guidelines

---

**Boilerplate Blueprint** - Empowering developers with AI-assisted project generation 🚀