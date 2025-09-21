# Product Requirements Document: Admin CMS Web Builder with CodeIgniter

## Project Overview

Develop a sophisticated drag-and-drop Admin CMS Web Builder using CodeIgniter 3.x/4.x framework with PHP 8.0. The system will enable users to visually create and manage websites through an intuitive interface similar to design.com's website builder, featuring modern UI/UX patterns and real-time preview capabilities.

## Technology Stack & Architecture

### Core Technologies

- **Framework**: CodeIgniter (latest stable version)
- **PHP Version**: 8.0+
- **Frontend Technologies**:
  - Modern JavaScript (ES6+) with modular architecture
  - CSS3 with CSS Grid and Flexbox for layouts
  - Tailwind CSS for utility-first styling
  - Vanilla JavaScript for UI interactions
- **Primary Drag & Drop Solution**: neodrag library (<https://next.neodrag.dev/>) - for ALL page builder drag-drop operations
- **Database**: MySQL 8.0 or MariaDB
- **Additional Frontend Libraries**:
  - TinyMCE or CKEditor 5 for rich text editing
  - SweetAlert2 for modern alerts
  - Chart.js for analytics dashboards

## CodeIgniter Project Structure & Organization

### Application Directory Structure

```
/application
├── /cache              # Page cache and compiled templates
├── /config             # Configuration files for CMS settings
│   ├── autoload.php    # Auto-load libraries and helpers
│   ├── database.php    # Database connection settings
│   ├── routes.php      # CMS routing configuration
│   └── cms_config.php  # Custom CMS settings
├── /controllers        # MVC Controllers
│   │   ├── Dashboard.php
│   │   ├── PageBuilder.php
│   │   ├── Templates.php
│   │   ├── Components.php
│   │   ├── Media.php
│   │   └── Settings.php
│   └── Preview.php     # Preview controller for built pages
├── /core               # Extended CI core classes
│   ├── MY_Controller.php
│   ├── MY_Model.php
│   └── MY_Loader.php
├── /helpers            # Custom helper functions
│   ├── cms_helper.php
│   ├── component_helper.php
│   └── builder_helper.php
├── /hooks              # Pre/post system execution hooks
├── /language           # Multi-language support files
├── /libraries          # Custom libraries
│   ├── Page_builder.php
│   ├── Component_manager.php
│   ├── Template_engine.php
│   ├── Asset_manager.php
│   └── Export_handler.php
├── /models             # Database models
│   ├── Page_model.php
│   ├── Component_model.php
│   ├── Template_model.php
│   ├── Media_model.php
│   └── Settings_model.php
├── /third_party        # Third-party integrations
├── /vendor             # Composer dependencies
├── /views              # View files
│   ├── admin/
│   │   ├── builder/
│   │   ├── templates/
│   │   ├── components/
│   │   └── layouts/
│   └── frontend/       # Generated page views
└── /widgets            # Reusable widget components
```

### Assets Directory Structure

```
/assets
├── /css
│   ├── /admin          # Admin panel styles
│   ├── /builder        # Page builder specific styles
│   └── /components     # Component-specific styles
├── /js
│   ├── /admin
│   ├── /builder
│   │   ├── drag-drop.js
│   │   ├── canvas.js
│   │   ├── properties.js
│   │   └── components.js
│   └── /vendor         # Third-party JS libraries
├── /images
│   ├── /templates      # Template thumbnails
│   ├── /components     # Component previews
│   └── /uploads        # User uploaded media
└── /fonts              # Custom web fonts
```

## Detailed Feature Specifications

### 1. Admin Dashboard & User Interface

#### Dashboard Overview

The main dashboard should provide a comprehensive overview of the CMS status with modern card-based widgets displaying:

- Total pages created and their status (draft/published)
- Recent editing activity with timestamps
- Quick access buttons to create new pages
- Template usage statistics
- System health indicators
- Recent media uploads

#### Navigation System

Implement a responsive sidebar navigation that:

- Collapses on mobile devices
- Shows active state for current section
- Includes icon-based menu items
- Supports nested menu structures for complex sites
- Has quick search functionality
- Remembers user's preferred state (expanded/collapsed)

### 2. Visual Page Builder Core

#### Canvas Area Specifications

The central canvas should function as a WYSIWYG editor where:

- Components can be dropped and arranged visually
- Real-time preview reflects exact front-end appearance
- Grid system overlay helps with alignment (toggleable)
- Responsive breakpoint preview (desktop/tablet/mobile)
- Zoom controls for detailed editing (50% to 200%)
- Ruler guides for precise positioning
- Component highlighting on hover
- Multi-select capability for bulk operations

#### Component Library Panel

Left sidebar containing categorized components:

- **Basic Elements**: Text, Image, Button, Divider, Spacer
- **Layout Components**: Container, Row, Column, Grid, Flexbox
- **Navigation**: Menu Bar, Breadcrumbs, Pagination, Tabs
- **Content Blocks**: Hero Section, Features, Testimonials, Team Members
- **Media Components**: Image Gallery, Video Player, Audio Player, Carousel
- **Forms**: Contact Form, Newsletter Signup, Survey Builder
- **E-commerce**: Product Grid, Shopping Cart, Checkout Form
- **Social**: Social Media Icons, Share Buttons, Comments Section
- **Advanced**: Custom HTML, Code Block, Embed Container

Each component should have:

- Preview thumbnail
- Drag handle for easy grabbing
- Quick info tooltip
- Usage frequency indicator

#### Properties Editor Panel

Right sidebar for editing selected components:

- **Content Tab**: Text editing, image selection, link management
- **Style Tab**: Colors, typography, spacing, borders, shadows
- **Advanced Tab**: CSS classes, custom CSS, visibility conditions
- **Responsive Tab**: Per-breakpoint settings
- **Animation Tab**: Entrance effects, scroll animations
- **Settings Tab**: Component-specific configurations

### 3. Component System Architecture

#### Neodrag Implementation Details

The entire drag-and-drop functionality for the page builder should be powered by **neodrag** (<https://next.neodrag.dev/>):

**Core Neodrag Features to Implement:**

- **Component Dragging**: From sidebar component library to canvas
- **Reordering**: Moving components up/down within the canvas
- **Nested Dragging**: Components inside containers/columns
- **Boundary Constraints**: Keep components within canvas bounds
- **Grid Snapping**: Align components to invisible grid
- **Multi-axis Control**: Some components only move vertically/horizontally
- **Clone Dragging**: Hold key to duplicate while dragging
- **Touch Support**: Mobile-friendly drag operations
- **Disabled States**: Lock certain components from being moved
- **Custom Drag Handles**: Specific grab areas on components
- **Drop Zones**: Visual indicators for valid drop areas
- **Collision Detection**: Prevent component overlap

**Neodrag Configuration Approach:**

- Initialize neodrag on component mount
- Set bounds to canvas container
- Configure grid size (e.g., 10px increments)
- Handle drag start/move/end events
- Update component positions in real-time
- Save final positions to database via AJAX

#### Base Component Structure

Every component in the system should extend from a base component class that provides:

- Unique identifier generation
- Property management system
- Validation framework
- Rendering pipeline
- Event handling system
- Versioning support

#### Component Categories and Examples

**Hero Sections**

- Classic hero with background image
- Video background hero
- Slider hero with multiple slides
- Split hero with image and text
- Gradient overlay hero

**Content Blocks**

- Article/blog post layout
- Feature boxes with icons
- Accordion/collapsible content
- Timeline components
- Pricing tables

**Interactive Elements**

- Modal/popup triggers
- Tooltip components
- Tab navigation systems
- Progress bars and counters
- Interactive maps

### 4. Template Management System

#### Template Categories

Organize templates by industry and use case:

- **Business**: Corporate, Agency, Consulting, Startup
- **E-commerce**: Fashion, Electronics, Food, General Store
- **Portfolio**: Photography, Design, Art, Personal
- **Services**: Restaurant, Gym, Salon, Medical
- **Events**: Wedding, Conference, Party, Webinar
- **Landing Pages**: Product Launch, Coming Soon, Lead Generation
- **Travel/Resort**: Hotel, Tourism, Booking (reference design.com example)

#### Template Features

- Live preview before selection
- One-click import with content
- Customization wizard after import
- Save custom templates from existing pages
- Template versioning and updates
- Template marketplace integration ready

### 5. Media & Asset Management

#### Media Library Features

- Standard HTML5 file upload interface (not drag-drop library)
- Multiple file selection capability
- Upload progress indicators
- Automatic image optimization
- Multiple size generation for responsive images
- Video thumbnail generation
- File type filtering and search
- Folder organization system
- Tag-based categorization
- Usage tracking (which pages use which assets)
- External media integration (Unsplash, Pexels)

#### Asset Processing

- Automatic WebP conversion
- Lazy loading implementation
- CDN preparation
- SVG optimization
- Image cropping and editing tools

### 6. Page Management System

#### Page Operations

- Create pages with SEO-friendly URLs
- Hierarchical page structure support
- Draft and revision system
- Scheduled publishing
- A/B testing capability
- Password protection option
- Custom 404 pages
- Maintenance mode pages

#### SEO Features

- Meta title and description editor
- Open Graph tags management
- Schema markup generator
- XML sitemap generation
- Robots.txt editor
- Canonical URL settings
- Multi-language SEO support

### 7. CodeIgniter-Specific Implementation Guidelines

#### Controller Structure

Controllers should be organized by functionality:

- Use base controller for common functionality
- Implement proper authentication checks
- Handle AJAX requests for builder operations
- Manage session data for builder state

#### Model Implementation

Models should follow active record pattern:

- Separate models for each major entity
- Use CodeIgniter's query builder
- Implement caching for frequently accessed data
- Handle database transactions for complex operations

#### Library Development

Custom libraries should handle:

- Component rendering logic
- Template parsing and compilation
- Asset optimization and management
- Export functionality
- Builder state management

#### Helper Functions

Create helpers for:

- Common HTML generation
- Asset path management
- Component data formatting
- Builder utility functions

### 8. Advanced Features

#### Collaboration Features

- Multi-user support with role management
- Real-time collaboration indicators
- Version control with rollback
- Comment system for feedback
- Activity logging

#### Integration Capabilities

- REST API for external integrations
- Webhook support for events
- Third-party service connections
- Social media integration
- Analytics integration (Google Analytics, etc.)

#### Performance Optimization

- Lazy loading of components
- Code splitting for JavaScript
- Progressive Web App features
- Server-side rendering option
- Built-in caching mechanisms

### 9. Security Considerations

#### Input Validation

- XSS prevention for all user inputs
- CSRF token implementation
- File upload validation and sanitization
- SQL injection prevention through query builder

#### Access Control

- Role-based permissions system
- IP whitelisting option
- Two-factor authentication support
- Session management
- API rate limiting

### 10. Export and Deployment

#### Export Options

- Static HTML/CSS/JS generation
- WordPress theme export
- CodeIgniter application export
- Progressive Web App export
- Email template export

#### Deployment Features

- One-click publishing to server
- FTP/SFTP integration
- Git deployment support
- Staging environment support
- Rollback capabilities

## Implementation Phases

### Phase 1: Foundation (Weeks 1-2)

- Set up CodeIgniter project structure
- Implement authentication system
- Create base controllers and models
- Design database schema
- Set up admin dashboard layout

### Phase 2: Core Builder (Weeks 3-5)

- Implement comprehensive drag-and-drop functionality using neodrag library
- Configure neodrag for component library to canvas dragging
- Set up neodrag for component reordering and nesting
- Create canvas rendering system
- Build component architecture
- Develop property editor
- Implement basic components (10-15)

### Phase 3: Component Library (Weeks 6-7)

- Develop full component suite (50+ components)
- Create component preview system
- Implement component variations
- Build component search and filtering

### Phase 4: Template System (Weeks 8-9)

- Design template structure
- Create template management interface
- Build template import/export
- Develop template customization wizard

### Phase 5: Advanced Features (Weeks 10-11)

- Implement media library
- Add SEO tools
- Create export functionality
- Build preview and publishing system

### Phase 6: Polish & Optimization (Weeks 12-14)

- UI/UX refinements
- Performance optimization
- Security hardening
- Documentation
- Testing and bug fixes

## Success Metrics

- Page builder loads in under 2 seconds
- Smooth 60fps drag-and-drop operations
- Support for 100+ components
- Mobile-responsive output for all templates
- 99.9% uptime for hosted solution
- Export time under 30 seconds for average site

## Future Enhancements

- AI-powered design suggestions
- Advanced animation builder
- E-commerce integration
- Membership site features
- Multi-language content management
- White-label solution options
- Plugin marketplace
- Advanced form builder with logic
- Email campaign builder
- API-first headless CMS mode
