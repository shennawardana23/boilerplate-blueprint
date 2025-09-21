# 🏗️ Boilerplate Blueprint - Architecture Diagrams

This document contains Mermaid.js diagrams illustrating the system architecture, data flow, and component interactions.

## 📊 System Architecture Overview

```mermaid
graph TB
    subgraph "User Layer"
        U[👤 User]
    end

    subgraph "Presentation Layer"
        FE[🌐 Vue.js Frontend<br/>Vue 3 + Vite]
        SW[📱 Service Worker<br/>PWA Support]
    end

    subgraph "API Gateway Layer"
        AG[🚪 API Gateway<br/>AWS Lambda]
        NG[🌐 Nginx<br/>Docker/K8s]
    end

    subgraph "Application Layer"
        LB[⚖️ Load Balancer<br/>AWS ALB]
        APP[🐹 Go Application<br/>Gin Framework]
    end

    subgraph "Business Logic Layer"
        SVC[🔧 Services Layer<br/>Project/Chat/Template]
        CTRL[🎮 Controllers Layer<br/>HTTP Handlers]
    end

    subgraph "Data Layer"
        CACHE[(💾 Redis Cache<br/>Optional)]
        DB[(🗄️ PostgreSQL<br/>Optional)]
        MEM[(🧠 In-Memory Store<br/>Current)]
    end

    subgraph "Infrastructure Layer"
        DKR[🐳 Docker Container]
        LBD[☁️ AWS Lambda]
        K8S[☸️ Kubernetes Pod]
    end

    U --> FE
    FE --> SW
    FE --> AG
    FE --> NG
    AG --> LB
    NG --> LB
    LB --> APP
    APP --> CTRL
    CTRL --> SVC
    SVC --> CACHE
    SVC --> DB
    SVC --> MEM
    APP --> DKR
    APP --> LBD
    APP --> K8S

    style FE fill:#e1f5fe
    style APP fill:#f3e5f5
    style SVC fill:#e8f5e8
    style DB fill:#fff3e0
```

## 🔄 Request Flow Control

### HTTP Request Journey

```mermaid
sequenceDiagram
    participant U as User
    participant FE as Frontend (Vue.js)
    participant AG as API Gateway
    participant APP as Go Application
    participant CTRL as Controller
    participant SVC as Service
    participant DB as Database/Store

    U->>FE: User Action (e.g., Create Project)
    FE->>AG: HTTP Request (POST /api/projects)
    AG->>APP: Forward Request
    APP->>CTRL: Route to Handler

    CTRL->>CTRL: Input Validation
    CTRL->>SVC: Business Logic Call

    SVC->>SVC: Process Request
    SVC->>DB: Data Operation
    DB-->>SVC: Data Response

    SVC-->>CTRL: Processed Result
    CTRL-->>APP: HTTP Response
    APP-->>AG: Response
    AG-->>FE: JSON Response
    FE-->>U: UI Update

    Note over CTRL,SVC: Error Handling<br/>Rollback if needed
```

### AWS Lambda Request Flow

```mermaid
sequenceDiagram
    participant U as User
    participant FE as Frontend
    participant AG as API Gateway
    participant LBD as Lambda Function
    participant RT as Router
    participant CTRL as Controller
    participant SVC as Service

    U->>FE: API Request
    FE->>AG: HTTP Request
    AG->>LBD: Invoke Lambda

    LBD->>RT: Route Request
    RT->>CTRL: Handler Call
    CTRL->>SVC: Service Logic
    SVC-->>CTRL: Response
    CTRL-->>RT: HTTP Response
    RT-->>LBD: Lambda Response
    LBD-->>AG: API Response
    AG-->>FE: JSON Response
```

## 🗂️ Data Lineage: Project Creation Flow

### Variable Tracing: `projectName` Parameter

```mermaid
flowchart TD
    subgraph "Input Layer"
        A[📝 User Input<br/>projectName: "my-app"]
        B[🌐 Frontend Form<br/>this.projectName]
    end

    subgraph "Network Layer"
        C[📡 HTTP Request<br/>POST /api/projects<br/>Body: {name: "my-app"}]
        D[🚪 API Gateway<br/>event.body.projectName]
    end

    subgraph "Application Layer"
        E[🎮 Controller<br/>req.Body → JSON Parse<br/>projectRequest.Name]
        F[🔧 Service<br/>ProjectService.CreateProject<br/>req.Name]
    end

    subgraph "Business Logic"
        G[📋 Validation<br/>validateProjectName(name)<br/>✓ Length, Format, Uniqueness]
        H[🏗️ Project Struct<br/>project := &Project{<br/>Name: name,<br/>...}]
    end

    subgraph "Data Layer"
        I[💾 In-Memory Store<br/>projects[projectID] = project]
        J[🗄️ Database (Future)<br/>INSERT INTO projects<br/>VALUES (name, ...)]
    end

    subgraph "Response Layer"
        K[📤 Service Response<br/>return project, nil]
        L[🎮 Controller Response<br/>c.JSON(200, project)]
        M[🌐 Frontend Update<br/>this.currentProject = response.data]
    end

    A --> B
    B --> C
    C --> D
    D --> E
    E --> F
    F --> G
    G --> H
    H --> I
    I --> J
    I --> K
    K --> L
    L --> M

    style A fill:#e3f2fd
    style B fill:#e3f2fd
    style G fill:#fff9c4
    style H fill:#fff9c4
    style I fill:#fff9c4
```

### Variable State Transitions

```mermaid
stateDiagram-v2
    [*] --> InputReceived: User types "my-app"
    InputReceived --> FrontendValidation: Vue.js validates format
    FrontendValidation --> NetworkTransmission: Axios POST request
    NetworkTransmission --> APIGateway: AWS API Gateway receives
    APIGateway --> LambdaInvocation: Lambda function invoked
    LambdaInvocation --> JSONParsing: Go JSON unmarshaling
    JSONParsing --> StructCreation: ProjectRequest struct
    StructCreation --> ServiceValidation: Business rule validation
    ServiceValidation --> EntityCreation: Project entity created
    EntityCreation --> Persistence: Stored in memory/database
    Persistence --> ResponseFormatting: JSON response created
    ResponseFormatting --> NetworkResponse: HTTP 200 response
    NetworkResponse --> FrontendUpdate: Vue.js state update
    FrontendUpdate --> [*]: UI reflects new project
```

## 🏛️ Component Architecture

### Service Component Diagram

```mermaid
graph TB
    subgraph "Entry Points"
        MAIN[📄 main.go<br/>Application Entry]
        LAMBDA[☁️ lambda_handler.go<br/>AWS Lambda Entry]
    end

    subgraph "HTTP Layer"
        ROUTER[🛣️ Router<br/>Gin Router]
        CORS[🔒 CORS Middleware<br/>Cross-Origin]
        LOG[📝 Logger Middleware<br/>Request Logging]
        AUTH[🔐 Auth Middleware<br/>Future JWT]
    end

    subgraph "Controller Layer"
        PROJ_CTRL[🎮 ProjectController<br/>handlers.go]
        CHAT_CTRL[💬 ChatController<br/>handlers.go]
        HEALTH_CTRL[❤️ HealthController<br/>handlers.go]
    end

    subgraph "Service Layer"
        PROJ_SVC[🔧 ProjectService<br/>project.go]
        CHAT_SVC[🤖 ChatService<br/>chat.go]
        TEMP_SVC[📄 TemplateService<br/>template.go]
    end

    subgraph "Model Layer"
        PROJ_MODEL[📋 Project Model<br/>project.go]
        CHAT_MODEL[💭 Chat Model<br/>chat.go]
    end

    subgraph "Infrastructure Layer"
        STORE[💾 In-Memory Store<br/>Thread-Safe Map]
        ZIP[📦 ZIP Generator<br/>archive/zip]
        UUID[🆔 UUID Generator<br/>google/uuid]
    end

    MAIN --> ROUTER
    LAMBDA --> ROUTER
    ROUTER --> CORS
    CORS --> LOG
    LOG --> AUTH
    AUTH --> PROJ_CTRL
    AUTH --> CHAT_CTRL
    AUTH --> HEALTH_CTRL

    PROJ_CTRL --> PROJ_SVC
    CHAT_CTRL --> CHAT_SVC
    HEALTH_CTRL --> TEMP_SVC

    PROJ_SVC --> PROJ_MODEL
    CHAT_SVC --> CHAT_MODEL
    TEMP_SVC --> PROJ_MODEL

    PROJ_SVC --> STORE
    CHAT_SVC --> STORE
    PROJ_SVC --> ZIP
    PROJ_SVC --> UUID

    style MAIN fill:#e8f5e8
    style LAMBDA fill:#e8f5e8
    style PROJ_SVC fill:#fff3e0
    style CHAT_SVC fill:#fff3e0
    style TEMP_SVC fill:#fff3e0
```

### Data Flow Architecture

```mermaid
graph LR
    subgraph "Input Sources"
        API[🌐 REST API<br/>HTTP Requests]
        UI[🖥️ Web UI<br/>Vue Components]
        CLI[💻 Command Line<br/>Future Feature]
    end

    subgraph "Processing Pipeline"
        VAL[✅ Validation<br/>Input Sanitization]
        AUTH[🔐 Authentication<br/>JWT/Future]
        CTRL[🎮 Controller<br/>Route Handling]
        SVC[🔧 Service<br/>Business Logic]
        MODEL[📋 Model<br/>Data Transformation]
    end

    subgraph "Storage Layer"
        MEM[🧠 In-Memory<br/>Current Default]
        CACHE[💾 Redis<br/>Optional Cache]
        DB[🗄️ PostgreSQL<br/>Optional DB]
    end

    subgraph "Output Destinations"
        RESP[📤 HTTP Response<br/>JSON API]
        FILE[📁 File Download<br/>ZIP Export]
        WS[🔌 WebSocket<br/>Real-time Updates]
    end

    API --> VAL
    UI --> VAL
    CLI --> VAL
    VAL --> AUTH
    AUTH --> CTRL
    CTRL --> SVC
    SVC --> MODEL
    MODEL --> MEM
    MODEL --> CACHE
    MODEL --> DB
    MEM --> RESP
    CACHE --> RESP
    DB --> RESP
    SVC --> FILE
    SVC --> WS

    style VAL fill:#e3f2fd
    style SVC fill:#fff3e0
    style MEM fill:#e8f5e8
```

## 🔄 CI/CD Pipeline Flow

### GitHub Actions Workflow

```mermaid
graph TD
    A[👥 Push/PR] --> B[Test & Build Job]
    A --> C[Security Scan Job]
    A --> D[Docker Build Job]

    B --> E{Quality Gate<br/>All Tests Pass?}
    C --> E
    D --> E

    E -->|✅ Pass| F[Deploy Staging]
    E -->|❌ Fail| G[Block Deployment]

    F --> H[Performance Test]
    H --> I{Performance OK?}
    I -->|✅| J[Deploy Production<br/>Manual Approval]
    I -->|❌| K[Rollback Staging]

    J --> L[Production Health Check]
    L --> M{Health OK?}
    M -->|✅| N[🎉 Success Notification]
    M -->|❌| O[Rollback Production]

    G --> P[❌ Failure Notification]
    K --> P
    O --> P

    style B fill:#e3f2fd
    style F fill:#e8f5e8
    style J fill:#fff3e0
    style N fill:#e8f5e8
    style P fill:#ffebee
```

## 🌐 Deployment Architecture

### Multi-Environment Deployment

```mermaid
graph TB
    subgraph "Development"
        DEV_FE[🌐 Frontend<br/>localhost:5173]
        DEV_BE[🐹 Backend<br/>localhost:8080]
        DEV_DB[(💾 Local Storage)]
    end

    subgraph "Staging"
        STG_AG[🚪 API Gateway<br/>Staging]
        STG_LBD[☁️ Lambda<br/>Staging]
        STG_CACHE[(💾 Redis<br/>Staging)]
    end

    subgraph "Production"
        PROD_AG[🚪 API Gateway<br/>Production]
        PROD_LBD[☁️ Lambda<br/>Production]
        PROD_CACHE[(💾 Redis<br/>Production)]
        PROD_DB[(🗄️ PostgreSQL<br/>Production)]
    end

    subgraph "Infrastructure"
        CF[🌐 CloudFront<br/>CDN]
        WAF[🛡️ WAF<br/>Security]
        MONITOR[📊 CloudWatch<br/>Monitoring]
    end

    DEV_FE --> DEV_BE
    DEV_BE --> DEV_DB

    STG_AG --> STG_LBD
    STG_LBD --> STG_CACHE

    PROD_AG --> PROD_LBD
    PROD_LBD --> PROD_CACHE
    PROD_LBD --> PROD_DB

    PROD_AG --> CF
    CF --> WAF
    PROD_LBD --> MONITOR

    style DEV_FE fill:#e3f2fd
    style STG_LBD fill:#fff3e0
    style PROD_LBD fill:#e8f5e8
```

## 🔐 Security Architecture

### Request Security Flow

```mermaid
flowchart LR
    subgraph "External"
        USER[👤 User Request]
        ATTACK[🚫 Malicious Request]
    end

    subgraph "Edge Security"
        WAF[🛡️ Web Application Firewall<br/>AWS WAF]
        RATE[⚡ Rate Limiting<br/>API Gateway]
        CORS[🔒 CORS Policy<br/>Origin Validation]
    end

    subgraph "Application Security"
        AUTH[🔐 Authentication<br/>JWT Validation]
        AUTHZ[👮 Authorization<br/>Role-based Access]
        VALID[✅ Input Validation<br/>Sanitization]
    end

    subgraph "Data Security"
        ENCRYPT[🔐 Data Encryption<br/>TLS in Transit]
        AUDIT[📝 Audit Logging<br/>CloudWatch]
        BACKUP[💾 Backup Security<br/>Encrypted Storage]
    end

    USER --> WAF
    ATTACK --> WAF
    WAF --> RATE
    RATE --> CORS
    CORS --> AUTH
    AUTH --> AUTHZ
    AUTHZ --> VALID
    VALID --> ENCRYPT
    ENCRYPT --> AUDIT
    AUDIT --> BACKUP

    style USER fill:#e8f5e8
    style ATTACK fill:#ffebee
    style WAF fill:#fff3e0
    style VALID fill:#e3f2fd
```

## 📈 Performance Monitoring

### Application Metrics Flow

```mermaid
graph LR
    subgraph "Application"
        APP[🐹 Go Application]
        METRICS[📊 Metrics Collection<br/>Custom Metrics]
    end

    subgraph "AWS Services"
        CW[📊 CloudWatch<br/>Metrics & Logs]
        XRAY[🔍 X-Ray<br/>Distributed Tracing]
        ALARMS[🚨 CloudWatch Alarms<br/>Alerting]
    end

    subgraph "Monitoring Stack"
        DASH[📊 Dashboards<br/>Grafana/Kibana]
        ALERTS[📢 Alert Manager<br/>Notifications]
        LOGS[📝 Log Aggregation<br/>ELK Stack]
    end

    subgraph "Response"
        DEV[👥 Development Team]
        OPS[⚙️ Operations Team]
        AUTO[🤖 Auto-scaling<br/>Lambda Concurrency]
    end

    APP --> METRICS
    METRICS --> CW
    METRICS --> XRAY
    CW --> ALARMS
    ALARMS --> AUTO
    CW --> DASH
    XRAY --> DASH
    ALARMS --> ALERTS
    ALERTS --> DEV
    ALERTS --> OPS

    style APP fill:#e3f2fd
    style ALARMS fill:#fff3e0
    style AUTO fill:#e8f5e8
```

## 🎯 Data Lineage Example: Chat Message Processing

### Message Processing Pipeline

```mermaid
flowchart TD
    A[💬 User Input<br/>message: "Create Go API"] --> B[🌐 Frontend<br/>chatStore.sendMessage()]
    B --> C[📡 HTTP POST<br/>/api/chat/message]
    C --> D[🚪 API Gateway<br/>Lambda Trigger]

    D --> E[☁️ Lambda Handler<br/>lambda_handler.go]
    E --> F[🛣️ Router<br/>gin.Router]
    F --> G[🎮 ChatController<br/>ProcessMessage()]

    G --> H[🔧 ChatService<br/>ProcessMessage()]
    H --> I[🤖 AI Processing<br/>generateRuleBasedResponse()]
    I --> J[📝 Message Creation<br/>ChatMessage struct]

    J --> K[💾 Storage<br/>conversations map]
    K --> L[📋 Response Formatting<br/>ChatResponse struct]
    L --> M[🎮 Controller Response<br/>c.JSON(200, response)]

    M --> N[☁️ Lambda Response<br/>events.APIGatewayV2HTTPResponse]
    N --> O[🚪 API Gateway Response]
    O --> P[🌐 Frontend Update<br/>messages.push(response)]

    style A fill:#e3f2fd
    style I fill:#fff3e0
    style J fill:#fff9c4
    style K fill:#fff9c4
    style P fill:#e8f5e8
```

## 📚 Component Interaction Details

### Service Dependencies

```mermaid
graph TD
    A[🐹 main.go] --> B[🎮 Controllers]
    A --> C[🔧 Services]
    A --> D[📋 Models]

    B --> C
    C --> D
    C --> E[💾 Storage Layer]
    C --> F[📦 Utilities]

    subgraph "Controllers"
        B1[ProjectController]
        B2[ChatController]
        B3[HealthController]
    end

    subgraph "Services"
        C1[ProjectService]
        C2[ChatService]
        C3[TemplateService]
    end

    subgraph "Models"
        D1[Project]
        D2[ProjectRequest]
        D3[ChatMessage]
        D4[ChatRequest]
    end

    subgraph "Storage"
        E1[In-Memory Map]
        E2[PostgreSQL]
        E3[Redis Cache]
    end

    subgraph "Utilities"
        F1[UUID Generator]
        F2[ZIP Creator]
        F3[JSON Handler]
        F4[Time Utilities]
    end

    B1 --> C1
    B2 --> C2
    B3 --> C3

    C1 --> D1
    C1 --> D2
    C2 --> D3
    C2 --> D4

    C1 --> E1
    C1 --> E2
    C1 --> E3

    C1 --> F1
    C1 --> F2
    C2 --> F3
    C --> F4

    style A fill:#e8f5e8
    style C fill:#fff3e0
    style E fill:#fff9c4
```

## 🔄 Error Handling Flow

### Exception Propagation

```mermaid
flowchart TD
    A[🎮 Controller] --> B{Input Validation}
    B -->|❌ Invalid| C[Return 400 Bad Request]
    B -->|✅ Valid| D[Call Service]

    D --> E{Service Logic}
    E -->|❌ Error| F[Log Error]
    F --> G[Return Error Response]
    E -->|✅ Success| H[Format Response]

    H --> I[Return 200 OK]

    C --> J[📊 Error Metrics]
    G --> J
    I --> K[📊 Success Metrics]

    J --> L[📊 CloudWatch Metrics]
    K --> L

    L --> M{Error Rate > Threshold?}
    M -->|Yes| N[🚨 Trigger Alert]
    M -->|No| O[Continue Normal Operation]

    style C fill:#ffebee
    style G fill:#ffebee
    style N fill:#ffebee
    style I fill:#e8f5e8
```

---

## 📖 Diagram Legend

| Symbol | Meaning |
|--------|---------|
| 👤 | User/Client |
| 🌐 | Frontend/Web |
| 🚪 | API Gateway |
| 🐹 | Go Application |
| 🎮 | Controller Layer |
| 🔧 | Service Layer |
| 📋 | Model/Data Layer |
| 💾 | Storage/Database |
| ☁️ | Cloud/AWS Lambda |
| 🐳 | Docker Container |
| ☸️ | Kubernetes |
| ⚖️ | Load Balancer |
| 🔒 | Security/Middleware |
| 📊 | Monitoring/Metrics |
| ✅ | Success/Validation |
| ❌ | Error/Failure |

## 🎯 Reading Guide for New Developers

### 1. **System Architecture Overview**
   - Start here to understand the high-level system components
   - Follow the data flow from user to database
   - Understand deployment options (Lambda vs Docker vs K8s)

### 2. **Request Flow Control**
   - Trace how HTTP requests journey through the system
   - Understand middleware layers and error handling
   - See how AWS Lambda differs from traditional server deployment

### 3. **Data Lineage Diagrams**
   - Follow specific data variables through the entire pipeline
   - Understand state transformations and validation steps
   - See how business logic processes user inputs

### 4. **Component Architecture**
   - Understand service dependencies and interactions
   - See how layers communicate with each other
   - Identify where to add new features or modify existing ones

### 5. **CI/CD Pipeline**
   - Understand automated testing and deployment
   - See quality gates and approval processes
   - Follow how changes move from development to production

### 6. **Error Handling Flow**
   - Understand how errors propagate through the system
   - See monitoring and alerting mechanisms
   - Learn debugging and troubleshooting approaches

This documentation provides a comprehensive visual guide to the Boilerplate Blueprint system architecture, making it easier for new developers to understand the codebase structure, data flow, and deployment processes.