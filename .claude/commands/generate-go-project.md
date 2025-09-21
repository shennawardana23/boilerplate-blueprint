# Generate Go Project Command

## Command
`/generate-go-project`

## Description
Quickly generate a Go project with Clean Architecture following project specifications.

## Usage
```
/generate-go-project [project-name] [framework] [database]
```

## Parameters
- `project-name` (required): Name of the project
- `framework` (optional): gin|chi|echo|standard (default: gin)
- `database` (optional): postgres|mysql|sqlite|mongodb (default: postgres)

## Examples
```bash
/generate-go-project my-api
/generate-go-project my-api gin postgres
/generate-go-project blog-service chi mysql
```

## Implementation
```bash
#!/bin/bash
PROJECT_NAME=${1:-"my-project"}
FRAMEWORK=${2:-"gin"}
DATABASE=${3:-"postgres"}

echo "Generating Go project: $PROJECT_NAME"
echo "Framework: $FRAMEWORK"
echo "Database: $DATABASE"

# Use Task tool to generate project with specifications
claude task "Generate a Go project named '$PROJECT_NAME' using $FRAMEWORK framework and $DATABASE database. Follow the exact Clean Architecture structure from project-description.md. Include all 17 utility packages and create a downloadable ZIP file."
```

## Output
- Complete Go project structure
- All utility packages implemented
- Configuration files
- Docker setup
- README with instructions
- ZIP file download link