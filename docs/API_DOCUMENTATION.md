# Boilerplate Blueprint - API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication
Currently, no authentication is required. All endpoints are publicly accessible.

## Response Format
All API responses follow a consistent format:

### Success Response
```json
{
  "success": true,
  "data": { ... },
  "message": "Optional success message"
}
```

### Error Response
```json
{
  "success": false,
  "error": "Error message describing what went wrong"
}
```

## Endpoints

### Health Check

#### GET /health
Check the health status of the API server.

**Response:**
```json
{
  "status": "healthy",
  "service": "boilerplate-blueprint",
  "version": "1.0.0"
}
```

---

### Templates

#### GET /templates
Get all available project templates.

**Response:**
```json
{
  "success": true,
  "templates": [
    {
      "language": "go",
      "name": "Go Clean Architecture",
      "description": "Complete Go project with Clean Architecture, Gin framework, and 17 utility packages",
      "options": [
        {
          "key": "framework",
          "label": "HTTP Framework",
          "type": "select",
          "required": true,
          "default": "gin",
          "options": ["gin", "chi", "echo", "standard"],
          "description": "Choose your HTTP framework"
        },
        {
          "key": "database",
          "label": "Database",
          "type": "select",
          "required": true,
          "default": "postgresql",
          "options": ["postgresql", "mysql", "sqlite", "mongodb"],
          "description": "Choose your database"
        },
        {
          "key": "authentication",
          "label": "Authentication",
          "type": "select",
          "required": true,
          "default": "jwt",
          "options": ["jwt", "oauth", "basic"],
          "description": "Choose your authentication method"
        }
      ]
    },
    {
      "language": "php",
      "name": "PHP CodeIgniter MVC",
      "description": "Complete PHP CodeIgniter project with MVC architecture and security features",
      "options": [
        {
          "key": "ci_version",
          "label": "CodeIgniter Version",
          "type": "select",
          "required": true,
          "default": "3",
          "options": ["3", "4"],
          "description": "Choose CodeIgniter version"
        },
        {
          "key": "database",
          "label": "Database",
          "type": "select",
          "required": true,
          "default": "postgresql",
          "options": ["postgresql", "mysql", "sqlite"],
          "description": "Choose your database"
        },
        {
          "key": "frontend",
          "label": "Frontend Framework",
          "type": "select",
          "required": true,
          "default": "bootstrap",
          "options": ["bootstrap", "tailwind", "custom"],
          "description": "Choose your frontend framework"
        }
      ]
    }
  ]
}
```

---

### Projects

#### POST /projects
Create a new project.

**Request Body:**
```json
{
  "name": "my-awesome-project",
  "language": "go",
  "description": "A web API for managing users",
  "options": {
    "framework": "gin",
    "database": "postgresql",
    "authentication": "jwt",
    "utilities": ["authentication", "cache", "common"]
  }
}
```

**Response:**
```json
{
  "success": true,
  "project": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "my-awesome-project",
    "language": "go",
    "description": "A web API for managing users",
    "options": {
      "framework": "gin",
      "database": "postgresql",
      "authentication": "jwt",
      "utilities": ["authentication", "cache", "common"]
    },
    "files": [],
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

**Validation Rules:**
- `name`: Required, non-empty string
- `language`: Required, must be "go" or "php"
- `description`: Optional string
- `options`: Optional object with language-specific options

#### GET /projects/:id
Get a project by its ID.

**Response:**
```json
{
  "success": true,
  "project": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "my-awesome-project",
    "language": "go",
    "description": "A web API for managing users",
    "options": {
      "framework": "gin",
      "database": "postgresql",
      "authentication": "jwt",
      "utilities": ["authentication", "cache", "common"]
    },
    "files": [
      {
        "path": "my-awesome-project/go.mod",
        "content": "module my-awesome-project\n\ngo 1.21\n\nrequire (\n\tgithub.com/gin-gonic/gin v1.9.1\n\tgithub.com/lib/pq v1.10.9\n)",
        "is_directory": false
      }
    ],
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:35:00Z"
  }
}
```

**Error Responses:**
- `404 Not Found`: Project with the given ID does not exist

#### POST /projects/:id/generate
Generate project files for an existing project.

**Response:**
```json
{
  "success": true,
  "project": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "my-awesome-project",
    "language": "go",
    "description": "A web API for managing users",
    "options": {
      "framework": "gin",
      "database": "postgresql",
      "authentication": "jwt",
      "utilities": ["authentication", "cache", "common"]
    },
    "files": [
      {
        "path": "my-awesome-project",
        "content": "",
        "is_directory": true
      },
      {
        "path": "my-awesome-project/go.mod",
        "content": "module my-awesome-project\n\ngo 1.21\n\nrequire (\n\tgithub.com/gin-gonic/gin v1.9.1\n\tgithub.com/lib/pq v1.10.9\n\tgithub.com/redis/go-redis/v9 v9.3.0\n\tgithub.com/golang-jwt/jwt/v5 v5.2.0\n\tgithub.com/google/uuid v1.6.0\n\tgithub.com/joho/godotenv v1.5.1\n)",
        "is_directory": false
      },
      {
        "path": "my-awesome-project/cmd",
        "content": "",
        "is_directory": true
      },
      {
        "path": "my-awesome-project/cmd/main.go",
        "content": "package main\n\nimport (\n\t\"log\"\n\t\"os\"\n\n\t\"my-awesome-project/internal/api\"\n\t\"my-awesome-project/internal/services\"\n\n\t\"github.com/gin-contrib/cors\"\n\t\"github.com/gin-gonic/gin\"\n\t\"github.com/joho/godotenv\"\n)\n\nfunc main() {\n\t// Load environment variables\n\tif err := godotenv.Load(); err != nil {\n\t\tlog.Println(\"No .env file found, using system environment variables\")\n\t}\n\n\t// Set Gin mode\n\tif os.Getenv(\"GIN_MODE\") == \"\" {\n\t\tgin.SetMode(gin.DebugMode)\n\t}\n\n\t// Initialize services\n\ttemplateService := services.NewTemplateService()\n\tprojectService := services.NewProjectService(templateService)\n\tchatService := services.NewChatService()\n\n\t// Initialize handlers\n\thandlers := api.NewHandlers(projectService, templateService, chatService)\n\n\t// Create Gin router\n\trouter := gin.Default()\n\n\t// Configure CORS\n\tconfig := cors.DefaultConfig()\n\tconfig.AllowOrigins = []string{\"http://localhost:3000\", \"http://localhost:5173\"}\n\tconfig.AllowMethods = []string{\"GET\", \"POST\", \"PUT\", \"DELETE\", \"OPTIONS\"}\n\tconfig.AllowHeaders = []string{\"Origin\", \"Content-Type\", \"Accept\", \"Authorization\"}\n\trouter.Use(cors.New(config))\n\n\t// Setup routes\n\tapi.SetupRoutes(router, handlers)\n\n\t// Serve static files (Vue.js build)\n\trouter.Static(\"/static\", \"./web/dist\")\n\trouter.StaticFile(\"/\", \"./web/dist/index.html\")\n\n\t// Start server\n\tport := os.Getenv(\"PORT\")\n\tif port == \"\" {\n\t\tport = \"8080\"\n\t}\n\n\tlog.Printf(\"🚀 Boilerplate Blueprint server starting on port %s\", port)\n\tif err := router.Run(\":\" + port); err != nil {\n\t\tlog.Fatal(\"Failed to start server:\", err)\n\t}\n}",
        "is_directory": false
      }
    ],
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:40:00Z"
  }
}
```

**Error Responses:**
- `404 Not Found`: Project with the given ID does not exist

#### GET /projects/:id/download
Download a project as a ZIP file.

**Response:**
- **Content-Type**: `application/zip`
- **Content-Disposition**: `attachment; filename="project-name-language.zip"`
- **Body**: Binary ZIP file content

**Error Responses:**
- `404 Not Found`: Project with the given ID does not exist
- `500 Internal Server Error`: Failed to generate ZIP file

---

### Chat

#### POST /chat/message
Send a message to the AI chat system.

**Request Body:**
```json
{
  "message": "I want to build a Go web API with authentication",
  "project_id": "550e8400-e29b-41d4-a716-446655440000",
  "context": "project setup"
}
```

**Response:**
```json
{
  "success": true,
  "message": {
    "id": "msg-12345",
    "role": "assistant",
    "content": "Great choice! Go is excellent for building high-performance web APIs. I can help you set up a Go project with Clean Architecture, including JWT authentication and all the necessary utility packages. What specific features do you need for your API?",
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "created_at": "2024-01-15T10:45:00Z"
  },
  "suggestions": [
    {
      "type": "language",
      "value": "go",
      "reason": "You mentioned Go in your message",
      "confidence": 0.9,
      "apply": true
    },
    {
      "type": "framework",
      "value": "gin",
      "reason": "Gin is a popular, fast HTTP framework for Go web applications",
      "confidence": 0.8,
      "apply": true
    },
    {
      "type": "authentication",
      "value": "jwt",
      "reason": "JWT is modern, stateless, and perfect for API authentication",
      "confidence": 0.85,
      "apply": false
    }
  ]
}
```

**Validation Rules:**
- `message`: Required, non-empty string
- `project_id`: Optional string
- `context`: Optional string

#### GET /chat/history
Get chat history for a project or general conversation.

**Query Parameters:**
- `project_id` (optional): Get history for specific project

**Response:**
```json
{
  "success": true,
  "history": {
    "project_id": "550e8400-e29b-41d4-a716-446655440000",
    "messages": [
      {
        "id": "msg-12345",
        "role": "user",
        "content": "I want to build a Go web API with authentication",
        "project_id": "550e8400-e29b-41d4-a716-446655440000",
        "created_at": "2024-01-15T10:44:00Z"
      },
      {
        "id": "msg-12346",
        "role": "assistant",
        "content": "Great choice! Go is excellent for building high-performance web APIs...",
        "project_id": "550e8400-e29b-41d4-a716-446655440000",
        "created_at": "2024-01-15T10:45:00Z"
      }
    ],
    "created_at": "2024-01-15T10:44:00Z",
    "updated_at": "2024-01-15T10:45:00Z"
  }
}
```

---

## Error Codes

### HTTP Status Codes
- `200 OK`: Request successful
- `400 Bad Request`: Invalid request data
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

### Common Error Messages
- `"unsupported language: {language}"`: Invalid language specified
- `"project not found: {id}"`: Project with given ID doesn't exist
- `"failed to generate project files"`: Error during file generation
- `"failed to create ZIP archive"`: Error during ZIP creation
- `"failed to generate AI response"`: Error in chat processing

## Rate Limiting
Currently, no rate limiting is implemented. This may be added in future versions.

## CORS
Cross-Origin Resource Sharing is enabled for the following origins:
- `http://localhost:3000`
- `http://localhost:5173`

Allowed methods: `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`
Allowed headers: `Origin`, `Content-Type`, `Accept`, `Authorization`

## Examples

### Creating a Go Project
```bash
curl -X POST http://localhost:8080/api/projects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-go-api",
    "language": "go",
    "description": "A REST API built with Go",
    "options": {
      "framework": "gin",
      "database": "postgresql",
      "authentication": "jwt"
    }
  }'
```

### Generating Project Files
```bash
curl -X POST http://localhost:8080/api/projects/550e8400-e29b-41d4-a716-446655440000/generate
```

### Downloading Project
```bash
curl -X GET http://localhost:8080/api/projects/550e8400-e29b-41d4-a716-446655440000/download \
  -o my-go-api.zip
```

### Sending Chat Message
```bash
curl -X POST http://localhost:8080/api/chat/message \
  -H "Content-Type: application/json" \
  -d '{
    "message": "I need help setting up database migrations",
    "project_id": "550e8400-e29b-41d4-a716-446655440000"
  }'
```

---

**Note**: This API is currently in development. Some endpoints and response formats may change in future versions.