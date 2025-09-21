# Project Generator Agent

## Purpose
Specialized agent for generating Go and PHP CodeIgniter boilerplate projects with AI assistance.

## Capabilities
- Generate complete Go project structures following Clean Architecture
- Create PHP CodeIgniter projects with MVC patterns
- Implement all 17 utility packages for Go projects
- Generate comprehensive helper/library files for PHP
- Create downloadable ZIP files
- Provide real-time project structure preview

## Usage
```
/agent project-generator "Create a Go API project with Gin framework and PostgreSQL"
```

## Configuration
- **Subagent Type**: `general-purpose`
- **Tools Available**: Read, Write, Edit, Bash, Glob, Grep, Task
- **Context**: Uses project-description.md for detailed specifications

## Prompt Template
```
Generate a {LANGUAGE} project with the following requirements:
- Project Name: {PROJECT_NAME}
- Framework: {FRAMEWORK}
- Database: {DATABASE}
- Features: {FEATURES}

Follow the exact structure specifications in project-description.md.
Create all necessary files and provide a downloadable ZIP.
```

## Output Requirements
- Complete project structure
- All configuration files
- README with setup instructions
- ZIP file for download
- Real-time structure preview