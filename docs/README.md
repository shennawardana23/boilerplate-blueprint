# 📚 Boilerplate Blueprint - Documentation Guide

Welcome to the Boilerplate Blueprint documentation! This guide helps you understand and work with our AI-powered boilerplate generator.

## 🎯 Quick Start for New Developers

### 1. **First 5 Minutes: System Overview**
- **What it is**: Web-based tool to generate Go and PHP project boilerplates
- **Tech Stack**: Vue.js frontend, Go backend, AWS Lambda deployment
- **Key Features**: AI chat assistance, real-time preview, ZIP downloads

### 2. **First 15 Minutes: Get Running**
```bash
# Clone and setup
git clone <repo-url>
cd boilerplate-blueprint
make deps
make dev

# Visit http://localhost:5173
```

### 3. **First Hour: Understand Architecture**
Follow the documentation in this order:

## 📖 Documentation Structure

### 🏗️ **Architecture & Design**
| Document | Purpose | When to Read |
|----------|---------|--------------|
| [`ARCHITECTURE_DIAGRAMS.md`](./ARCHITECTURE_DIAGRAMS.md) | Visual system overview | **Start here** - understand the big picture |
| [`FLOW_DIAGRAMS.md`](./FLOW_DIAGRAMS.md) | Detailed data flow | When debugging or adding features |
| [`../README.md`](../README.md) | Project overview | **Essential** - setup and features |

### 🔧 **Development & Deployment**
| Document | Purpose | When to Read |
|----------|---------|--------------|
| [`../DEVELOPMENT_GUIDE.md`](../DEVELOPMENT_GUIDE.md) | Development workflow | **Essential** - coding standards, debugging |
| [`../API_DOCUMENTATION.md`](../API_DOCUMENTATION.md) | API reference | When building integrations |
| [`../DEPLOYMENT.md`](../DEPLOYMENT.md) | Deployment options | When deploying to production |

### 📋 **Project Management**
| Document | Purpose | When to Read |
|----------|---------|--------------|
| [`../PROJECT_ONBOARDING.md`](../PROJECT_ONBOARDING.md) | Component details | When deep-diving into code |
| [`../INFRASTRUCTURE_ANALYSIS.md`](../INFRASTRUCTURE_ANALYSIS.md) | Performance & scaling | When optimizing or planning growth |

## 🔍 Key Architecture Concepts

### 🏛️ **System Layers**
```
User Interface (Vue.js)
    ↓
API Gateway (AWS Lambda)
    ↓
Application Layer (Go/Gin)
├── Controller Layer (HTTP handlers)
├── Service Layer (Business logic)
└── Model Layer (Data structures)
    ↓
Storage Layer (In-memory/PostgreSQL)
```

### 🔄 **Request Flow Pattern**
```
User Action → Vue Component → HTTP Request → API Gateway → Lambda → Controller → Service → Model → Storage → Response
```

### 🗂️ **Data Flow Pattern**
```
Input Validation → Business Logic → Entity Creation → Persistence → Response Formatting
```

## 🎯 Common Developer Tasks

### "I want to add a new project template"

1. **Read**: [`../DEVELOPMENT_GUIDE.md`](../DEVELOPMENT_GUIDE.md) - Component architecture
2. **Read**: [`FLOW_DIAGRAMS.md`](./FLOW_DIAGRAMS.md) - Data flow patterns
3. **Code**: Modify `internal/services/template.go`
4. **Test**: Add tests in `tests/services/template_service_test.go`
5. **Deploy**: Follow [`../DEPLOYMENT.md`](../DEPLOYMENT.md)

### "I want to add a new API endpoint"

1. **Read**: [`../API_DOCUMENTATION.md`](../API_DOCUMENTATION.md) - API patterns
2. **Read**: [`FLOW_DIAGRAMS.md`](./FLOW_DIAGRAMS.md) - Request flow
3. **Code**: Add handler in `internal/api/handlers.go`
4. **Code**: Add route in `internal/api/routes.go`
5. **Test**: Add tests in `tests/api/handlers_test.go`

### "I want to optimize performance"

1. **Read**: [`../INFRASTRUCTURE_ANALYSIS.md`](../INFRASTRUCTURE_ANALYSIS.md) - Performance analysis
2. **Read**: [`FLOW_DIAGRAMS.md`](./FLOW_DIAGRAMS.md) - Bottleneck identification
3. **Monitor**: Use CloudWatch metrics
4. **Optimize**: Focus on identified bottlenecks

### "I want to deploy to production"

1. **Read**: [`../DEPLOYMENT.md`](../DEPLOYMENT.md) - Deployment options
2. **Choose**: Lambda, Docker, or Kubernetes
3. **Configure**: Environment variables and secrets
4. **Deploy**: Use provided scripts and CI/CD

## 🛠️ Development Tools

### Local Development
```bash
make dev          # Start both frontend & backend
make build        # Build the application
make test         # Run all tests
make clean        # Clean build artifacts
```

### Code Quality
```bash
make fmt          # Format Go code
make lint         # Run linting
make vet          # Static analysis
make check-all    # All quality checks
```

### Deployment
```bash
make lambda-deploy    # Deploy to AWS Lambda
make docker-deploy    # Deploy with Docker
make docker-run       # Run locally with Docker
```

## 🔍 Understanding the Diagrams

### 📊 **Architecture Diagrams**
- **System Overview**: High-level component relationships
- **Request Flow**: How HTTP requests travel through the system
- **Data Lineage**: How data transforms from input to storage
- **Component Architecture**: Detailed service interactions

### 🔄 **Flow Diagrams**
- **Controller to Database**: Complete request journey
- **Data Lineage**: Variable tracing through layers
- **Error Handling**: Exception propagation paths
- **Performance**: Latency and memory usage breakdown

### 🎨 **Diagram Types Used**
- **Sequence Diagrams**: Show time-ordered interactions
- **Flowcharts**: Show decision points and data transformation
- **State Diagrams**: Show state changes over time
- **Pie Charts**: Show resource distribution
- **Graph Diagrams**: Show component relationships

## 🎯 Best Practices

### Code Organization
- **Controllers**: Handle HTTP, validation, responses
- **Services**: Contain business logic, call repositories
- **Models**: Define data structures and validation
- **Tests**: Mirror source structure in `tests/` directory

### Error Handling
- **Validation**: Early in request pipeline
- **Logging**: Structured logging with context
- **Metrics**: Error rates and performance monitoring
- **Responses**: Consistent error response format

### Performance
- **Caching**: Consider Redis for frequently accessed data
- **Database**: Use connection pooling
- **Async**: Background processing for heavy operations
- **CDN**: Static assets via CloudFront

## 🚨 Troubleshooting

### Common Issues

#### "Tests are failing with LC_UUID error"
**Cause**: macOS system issue with test binaries
**Solution**: Run tests individually or use CI/CD pipeline

#### "Lambda deployment fails"
**Cause**: Missing AWS credentials or permissions
**Solution**: Check `aws configure` and IAM policies

#### "Frontend build fails"
**Cause**: Missing dependencies or Node.js version
**Solution**: Run `npm install` and check Node.js version

#### "Database connection fails"
**Cause**: Missing environment variables
**Solution**: Check `DATABASE_URL` and connection settings

### Getting Help

1. **Check Documentation**: Most answers are in these docs
2. **Run Tests**: `make test` to verify local setup
3. **Check Logs**: Use `make lambda-logs` for deployed issues
4. **GitHub Issues**: For bugs and feature requests

## 📈 Learning Path for New Developers

### Week 1: Getting Started
- [ ] Read project README
- [ ] Set up local development environment
- [ ] Run the application locally
- [ ] Make a small UI change

### Week 2: Understanding Architecture
- [ ] Read architecture diagrams
- [ ] Follow a request through the flow diagrams
- [ ] Understand service layer responsibilities
- [ ] Add a simple feature (e.g., new validation)

### Week 3: Deep Dive
- [ ] Read API documentation
- [ ] Understand data models
- [ ] Add a new API endpoint
- [ ] Write comprehensive tests

### Week 4: Production Ready
- [ ] Read deployment documentation
- [ ] Set up CI/CD pipeline
- [ ] Deploy to staging environment
- [ ] Monitor and optimize performance

## 🎉 Welcome Aboard!

This documentation is designed to get you productive quickly while providing deep technical reference when needed. Start with the basics, use the diagrams to understand complex flows, and refer to specific guides for detailed implementation.

**Happy coding! 🚀**

---

*Last updated: 2024-09-21 | Boilerplate Blueprint v1.0.0*