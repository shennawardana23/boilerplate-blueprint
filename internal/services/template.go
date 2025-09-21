package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"boilerplate-blueprint/internal/models"
)

type TemplateService struct {
	goTemplates  map[string]*template.Template
	phpTemplates map[string]*template.Template
}

func NewTemplateService() *TemplateService {
	return &TemplateService{
		goTemplates:  make(map[string]*template.Template),
		phpTemplates: make(map[string]*template.Template),
	}
}

func (s *TemplateService) GetAvailableTemplates() []models.TemplateInfo {
	return []models.TemplateInfo{
		{
			Language:    models.LanguageGo,
			Name:        "Go Clean Architecture",
			Description: "Complete Go project with Clean Architecture, Gin framework, and 17 utility packages",
			Options: []models.TemplateOption{
				{
					Key:      "framework",
					Label:    "HTTP Framework",
					Type:     "select",
					Required: true,
					Default:  "gin",
					Options:  []string{"gin", "chi", "echo", "standard"},
				},
				{
					Key:      "database",
					Label:    "Database",
					Type:     "select",
					Required: true,
					Default:  "postgresql",
					Options:  []string{"postgresql", "mysql", "sqlite", "mongodb"},
				},
				{
					Key:      "authentication",
					Label:    "Authentication",
					Type:     "select",
					Required: true,
					Default:  "jwt",
					Options:  []string{"jwt", "oauth", "basic"},
				},
			},
		},
		{
			Language:    models.LanguagePHP,
			Name:        "PHP CodeIgniter MVC",
			Description: "Complete PHP CodeIgniter project with MVC architecture and security features",
			Options: []models.TemplateOption{
				{
					Key:      "ci_version",
					Label:    "CodeIgniter Version",
					Type:     "select",
					Required: true,
					Default:  "3",
					Options:  []string{"3", "4"},
				},
				{
					Key:      "database",
					Label:    "Database",
					Type:     "select",
					Required: true,
					Default:  "postgresql",
					Options:  []string{"postgresql", "mysql", "sqlite"},
				},
				{
					Key:      "frontend",
					Label:    "Frontend Framework",
					Type:     "select",
					Required: true,
					Default:  "bootstrap",
					Options:  []string{"bootstrap", "tailwind", "custom"},
				},
			},
		},
	}
}

func (s *TemplateService) GenerateGoProject(project *models.Project) ([]models.ProjectFile, error) {
	var files []models.ProjectFile

	// Template data
	data := map[string]interface{}{
		"ProjectName":    project.Name,
		"Description":    project.Description,
		"Framework":      project.Options.Framework,
		"Database":       project.Options.Database,
		"Authentication": project.Options.Authentication,
		"Utilities":      project.Options.Utilities,
		"PackageName":    strings.ToLower(strings.ReplaceAll(project.Name, " ", "-")),
	}

	// Generate directory structure first
	files = append(files, s.createGoDirectoryStructure(data)...)

	// Generate core files
	files = append(files, s.generateGoModFile(data))
	files = append(files, s.generateGoMainFile(data))
	files = append(files, s.generateGoMakefile(data))
	files = append(files, s.generateGoDockerfile(data))
	files = append(files, s.generateGoReadme(data))
	files = append(files, s.generateGoGitignore(data))
	files = append(files, s.generateGoEnvFiles(data)...)

	// Generate application structure
	files = append(files, s.generateGoConfigFiles(data)...)
	files = append(files, s.generateGoMiddleware(data)...)
	files = append(files, s.generateGoControllers(data)...)
	files = append(files, s.generateGoServices(data)...)
	files = append(files, s.generateGoRepositories(data)...)
	files = append(files, s.generateGoModels(data)...)
	files = append(files, s.generateGoUtilities(data)...)
	files = append(files, s.generateGoRoutes(data))

	return files, nil
}

func (s *TemplateService) GeneratePHPProject(project *models.Project) ([]models.ProjectFile, error) {
	var files []models.ProjectFile

	// Template data
	data := map[string]interface{}{
		"ProjectName": project.Name,
		"Description": project.Description,
		"CIVersion":   project.Options.CIVersion,
		"Database":    project.Options.Database,
		"Frontend":    project.Options.Frontend,
		"Features":    project.Options.Features,
	}

	// Generate directory structure first
	files = append(files, s.createPHPDirectoryStructure(data)...)

	// Generate core files
	files = append(files, s.generatePHPIndexFile(data))
	files = append(files, s.generatePHPComposerFile(data))
	files = append(files, s.generatePHPReadme(data))
	files = append(files, s.generatePHPGitignore(data))

	// Generate application structure
	files = append(files, s.generatePHPConfigFiles(data)...)
	files = append(files, s.generatePHPControllers(data)...)
	files = append(files, s.generatePHPModels(data)...)
	files = append(files, s.generatePHPViews(data)...)
	files = append(files, s.generatePHPHelpers(data)...)
	files = append(files, s.generatePHPLibraries(data)...)
	files = append(files, s.generatePHPCore(data)...)

	return files, nil
}

func (s *TemplateService) CreateZIPArchive(project *models.Project) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, file := range project.Files {
		if file.IsDirectory {
			// Create directory entry
			_, err := zipWriter.Create(file.Path + "/")
			if err != nil {
				return nil, fmt.Errorf("failed to create directory %s: %w", file.Path, err)
			}
		} else {
			// Create file entry
			writer, err := zipWriter.Create(file.Path)
			if err != nil {
				return nil, fmt.Errorf("failed to create file %s: %w", file.Path, err)
			}

			_, err = writer.Write([]byte(file.Content))
			if err != nil {
				return nil, fmt.Errorf("failed to write file %s: %w", file.Path, err)
			}
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close ZIP writer: %w", err)
	}

	return buf.Bytes(), nil
}

// Helper functions for creating project structure

func (s *TemplateService) createGoDirectoryStructure(data map[string]interface{}) []models.ProjectFile {
	projectName := data["ProjectName"].(string)
	dirs := []string{
		fmt.Sprintf("%s", projectName),
		fmt.Sprintf("%s/cmd", projectName),
		fmt.Sprintf("%s/internal", projectName),
		fmt.Sprintf("%s/internal/app", projectName),
		fmt.Sprintf("%s/internal/app/database", projectName),
		fmt.Sprintf("%s/internal/app/middleware", projectName),
		fmt.Sprintf("%s/internal/controller", projectName),
		fmt.Sprintf("%s/internal/service", projectName),
		fmt.Sprintf("%s/internal/repository", projectName),
		fmt.Sprintf("%s/internal/entity", projectName),
		fmt.Sprintf("%s/internal/model/api", projectName),
		fmt.Sprintf("%s/internal/converter", projectName),
		fmt.Sprintf("%s/internal/routes", projectName),
		fmt.Sprintf("%s/internal/util", projectName),
		fmt.Sprintf("%s/scripts", projectName),
		fmt.Sprintf("%s/tests", projectName),
		fmt.Sprintf("%s/api", projectName),
	}

	var files []models.ProjectFile
	for _, dir := range dirs {
		files = append(files, models.ProjectFile{
			Path:        dir,
			Content:     "",
			IsDirectory: true,
		})
	}

	return files
}

func (s *TemplateService) createPHPDirectoryStructure(data map[string]interface{}) []models.ProjectFile {
	projectName := data["ProjectName"].(string)
	dirs := []string{
		fmt.Sprintf("%s", projectName),
		fmt.Sprintf("%s/application", projectName),
		fmt.Sprintf("%s/application/cache", projectName),
		fmt.Sprintf("%s/application/config", projectName),
		fmt.Sprintf("%s/application/controllers", projectName),
		fmt.Sprintf("%s/application/core", projectName),
		fmt.Sprintf("%s/application/helpers", projectName),
		fmt.Sprintf("%s/application/libraries", projectName),
		fmt.Sprintf("%s/application/models", projectName),
		fmt.Sprintf("%s/application/views", projectName),
		fmt.Sprintf("%s/application/widgets", projectName),
		fmt.Sprintf("%s/assets", projectName),
		fmt.Sprintf("%s/assets/css", projectName),
		fmt.Sprintf("%s/assets/js", projectName),
		fmt.Sprintf("%s/assets/fonts", projectName),
		fmt.Sprintf("%s/assets/plugins", projectName),
		fmt.Sprintf("%s/system", projectName),
		fmt.Sprintf("%s/vendor", projectName),
	}

	var files []models.ProjectFile
	for _, dir := range dirs {
		files = append(files, models.ProjectFile{
			Path:        dir,
			Content:     "",
			IsDirectory: true,
		})
	}

	return files
}

// Placeholder functions for file generation (will be implemented next)

func (s *TemplateService) generateGoModFile(data map[string]interface{}) models.ProjectFile {
	content := fmt.Sprintf(`module %s

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/lib/pq v1.10.9
	github.com/redis/go-redis/v9 v9.3.0
	github.com/golang-jwt/jwt/v5 v5.2.0
	golang.org/x/crypto v0.40.0
	github.com/go-playground/validator/v10 v10.19.0
	github.com/sirupsen/logrus v1.9.3
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/rs/cors v1.10.1
)`, data["PackageName"])

	return models.ProjectFile{
		Path:        filepath.Join(data["ProjectName"].(string), "go.mod"),
		Content:     content,
		IsDirectory: false,
	}
}

// More generator functions will be added in subsequent steps

func (s *TemplateService) generateGoMainFile(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "cmd", "main.go"), Content: "// Main file placeholder", IsDirectory: false}
}

func (s *TemplateService) generateGoMakefile(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "Makefile"), Content: "# Makefile placeholder", IsDirectory: false}
}

func (s *TemplateService) generateGoDockerfile(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "Dockerfile"), Content: "# Dockerfile placeholder", IsDirectory: false}
}

func (s *TemplateService) generateGoReadme(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "README.md"), Content: "# README placeholder", IsDirectory: false}
}

func (s *TemplateService) generateGoGitignore(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), ".gitignore"), Content: "# Gitignore placeholder", IsDirectory: false}
}

func (s *TemplateService) generateGoEnvFiles(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{
		{Path: filepath.Join(data["ProjectName"].(string), ".env"), Content: "# Environment variables", IsDirectory: false},
		{Path: filepath.Join(data["ProjectName"].(string), ".env.example"), Content: "# Environment example", IsDirectory: false},
	}
}

func (s *TemplateService) generateGoConfigFiles(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoMiddleware(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoControllers(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoServices(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoRepositories(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoModels(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoUtilities(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generateGoRoutes(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "internal", "routes", "router.go"), Content: "// Router placeholder", IsDirectory: false}
}

// PHP file generation functions (placeholders for now)

func (s *TemplateService) generatePHPIndexFile(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "index.php"), Content: "<?php // PHP index", IsDirectory: false}
}

func (s *TemplateService) generatePHPComposerFile(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "composer.json"), Content: "{}", IsDirectory: false}
}

func (s *TemplateService) generatePHPReadme(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), "README.md"), Content: "# PHP README", IsDirectory: false}
}

func (s *TemplateService) generatePHPGitignore(data map[string]interface{}) models.ProjectFile {
	return models.ProjectFile{Path: filepath.Join(data["ProjectName"].(string), ".gitignore"), Content: "# PHP Gitignore", IsDirectory: false}
}

func (s *TemplateService) generatePHPConfigFiles(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generatePHPControllers(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generatePHPModels(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generatePHPViews(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generatePHPHelpers(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generatePHPLibraries(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}

func (s *TemplateService) generatePHPCore(data map[string]interface{}) []models.ProjectFile {
	return []models.ProjectFile{}
}
