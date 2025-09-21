# This is the core .md file of the project

<idea>
Build a simple web-based Create Boilerplate folder,file for Go language and PHP Codeigniter powered by AI. That allows users to spin up a Go project with the corresponding structure seamlessly. It also gives the option to integrate with one of the more popular Go frameworks reference fo Go is https://go-blueprint.dev/ and make it there is the dropdwon for Go and PHP first on page then show each page. And AI should respond with result or update the Data folder structure, file and etc. That generates a **web-based boilerplate generator application** similar to [Go Blueprint](https://go-blueprint.dev/). This application allows users to create Go and PHP CodeIgniter projects with AI assistance through a chat interface, seamlessly generating project structures, files, and configurations.
</idea>

## 🎯 **Project Overview**

You are an AI assistant that generates a **web-based boilerplate generator application** similar to [Go Blueprint](https://go-blueprint.dev/). This application allows users to create Go and PHP CodeIgniter projects with AI assistance through a chat interface, seamlessly generating project structures, files, and configurations.

## 🚀 **Core Generation Prompts**

### **1. Web Application Foundation Prompt**

```
Create a web-based boilerplate generator application named `{PROJECT_NAME}` with the following core features:

**Core Requirements:**
- **Language Selection**: Dropdown for Go and PHP CodeIgniter
- **AI Chat Integration**: Chat interface for project customization
- **Project Generation**: Dynamic folder and file creation
- **Real-time Updates**: Live preview of project structure
- **Download Functionality**: ZIP file generation for completed projects

**Technology Stack:**
- **Frontend**: React/Next.js with TypeScript
- **Backend**: Go with Gin Framework (following `prompt-go-project.md` structure)
- **AI Integration**: OpenAI API or similar for chat responses
- **Database**: PostgreSQL for storing project templates and user sessions
- **File System**: Dynamic file generation and ZIP creation

**Expected Output:**
- Complete web application with AI chat interface
- Project structure generator for Go and PHP
- Real-time project preview
- ZIP download functionality
- **DOWNLOADABLE ZIP FILE** with the complete application
```

### **2. Go Project Generator Integration Prompt**

```
Integrate Go project generation into the `{PROJECT_NAME}` boilerplate generator:

**Go Project Features:**
- **Framework Selection**: Gin, Chi, Echo, or Standard Library
- **Database Options**: PostgreSQL, MySQL, SQLite, or MongoDB
- **Project Structure**: Follow `prompt-go-project.md` Clean Architecture exactly:

**Required Directory Structure:**
```

{PROJECT_NAME}/
├── .git/
├── .gitignore
├── .air.toml
├── .env
├── .env.example
├── Dockerfile
├── Dockerfile_golang
├── Makefile
├── README.md
├── buildspec.yml
├── go.mod
├── go.sum
├── main
├── cmd/
│   ├── main.go
│   ├── config/
│   │   ├── config.go
│   │   └── environment.go
│   └── docs/
│       └── docs.go
├── internal/
│   ├── app/
│   │   ├── database/
│   │   │   ├── database.go
│   │   │   ├── postgres.go
│   │   │   └── redis.go
│   │   ├── middleware/
│   │   │   ├── http_middleware.go
│   │   │   ├── http_jwt_auth.go
│   │   │   ├── http_basic_auth.go
│   │   │   ├── http_whitelist.go
│   │   │   ├── security_headers.go
│   │   │   ├── cors.go
│   │   │   ├── logger.go
│   │   │   ├── session.go
│   │   │   └── http_pre_handle.go
│   │   ├── sentry/
│   │   │   └── sentry.go
│   │   └── newrelic/
│   │       └── newrelic.go
│   ├── controller/
│   │   ├── base_controller.go
│   │   ├── user_controller.go
│   │   ├── project_controller.go
│   │   └── health_controller.go
│   ├── service/
│   │   ├── base_service.go
│   │   ├── user_service.go
│   │   ├── project_service.go
│   │   └── health_service.go
│   ├── repository/
│   │   ├── base_repository.go
│   │   ├── user_repository.go
│   │   ├── project_repository.go
│   │   └── mocks/
│   ├── entity/
│   │   ├── user.go
│   │   ├── project.go
│   │   └── base_entity.go
│   ├── model/
│   │   └── api/
│   │       ├── api_request.go
│   │       ├── api_response.go
│   │       ├── user_request.go
│   │       ├── user_response.go
│   │       ├── project_request.go
│   │       └── project_response.go
│   ├── converter/
│   │   ├── user_converter.go
│   │   └── project_converter.go
│   ├── routes/
│   │   └── router.go
│   ├── outbound/
│   │   ├── google_chat.go
│   │   └── email.go
│   ├── state/
│   │   └── state.go
│   ├── util/
│   └── docs/
├── scripts/
├── tests/
└── api/

```

**Required Utility Packages (ALL 17):**
```

internal/util/
├── authentication/     # JWT token management
├── cache/             # Redis caching operations
├── common/            # General utilities
├── constants/         # Application constants
├── converter/         # Data transformation
├── date/              # Date/time utilities
├── datatype/          # Custom data structures
├── encryption/        # Security utilities
├── exception/         # Error handling
├── exceptioncode/     # Error codes
├── helper/            # Helper functions
├── httphelper/        # HTTP utilities
├── json/              # JSON operations
├── logger/            # Structured logging
├── password/          # Password utilities
├── queryhelper/       # SQL query builder
├── sort/              # Sorting utilities
├── template/          # Template processing
├── validator/         # Input validation
└── alert/             # Notifications

```

**Required Dependencies (go.mod):**
```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
    github.com/redis/go-redis/v9 v9.3.0
    github.com/golang-jwt/jwt/v5 v5.2.0
    golang.org/x/crypto v0.40.0
    github.com/go-playground/validator/v10 v10.19.0
    github.com/sirupsen/logrus v1.9.3
    github.com/getsentry/sentry-go v0.25.0
    github.com/newrelic/go-agent/v3 v3.28.0
    github.com/aws/aws-lambda-go v1.41.0
    github.com/apex/gateway/v2 v2.0.0
    github.com/google/uuid v1.6.0
    github.com/joho/godotenv v1.5.1
    github.com/rs/cors v1.10.1
    github.com/stretchr/testify v1.10.0
    github.com/golang/mock v1.6.0
)
```

**Implementation Patterns:**

- **Dependency Injection**: Use constructor functions with interfaces
- **Clean Architecture**: Controllers → Services → Repositories → Entities
- **Error Handling**: Explicit error checking with wrapped errors
- **Context Usage**: Request-scoped values and cancellation
- **Interface Design**: Small, purpose-specific interfaces
- **Testing**: Table-driven tests with mocks

**Expected Output:**

- Go project generator with AI chat
- Dynamic project structure creation following exact patterns
- Real-time code generation with all 17 utility packages
- **DOWNLOADABLE ZIP FILE** with Go generator

```

### **3. PHP CodeIgniter Generator Integration Prompt**

```

Integrate PHP CodeIgniter project generation into the `{PROJECT_NAME}` boilerplate generator:

**CodeIgniter Features:**

- **Version Selection**: CodeIgniter 3 or 4
- **Database Options**: MySQL, PostgreSQL, SQLite
- **Project Structure**: MVC architecture with proper organization

**Required Directory Structure:**

```
PROJECT_NAME/
├── application/
│   ├── cache/
│   │   ├── index.html
│   │   └── sessions/
│   ├── config/
│   │   ├── autoload.php
│   │   ├── config.php
│   │   ├── constants.php
│   │   ├── database.php
│   │   ├── routes.php
│   │   └── index.html
│   ├── controllers/
│   │   ├── api/
│   │   ├── Command.php
│   │   ├── Dashboard.php
│   │   ├── Home.php
│   │   ├── Login.php
│   │   ├── User.php
│   │   └── index.html
│   ├── core/
│   │   ├── MY_Controller.php
│   │   └── index.html
│   ├── helpers/
│   │   ├── button_helper.php
│   │   ├── common_helper.php
│   │   ├── debug_helper.php
│   │   ├── memory_helper.php
│   │   ├── privilege_helper.php
│   │   ├── request_helper.php
│   │   ├── template_prototype_helper.php
│   │   ├── upload_helper.php
│   │   ├── uuid_helper.php
│   │   ├── validation_helper.php
│   │   └── index.html
│   ├── libraries/
│   │   ├── Awsconnector.php
│   │   ├── Breadcrumbs.php
│   │   ├── Connectrediscluster.php
│   │   ├── Template.php
│   │   ├── Usermanagement.php
│   │   └── index.html
│   ├── models/
│   │   ├── App_model.php
│   │   ├── Dashboard_model.php
│   │   ├── User_model.php
│   │   └── index.html
│   ├── views/
│   │   ├── dashboard.php
│   │   ├── login.php
│   │   ├── user/
│   │   ├── widgets/
│   │   └── index.html
│   └── widgets/
│       ├── footer.php
│       ├── header.php
│       ├── sidebar.php
│       └── index.html
├── assets/
│   ├── css/
│   ├── js/
│   ├── fonts/
│   ├── plugins/
│   └── scss/
├── system/
├── vendor/
├── composer.json
├── composer.lock
├── index.php
└── README.md
```

**Required Implementation Patterns:**

**Base Controller (MY_Controller.php):**

- Security headers (X-Frame-Options, X-Content-Type-Options, X-XSS-Protection)
- Session security enhancement
- Access validation
- Output formatting methods
- Notification generation

**Configuration Files:**

- Environment-based configuration
- Database connection management
- Session configuration
- Security settings
- Route definitions

**Helper Functions:**

- Common utilities (generate_copy_url_input, generate_image_preview)
- Button generation (generate_action_buttons)
- Form helpers
- Validation helpers
- Debug helpers

**Library Classes:**

- Template management system
- User management
- AWS connector
- Breadcrumb navigation
- Redis cluster connection

**Model Structure:**

- Base model with CRUD operations
- User management model
- Dashboard data model
- UUID generation
- Data validation

**Controller Implementation:**

- Authentication controllers (Login, Dashboard)
- User management controllers
- API controllers
- Command-line controllers
- Proper inheritance from MY_Controller

**View Templates:**

- Login forms with validation
- Dashboard layouts
- User management interfaces
- Widget components (header, sidebar, footer)
- Responsive design

**Composer Configuration:**

```json
{
    "require": {
        "php": ">=7.4",
        "predis/predis": "^1.1",
        "aws/aws-sdk-php": "^3.293",
        "ramsey/uuid": "^4.2",
        "vlucas/phpdotenv": "^5.6",
        "guzzlehttp/guzzle": "^7.0",
        "league/flysystem": "^2.0",
        "league/flysystem-aws-s3-v3": "^2.0",
        "monolog/monolog": "^2.0"
    }
}
```

**Security Features:**

- CSRF protection
- Input validation and sanitization
- Session security
- SQL injection prevention
- XSS protection
- Secure headers

**Expected Output:**

- PHP CodeIgniter generator with AI chat
- Dynamic project structure creation with MVC patterns
- Real-time code generation with security features
- **DOWNLOADABLE ZIP FILE** with PHP generator

```

### **4. AI Chat Interface Prompt**

```

Implement the AI chat interface for the `{PROJECT_NAME}` boilerplate generator:

**Chat Interface Requirements:**

- **Real-time Communication**: WebSocket or Server-Sent Events
- **Context Awareness**: Remember project requirements and structure
- **Code Generation**: AI-assisted file and folder creation
- **Project Preview**: Live updates to project structure
- **Error Handling**: AI suggestions for invalid configurations
- **History Management**: Chat conversation persistence

**AI Integration Features:**

- **Project Analysis**: Understand user requirements
- **Structure Suggestion**: Recommend optimal project organization
- **Code Generation**: Create boilerplate code based on requirements
- **Dependency Management**: Suggest appropriate packages/libraries
- **Best Practices**: Recommend coding standards and patterns

**AI Response Patterns:**

**For Go Projects:**

- Suggest appropriate utility packages based on requirements
- Recommend database and framework combinations
- Provide Clean Architecture guidance
- Suggest testing strategies
- Recommend security implementations

**For PHP CodeIgniter Projects:**

- Suggest MVC structure organization
- Recommend helper and library combinations
- Provide authentication patterns
- Suggest database optimization
- Recommend security implementations

**Expected Output:**

- Complete AI chat interface
- Real-time project generation
- Context-aware AI responses
- **DOWNLOADABLE ZIP FILE** with AI chat system

```

### **5. Project Structure Generator Prompt**

```

Implement the dynamic project structure generator for the `{PROJECT_NAME}` application:

**Generator Features:**

- **Real-time Preview**: Live project structure visualization
- **Dynamic Creation**: Add/remove folders and files
- **Template System**: Pre-built project templates
- **Customization**: User-defined project structures
- **Validation**: Ensure valid project configurations
- **Export Options**: Multiple output formats (ZIP, GitHub repo, etc.)

**Project Management:**

- **Template Library**: Go and PHP project templates
- **User Projects**: Save and manage user-generated projects
- **Version Control**: Track project structure changes
- **Collaboration**: Share projects with team members
- **Backup/Restore**: Project backup and recovery

**Template System Requirements:**

**Go Project Templates:**

- **Basic API**: Minimal Go API with Gin
- **Full Stack**: Complete web application with all utilities
- **Microservice**: Service-oriented architecture
- **CLI Tool**: Command-line application
- **Library**: Reusable Go package

**PHP CodeIgniter Templates:**

- **Basic MVC**: Simple MVC application
- **Admin Panel**: Full admin interface
- **API Backend**: RESTful API structure
- **E-commerce**: Shopping cart application
- **CMS**: Content management system

**Expected Output:**

- Dynamic project structure generator
- Real-time preview system
- Template management system
- **DOWNLOADABLE ZIP FILE** with generator

```

### **6. Frontend User Interface Prompt**

```

Create the frontend user interface for the `{PROJECT_NAME}` boilerplate generator:

**UI Requirements:**

- **Language Selection**: Clean dropdown for Go vs PHP
- **Project Configuration**: Form-based project setup
- **AI Chat Panel**: Integrated chat interface
- **Project Preview**: Visual project structure tree
- **Real-time Updates**: Live project structure changes
- **Download Section**: Project export and download

**Design Principles:**

- **Minimal & Clean**: Similar to Go Blueprint styling
- **Responsive Design**: Mobile and desktop compatibility
- **Dark/Light Mode**: Theme switching capability
- **Accessibility**: WCAG compliance and keyboard navigation
- **Performance**: Fast loading and smooth interactions

**UI Components:**

**Language Selection Panel:**

- Dropdown with Go and PHP CodeIgniter options
- Visual indicators for each language
- Quick start templates for each option

**Project Configuration Forms:**

**Go Project Form:**

- Project name and description
- Framework selection (Gin, Chi, Echo, Standard)
- Database selection (PostgreSQL, MySQL, SQLite, MongoDB)
- Authentication options (JWT, OAuth, Basic Auth)
- Deployment options (Lambda, Docker, Traditional)
- Utility package selection (all 17 packages)

**PHP CodeIgniter Form:**

- Project name and description
- CodeIgniter version (3 or 4)
- Database selection (MySQL, PostgreSQL, SQLite)
- Authentication system (Built-in, Custom)
- API structure (RESTful, GraphQL)
- Frontend framework (Bootstrap, Tailwind, Custom)

**AI Chat Interface:**

- Real-time chat window
- Project context display
- Code suggestion panel
- Error handling display
- Chat history management

**Project Structure Preview:**

- Interactive tree view
- File/folder icons
- Real-time updates
- Drag-and-drop organization
- Collapsible sections

**Expected Output:**

- Complete frontend application
- Responsive design system
- Theme management
- **DOWNLOADABLE ZIP FILE** with frontend

```

### **7. Backend API & File Generation Prompt**

```

Implement the backend API and file generation system for the `{PROJECT_NAME}` application:

**Backend Requirements:**

- **Project Generation API**: Create project structures and files
- **File System Operations**: Dynamic file creation and management
- **ZIP Generation**: Create downloadable project archives
- **AI Integration**: OpenAI API integration for chat responses
- **Template Management**: Store and retrieve project templates
- **User Management**: Project history and user sessions

**File Generation Features:**

- **Dynamic Templates**: Generate files based on user requirements
- **Code Formatting**: Proper indentation and syntax
- **Dependency Management**: Generate appropriate config files
- **Documentation**: Auto-generate README and documentation
- **Testing Setup**: Create test files and configurations

**API Endpoints:**

**Project Management:**

```
POST /api/projects/create          # Create new project
GET  /api/projects/{id}           # Get project details
PUT  /api/projects/{id}           # Update project
DELETE /api/projects/{id}         # Delete project
GET  /api/projects                # List user projects
POST /api/projects/{id}/generate  # Generate project files
POST /api/projects/{id}/download  # Create ZIP download
```

**Template Management:**

```
GET  /api/templates               # List available templates
GET  /api/templates/{id}          # Get template details
POST /api/templates               # Create custom template
PUT  /api/templates/{id}          # Update template
DELETE /api/templates/{id}        # Delete template
```

**AI Chat Integration:**

```
POST /api/chat/message            # Send chat message
GET  /api/chat/history/{id}      # Get chat history
POST /api/chat/generate           # Generate code from chat
```

**File Generation Engine:**

**Go Project Generation:**

- Generate `go.mod` with correct dependencies
- Create directory structure following Clean Architecture
- Generate all 17 utility packages with proper interfaces
- Create controllers, services, repositories with DI
- Generate middleware and authentication
- Create database configuration and migrations
- Generate Docker and deployment files
- Create comprehensive README and documentation

**PHP CodeIgniter Generation:**

- Generate `composer.json` with dependencies
- Create MVC directory structure
- Generate base controller with security features
- Create helper and library files
- Generate configuration files
- Create view templates
- Generate authentication system
- Create database models and migrations

**ZIP Generation System:**

- Create temporary project directory
- Generate all project files
- Validate file structure
- Create ZIP archive
- Clean up temporary files
- Provide download link

**Expected Output:**

- Complete backend API system
- File generation engine
- ZIP creation system
- **DOWNLOADABLE ZIP FILE** with backend

```

## 📦 **Auto-Download Instructions for AI Agents**

### **ZIP File Generation Requirements**

**IMPORTANT: You MUST generate a downloadable ZIP file for every project generation.**

**ZIP File Structure:**

```

{PROJECT_NAME}.zip
├── {PROJECT_NAME}/
│   ├── [all application files and directories]
│   └── [complete web application structure]
└── README-INSTALLATION.md

```

**ZIP File Contents:**

1. **Complete web application** with AI chat integration
2. **Go project generator** following `prompt-go-project.md`
3. **PHP CodeIgniter generator** with MVC structure
4. **AI chat interface** for project customization
5. **Project structure generator** with real-time preview
6. **Installation and setup instructions**

**ZIP File Naming Convention:**

- Format: `{PROJECT_NAME}-{FEATURES}-{DATE}.zip`
- Example: `boilerplate-generator-go-php-ai-2024-01-15.zip`

**Download Instructions for Users:**

```

🎉 Your AI-Powered Boilerplate Generator is Ready!

📦 **Download Instructions:**

1. Click the download link below
2. Extract the ZIP file to your desired location
3. Follow the README-INSTALLATION.md for setup
4. Run the application and start generating projects

🚀 **Quick Start:**
cd {PROJECT_NAME}
make dev
open <http://localhost:3000>

📚 **Features:**

- Go project generation with Clean Architecture
- PHP CodeIgniter project generation
- AI chat for project customization
- Real-time project preview
- ZIP download functionality

```

## 🎯 **Success Criteria for AI Agents**

### **1. Always Generate Downloadable ZIP Files**

- ✅ **REQUIRED**: Every project generation must include a ZIP file
- ✅ **REQUIRED**: ZIP file must contain complete working application
- ✅ **REQUIRED**: Include installation and setup instructions

### **2. Follow Go Project Structure Exactly**

- ✅ **REQUIRED**: Use exact directory structure from `prompt-go-project.md`
- ✅ **REQUIRED**: Implement all 17 utility packages
- ✅ **REQUIRED**: Follow code examples and patterns
- ✅ **REQUIRED**: Use dependency injection patterns

### **3. Follow PHP CodeIgniter Structure Exactly**

- ✅ **REQUIRED**: Use exact directory structure from `prompt-php-project.md`
- ✅ **REQUIRED**: Implement all helper and library files
- ✅ **REQUIRED**: Follow MVC patterns and security features
- ✅ **REQUIRED**: Include all configuration files

### **4. Ensure Web Application Functionality**

- ✅ **REQUIRED**: Language selection (Go vs PHP)
- ✅ **REQUIRED**: AI chat integration
- ✅ **REQUIRED**: Project structure generator
- ✅ **REQUIRED**: Real-time preview
- ✅ **REQUIRED**: ZIP download functionality

### **5. Production Readiness**

- ✅ **REQUIRED**: Proper error handling
- ✅ **REQUIRED**: Comprehensive logging
- ✅ **REQUIRED**: Security features
- ✅ **REQUIRED**: Testing coverage
- ✅ **REQUIRED**: Documentation

## 🚀 **Usage Instructions for AI Agents**

### **Step-by-Step Generation Process**

1. **Start with Web Application Foundation** (Prompt #1)
   - Generate basic web application structure
   - Set up AI chat integration
   - Create ZIP file with foundation

2. **Add Go Project Generator** (Prompt #2)
   - Implement Go project generation
   - Integrate with `prompt-go-project.md` structure
   - Create ZIP file with Go generator

3. **Add PHP CodeIgniter Generator** (Prompt #3)
   - Implement PHP project generation
   - Create MVC structure templates
   - Create ZIP file with PHP generator

4. **Implement AI Chat Interface** (Prompt #4)
   - Add real-time chat functionality
   - Integrate AI responses
   - Create ZIP file with chat system

5. **Add Project Structure Generator** (Prompt #5)
   - Implement dynamic structure creation
   - Add real-time preview
   - Create ZIP file with generator

6. **Create Frontend UI** (Prompt #6)
   - Build user interface
   - Add language selection
   - Create ZIP file with frontend

7. **Implement Backend API** (Prompt #7)
   - Add file generation API
   - Implement ZIP creation
   - Create final ZIP file

### **Quality Assurance Checklist**

Before generating ZIP file, ensure:

- [ ] All code compiles without errors
- [ ] Go projects follow `prompt-go-project.md` structure exactly
- [ ] PHP projects follow `prompt-php-project.md` structure exactly
- [ ] AI chat interface is functional
- [ ] Project structure generator works
- [ ] ZIP download functionality is implemented
- [ ] Real-time preview is working
- [ ] Language selection works properly
- [ ] Installation instructions are clear
- [ ] All 17 Go utility packages are implemented
- [ ] All PHP helper and library files are implemented
- [ ] Security features are properly implemented
- [ ] Testing setup is complete

## 📚 **Reference Resources**

- [Go Blueprint](https://go-blueprint.dev/) - Project structure inspiration
- [Go Blueprint GitHub](https://github.com/Melkeydev/go-blueprint) - Source code
- [CodeIgniter Documentation](https://codeigniter.com/) - PHP framework docs
- [OpenAI API](https://platform.openai.com/) - AI integration
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Architecture principles

---

**Note**: This prompt file is designed to generate a web-based AI-powered boilerplate generator that creates both Go and PHP CodeIgniter projects. The application should be similar to Go Blueprint but with AI chat integration for project customization. Every generation MUST include a downloadable ZIP file with the complete working application. Follow the exact structures and patterns from both `prompt-go-project.md` and `prompt-php-project.md` for accurate implementation.