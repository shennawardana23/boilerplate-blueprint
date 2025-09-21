#!/bin/bash

# Docker Deployment Script for Boilerplate Blueprint
# Usage: ./scripts/deploy/docker.sh [action] [tag]

set -e

# Configuration
SERVICE_NAME="boilerplate-blueprint"
DOCKER_IMAGE="${SERVICE_NAME}"
TAG=${2:-latest}
ACTION=${1:-build}

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Build Docker image
build_image() {
    log_info "Building Docker image: ${DOCKER_IMAGE}:${TAG}"

    # Build the Go application first
    log_info "Building Go application..."
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/main.go

    # Build the frontend
    if [ -d "web" ]; then
        log_info "Building frontend..."
        cd web && npm run build && cd ..
    fi

    # Build Docker image
    docker build -t ${DOCKER_IMAGE}:${TAG} .

    if [ $? -eq 0 ]; then
        log_success "Docker image built successfully: ${DOCKER_IMAGE}:${TAG}"
    else
        log_error "Failed to build Docker image"
        exit 1
    fi
}

# Push Docker image
push_image() {
    log_info "Pushing Docker image: ${DOCKER_IMAGE}:${TAG}"

    # Check if registry is specified
    if [ -n "${DOCKER_REGISTRY}" ]; then
        FULL_IMAGE="${DOCKER_REGISTRY}/${DOCKER_IMAGE}:${TAG}"
        docker tag ${DOCKER_IMAGE}:${TAG} ${FULL_IMAGE}
        docker push ${FULL_IMAGE}
        log_success "Image pushed to registry: ${FULL_IMAGE}"
    else
        docker push ${DOCKER_IMAGE}:${TAG}
        log_success "Image pushed: ${DOCKER_IMAGE}:${TAG}"
    fi
}

# Run Docker container locally
run_local() {
    log_info "Running Docker container locally"

    # Stop any existing container
    docker stop ${SERVICE_NAME} 2>/dev/null || true
    docker rm ${SERVICE_NAME} 2>/dev/null || true

    # Run new container
    docker run -d \
        --name ${SERVICE_NAME} \
        -p 8080:8080 \
        --restart unless-stopped \
        ${DOCKER_IMAGE}:${TAG}

    log_success "Container started successfully"
    log_info "Application available at: http://localhost:8080"
}

# Deploy to production
deploy_production() {
    log_warning "Production deployment not implemented in this script"
    log_info "For production deployment, consider:"
    log_info "  - AWS ECS/Fargate"
    log_info "  - Google Cloud Run"
    log_info "  - Azure Container Instances"
    log_info "  - Kubernetes"
    log_info "  - Docker Compose with reverse proxy"
}

# Show container logs
show_logs() {
    log_info "Showing container logs for ${SERVICE_NAME}"
    docker logs -f ${SERVICE_NAME}
}

# Show container status
show_status() {
    log_info "Container status:"
    docker ps -f name=${SERVICE_NAME} --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
}

# Stop and remove container
stop_container() {
    log_info "Stopping and removing container: ${SERVICE_NAME}"
    docker stop ${SERVICE_NAME} 2>/dev/null || true
    docker rm ${SERVICE_NAME} 2>/dev/null || true
    log_success "Container stopped and removed"
}

# Clean up Docker resources
cleanup() {
    log_info "Cleaning up Docker resources..."
    docker image prune -f
    docker container prune -f
    log_success "Cleanup completed"
}

# Show usage
show_usage() {
    echo "Docker Deployment Script for Boilerplate Blueprint"
    echo ""
    echo "Usage: $0 [action] [tag]"
    echo ""
    echo "Actions:"
    echo "  build       Build Docker image"
    echo "  push        Push Docker image to registry"
    echo "  run         Run container locally"
    echo "  deploy      Deploy to production (placeholder)"
    echo "  logs        Show container logs"
    echo "  status      Show container status"
    echo "  stop        Stop and remove container"
    echo "  cleanup     Clean up Docker resources"
    echo "  all         Build, run, and show status"
    echo ""
    echo "Arguments:"
    echo "  tag         Docker image tag [default: latest]"
    echo ""
    echo "Environment Variables:"
    echo "  DOCKER_REGISTRY    Docker registry URL (e.g., your-registry.com)"
    echo ""
    echo "Examples:"
    echo "  $0 build                    # Build image with latest tag"
    echo "  $0 build v1.0.0            # Build image with v1.0.0 tag"
    echo "  $0 run                      # Run container locally"
    echo "  DOCKER_REGISTRY=your-registry.com $0 push v1.0.0"
    echo "  $0 all                      # Build and run locally"
}

# Main deployment process
main() {
    case "${ACTION}" in
        build)
            build_image
            ;;
        push)
            push_image
            ;;
        run)
            run_local
            ;;
        deploy)
            deploy_production
            ;;
        logs)
            show_logs
            ;;
        status)
            show_status
            ;;
        stop)
            stop_container
            ;;
        cleanup)
            cleanup
            ;;
        all)
            build_image
            run_local
            sleep 2
            show_status
            ;;
        --help|-h)
            show_usage
            exit 0
            ;;
        *)
            log_error "Unknown action: ${ACTION}"
            echo ""
            show_usage
            exit 1
            ;;
    esac
}

# Run main function
main "$@"