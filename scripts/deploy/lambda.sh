#!/bin/bash

# AWS Lambda Deployment Script for Boilerplate Blueprint
# Usage: ./scripts/deploy/lambda.sh [stage] [region]

set -e

# Configuration
STAGE=${1:-dev}
REGION=${2:-us-east-1}
SERVICE_NAME="boilerplate-blueprint"
BINARY_NAME="main"

echo "🚀 Deploying $SERVICE_NAME to AWS Lambda"
echo "📍 Stage: $STAGE"
echo "🌍 Region: $REGION"

# Check prerequisites
check_prerequisites() {
    echo "🔍 Checking prerequisites..."

    if ! command -v aws &> /dev/null; then
        echo "❌ AWS CLI is not installed. Please install it first."
        exit 1
    fi

    if ! command -v serverless &> /dev/null; then
        echo "❌ Serverless Framework is not installed. Please install it first."
        echo "   npm install -g serverless"
        exit 1
    fi

    if ! aws sts get-caller-identity &> /dev/null; then
        echo "❌ AWS CLI is not configured. Please run 'aws configure' first."
        exit 1
    fi

    echo "✅ Prerequisites check passed"
}

# Build the application
build_application() {
    echo "🔨 Building Go application for Lambda..."

    # Build for Linux (required for Lambda)
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
        -ldflags="-s -w" \
        -tags lambda \
        -o $BINARY_NAME \
        ./cmd/main.go

    if [ ! -f "$BINARY_NAME" ]; then
        echo "❌ Build failed"
        exit 1
    fi

    echo "✅ Application built successfully"
}

# Build frontend
build_frontend() {
    echo "🔨 Building frontend..."

    if [ ! -d "web" ]; then
        echo "⚠️  Web directory not found, skipping frontend build"
        return
    fi

    cd web
    npm run build

    if [ ! -d "dist" ]; then
        echo "❌ Frontend build failed"
        exit 1
    fi

    cd ..
    echo "✅ Frontend built successfully"
}

# Deploy to Lambda
deploy_lambda() {
    echo "☁️  Deploying to AWS Lambda..."

    # Deploy using Serverless Framework
    serverless deploy --stage $STAGE --region $REGION

    if [ $? -eq 0 ]; then
        echo "✅ Deployment successful!"

        # Get the API Gateway URL
        API_URL=$(serverless info --stage $STAGE --region $REGION | grep "HttpApiUrl" | awk '{print $2}')

        echo ""
        echo "🎉 Deployment Complete!"
        echo "🌐 API URL: $API_URL"
        echo "📊 Stage: $STAGE"
        echo "🌍 Region: $REGION"
        echo ""
        echo "📋 Useful commands:"
        echo "  • View logs: serverless logs --function api --stage $STAGE --region $REGION"
        echo "  • Remove deployment: serverless remove --stage $STAGE --region $REGION"
        echo "  • Update function: serverless deploy function --function api --stage $STAGE --region $REGION"
    else
        echo "❌ Deployment failed"
        exit 1
    fi
}

# Cleanup
cleanup() {
    echo "🧹 Cleaning up temporary files..."
    rm -f $BINARY_NAME
    echo "✅ Cleanup complete"
}

# Main deployment process
main() {
    echo "🏁 Starting deployment process..."

    check_prerequisites
    build_application
    build_frontend
    deploy_lambda
    cleanup

    echo ""
    echo "🎊 All done! Your Boilerplate Blueprint is now running on AWS Lambda!"
}

# Handle script arguments
case "${1:-}" in
    --help|-h)
        echo "AWS Lambda Deployment Script"
        echo ""
        echo "Usage: $0 [stage] [region]"
        echo ""
        echo "Arguments:"
        echo "  stage    Deployment stage (dev, staging, prod) [default: dev]"
        echo "  region   AWS region [default: us-east-1]"
        echo ""
        echo "Examples:"
        echo "  $0                    # Deploy to dev stage in us-east-1"
        echo "  $0 staging           # Deploy to staging stage in us-east-1"
        echo "  $0 prod us-west-2    # Deploy to prod stage in us-west-2"
        exit 0
        ;;
    *)
        main "$@"
        ;;
esac