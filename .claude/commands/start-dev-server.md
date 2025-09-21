# Start Development Server Command

## Command
`/start-dev-server`

## Description
Start the boilerplate generator web application in development mode.

## Usage
```
/start-dev-server [port]
```

## Parameters
- `port` (optional): Port number (default: 8080)

## Examples
```bash
/start-dev-server
/start-dev-server 3000
```

## Implementation
```bash
#!/bin/bash
PORT=${1:-8080}

echo "Starting Boilerplate Generator on port $PORT"

# Check if main.go exists
if [ ! -f "main.go" ]; then
    echo "Error: main.go not found. Please generate the project first."
    exit 1
fi

# Install dependencies if needed
if [ ! -f "go.sum" ]; then
    echo "Installing Go dependencies..."
    go mod tidy
fi

# Start the server
echo "Server starting on http://localhost:$PORT"
go run main.go
```

## Prerequisites
- Go installed
- Project structure generated
- Dependencies installed

## Output
- Development server running
- Web interface accessible
- Hot reload enabled (if using air)
- Console logs displayed