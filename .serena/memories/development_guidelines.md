# Development Guidelines

## Project Philosophy
This project follows a **quick prototype approach** prioritizing:
- Getting a working MVP quickly
- Simple, understandable code over complex architectures
- Iterative development and improvement
- Real functionality over perfect design

## Design Patterns and Principles

### Go Backend Patterns
- **Service Layer Pattern**: Business logic separated from HTTP handlers
- **Constructor Functions**: `NewXxxService()` functions for dependency injection
- **Interface Segregation**: Small, focused interfaces for testability
- **Error Wrapping**: Proper error handling with context
- **Environment Configuration**: `.env` files with godotenv

### Frontend Patterns (Vue.js)
- **Composition API**: Preferred over Options API for new components
- **Composables**: Reusable logic extracted into composable functions
- **Reactive Data**: `ref()` and `reactive()` for state management
- **Component Composition**: Small, focused, reusable components

## Key Requirements to Follow

### Project Generation Must Include
1. **Exact Directory Structures**: Follow patterns from `project-description.md`
2. **All 17 Go Utility Packages**: When generating Go projects
3. **Complete MVC Structure**: When generating PHP CodeIgniter projects
4. **AI Chat Integration**: Real-time chat for project customization
5. **ZIP Download**: Every generation must produce downloadable ZIP

### Technology Decisions
- **Backend**: Go + Gin (avoid complex frameworks initially)
- **Frontend**: Vue.js + TailwindCSS (quick UI development)
- **Storage**: In-memory for prototype (no database dependency)
- **AI Integration**: Direct API calls (OpenAI or similar)
- **File Generation**: Go templates for dynamic code generation

## Security Considerations
- **CORS Configuration**: Proper cross-origin setup
- **Input Validation**: Validate all user inputs
- **File System Safety**: Secure file operations and paths
- **Environment Variables**: Sensitive data in .env files
- **Error Handling**: Don't expose internal errors to users

## Performance Guidelines
- **Concurrent Operations**: Use Go's concurrency for file operations
- **Template Caching**: Cache compiled templates
- **Static File Serving**: Efficient serving of Vue.js build
- **Memory Management**: Clean up temporary files and resources

## Testing Strategy
- **Unit Tests**: Test services and business logic
- **Integration Tests**: Test API endpoints
- **Frontend Testing**: Component testing (when implemented)
- **Manual Testing**: UI flows and file generation

## Migration Path (Future)
Current prototype can be enhanced with:
- PostgreSQL for persistence
- Redis for caching
- WebSocket for real-time updates
- React/Next.js migration if needed
- Docker containerization
- CI/CD pipeline