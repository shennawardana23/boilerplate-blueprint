# System Commands (Darwin/macOS)

## File Operations
```bash
# Directory listing
ls -la                     # List all files with details
find . -name "*.go"        # Find Go files recursively
find . -type d             # Find directories only

# File searching
grep -r "pattern" .        # Search for pattern in files
grep -r --include="*.go"   # Search only in Go files
ripgrep "pattern" .        # Faster alternative (if installed)

# File operations
cat filename               # Display file contents
head -n 20 filename        # First 20 lines
tail -n 20 filename        # Last 20 lines
tail -f logfile           # Follow log file updates
```

## Process Management
```bash
# Process monitoring
ps aux | grep "go"         # Find Go processes
lsof -i :8080             # Check what's using port 8080
kill -9 PID               # Force kill process
pkill -f "process-name"   # Kill by process name

# System monitoring
top                       # System monitor
htop                      # Better system monitor (if installed)
df -h                     # Disk usage
du -sh directory          # Directory size
```

## Network Operations
```bash
# Port and network checking
netstat -an | grep 8080   # Check port 8080 status
curl http://localhost:8080/health  # Test endpoint
curl -X POST -H "Content-Type: application/json" -d '{"key":"value"}' http://localhost:8080/api/endpoint

# DNS and connectivity
ping google.com           # Network connectivity test
nslookup domain.com       # DNS lookup
```

## Git Operations
```bash
# Repository status
git status                # Working tree status
git log --oneline -10     # Last 10 commits
git branch -a             # All branches
git remote -v             # Remote repositories

# File operations
git add .                 # Stage all changes
git commit -m "message"   # Commit changes
git push origin branch    # Push to remote
git pull origin branch    # Pull from remote

# Useful Git commands
git diff                  # Show unstaged changes
git diff --staged         # Show staged changes
git stash                 # Stash changes
git stash pop             # Apply stashed changes
```

## Package Management
```bash
# Go modules
go mod init module-name   # Initialize Go module
go mod tidy               # Clean up dependencies
go mod download           # Download dependencies
go mod verify             # Verify dependencies

# Node.js/NPM
npm install               # Install dependencies
npm update               # Update dependencies
npm audit                # Security audit
npm run script-name      # Run package.json script
```

## Environment Management
```bash
# Environment variables
env                      # Show all environment variables
echo $GOPATH            # Show specific variable
export VAR=value        # Set environment variable
source .env             # Load .env file (with appropriate tool)

# Path management
which go                # Find Go binary location
echo $PATH             # Show PATH variable
export PATH=$PATH:/new/path  # Add to PATH
```

## Development Utilities
```bash
# File watching
fswatch -o . | xargs -n1 -I{} echo "Files changed"  # Watch for file changes

# JSON processing
cat file.json | jq '.'   # Pretty print JSON (if jq installed)

# Archive operations  
tar -czf archive.tar.gz directory/  # Create tar.gz
unzip file.zip           # Extract zip file
zip -r archive.zip directory/       # Create zip file

# Text processing
sed 's/old/new/g' file   # Replace text in file
awk '{print $1}' file    # Print first column
sort file                # Sort file contents
uniq file               # Remove duplicates
```