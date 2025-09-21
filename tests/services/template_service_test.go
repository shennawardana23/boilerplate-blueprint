package services_test

import (
	"testing"

	"boilerplate-blueprint/internal/models"
	"boilerplate-blueprint/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTemplateService(t *testing.T) {
	service := services.NewTemplateService()

	assert.NotNil(t, service)
}

func TestTemplateService_GetAvailableTemplates(t *testing.T) {
	service := services.NewTemplateService()

	templates := service.GetAvailableTemplates()

	assert.Len(t, templates, 2)

	// Check Go template
	goTemplate := templates[0]
	assert.Equal(t, models.LanguageGo, goTemplate.Language)
	assert.Equal(t, "Go Clean Architecture", goTemplate.Name)
	assert.Contains(t, goTemplate.Description, "Clean Architecture")
	assert.NotEmpty(t, goTemplate.Options)

	// Check PHP template
	phpTemplate := templates[1]
	assert.Equal(t, models.LanguagePHP, phpTemplate.Language)
	assert.Equal(t, "PHP CodeIgniter MVC", phpTemplate.Name)
	assert.Contains(t, phpTemplate.Description, "MVC")
	assert.NotEmpty(t, phpTemplate.Options)
}

func TestTemplateService_GetAvailableTemplates_GoOptions(t *testing.T) {
	service := services.NewTemplateService()

	templates := service.GetAvailableTemplates()
	goTemplate := templates[0]

	// Check framework option
	var frameworkOption *models.TemplateOption
	for _, option := range goTemplate.Options {
		if option.Key == "framework" {
			frameworkOption = &option
			break
		}
	}
	assert.NotNil(t, frameworkOption)
	assert.Equal(t, "HTTP Framework", frameworkOption.Label)
	assert.Equal(t, "select", frameworkOption.Type)
	assert.True(t, frameworkOption.Required)
	assert.Equal(t, "gin", frameworkOption.Default)
	assert.Contains(t, frameworkOption.Options, "gin")
	assert.Contains(t, frameworkOption.Options, "chi")
	assert.Contains(t, frameworkOption.Options, "echo")
	assert.Contains(t, frameworkOption.Options, "standard")

	// Check database option
	var databaseOption *models.TemplateOption
	for _, option := range goTemplate.Options {
		if option.Key == "database" {
			databaseOption = &option
			break
		}
	}
	assert.NotNil(t, databaseOption)
	assert.Equal(t, "Database", databaseOption.Label)
	assert.Equal(t, "select", databaseOption.Type)
	assert.True(t, databaseOption.Required)
	assert.Equal(t, "postgresql", databaseOption.Default)
	assert.Contains(t, databaseOption.Options, "postgresql")
	assert.Contains(t, databaseOption.Options, "mysql")
	assert.Contains(t, databaseOption.Options, "sqlite")
	assert.Contains(t, databaseOption.Options, "mongodb")

	// Check authentication option
	var authOption *models.TemplateOption
	for _, option := range goTemplate.Options {
		if option.Key == "authentication" {
			authOption = &option
			break
		}
	}
	assert.NotNil(t, authOption)
	assert.Equal(t, "Authentication", authOption.Label)
	assert.Equal(t, "select", authOption.Type)
	assert.True(t, authOption.Required)
	assert.Equal(t, "jwt", authOption.Default)
	assert.Contains(t, authOption.Options, "jwt")
	assert.Contains(t, authOption.Options, "oauth")
	assert.Contains(t, authOption.Options, "basic")
}

func TestTemplateService_GenerateGoProject(t *testing.T) {
	service := services.NewTemplateService()

	project := &models.Project{
		Name:        "test-go-project",
		Language:    models.LanguageGo,
		Description: "A test Go project",
		Options: models.ProjectOptions{
			Framework:      "gin",
			Database:       "postgresql",
			Authentication: "jwt",
			Utilities:      []string{"authentication", "cache"},
		},
	}

	files, err := service.GenerateGoProject(project)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// Verify directory structure
	filePaths := make(map[string]bool)
	for _, file := range files {
		filePaths[file.Path] = true
	}

	// Check main directories
	assert.True(t, filePaths["test-go-project"])
	assert.True(t, filePaths["test-go-project/cmd"])
	assert.True(t, filePaths["test-go-project/internal"])
	assert.True(t, filePaths["test-go-project/internal/app"])
	assert.True(t, filePaths["test-go-project/internal/app/database"])
	assert.True(t, filePaths["test-go-project/internal/app/middleware"])
	assert.True(t, filePaths["test-go-project/internal/controller"])
	assert.True(t, filePaths["test-go-project/internal/service"])
	assert.True(t, filePaths["test-go-project/internal/repository"])
	assert.True(t, filePaths["test-go-project/internal/entity"])
	assert.True(t, filePaths["test-go-project/internal/model/api"])
	assert.True(t, filePaths["test-go-project/internal/converter"])
	assert.True(t, filePaths["test-go-project/internal/routes"])
	assert.True(t, filePaths["test-go-project/internal/util"])
	assert.True(t, filePaths["test-go-project/scripts"])
	assert.True(t, filePaths["test-go-project/tests"])
	assert.True(t, filePaths["test-go-project/api"])

	// Check core files
	assert.True(t, filePaths["test-go-project/go.mod"])
	assert.True(t, filePaths["test-go-project/cmd/main.go"])
	assert.True(t, filePaths["test-go-project/Makefile"])
	assert.True(t, filePaths["test-go-project/Dockerfile"])
	assert.True(t, filePaths["test-go-project/README.md"])
	assert.True(t, filePaths["test-go-project/.gitignore"])
	assert.True(t, filePaths["test-go-project/.env"])
	assert.True(t, filePaths["test-go-project/.env.example"])
	assert.True(t, filePaths["test-go-project/internal/routes/router.go"])
}

func TestTemplateService_GenerateGoProject_goModContent(t *testing.T) {
	service := services.NewTemplateService()

	project := &models.Project{
		Name:        "my-go-app",
		Language:    models.LanguageGo,
		Description: "My Go application",
		Options: models.ProjectOptions{
			Framework: "gin",
			Database:  "postgresql",
		},
	}

	files, err := service.GenerateGoProject(project)
	require.NoError(t, err)

	// Find go.mod file
	var goModFile *models.ProjectFile
	for _, file := range files {
		if file.Path == "my-go-app/go.mod" {
			goModFile = &file
			break
		}
	}

	require.NotNil(t, goModFile)
	assert.Contains(t, goModFile.Content, "module my-go-app")
	assert.Contains(t, goModFile.Content, "go 1.21")
	assert.Contains(t, goModFile.Content, "github.com/gin-gonic/gin")
	assert.Contains(t, goModFile.Content, "github.com/lib/pq")
	assert.Contains(t, goModFile.Content, "github.com/redis/go-redis/v9")
	assert.Contains(t, goModFile.Content, "github.com/golang-jwt/jwt/v5")
}

func TestTemplateService_GeneratePHPProject(t *testing.T) {
	service := services.NewTemplateService()

	project := &models.Project{
		Name:        "test-php-project",
		Language:    models.LanguagePHP,
		Description: "A test PHP project",
		Options: models.ProjectOptions{
			CIVersion: "3",
			Database:  "mysql",
			Frontend:  "bootstrap",
			Features:  []string{"authentication", "dashboard"},
		},
	}

	files, err := service.GeneratePHPProject(project)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// Verify directory structure
	filePaths := make(map[string]bool)
	for _, file := range files {
		filePaths[file.Path] = true
	}

	// Check main directories
	assert.True(t, filePaths["test-php-project"])
	assert.True(t, filePaths["test-php-project/application"])
	assert.True(t, filePaths["test-php-project/application/cache"])
	assert.True(t, filePaths["test-php-project/application/config"])
	assert.True(t, filePaths["test-php-project/application/controllers"])
	assert.True(t, filePaths["test-php-project/application/core"])
	assert.True(t, filePaths["test-php-project/application/helpers"])
	assert.True(t, filePaths["test-php-project/application/libraries"])
	assert.True(t, filePaths["test-php-project/application/models"])
	assert.True(t, filePaths["test-php-project/application/views"])
	assert.True(t, filePaths["test-php-project/application/widgets"])
	assert.True(t, filePaths["test-php-project/assets"])
	assert.True(t, filePaths["test-php-project/assets/css"])
	assert.True(t, filePaths["test-php-project/assets/js"])
	assert.True(t, filePaths["test-php-project/assets/fonts"])
	assert.True(t, filePaths["test-php-project/assets/plugins"])
	assert.True(t, filePaths["test-php-project/system"])
	assert.True(t, filePaths["test-php-project/vendor"])

	// Check core files
	assert.True(t, filePaths["test-php-project/index.php"])
	assert.True(t, filePaths["test-php-project/composer.json"])
	assert.True(t, filePaths["test-php-project/README.md"])
	assert.True(t, filePaths["test-php-project/.gitignore"])
}

func TestTemplateService_CreateZIPArchive(t *testing.T) {
	service := services.NewTemplateService()

	project := &models.Project{
		Name:     "test-zip-project",
		Language: models.LanguageGo,
		Files: []models.ProjectFile{
			{
				Path:        "test-zip-project",
				Content:     "",
				IsDirectory: true,
			},
			{
				Path:        "test-zip-project/go.mod",
				Content:     "module test-zip-project\ngo 1.21",
				IsDirectory: false,
			},
			{
				Path:        "test-zip-project/cmd",
				Content:     "",
				IsDirectory: true,
			},
			{
				Path:        "test-zip-project/cmd/main.go",
				Content:     "package main\n\nfunc main() {\n\t// Main function\n}",
				IsDirectory: false,
			},
		},
	}

	zipData, err := service.CreateZIPArchive(project)

	require.NoError(t, err)
	assert.NotEmpty(t, zipData)
	assert.Greater(t, len(zipData), 0)
}

func TestTemplateService_CreateZIPArchive_EmptyProject(t *testing.T) {
	service := services.NewTemplateService()

	project := &models.Project{
		Name:     "empty-project",
		Language: models.LanguageGo,
		Files:    []models.ProjectFile{},
	}

	zipData, err := service.CreateZIPArchive(project)

	require.NoError(t, err)
	assert.NotEmpty(t, zipData)
}
