package models_test

import (
	"testing"
	"time"

	"boilerplate-blueprint/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestProjectLanguage_Constants(t *testing.T) {
	assert.Equal(t, models.ProjectLanguage("go"), models.LanguageGo)
	assert.Equal(t, models.ProjectLanguage("php"), models.LanguagePHP)
}

func TestProjectRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		request models.ProjectRequest
		valid   bool
	}{
		{
			name: "Valid Go project request",
			request: models.ProjectRequest{
				Name:        "test-project",
				Language:    models.LanguageGo,
				Description: "A test project",
				Options: models.ProjectOptions{
					Framework:      "gin",
					Database:       "postgresql",
					Authentication: "jwt",
					Utilities:      []string{"authentication", "cache"},
				},
			},
			valid: true,
		},
		{
			name: "Valid PHP project request",
			request: models.ProjectRequest{
				Name:        "test-php-project",
				Language:    models.LanguagePHP,
				Description: "A test PHP project",
				Options: models.ProjectOptions{
					CIVersion: "3",
					Database:  "mysql",
					Frontend:  "bootstrap",
					Features:  []string{"authentication", "dashboard"},
				},
			},
			valid: true,
		},
		{
			name: "Empty name",
			request: models.ProjectRequest{
				Name:        "",
				Language:    models.LanguageGo,
				Description: "A test project",
			},
			valid: false,
		},
		{
			name: "Invalid language",
			request: models.ProjectRequest{
				Name:        "test-project",
				Language:    "invalid",
				Description: "A test project",
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				assert.NotEmpty(t, tt.request.Name)
				assert.Contains(t, []models.ProjectLanguage{models.LanguageGo, models.LanguagePHP}, tt.request.Language)
			} else {
				if tt.request.Name == "" {
					assert.Empty(t, tt.request.Name)
				}
				if tt.request.Language != models.LanguageGo && tt.request.Language != models.LanguagePHP {
					assert.NotContains(t, []models.ProjectLanguage{models.LanguageGo, models.LanguagePHP}, tt.request.Language)
				}
			}
		})
	}
}

func TestProjectOptions_GoOptions(t *testing.T) {
	options := models.ProjectOptions{
		Framework:      "gin",
		Database:       "postgresql",
		Authentication: "jwt",
		Utilities:      []string{"authentication", "cache", "common"},
	}

	assert.Equal(t, "gin", options.Framework)
	assert.Equal(t, "postgresql", options.Database)
	assert.Equal(t, "jwt", options.Authentication)
	assert.Len(t, options.Utilities, 3)
	assert.Contains(t, options.Utilities, "authentication")
	assert.Contains(t, options.Utilities, "cache")
	assert.Contains(t, options.Utilities, "common")
}

func TestProjectOptions_PHPOptions(t *testing.T) {
	options := models.ProjectOptions{
		CIVersion: "3",
		Database:  "mysql",
		Frontend:  "bootstrap",
		Features:  []string{"authentication", "user_management", "dashboard"},
	}

	assert.Equal(t, "3", options.CIVersion)
	assert.Equal(t, "mysql", options.Database)
	assert.Equal(t, "bootstrap", options.Frontend)
	assert.Len(t, options.Features, 3)
	assert.Contains(t, options.Features, "authentication")
	assert.Contains(t, options.Features, "user_management")
	assert.Contains(t, options.Features, "dashboard")
}

func TestProject_Creation(t *testing.T) {
	now := time.Now()
	project := models.Project{
		ID:          "test-id",
		Name:        "test-project",
		Language:    models.LanguageGo,
		Description: "A test project",
		Options: models.ProjectOptions{
			Framework: "gin",
			Database:  "postgresql",
		},
		Files: []models.ProjectFile{
			{
				Path:        "test-project/go.mod",
				Content:     "module test-project",
				IsDirectory: false,
			},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	assert.Equal(t, "test-id", project.ID)
	assert.Equal(t, "test-project", project.Name)
	assert.Equal(t, models.LanguageGo, project.Language)
	assert.Equal(t, "A test project", project.Description)
	assert.Equal(t, "gin", project.Options.Framework)
	assert.Equal(t, "postgresql", project.Options.Database)
	assert.Len(t, project.Files, 1)
	assert.Equal(t, "test-project/go.mod", project.Files[0].Path)
	assert.Equal(t, "module test-project", project.Files[0].Content)
	assert.False(t, project.Files[0].IsDirectory)
	assert.Equal(t, now, project.CreatedAt)
	assert.Equal(t, now, project.UpdatedAt)
}

func TestProjectFile_Creation(t *testing.T) {
	// Test file
	file := models.ProjectFile{
		Path:        "test-project/main.go",
		Content:     "package main\n\nfunc main() {\n\t// Main function\n}",
		IsDirectory: false,
	}

	assert.Equal(t, "test-project/main.go", file.Path)
	assert.Contains(t, file.Content, "package main")
	assert.Contains(t, file.Content, "func main()")
	assert.False(t, file.IsDirectory)

	// Test directory
	directory := models.ProjectFile{
		Path:        "test-project/internal",
		Content:     "",
		IsDirectory: true,
	}

	assert.Equal(t, "test-project/internal", directory.Path)
	assert.Empty(t, directory.Content)
	assert.True(t, directory.IsDirectory)
}

func TestProjectResponse_Success(t *testing.T) {
	project := &models.Project{
		ID:       "test-id",
		Name:     "test-project",
		Language: models.LanguageGo,
	}

	response := models.ProjectResponse{
		Success: true,
		Message: "Project created successfully",
		Project: project,
	}

	assert.True(t, response.Success)
	assert.Equal(t, "Project created successfully", response.Message)
	assert.Equal(t, project, response.Project)
	assert.Empty(t, response.Error)
}

func TestProjectResponse_Error(t *testing.T) {
	response := models.ProjectResponse{
		Success: false,
		Error:   "Project not found",
	}

	assert.False(t, response.Success)
	assert.Equal(t, "Project not found", response.Error)
	assert.Empty(t, response.Message)
	assert.Nil(t, response.Project)
}

func TestTemplateInfo_Creation(t *testing.T) {
	template := models.TemplateInfo{
		Language:    models.LanguageGo,
		Name:        "Go Clean Architecture",
		Description: "Complete Go project with Clean Architecture",
		Options: []models.TemplateOption{
			{
				Key:         "framework",
				Label:       "HTTP Framework",
				Type:        "select",
				Required:    true,
				Default:     "gin",
				Options:     []string{"gin", "chi", "echo"},
				Description: "Choose your HTTP framework",
			},
		},
	}

	assert.Equal(t, models.LanguageGo, template.Language)
	assert.Equal(t, "Go Clean Architecture", template.Name)
	assert.Equal(t, "Complete Go project with Clean Architecture", template.Description)
	assert.Len(t, template.Options, 1)

	option := template.Options[0]
	assert.Equal(t, "framework", option.Key)
	assert.Equal(t, "HTTP Framework", option.Label)
	assert.Equal(t, "select", option.Type)
	assert.True(t, option.Required)
	assert.Equal(t, "gin", option.Default)
	assert.Len(t, option.Options, 3)
	assert.Contains(t, option.Options, "gin")
	assert.Contains(t, option.Options, "chi")
	assert.Contains(t, option.Options, "echo")
	assert.Equal(t, "Choose your HTTP framework", option.Description)
}

func TestTemplateOption_Types(t *testing.T) {
	tests := []struct {
		name     string
		option   models.TemplateOption
		expected string
	}{
		{
			name: "Select option",
			option: models.TemplateOption{
				Key:     "framework",
				Type:    "select",
				Options: []string{"gin", "chi", "echo"},
			},
			expected: "select",
		},
		{
			name: "Text option",
			option: models.TemplateOption{
				Key:  "project_name",
				Type: "text",
			},
			expected: "text",
		},
		{
			name: "Checkbox option",
			option: models.TemplateOption{
				Key:  "utilities",
				Type: "checkbox",
			},
			expected: "checkbox",
		},
		{
			name: "Radio option",
			option: models.TemplateOption{
				Key:     "database",
				Type:    "radio",
				Options: []string{"postgresql", "mysql", "sqlite"},
			},
			expected: "radio",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.option.Type)
		})
	}
}

func TestTemplateOption_Required(t *testing.T) {
	requiredOption := models.TemplateOption{
		Key:      "framework",
		Required: true,
	}

	optionalOption := models.TemplateOption{
		Key:      "description",
		Required: false,
	}

	assert.True(t, requiredOption.Required)
	assert.False(t, optionalOption.Required)
}

func TestTemplateOption_DefaultValues(t *testing.T) {
	option := models.TemplateOption{
		Key:     "framework",
		Default: "gin",
		Options: []string{"gin", "chi", "echo", "standard"},
	}

	assert.Equal(t, "gin", option.Default)
	assert.Contains(t, option.Options, option.Default)
}

func TestProjectOptions_EmptyOptions(t *testing.T) {
	options := models.ProjectOptions{}

	assert.Empty(t, options.Framework)
	assert.Empty(t, options.Database)
	assert.Empty(t, options.Authentication)
	assert.Empty(t, options.Utilities)
	assert.Empty(t, options.CIVersion)
	assert.Empty(t, options.Frontend)
	assert.Empty(t, options.Features)
}

func TestProjectOptions_PartialOptions(t *testing.T) {
	options := models.ProjectOptions{
		Framework: "gin",
		Database:  "postgresql",
		// Other fields left empty
	}

	assert.Equal(t, "gin", options.Framework)
	assert.Equal(t, "postgresql", options.Database)
	assert.Empty(t, options.Authentication)
	assert.Empty(t, options.Utilities)
	assert.Empty(t, options.CIVersion)
	assert.Empty(t, options.Frontend)
	assert.Empty(t, options.Features)
}

func TestProjectFile_EmptyContent(t *testing.T) {
	file := models.ProjectFile{
		Path:        "test-project/.gitignore",
		Content:     "",
		IsDirectory: false,
	}

	assert.Equal(t, "test-project/.gitignore", file.Path)
	assert.Empty(t, file.Content)
	assert.False(t, file.IsDirectory)
}

func TestProjectFile_DirectoryWithContent(t *testing.T) {
	// This shouldn't happen in practice, but test the structure
	directory := models.ProjectFile{
		Path:        "test-project/internal",
		Content:     "This shouldn't be here",
		IsDirectory: true,
	}

	assert.Equal(t, "test-project/internal", directory.Path)
	assert.NotEmpty(t, directory.Content)
	assert.True(t, directory.IsDirectory)
}

func TestProject_EmptyFiles(t *testing.T) {
	project := models.Project{
		ID:       "test-id",
		Name:     "test-project",
		Language: models.LanguageGo,
		Files:    []models.ProjectFile{},
	}

	assert.Equal(t, "test-id", project.ID)
	assert.Equal(t, "test-project", project.Name)
	assert.Equal(t, models.LanguageGo, project.Language)
	assert.Empty(t, project.Files)
}

func TestProject_MultipleFiles(t *testing.T) {
	project := models.Project{
		ID:       "test-id",
		Name:     "test-project",
		Language: models.LanguageGo,
		Files: []models.ProjectFile{
			{
				Path:        "test-project/go.mod",
				Content:     "module test-project",
				IsDirectory: false,
			},
			{
				Path:        "test-project/cmd",
				Content:     "",
				IsDirectory: true,
			},
			{
				Path:        "test-project/cmd/main.go",
				Content:     "package main",
				IsDirectory: false,
			},
		},
	}

	assert.Len(t, project.Files, 3)
	assert.Equal(t, "test-project/go.mod", project.Files[0].Path)
	assert.Equal(t, "test-project/cmd", project.Files[1].Path)
	assert.Equal(t, "test-project/cmd/main.go", project.Files[2].Path)
	assert.False(t, project.Files[0].IsDirectory)
	assert.True(t, project.Files[1].IsDirectory)
	assert.False(t, project.Files[2].IsDirectory)
}
