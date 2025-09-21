# Claude Configuration

This directory contains Claude Code configuration files for the Boilerplate Generator project.

## Files

- `preferences.json` - Claude preferences and settings
- `project-context.md` - Project-specific context and guidelines
- `project-description.md` - Complete project requirements and prompts
- `CLAUDE.md` - This configuration overview

## Development Notes

When working on this project, Claude should:

1. **Follow the quick prototype approach** - prioritize getting a working MVP
2. **Use Go backend with Gin** - avoid complex architectures initially  
3. **Generate proper file structures** - follow the exact patterns in project-description.md
4. **Create downloadable ZIP files** - every generation must produce a ZIP
5. **Implement AI chat functionality** - even if basic initially
6. **Support both Go and PHP** - with proper template systems for each

## Key Requirements

- Must generate exact directory structures as specified
- Must include all 17 utility packages for Go projects
- Must include comprehensive MVC structure for PHP projects
- Must provide real-time project structure preview
- Must create downloadable ZIP files with complete projects

## Reference Files

- See `project-description.md` for complete requirements
- See root `prompt-go-project.md` for Go specifications  
- See root `prompt-php-project.md` for PHP specifications