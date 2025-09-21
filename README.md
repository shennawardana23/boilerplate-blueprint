# 🚀 Boilerplate Blueprint

**AI-Powered Web-Based Boilerplate Generator**

Create Go and PHP CodeIgniter projects with AI assistance through an intuitive chat interface. Generate complete project structures, files, and configurations in seconds.

## ✨ Features

- 🤖 **AI Chat Interface** - Get intelligent suggestions and project recommendations
- 🐹 **Go Project Generation** - Clean Architecture with 17 utility packages
- 🐘 **PHP CodeIgniter** - MVC structure with security features
- 📦 **ZIP Downloads** - Instant project exports
- 🎨 **Modern UI** - Clean, responsive Vue.js interface
- ⚡ **Fast Generation** - Real-time project structure creation
- ☁️ **Cloud Deployment** - AWS Lambda, Docker, Kubernetes support
- 🔄 **CI/CD Pipeline** - Automated testing and deployment

## 🏗️ Architecture

- **Backend**: Go 1.21 + Gin Framework (Clean Architecture)
- **Frontend**: Vue.js 3 + Vite + Tailwind CSS
- **AI**: Rule-based chat system (OpenAI integration ready)
- **Database**: In-memory storage (no database required for prototype)
- **Deployment**: Multi-platform support (Local, AWS Lambda, Docker, K8s)

### Project Structure

```
boilerplate-blueprint/
├── cmd/main.go                    # Application entry point
├── internal/
│   ├── api/ (handlers, routes)    # REST API layer
│   │   └── lambda_handler.go      # AWS Lambda support
│   ├── models/ (project, chat)    # Data structures
│   └── services/ (3 services)     # Business logic
├── web/                          # Vue.js frontend
│   ├── src/ (14 Vue files)
│   └── test/ (frontend tests)
├── scripts/deploy/               # Deployment scripts
│   ├── lambda.sh                 # AWS Lambda deployment
│   └── docker.sh                 # Docker deployment
├── .github/workflows/            # CI/CD pipelines
├── tests/ (6 Go test files)       # Comprehensive tests
├── serverless.yml                # AWS Lambda config
├── docker-compose.yml            # Local Docker setup
├── Dockerfile                    # Container definition
└── Makefile                      # Development commands
```

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Node.js 18+
- Git
- Docker (for container deployment)
- AWS CLI (for Lambda deployment)

### Local Development

```bash
# Clone repository
git clone <repository-url>
cd boilerplate-blueprint

# Install dependencies
make deps

# Start development servers
make dev
```

Visit **http://localhost:5173** in your browser!

### Manual Setup

```bash
# Backend
go mod tidy
go run cmd/main.go

# Frontend (new terminal)
cd web
npm install
npm run dev
```

## ☁️ Deployment Options

### AWS Lambda Deployment

```bash
# Install Serverless Framework
npm install -g serverless

# Configure AWS credentials
aws configure

# Deploy to Lambda
make lambda-deploy

# View deployment info
make lambda-info

# View logs
make lambda-logs

# Remove deployment
make lambda-remove
```

### Docker Deployment

```bash
# Build and run locally
make docker-deploy

# View status
make docker-status

# View logs
make docker-logs

# Stop container
make docker-stop
```

### Docker Compose (Local)

```bash
# Start with Docker Compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

### Kubernetes (Future)

```bash
# Deploy to Kubernetes (when implemented)
make k8s-deploy
```

## 🧪 Testing & Quality

```bash
# Run all tests
make test

# Run frontend tests only
make test-frontend

# Run with coverage
make test-coverage

# Code quality checks
make check-all

# Run linting
make lint
```

## 📋 Usage

1. **Select Language**: Choose Go or PHP CodeIgniter
2. **Configure Project**: Set options and preferences
3. **AI Chat**: Get suggestions and customize your project
4. **Generate**: Create project structure and files
5. **Download**: Get your project as a ZIP file

## 🛠️ Development

### Available Commands

```bash
# Development
make help           # Show all available commands
make dev            # Start both backend & frontend
make run            # Start backend server only
make frontend-dev   # Start frontend development

# Building
make build          # Build application binary
make build-lambda   # Build for AWS Lambda
make docker-build   # Build Docker image

# Testing
make test           # Run all Go tests
make test-coverage  # Run tests with coverage
make test-frontend  # Run frontend tests

# Deployment
make lambda-deploy  # Deploy to AWS Lambda
make docker-deploy  # Deploy with Docker
make lambda-remove  # Remove Lambda deployment

# Code Quality
make fmt            # Format Go code
make lint           # Run linting
make vet            # Run static analysis
make check-all      # All quality checks

# Cleanup
make clean          # Clean build artifacts
make clean-docker   # Clean Docker resources
make clean-all      # Clean everything
```

### Environment Variables

```bash
# Server Configuration
PORT=8080                    # Server port
GIN_MODE=debug              # Gin mode (debug/release)

# AWS Lambda (when applicable)
LAMBDA_STAGE=dev            # Deployment stage
LAMBDA_REGION=us-east-1     # AWS region

# Docker
DOCKER_IMAGE=boilerplate-blueprint
DOCKER_TAG=latest
DOCKER_REGISTRY=            # Optional registry URL
```

## 🔄 CI/CD Pipeline

### GitHub Actions

The project includes comprehensive CI/CD with:

- **Automated Testing**: Go and JavaScript tests on every push/PR
- **Security Scanning**: Gosec and Trivy vulnerability scanning
- **Multi-Environment Deployment**: Staging and production environments
- **Performance Testing**: Load testing on staging deployments
- **Docker Build**: Automated container image building
- **Slack Notifications**: Deployment status notifications

### Pipeline Stages

1. **Test & Build** - Code quality, testing, and building
2. **Security Scan** - Vulnerability scanning
3. **Docker Build** - Container image creation
4. **Deploy Staging** - Automated staging deployment
5. **Performance Test** - Load testing on staging
6. **Deploy Production** - Production deployment (manual approval)
7. **Notifications** - Slack notifications

### Deployment Environments

- **Development**: Local development with hot reload
- **Staging**: Automated deployment from `develop` branch
- **Production**: Manual deployment from `main` branch

## 🎯 Project Templates

### Go Projects
- **Clean Architecture** with dependency injection
- **17 Utility Packages**: authentication, cache, database, logging, etc.
- **Framework Options**: Gin, Chi, Echo, Standard Library
- **Database Support**: PostgreSQL, MySQL, SQLite, MongoDB

### PHP CodeIgniter Projects
- **MVC Architecture** with proper separation
- **Security Features**: CSRF, XSS protection, input validation
- **Helper Libraries**: Template, authentication, database utilities
- **Database Support**: MySQL, PostgreSQL, SQLite

## 📊 Performance & Resources

### Local Development
- **Startup Time**: < 1 second
- **Memory Usage**: ~25-50MB RAM
- **Response Time**: < 100ms for API calls
- **Concurrent Users**: Supports 100+ simultaneous users

### AWS Lambda
- **Cold Start**: 1-3 seconds
- **Warm Response**: < 100ms
- **Memory**: 1024MB allocated
- **Timeout**: 30 seconds
- **Cost**: Pay-per-use pricing

### Docker
- **Image Size**: ~200MB (including Go and Node.js)
- **Memory Usage**: ~50-100MB RAM
- **Startup Time**: 2-5 seconds

## 🔒 Security

- **Input Validation**: All API inputs validated
- **CORS Configuration**: Properly configured for web access
- **No Database**: No SQL injection risk
- **Stateless Design**: No session vulnerabilities
- **HTTPS Ready**: TLS support in production deployments

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make changes with tests
4. Run `make check-all`
5. Commit changes (`git commit -m 'Add amazing feature'`)
6. Push to branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Guidelines

- Follow Go and Vue.js best practices
- Add tests for new features
- Update documentation
- Ensure all CI checks pass
- Use conventional commit messages

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by [Go Blueprint](https://go-blueprint.dev/)
- Built with Go, Vue.js, and modern web technologies
- AI-powered project generation with extensible architecture

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/your-org/boilerplate-blueprint/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-org/boilerplate-blueprint/discussions)
- **Documentation**: See [docs/](docs/) directory

---

**Made with ❤️ for developers who want to build amazing things faster!**

🎉 **Ready to generate your next project?** `make dev`