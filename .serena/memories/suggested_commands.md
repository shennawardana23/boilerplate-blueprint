# Suggested Commands

## Development Workflow

### Backend (Go)
```bash
# Start development server
make run                    # Start Go server on port 8080

# Development with hot reload
make run-air               # Start with Air for hot reloading

# Testing
make test                  # Run all tests
make test-coverage         # Run tests with coverage report

# Code Quality
make fmt                   # Format Go code
make vet                   # Run go vet
make lint                  # Run golangci-lint (if installed)
make check-all             # Run all quality checks

# Build
make build                 # Build for current platform
make build-linux           # Build for Linux/AWS Lambda
make build-all             # Build for all platforms
```

### Frontend (Vue.js)
```bash
# Navigate to frontend directory
cd web

# Install dependencies
npm install

# Development server
npm run dev                # Start Vite dev server (usually port 5173)

# Build for production
npm run build              # Create production build in dist/

# Preview production build
npm run preview            # Preview built files

# Linting
npm run lint               # Run ESLint
```

### Full Stack Development
```bash
# Terminal 1: Backend
make run-air               # Hot reload Go server

# Terminal 2: Frontend  
cd web && npm run dev      # Vite dev server with HMR
```

## Database & Infrastructure
```bash
# Database migrations (if implemented)
make migrate-up            # Run migrations
make migrate-down          # Rollback migrations
make migrate-status        # Check migration status

# Docker operations
make docker-build          # Build Docker image
make docker-run            # Run in container
```

## Utility Commands
```bash
# Project setup
make setup                 # Complete initial setup
make deps                  # Install/update Go dependencies
make clean                 # Clean build artifacts

# Environment check
make env-check             # Check environment variables
make deps-check            # Check required tools

# Documentation
make swagger               # Generate API documentation
make help                  # Show all available commands
```