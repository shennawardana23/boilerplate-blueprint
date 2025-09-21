package services_test

import (
	"fmt"
	"testing"

	"boilerplate-blueprint/internal/models"
	"boilerplate-blueprint/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProjectService(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	assert.NotNil(t, service)
}

func TestProjectService_CreateProject_Go(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	req := &models.ProjectRequest{
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

	project, err := service.CreateProject(req)

	require.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, req.Name, project.Name)
	assert.Equal(t, req.Language, project.Language)
	assert.Equal(t, req.Description, project.Description)
	assert.Equal(t, req.Options.Framework, project.Options.Framework)
	assert.Equal(t, req.Options.Database, project.Options.Database)
	assert.Equal(t, req.Options.Authentication, project.Options.Authentication)
	assert.Equal(t, req.Options.Utilities, project.Options.Utilities)
	assert.NotEmpty(t, project.ID)
	assert.False(t, project.CreatedAt.IsZero())
	assert.False(t, project.UpdatedAt.IsZero())
}

func TestProjectService_CreateProject_PHP(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	req := &models.ProjectRequest{
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

	project, err := service.CreateProject(req)

	require.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, req.Name, project.Name)
	assert.Equal(t, req.Language, project.Language)
	assert.Equal(t, req.Description, project.Description)
	assert.Equal(t, req.Options.CIVersion, project.Options.CIVersion)
	assert.Equal(t, req.Options.Database, project.Options.Database)
	assert.Equal(t, req.Options.Frontend, project.Options.Frontend)
	assert.Equal(t, req.Options.Features, project.Options.Features)
}

func TestProjectService_CreateProject_InvalidLanguage(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	req := &models.ProjectRequest{
		Name:        "test-project",
		Language:    "invalid",
		Description: "A test project",
	}

	project, err := service.CreateProject(req)

	assert.Error(t, err)
	assert.Nil(t, project)
	assert.Contains(t, err.Error(), "unsupported language")
}

func TestProjectService_GetProject(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	// Create a project first
	req := &models.ProjectRequest{
		Name:        "test-project",
		Language:    models.LanguageGo,
		Description: "A test project",
	}

	createdProject, err := service.CreateProject(req)
	require.NoError(t, err)

	// Get the project
	retrievedProject, err := service.GetProject(createdProject.ID)

	require.NoError(t, err)
	assert.Equal(t, createdProject.ID, retrievedProject.ID)
	assert.Equal(t, createdProject.Name, retrievedProject.Name)
	assert.Equal(t, createdProject.Language, retrievedProject.Language)
}

func TestProjectService_GetProject_NotFound(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	project, err := service.GetProject("non-existent-id")

	assert.Error(t, err)
	assert.Nil(t, project)
	assert.Contains(t, err.Error(), "project not found")
}

func TestProjectService_GenerateProjectFiles_Go(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	// Create a Go project
	req := &models.ProjectRequest{
		Name:        "test-go-project",
		Language:    models.LanguageGo,
		Description: "A test Go project",
		Options: models.ProjectOptions{
			Framework: "gin",
			Database:  "postgresql",
		},
	}

	project, err := service.CreateProject(req)
	require.NoError(t, err)

	// Generate files
	files, err := service.GenerateProjectFiles(project)

	require.NoError(t, err)
	assert.NotEmpty(t, files)

	// Check that project was updated with files
	assert.NotEmpty(t, project.Files)
	assert.False(t, project.UpdatedAt.IsZero())

	// Verify some expected files exist
	filePaths := make(map[string]bool)
	for _, file := range files {
		filePaths[file.Path] = true
	}

	assert.True(t, filePaths["test-go-project/go.mod"])
	assert.True(t, filePaths["test-go-project/cmd/main.go"])
	assert.True(t, filePaths["test-go-project/Makefile"])
	assert.True(t, filePaths["test-go-project/README.md"])
}

func TestProjectService_CreateProjectZIP(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	// Create a project
	req := &models.ProjectRequest{
		Name:        "test-zip-project",
		Language:    models.LanguageGo,
		Description: "A test project for ZIP",
	}

	project, err := service.CreateProject(req)
	require.NoError(t, err)

	// Create ZIP
	zipData, filename, err := service.CreateProjectZIP(project.ID)

	require.NoError(t, err)
	assert.NotEmpty(t, zipData)
	assert.Equal(t, "test-zip-project-go.zip", filename)
	assert.Greater(t, len(zipData), 0)
}

func TestProjectService_ListProjects(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	// Initially should be empty
	projects := service.ListProjects()
	assert.Empty(t, projects)

	// Create some projects
	req1 := &models.ProjectRequest{
		Name:     "project1",
		Language: models.LanguageGo,
	}
	req2 := &models.ProjectRequest{
		Name:     "project2",
		Language: models.LanguagePHP,
	}

	_, err := service.CreateProject(req1)
	require.NoError(t, err)
	_, err = service.CreateProject(req2)
	require.NoError(t, err)

	// List projects
	projects = service.ListProjects()
	assert.Len(t, projects, 2)

	// Verify projects exist
	projectNames := make(map[string]bool)
	for _, project := range projects {
		projectNames[project.Name] = true
	}
	assert.True(t, projectNames["project1"])
	assert.True(t, projectNames["project2"])
}

func TestProjectService_ConcurrentAccess(t *testing.T) {
	templateService := services.NewTemplateService()
	service := services.NewProjectService(templateService)

	// Test concurrent project creation
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			req := &models.ProjectRequest{
				Name:     fmt.Sprintf("concurrent-project-%d", i),
				Language: models.LanguageGo,
			}
			_, err := service.CreateProject(req)
			assert.NoError(t, err)
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify all projects were created
	projects := service.ListProjects()
	assert.Len(t, projects, 10)
}
