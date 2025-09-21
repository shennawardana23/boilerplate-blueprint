#!/bin/bash
# Post-generation hook for boilerplate projects

echo "🔧 Post-generation processing starting..."

PROJECT_DIR="generated/${PROJECT_NAME}"

# Verify project structure was created
if [ ! -d "$PROJECT_DIR" ]; then
    echo "❌ Error: Project directory not found at $PROJECT_DIR"
    exit 1
fi

# Language-specific post-processing
if [ "$LANGUAGE" = "go" ]; then
    echo "🐹 Processing Go project..."
    
    # Verify go.mod exists
    if [ ! -f "$PROJECT_DIR/go.mod" ]; then
        echo "❌ Error: go.mod not found in Go project"
        exit 1
    fi
    
    # Verify main.go exists  
    if [ ! -f "$PROJECT_DIR/main.go" ] && [ ! -f "$PROJECT_DIR/cmd/main.go" ]; then
        echo "❌ Error: main.go not found in Go project"
        exit 1
    fi
    
    # Check for required utility packages
    UTIL_PACKAGES=("authentication" "cache" "common" "constants" "converter" "date" "datatype" "encryption" "exception" "exceptioncode" "helper" "httphelper" "json" "logger" "password" "queryhelper" "sort" "template" "validator" "alert")
    
    MISSING_PACKAGES=()
    for package in "${UTIL_PACKAGES[@]}"; do
        if [ ! -d "$PROJECT_DIR/internal/util/$package" ]; then
            MISSING_PACKAGES+=("$package")
        fi
    done
    
    if [ ${#MISSING_PACKAGES[@]} -gt 0 ]; then
        echo "⚠️  Warning: Missing utility packages: ${MISSING_PACKAGES[*]}"
    else
        echo "✅ All 17 utility packages present"
    fi
    
    # Try to build the project
    cd "$PROJECT_DIR"
    if go mod tidy && go build .; then
        echo "✅ Go project builds successfully"
    else
        echo "⚠️  Warning: Go project has build issues"
    fi
    cd - > /dev/null
fi

if [ "$LANGUAGE" = "php" ]; then
    echo "🐘 Processing PHP project..."
    
    # Verify index.php exists
    if [ ! -f "$PROJECT_DIR/index.php" ]; then
        echo "❌ Error: index.php not found in PHP project"
        exit 1
    fi
    
    # Verify application structure
    if [ ! -d "$PROJECT_DIR/application" ]; then
        echo "❌ Error: application directory not found"
        exit 1
    fi
    
    # Check for required MVC components
    REQUIRED_DIRS=("controllers" "models" "views" "config" "helpers" "libraries")
    MISSING_DIRS=()
    
    for dir in "${REQUIRED_DIRS[@]}"; do
        if [ ! -d "$PROJECT_DIR/application/$dir" ]; then
            MISSING_DIRS+=("$dir")
        fi
    done
    
    if [ ${#MISSING_DIRS[@]} -gt 0 ]; then
        echo "⚠️  Warning: Missing MVC directories: ${MISSING_DIRS[*]}"
    else
        echo "✅ Complete MVC structure present"
    fi
    
    # Verify composer.json if present
    if [ -f "$PROJECT_DIR/composer.json" ]; then
        echo "✅ Composer configuration found"
    fi
fi

# Generate README if missing
if [ ! -f "$PROJECT_DIR/README.md" ]; then
    echo "📝 Generating README.md..."
    cat > "$PROJECT_DIR/README.md" << EOF
# $PROJECT_NAME

Generated with AI-Powered Boilerplate Generator

## Language: $LANGUAGE
$([ -n "$FRAMEWORK" ] && echo "## Framework: $FRAMEWORK")
$([ -n "$VERSION" ] && echo "## Version: $VERSION")
$([ -n "$DATABASE" ] && echo "## Database: $DATABASE")

## Quick Start

1. Navigate to the project directory
2. Install dependencies
3. Configure environment variables
4. Run the application

Generated on: $(date)
EOF
fi

# Set proper permissions
find "$PROJECT_DIR" -type f -name "*.sh" -exec chmod +x {} \;

# Calculate project size
PROJECT_SIZE=$(du -sh "$PROJECT_DIR" | cut -f1)
FILE_COUNT=$(find "$PROJECT_DIR" -type f | wc -l)

echo "✅ Post-generation processing completed"
echo "📁 Project size: $PROJECT_SIZE"
echo "📄 Files created: $FILE_COUNT"
echo "📍 Location: $PROJECT_DIR"

exit 0