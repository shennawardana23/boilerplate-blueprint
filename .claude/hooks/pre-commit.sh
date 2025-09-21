#!/bin/bash
# Pre-commit hook for boilerplate generator project

echo "🔍 Pre-commit validation starting..."

# Check for Go files and run gofmt
if find . -name "*.go" -not -path "./.git/*" | head -1 | grep -q .; then
    echo "🐹 Checking Go code formatting..."
    
    # Check if gofmt would make changes
    UNFORMATTED=$(gofmt -l . | grep -v ".git")
    if [ -n "$UNFORMATTED" ]; then
        echo "❌ Error: The following Go files need formatting:"
        echo "$UNFORMATTED"
        echo "Run 'gofmt -w .' to fix formatting"
        exit 1
    fi
    
    # Run go vet if go.mod exists
    if [ -f "go.mod" ]; then
        echo "🔍 Running go vet..."
        if ! go vet ./...; then
            echo "❌ Error: go vet found issues"
            exit 1
        fi
    fi
    
    echo "✅ Go code formatting and vet checks passed"
fi

# Check for PHP files and basic syntax
if find . -name "*.php" -not -path "./.git/*" | head -1 | grep -q .; then
    echo "🐘 Checking PHP syntax..."
    
    # Check PHP syntax for all PHP files
    find . -name "*.php" -not -path "./.git/*" -exec php -l {} \; | grep -v "No syntax errors detected" | grep -v "^$"
    
    if [ $? -eq 0 ]; then
        echo "❌ Error: PHP syntax errors found"
        exit 1
    fi
    
    echo "✅ PHP syntax checks passed"
fi

# Check for large files (>10MB)
LARGE_FILES=$(find . -type f -size +10M -not -path "./.git/*")
if [ -n "$LARGE_FILES" ]; then
    echo "⚠️  Warning: Large files detected:"
    echo "$LARGE_FILES"
fi

# Check for sensitive files
SENSITIVE_PATTERNS=("*.key" "*.pem" "*.p12" ".env" "password" "secret" "token")
for pattern in "${SENSITIVE_PATTERNS[@]}"; do
    FOUND=$(find . -name "$pattern" -not -path "./.git/*")
    if [ -n "$FOUND" ]; then
        echo "⚠️  Warning: Potential sensitive files found:"
        echo "$FOUND"
        echo "Please ensure these files are properly protected or excluded"
    fi
done

# Check for proper README.md
if [ ! -f "README.md" ]; then
    echo "⚠️  Warning: README.md not found. Consider adding project documentation"
fi

# Validate .gitignore patterns
if [ -f ".gitignore" ]; then
    # Check if common patterns are present
    COMMON_PATTERNS=("*.log" "*.tmp" ".env" "node_modules/" "vendor/")
    for pattern in "${COMMON_PATTERNS[@]}"; do
        if ! grep -q "$pattern" .gitignore; then
            echo "💡 Suggestion: Consider adding '$pattern' to .gitignore"
        fi
    done
fi

# Check for TODO/FIXME comments
TODO_COUNT=$(find . -type f \( -name "*.go" -o -name "*.php" -o -name "*.js" -o -name "*.md" \) -not -path "./.git/*" -exec grep -l "TODO\|FIXME" {} \; | wc -l)
if [ $TODO_COUNT -gt 0 ]; then
    echo "📋 Info: Found $TODO_COUNT files with TODO/FIXME comments"
fi

echo "✅ Pre-commit validation completed"

exit 0