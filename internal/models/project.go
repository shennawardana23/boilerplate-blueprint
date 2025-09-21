package models

import "time"

// ProjectLanguage represents the supported programming languages
type ProjectLanguage string

const (
	LanguageGo  ProjectLanguage = "go"
	LanguagePHP ProjectLanguage = "php"
)

// ProjectRequest represents the request for creating a new project
type ProjectRequest struct {
	Name        string          `json:"name" binding:"required"`
	Language    ProjectLanguage `json:"language" binding:"required"`
	Description string          `json:"description"`
	Options     ProjectOptions  `json:"options"`
}

// ProjectOptions contains language-specific configuration options
type ProjectOptions struct {
	// Go-specific options
	Framework      string   `json:"framework,omitempty"`      // gin, chi, echo, standard
	Database       string   `json:"database,omitempty"`       // postgresql, mysql, sqlite, mongodb
	Authentication string   `json:"authentication,omitempty"` // jwt, oauth, basic
	Utilities      []string `json:"utilities,omitempty"`      // Selected utility packages

	// PHP-specific options
	CIVersion string   `json:"ci_version,omitempty"` // 3, 4
	Frontend  string   `json:"frontend,omitempty"`   // bootstrap, tailwind, custom
	Features  []string `json:"features,omitempty"`   // Selected features
}

// Project represents a generated project
type Project struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Language    ProjectLanguage `json:"language"`
	Description string          `json:"description"`
	Options     ProjectOptions  `json:"options"`
	Files       []ProjectFile   `json:"files"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// ProjectFile represents a file in the generated project
type ProjectFile struct {
	Path        string `json:"path"`
	Content     string `json:"content"`
	IsDirectory bool   `json:"is_directory"`
}

// ProjectResponse represents the API response for project operations
type ProjectResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
	Project *Project `json:"project,omitempty"`
	Error   string   `json:"error,omitempty"`
}

// TemplateInfo represents information about available templates
type TemplateInfo struct {
	Language    ProjectLanguage  `json:"language"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Options     []TemplateOption `json:"options"`
}

// TemplateOption represents a configurable option for a template
type TemplateOption struct {
	Key         string   `json:"key"`
	Label       string   `json:"label"`
	Type        string   `json:"type"` // text, select, checkbox, radio
	Required    bool     `json:"required"`
	Default     string   `json:"default,omitempty"`
	Options     []string `json:"options,omitempty"` // For select/radio types
	Description string   `json:"description,omitempty"`
}
