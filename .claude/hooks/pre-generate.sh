#!/bin/bash
# Pre-generation hook for boilerplate projects

echo "🔍 Pre-generation validation starting..."

# Validate project name
if [ -z "$PROJECT_NAME" ]; then
    echo "❌ Error: PROJECT_NAME environment variable is required"
    exit 1
fi

# Check for invalid characters in project name
if [[ ! $PROJECT_NAME =~ ^[a-zA-Z][a-zA-Z0-9_-]*$ ]]; then
    echo "❌ Error: Project name must start with a letter and contain only letters, numbers, hyphens, and underscores"
    exit 1
fi

# Validate language selection
if [ -z "$LANGUAGE" ]; then
    echo "❌ Error: LANGUAGE environment variable is required (go|php)"
    exit 1
fi

if [[ "$LANGUAGE" != "go" && "$LANGUAGE" != "php" ]]; then
    echo "❌ Error: LANGUAGE must be either 'go' or 'php'"
    exit 1
fi

# Language-specific validations
if [ "$LANGUAGE" = "go" ]; then
    # Validate Go framework
    if [ -n "$FRAMEWORK" ] && [[ ! "$FRAMEWORK" =~ ^(gin|chi|echo|standard)$ ]]; then
        echo "❌ Error: Go FRAMEWORK must be one of: gin, chi, echo, standard"
        exit 1
    fi
    
    # Validate Go database
    if [ -n "$DATABASE" ] && [[ ! "$DATABASE" =~ ^(postgres|mysql|sqlite|mongodb)$ ]]; then
        echo "❌ Error: Go DATABASE must be one of: postgres, mysql, sqlite, mongodb"
        exit 1
    fi
fi

if [ "$LANGUAGE" = "php" ]; then
    # Validate PHP version
    if [ -n "$VERSION" ] && [[ ! "$VERSION" =~ ^(ci3|ci4)$ ]]; then
        echo "❌ Error: PHP VERSION must be one of: ci3, ci4"
        exit 1
    fi
    
    # Validate PHP database
    if [ -n "$DATABASE" ] && [[ ! "$DATABASE" =~ ^(mysql|postgres|sqlite)$ ]]; then
        echo "❌ Error: PHP DATABASE must be one of: mysql, postgres, sqlite"
        exit 1
    fi
fi

# Check available disk space (minimum 100MB)
AVAILABLE_SPACE=$(df . | awk 'NR==2 {print $4}')
if [ "$AVAILABLE_SPACE" -lt 102400 ]; then
    echo "⚠️  Warning: Low disk space. At least 100MB recommended for project generation"
fi

echo "✅ Pre-generation validation passed"
echo "📋 Project: $PROJECT_NAME ($LANGUAGE)"
[ -n "$FRAMEWORK" ] && echo "🔧 Framework: $FRAMEWORK"
[ -n "$VERSION" ] && echo "🔧 Version: $VERSION"  
[ -n "$DATABASE" ] && echo "💾 Database: $DATABASE"

exit 0