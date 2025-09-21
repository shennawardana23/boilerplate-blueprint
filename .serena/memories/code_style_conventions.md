# Code Style and Conventions

## Go Code Style
- **Package Naming**: Uses kebab-case for module name (`boilerplate-blueprint`)
- **Directory Structure**: Clean Architecture pattern with `internal/` for private packages
- **Import Organization**: Standard library imports first, then external packages, then local packages
- **Error Handling**: Standard Go error handling with explicit checks
- **Logging**: Uses standard `log` package for basic logging
- **Configuration**: Environment variables loaded via godotenv

## Go File Organization
```
internal/
├── api/           # HTTP handlers and routes
├── services/      # Business logic layer
└── models/        # Data structures and types
```

## Frontend Code Style (Vue.js)
- **Framework**: Vue 3 with Composition API
- **Module System**: ES6 modules
- **Styling**: TailwindCSS utility classes
- **Component Organization**: Single File Components (.vue)
- **State Management**: Pinia stores
- **Routing**: Vue Router with history mode

## Naming Conventions
- **Go**: 
  - Functions/Methods: PascalCase for public, camelCase for private
  - Variables: camelCase
  - Constants: UPPER_SNAKE_CASE
  - Packages: lowercase, single word preferred
- **Vue.js**:
  - Components: PascalCase
  - Props/Data: camelCase
  - Events: kebab-case
  - CSS Classes: TailwindCSS utilities

## Development Patterns
- **Go**: 
  - Dependency injection through constructor functions
  - Interface-based design for testability
  - Service layer pattern for business logic
- **Vue.js**:
  - Composition API over Options API
  - Reactive data with ref/reactive
  - Composables for reusable logic