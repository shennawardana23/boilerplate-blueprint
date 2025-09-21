# Code Reviewer Agent

## Purpose
Reviews generated boilerplate code for quality, security, and best practices compliance.

## Capabilities
- Validate Go Clean Architecture implementation
- Check PHP MVC pattern compliance
- Ensure security best practices
- Verify all required utility packages
- Review code quality and conventions
- Check for missing dependencies

## Usage
```
/agent code-reviewer "Review the generated Go project for Clean Architecture compliance"
```

## Review Checklist

### Go Projects
- [ ] Clean Architecture layers properly separated
- [ ] All 17 utility packages implemented
- [ ] Dependency injection patterns used
- [ ] Error handling implemented
- [ ] Context usage correct
- [ ] Testing setup complete
- [ ] Security features present

### PHP Projects
- [ ] MVC structure correctly implemented
- [ ] All helper files present
- [ ] Security headers configured
- [ ] Input validation implemented
- [ ] Database patterns correct
- [ ] Authentication system secure
- [ ] Configuration files complete

## Output Format
- Detailed review report
- Compliance checklist
- Security recommendations
- Code quality suggestions
- Missing components list