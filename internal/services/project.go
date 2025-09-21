package services

import (
	"fmt"
	"sync"
	"time"

	"boilerplate-blueprint/internal/models"

	"github.com/google/uuid"
)

type ProjectService struct {
	projects        map[string]*models.Project
	templateService *TemplateService
	mu              sync.RWMutex
}

func NewProjectService(templateService *TemplateService) *ProjectService {
	return &ProjectService{
		projects:        make(map[string]*models.Project),
		templateService: templateService,
	}
}

func (s *ProjectService) CreateProject(req *models.ProjectRequest) (*models.Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate language
	if req.Language != models.LanguageGo && req.Language != models.LanguagePHP {
		return nil, fmt.Errorf("unsupported language: %s", req.Language)
	}

	// Create new project
	project := &models.Project{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Language:    req.Language,
		Description: req.Description,
		Options:     req.Options,
		Files:       []models.ProjectFile{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Set default options based on language
	if err := s.setDefaultOptions(project); err != nil {
		return nil, fmt.Errorf("failed to set default options: %w", err)
	}

	// Store project
	s.projects[project.ID] = project

	return project, nil
}

func (s *ProjectService) GetProject(projectID string) (*models.Project, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	project, exists := s.projects[projectID]
	if !exists {
		return nil, fmt.Errorf("project not found: %s", projectID)
	}

	return project, nil
}

func (s *ProjectService) GenerateProjectFiles(project *models.Project) ([]models.ProjectFile, error) {
	var files []models.ProjectFile
	var err error

	switch project.Language {
	case models.LanguageGo:
		files, err = s.templateService.GenerateGoProject(project)
	case models.LanguagePHP:
		files, err = s.templateService.GeneratePHPProject(project)
	default:
		return nil, fmt.Errorf("unsupported language: %s", project.Language)
	}

	if err != nil {
		return nil, err
	}

	// Update project with generated files
	s.mu.Lock()
	project.Files = files
	project.UpdatedAt = time.Now()
	s.mu.Unlock()

	return files, nil
}

func (s *ProjectService) CreateProjectZIP(projectID string) ([]byte, string, error) {
	project, err := s.GetProject(projectID)
	if err != nil {
		return nil, "", err
	}

	// Generate files if not already generated
	if len(project.Files) == 0 {
		_, err = s.GenerateProjectFiles(project)
		if err != nil {
			return nil, "", fmt.Errorf("failed to generate project files: %w", err)
		}
	}

	// Create ZIP file
	zipData, err := s.templateService.CreateZIPArchive(project)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create ZIP archive: %w", err)
	}

	filename := fmt.Sprintf("%s-%s.zip", project.Name, project.Language)
	return zipData, filename, nil
}

func (s *ProjectService) setDefaultOptions(project *models.Project) error {
	switch project.Language {
	case models.LanguageGo:
		// Set default Go options if not specified
		if project.Options.Framework == "" {
			project.Options.Framework = "gin"
		}
		if project.Options.Database == "" {
			project.Options.Database = "postgresql"
		}
		if project.Options.Authentication == "" {
			project.Options.Authentication = "jwt"
		}
		if len(project.Options.Utilities) == 0 {
			// Default to all utility packages
			project.Options.Utilities = []string{
				"authentication", "cache", "common", "constants", "converter",
				"date", "datatype", "encryption", "exception", "exceptioncode",
				"helper", "httphelper", "json", "logger", "password",
				"queryhelper", "sort", "template", "validator", "alert",
			}
		}

	case models.LanguagePHP:
		// Set default PHP options if not specified
		if project.Options.CIVersion == "" {
			project.Options.CIVersion = "3"
		}
		if project.Options.Database == "" {
			project.Options.Database = "postgresql"
		}
		if project.Options.Frontend == "" {
			project.Options.Frontend = "bootstrap"
		}
		if len(project.Options.Features) == 0 {
			project.Options.Features = []string{"authentication", "user_management", "dashboard"}
		}
	}

	return nil
}

func (s *ProjectService) ListProjects() []*models.Project {
	s.mu.RLock()
	defer s.mu.RUnlock()

	projects := make([]*models.Project, 0, len(s.projects))
	for _, project := range s.projects {
		projects = append(projects, project)
	}

	return projects
}
