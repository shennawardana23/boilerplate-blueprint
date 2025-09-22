# 🚀 Boilerplate Blueprint - Deployment Guide

This guide covers all deployment options for Boilerplate Blueprint, from local development to cloud production environments.

## 📋 Deployment Options Overview

| Platform | Use Case | Setup Time | Scaling | Cost |
|----------|----------|------------|---------|------|
| Local | Development | < 5 min | Single user | Free |
| Docker | Development/Production | < 10 min | Manual | Low |
| AWS Lambda | Serverless Production | < 15 min | Auto | Pay-per-use |
| Kubernetes | Enterprise Production | 30+ min | Auto | Variable |

## 🏠 Local Development

### Quick Start
```bash
make dev
```
This starts both backend (port 8080) and frontend (port 5173) with hot reload.

### Manual Setup
```bash
# Backend
go run cmd/main.go

# Frontend (new terminal)
cd web && npm run dev
```

### Docker Compose
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## ☁️ AWS Lambda Deployment

### Prerequisites
```bash
# Install Serverless Framework
npm install -g serverless

# Configure AWS credentials
aws configure

# Or set environment variables
export AWS_ACCESS_KEY_ID=your-key
export AWS_SECRET_ACCESS_KEY=your-secret
```

### Deploy to Development
```bash
# Deploy to dev environment
make lambda-deploy LAMBDA_STAGE=dev

# View deployment info
make lambda-info

# Check health
curl https://your-api-id.execute-api.us-east-1.amazonaws.com/dev/api/health
```

### Deploy to Production
```bash
# Deploy to production
make lambda-deploy LAMBDA_STAGE=prod LAMBDA_REGION=us-east-1

# Update function code only (faster)
make lambda-deploy LAMBDA_STAGE=prod
```

### Environment Variables
Create `serverless.yml` environment sections:

```yaml
environment:
  GIN_MODE: release
  # Add your custom env vars here
  DATABASE_URL: ${env:DATABASE_URL}
  REDIS_URL: ${env:REDIS_URL}
```

### Custom Domain
```yaml
# In serverless.yml
custom:
  customDomain:
    domainName: api.yourdomain.com
    basePath: ''
    stage: prod
    createRoute53Record: true
```

### Monitoring
```bash
# View logs
make lambda-logs LAMBDA_STAGE=prod

# Enable CloudWatch alarms
# Add to serverless.yml:
plugins:
  - serverless-plugin-aws-alerts

custom:
  alerts:
    - functionErrors
    - functionDuration
    - functionInvocations
    - functionThrottles
```

## 🐳 Docker Deployment

### Build and Run Locally
```bash
# Build image
make docker-build

# Run container
make docker-run

# View status
make docker-status

# View logs
make docker-logs
```

### Docker Registry Deployment
```bash
# Build for registry
make docker-build DOCKER_REGISTRY=your-registry.com

# Push to registry
make docker-push DOCKER_REGISTRY=your-registry.com

# Run from registry
docker run -d -p 8080:8080 your-registry.com/boilerplate-blueprint:latest
```

### Docker Compose Production
```yaml
version: '3.8'

services:
  boilerplate-blueprint:
    image: your-registry.com/boilerplate-blueprint:latest
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/api/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  # Add Redis for caching
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    restart: unless-stopped

  # Add PostgreSQL for persistence
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: boilerplate_blueprint
      POSTGRES_USER: boilerplate
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  redis_data:
  postgres_data:
```

## ☸️ Kubernetes Deployment

### Prerequisites
```bash
# Install kubectl and helm
# Set up kubeconfig for your cluster
```

### Basic Deployment
```bash
# Create namespace
kubectl create namespace boilerplate-blueprint

# Apply manifests
kubectl apply -f k8s/

# Check deployment
kubectl get pods -n boilerplate-blueprint
kubectl get services -n boilerplate-blueprint
```

### Helm Chart (Recommended)
```bash
# Add helm repo (when available)
helm repo add boilerplate-blueprint https://charts.yourdomain.com
helm install boilerplate-blueprint boilerplate-blueprint/boilerplate-blueprint

# Or install from local chart
helm install boilerplate-blueprint ./helm/boilerplate-blueprint
```

### Sample Kubernetes Manifests

**deployment.yaml:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: boilerplate-blueprint
  namespace: boilerplate-blueprint
spec:
  replicas: 3
  selector:
    matchLabels:
      app: boilerplate-blueprint
  template:
    metadata:
      labels:
        app: boilerplate-blueprint
    spec:
      containers:
      - name: boilerplate-blueprint
        image: your-registry.com/boilerplate-blueprint:latest
        ports:
        - containerPort: 8080
        env:
        - name: GIN_MODE
          value: "release"
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /api/health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /api/health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

**service.yaml:**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: boilerplate-blueprint
  namespace: boilerplate-blueprint
spec:
  selector:
    app: boilerplate-blueprint
  ports:
  - port: 80
    targetPort: 8080
  type: ClusterIP
```

**ingress.yaml:**
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: boilerplate-blueprint
  namespace: boilerplate-blueprint
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
  - hosts:
    - api.yourdomain.com
    secretName: boilerplate-blueprint-tls
  rules:
  - host: api.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: boilerplate-blueprint
            port:
              number: 80
```

## 🔄 CI/CD Pipeline

### GitHub Actions Setup

The project includes a comprehensive CI/CD pipeline:

1. **Test & Build** - Code quality and testing
2. **Security Scan** - Vulnerability scanning
3. **Docker Build** - Container image creation
4. **Deploy Staging** - Automated staging deployment
5. **Performance Test** - Load testing
6. **Deploy Production** - Manual production deployment

### Required Secrets

Set these in your GitHub repository settings:

```bash
# AWS Credentials
AWS_ACCESS_KEY_ID_STAGING
AWS_SECRET_ACCESS_KEY_STAGING
AWS_ACCESS_KEY_ID_PRODUCTION
AWS_SECRET_ACCESS_KEY_PRODUCTION

# Docker Registry
DOCKER_USERNAME
DOCKER_PASSWORD

# Slack Notifications
SLACK_WEBHOOK_URL
```

### Customizing the Pipeline

Edit `.github/workflows/ci-cd.yml`:

```yaml
# Change deployment regions
env:
  AWS_REGION: us-west-2

# Modify test commands
- name: Run tests
  run: make test-coverage

# Add custom deployment steps
- name: Custom deployment step
  run: |
    echo "Custom deployment logic here"
```

## 📊 Monitoring & Observability

### AWS Lambda Monitoring
- **CloudWatch Logs**: Automatic logging
- **CloudWatch Metrics**: Performance metrics
- **X-Ray**: Distributed tracing (optional)

### Docker Monitoring
```bash
# Container metrics
docker stats boilerplate-blueprint

# Application logs
make docker-logs
```

### Kubernetes Monitoring
```bash
# Pod metrics
kubectl top pods -n boilerplate-blueprint

# Application logs
kubectl logs -f deployment/boilerplate-blueprint -n boilerplate-blueprint

# Resource usage
kubectl describe pod <pod-name> -n boilerplate-blueprint
```

### Health Checks

All deployments include health check endpoints:

```bash
# Health check
GET /api/health

# Response
{
  "status": "healthy",
  "service": "boilerplate-blueprint",
  "version": "1.0.0"
}
```

## 🔧 Troubleshooting

### AWS Lambda Issues

**Cold Start Problems:**
```bash
# Increase memory allocation in serverless.yml
memorySize: 2048  # Increase from 1024

# Or provision concurrency
reservedConcurrency: 5
```

**Timeout Issues:**
```yaml
# Increase timeout in serverless.yml
timeout: 60  # Increase from 30 seconds
```

**Environment Variables:**
```bash
# Check Lambda environment
serverless logs --function api --stage prod --region us-east-1
```

### Docker Issues

**Port Conflicts:**
```bash
# Change host port
docker run -p 3000:8080 boilerplate-blueprint

# Or use docker-compose override
version: '3.8'
services:
  boilerplate-blueprint:
    ports:
      - "3000:8080"
```

**Memory Issues:**
```yaml
# Limit container memory
deploy:
  resources:
    limits:
      memory: 512M
    reservations:
      memory: 256M
```

### Kubernetes Issues

**Pod Crashes:**
```bash
# Check pod logs
kubectl logs <pod-name> -n boilerplate-blueprint

# Describe pod for events
kubectl describe pod <pod-name> -n boilerplate-blueprint
```

**Service Issues:**
```bash
# Check service endpoints
kubectl get endpoints -n boilerplate-blueprint

# Test service connectivity
kubectl run test --image=curlimages/curl --rm -it --restart=Never -- curl http://boilerplate-blueprint:80/api/health
```

## 🚀 Performance Optimization

### AWS Lambda
- **Memory Allocation**: Increase for better CPU performance
- **Provisioned Concurrency**: Reduce cold starts
- **Code Optimization**: Minimize bundle size
- **Connection Reuse**: Use connection pooling

### Docker
- **Multi-stage Builds**: Reduce image size
- **Layer Caching**: Optimize Dockerfile order
- **Resource Limits**: Set appropriate CPU/memory limits
- **Health Checks**: Implement proper health checks

### Kubernetes
- **Horizontal Pod Autoscaling**: Auto-scale based on CPU/memory
- **Resource Requests/Limits**: Proper resource allocation
- **Readiness/Liveness Probes**: Proper health checks
- **ConfigMaps/Secrets**: Externalize configuration

## 🔒 Security Considerations

### AWS Lambda Security
```yaml
# IAM permissions (principle of least privilege)
iam:
  role:
    statements:
      - Effect: Allow
        Action:
          - logs:CreateLogGroup
          - logs:CreateLogStream
          - logs:PutLogEvents
        Resource: "arn:aws:logs:*:*:*"
```

### Container Security
```dockerfile
# Use non-root user
FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
```

### Network Security
- Use HTTPS in production
- Implement API rate limiting
- Use security groups/firewalls
- Regular security updates

## 📈 Scaling Strategies

### AWS Lambda Scaling
- **Concurrent Executions**: Default 1000 (configurable)
- **Reserved Concurrency**: Guarantee capacity
- **Auto-scaling**: Automatic based on load

### Docker Scaling
- **Docker Swarm**: Built-in orchestration
- **Load Balancer**: Distribute traffic
- **Multiple Containers**: Horizontal scaling

### Kubernetes Scaling
- **Horizontal Pod Autoscaler**: CPU/memory based scaling
- **Cluster Autoscaler**: Node-level scaling
- **Load Balancing**: Service mesh (Istio/Linkerd)

## 💰 Cost Optimization

### AWS Lambda Costs
- **Request Pricing**: $0.20 per 1M requests
- **Duration Pricing**: $0.0000166667 per GB-second
- **Free Tier**: 1M requests + 400,000 GB-seconds/month

### Docker Costs
- **Compute**: EC2/ECS pricing
- **Storage**: EBS/EFS pricing
- **Network**: Data transfer costs

### Kubernetes Costs
- **Compute**: Node costs (EC2)
- **Storage**: Persistent volume costs
- **Network**: Load balancer costs

## 📞 Support & Resources

### Useful Commands

```bash
# Check deployment status
make lambda-info
make docker-status

# View logs
make lambda-logs
make docker-logs

# Health checks
curl http://localhost:8080/api/health

# Performance testing
make load-test
```

### Documentation Links
- [AWS Lambda Documentation](https://docs.aws.amazon.com/lambda/)
- [Serverless Framework](https://www.serverless.com/framework/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)

---

## 🎯 Deployment Checklist

### Pre-deployment
- [ ] Environment variables configured
- [ ] AWS credentials set up
- [ ] Docker registry access configured
- [ ] Domain/DNS configured
- [ ] SSL certificates ready

### Deployment
- [ ] Code tested and built successfully
- [ ] Security scans passed
- [ ] Performance benchmarks met
- [ ] Rollback plan prepared
- [ ] Monitoring/alerts configured

### Post-deployment
- [ ] Health checks passing
- [ ] Application accessible
- [ ] Logs monitoring active
- [ ] Performance monitoring active
- [ ] Backup/restore tested

**Happy Deploying! 🚀**