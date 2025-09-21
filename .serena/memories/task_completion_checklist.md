# Task Completion Checklist

## Before Marking Tasks Complete

### Go Backend Development
- [ ] **Code Quality**: Run `make fmt` to format code
- [ ] **Static Analysis**: Run `make vet` for go vet checks
- [ ] **Linting**: Run `make lint` (if golangci-lint available)
- [ ] **Testing**: Run `make test` to ensure all tests pass
- [ ] **Build Verification**: Run `make build` to ensure compilation

### Frontend Development
- [ ] **Linting**: Run `npm run lint` in web/ directory
- [ ] **Build Check**: Run `npm run build` to verify production build
- [ ] **Type Safety**: Ensure no JavaScript errors in console

### Full Stack Integration
- [ ] **Backend Running**: Verify `make run` starts successfully
- [ ] **Frontend Running**: Verify `cd web && npm run dev` starts successfully
- [ ] **CORS Configuration**: Verify frontend can communicate with backend
- [ ] **API Endpoints**: Test API endpoints return expected responses

### Quality Assurance
- [ ] **Comprehensive Check**: Run `make check-all` for complete quality validation
- [ ] **Test Coverage**: Run `make test-coverage` and review coverage report
- [ ] **Error Handling**: Verify proper error responses and logging
- [ ] **Environment Variables**: Check `.env` configuration is working

### Production Readiness
- [ ] **Build Artifacts**: Generate production builds with `make build-all`
- [ ] **Documentation**: Update relevant documentation if needed
- [ ] **Dependencies**: Verify `go mod tidy` and `npm install` work clean

## Commands to Run After Task Completion
```bash
# Always run these after significant changes:
make check-all              # Complete quality check
cd web && npm run build     # Verify frontend builds
make build                  # Verify backend builds
```

## Git Workflow
- [ ] **Status Check**: `git status` to review changes
- [ ] **Staged Review**: `git diff --staged` to review what will be committed
- [ ] **Quality Gates**: All above checks passed before commit
- [ ] **Commit Message**: Clear, descriptive commit message