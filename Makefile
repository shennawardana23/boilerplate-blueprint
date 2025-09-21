# Boilerplate Blueprint - Development Tools

.PHONY: help build run test clean deps fmt lint check-all dev frontend frontend-dev lambda-deploy lambda-remove docker-build docker-run docker-push docker-deploy k8s-deploy k8s-remove

# =============================================================================
# VARIABLES
# =============================================================================
BINARY_NAME=boilerplate-blueprint
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev-$(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')")
BUILD_TIME=$(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS=-ldflags="-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# AWS Lambda configuration
LAMBDA_STAGE?=dev
LAMBDA_REGION?=us-east-1

# Docker configuration
DOCKER_IMAGE?=boilerplate-blueprint
DOCKER_TAG?=latest
DOCKER_REGISTRY?=

# =============================================================================
# DEFAULT TARGET
# =============================================================================
help: ## Show help message
	@echo '🚀 Boilerplate Blueprint - AI-Powered Project Generator'
	@echo '======================================================='
	@echo 'Version: $(VERSION)'
	@echo 'Build Time: $(BUILD_TIME)'
	@echo ''
	@echo '📋 Available Commands:'
	@echo ''
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ''
	@echo '🔧 Quick Start:'
	@echo '  make dev              - Start both backend & frontend'
	@echo '  make build            - Build the application'
	@echo '  make lambda-deploy    - Deploy to AWS Lambda'
	@echo '  make docker-deploy    - Deploy with Docker'
	@echo ''
	@echo '📦 Deployment Options:'
	@echo '  • AWS Lambda: make lambda-deploy'
	@echo '  • Docker: make docker-deploy'
	@echo '  • Kubernetes: make k8s-deploy'

# =============================================================================
# DEVELOPMENT
# =============================================================================
dev: ## Start both backend and frontend in development mode
	@echo "🚀 Starting development servers..."
	@echo "📡 Backend will be available at: http://localhost:8080"
	@echo "🌐 Frontend will be available at: http://localhost:5173"
	@echo ""
	@echo "Starting backend..."
	@make run &
	@echo "Starting frontend..."
	@cd web && npm run dev &
	@echo ""
	@echo "✅ Both servers started! Press Ctrl+C to stop all servers."

run: ## Start the backend server
	@echo "🚀 Starting Boilerplate Blueprint server..."
	@echo "📡 Server will be available at: http://localhost:8080"
	@echo "🌐 Frontend will be available at: http://localhost:5173"
	go run cmd/main.go

frontend: ## Build the frontend for production
	@echo "🔨 Building frontend..."
	cd web && npm run build
	@echo "✅ Frontend built successfully"

frontend-dev: ## Start frontend in development mode
	@echo "🚀 Starting frontend development server..."
	cd web && npm run dev

# =============================================================================
# BUILD
# =============================================================================
build: ## Build the application binary
	@echo "🔨 Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) cmd/main.go
	@echo "✅ Binary built: $(BINARY_NAME)"
	@echo "📊 Binary size: $$(ls -lh $(BINARY_NAME) | awk '{print $$5}')"

build-lambda: ## Build for AWS Lambda
	@echo "🔨 Building for AWS Lambda..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -tags lambda -o main cmd/main.go
	@echo "✅ Lambda binary built: main"

build-optimized: ## Build optimized binary for production
	@echo "🔨 Building optimized $(BINARY_NAME)..."
	CGO_ENABLED=0 go build $(LDFLAGS) -a -installsuffix cgo -o $(BINARY_NAME) cmd/main.go
	@echo "✅ Optimized binary built: $(BINARY_NAME)"
	@echo "📊 Binary size: $$(ls -lh $(BINARY_NAME) | awk '{print $$5}')"

# =============================================================================
# TESTING
# =============================================================================
test: ## Run all tests
	@echo "🧪 Running all tests..."
	go test -v -race -timeout=30s ./tests/...
	@echo "✅ All tests completed"

test-coverage: ## Run tests with coverage report
	@echo "🧪 Running tests with coverage..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"
	@echo "📈 Coverage summary:"
	@go tool cover -func=coverage.out | tail -1

test-frontend: ## Run frontend tests
	@echo "🧪 Running frontend tests..."
	cd web && npm test

# =============================================================================
# CODE QUALITY
# =============================================================================
deps: ## Install and update all dependencies
	@echo "📦 Installing Go dependencies..."
	go mod download
	go mod tidy
	go mod verify
	@echo "📦 Installing frontend dependencies..."
	cd web && npm install
	@echo "✅ All dependencies installed"

fmt: ## Format all Go code
	@echo "🎨 Formatting Go code..."
	go fmt ./...
	@echo "✅ Code formatting complete"

lint: ## Run linting (requires golangci-lint)
	@echo "🔍 Running linting checks..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run --timeout=5m; \
	else \
		echo "⚠️  golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi
	@echo "✅ Linting complete"

vet: ## Run go vet for static analysis
	@echo "🔍 Running go vet..."
	go vet ./...
	@echo "✅ Go vet complete"

check-all: fmt vet test ## Run all quality checks
	@echo "✅ All quality checks completed"

# =============================================================================
# AWS LAMBDA DEPLOYMENT
# =============================================================================
lambda-deploy: ## Deploy to AWS Lambda
	@echo "☁️  Deploying to AWS Lambda..."
	@echo "📍 Stage: $(LAMBDA_STAGE)"
	@echo "🌍 Region: $(LAMBDA_REGION)"
	./scripts/deploy/lambda.sh $(LAMBDA_STAGE) $(LAMBDA_REGION)

lambda-remove: ## Remove AWS Lambda deployment
	@echo "🗑️  Removing AWS Lambda deployment..."
	@echo "📍 Stage: $(LAMBDA_STAGE)"
	@echo "🌍 Region: $(LAMBDA_REGION)"
	@if command -v serverless >/dev/null 2>&1; then \
		serverless remove --stage $(LAMBDA_STAGE) --region $(LAMBDA_REGION); \
	else \
		echo "⚠️  Serverless Framework not found. Install with: npm install -g serverless"; \
		exit 1; \
	fi

lambda-logs: ## View AWS Lambda logs
	@echo "📋 Viewing AWS Lambda logs..."
	@echo "📍 Stage: $(LAMBDA_STAGE)"
	serverless logs --function api --stage $(LAMBDA_STAGE) --region $(LAMBDA_REGION) --tail

lambda-info: ## Get AWS Lambda deployment info
	@echo "ℹ️  AWS Lambda deployment info..."
	serverless info --stage $(LAMBDA_STAGE) --region $(LAMBDA_REGION)

# =============================================================================
# DOCKER DEPLOYMENT
# =============================================================================
docker-build: ## Build Docker image
	@echo "🐳 Building Docker image: $(DOCKER_IMAGE):$(DOCKER_TAG)"
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "✅ Docker image built successfully"

docker-run: ## Run Docker container locally
	@echo "🐳 Running Docker container..."
	docker run -d --name $(BINARY_NAME) -p 8080:8080 --restart unless-stopped $(DOCKER_IMAGE):$(DOCKER_TAG)
	@echo "✅ Container started: http://localhost:8080"

docker-stop: ## Stop Docker container
	@echo "🐳 Stopping Docker container..."
	docker stop $(BINARY_NAME) 2>/dev/null || true
	docker rm $(BINARY_NAME) 2>/dev/null || true
	@echo "✅ Container stopped"

docker-push: ## Push Docker image to registry
	@echo "🐳 Pushing Docker image..."
	@if [ -n "$(DOCKER_REGISTRY)" ]; then \
		docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(DOCKER_TAG); \
		docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(DOCKER_TAG); \
		echo "✅ Image pushed to: $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(DOCKER_TAG)"; \
	else \
		docker push $(DOCKER_IMAGE):$(DOCKER_TAG); \
		echo "✅ Image pushed: $(DOCKER_IMAGE):$(DOCKER_TAG)"; \
	fi

docker-logs: ## View Docker container logs
	@echo "📋 Docker container logs..."
	docker logs -f $(BINARY_NAME)

docker-deploy: ## Build and deploy with Docker
	@echo "🚀 Docker deployment..."
	make docker-build
	make docker-run
	@echo "⏳ Waiting for container to be ready..."
	@sleep 5
	@make docker-status

docker-status: ## Show Docker container status
	@echo "📊 Docker container status:"
	@docker ps -f name=$(BINARY_NAME) --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

# =============================================================================
# KUBERNETES DEPLOYMENT
# =============================================================================
k8s-deploy: ## Deploy to Kubernetes
	@echo "☸️  Deploying to Kubernetes..."
	@if ! command -v kubectl >/dev/null 2>&1; then \
		echo "❌ kubectl not found. Please install Kubernetes CLI."; \
		exit 1; \
	fi
	@echo "⚠️  Kubernetes deployment not yet implemented"
	@echo "📋 To implement:"
	@echo "  1. Create k8s/ directory with manifests"
	@echo "  2. Add Deployment, Service, and Ingress YAML files"
	@echo "  3. Run: kubectl apply -f k8s/"

k8s-remove: ## Remove Kubernetes deployment
	@echo "☸️  Removing Kubernetes deployment..."
	@echo "⚠️  Kubernetes removal not yet implemented"

# =============================================================================
# CLEANUP
# =============================================================================
clean: ## Clean all build artifacts
	@echo "🧹 Cleaning build artifacts..."
	rm -f $(BINARY_NAME) main
	rm -f coverage.out coverage.html
	rm -rf web/dist
	@echo "✅ Cleanup complete"

clean-docker: ## Clean Docker resources
	@echo "🧹 Cleaning Docker resources..."
	docker stop $(BINARY_NAME) 2>/dev/null || true
	docker rm $(BINARY_NAME) 2>/dev/null || true
	docker image rm $(DOCKER_IMAGE):$(DOCKER_TAG) 2>/dev/null || true
	docker system prune -f
	@echo "✅ Docker cleanup complete"

clean-all: clean clean-docker ## Clean everything
	@echo "🧹 Complete cleanup..."
	go clean -cache
	go clean -testcache
	go clean -modcache
	rm -rf web/node_modules web/package-lock.json
	@echo "✅ Complete cleanup finished"

# =============================================================================
# UTILITIES
# =============================================================================
version: ## Show version information
	@echo "📋 Version Information:"
	@echo "  Version: $(VERSION)"
	@echo "  Go Version: $(shell go version)"
	@echo "  Build Time: $(BUILD_TIME)"
	@echo "  Docker Image: $(DOCKER_IMAGE):$(DOCKER_TAG)"

size: ## Show binary size and project statistics
	@echo "📊 Project Statistics:"
	@echo "  Go files: $$(find . -name '*.go' -not -path './vendor/*' | wc -l)"
	@echo "  Test files: $$(find . -name '*_test.go' | wc -l)"
	@echo "  Vue files: $$(find web -name '*.vue' | wc -l)"
	@echo "  Total lines of Go code: $$(find . -name '*.go' -not -path './vendor/*' -exec wc -l {} \; | awk '{sum += $$1} END {print sum}')"
	@if [ -f $(BINARY_NAME) ]; then \
		echo "  Binary size: $$(ls -lh $(BINARY_NAME) | awk '{print $$5}')"; \
	fi
	@if [ -f main ]; then \
		echo "  Lambda binary size: $$(ls -lh main | awk '{print $$5}')"; \
	fi

health-check: ## Run health check against running server
	@echo "🏥 Running health check..."
	@if curl -f -s http://localhost:8080/api/health > /dev/null; then \
		echo "✅ Server is healthy"; \
	else \
		echo "❌ Server health check failed"; \
		exit 1; \
	fi

load-test: ## Run basic load test (requires hey)
	@echo "⚡ Running load test..."
	@if command -v hey >/dev/null 2>&1; then \
		hey -n 100 -c 10 http://localhost:8080/api/health; \
	else \
		echo "⚠️  hey not found. Install with: go install github.com/rakyll/hey@latest"; \
	fi

# =============================================================================
# DEVELOPMENT SETUP
# =============================================================================
setup: deps ## Complete initial development setup
	@echo "🎉 Setup complete!"
	@echo ""
	@echo "📋 Next steps:"
	@echo "  1. Run 'make dev' to start development servers"
	@echo "  2. Visit http://localhost:5173 in your browser"
	@echo "  3. Start building amazing boilerplate projects!"
	@echo ""
	@echo "🚀 Deployment Options:"
	@echo "  • Local: make dev"
	@echo "  • AWS Lambda: make lambda-deploy"
	@echo "  • Docker: make docker-deploy"
	@echo "  • Kubernetes: make k8s-deploy (future)"
	@echo ""
	@echo "📚 Useful commands:"
	@echo "  make test           - Run all tests"
	@echo "  make build          - Build the application"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make help           - Show all commands"