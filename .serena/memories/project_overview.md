# Boilerplate Blueprint - Project Overview

## Purpose
This is a web-based boilerplate generator application that creates Go and PHP CodeIgniter projects with AI assistance. Similar to [Go Blueprint](https://go-blueprint.dev/), it allows users to:

- Select between Go and PHP CodeIgniter projects
- Use AI chat interface for project customization
- Generate complete project structures with boilerplate code
- Download projects as ZIP files

## Architecture
**Backend**: Go with Gin framework
- Serves API endpoints for project generation
- Handles AI chat integration
- Creates ZIP files for download
- Template-based code generation

**Frontend**: Vue.js 3 with Vite
- Modern SPA with Vue Router and Pinia state management
- TailwindCSS for styling
- Headless UI components

## Key Features
- Language selection (Go vs PHP)
- AI-powered project customization chat
- Real-time project structure preview
- Dynamic file and folder generation
- ZIP download functionality
- Template management system

## Current Status
This appears to be a working prototype with basic functionality implemented. The project follows the requirements outlined in `project-description.md` for creating a comprehensive boilerplate generation system.