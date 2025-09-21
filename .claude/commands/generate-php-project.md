# Generate PHP Project Command

## Command
`/generate-php-project`

## Description
Quickly generate a PHP CodeIgniter project with MVC architecture following project specifications.

## Usage
```
/generate-php-project [project-name] [version] [database]
```

## Parameters
- `project-name` (required): Name of the project
- `version` (optional): ci3|ci4 (default: ci4)
- `database` (optional): mysql|postgres|sqlite (default: mysql)

## Examples
```bash
/generate-php-project my-webapp
/generate-php-project my-webapp ci4 mysql
/generate-php-project admin-panel ci3 postgres
```

## Implementation
```bash
#!/bin/bash
PROJECT_NAME=${1:-"my-project"}
VERSION=${2:-"ci4"}
DATABASE=${3:-"mysql"}

echo "Generating PHP CodeIgniter project: $PROJECT_NAME"
echo "Version: $VERSION"
echo "Database: $DATABASE"

# Use Task tool to generate project with specifications
claude task "Generate a PHP CodeIgniter project named '$PROJECT_NAME' using version $VERSION and $DATABASE database. Follow the exact MVC structure from project-description.md. Include all helper files, libraries, and security features. Create a downloadable ZIP file."
```

## Output
- Complete MVC project structure
- All helper and library files
- Security configurations
- Database models
- View templates
- Composer configuration
- ZIP file download link