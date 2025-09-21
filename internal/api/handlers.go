package api

import (
	"fmt"
	"net/http"

	"boilerplate-blueprint/internal/models"
	"boilerplate-blueprint/internal/services"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	projectService  *services.ProjectService
	templateService *services.TemplateService
	chatService     *services.ChatService
}

func NewHandlers(projectService *services.ProjectService, templateService *services.TemplateService, chatService *services.ChatService) *Handlers {
	return &Handlers{
		projectService:  projectService,
		templateService: templateService,
		chatService:     chatService,
	}
}

// Health check endpoint
// @Summary Health check
// @Description Health check endpoint
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/health [get]
func (h *Handlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "boilerplate-blueprint",
		"version": "1.0.0",
	})
}

// Get available templates
// @Summary Get templates
// @Description Get available project templates
// @Tags Templates
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/templates [get]
func (h *Handlers) GetTemplates(c *gin.Context) {
	templates := h.templateService.GetAvailableTemplates()
	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"templates": templates,
	})
}

// Create a new project
// @Summary Create project
// @Description Create a new project
// @Tags Projects
// @Accept json
// @Produce json
// @Param request body models.ProjectRequest true "Project creation request"
// @Success 201 {object} models.ProjectResponse
// @Failure 400 {object} models.ProjectResponse
// @Failure 500 {object} models.ProjectResponse
// @Router /api/projects [post]
func (h *Handlers) CreateProject(c *gin.Context) {
	var req models.ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ProjectResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	project, err := h.projectService.CreateProject(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ProjectResponse{
			Success: false,
			Error:   "Failed to create project: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.ProjectResponse{
		Success: true,
		Message: "Project created successfully",
		Project: project,
	})
}

// Get project by ID
func (h *Handlers) GetProject(c *gin.Context) {
	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, models.ProjectResponse{
			Success: false,
			Error:   "Project ID is required",
		})
		return
	}

	project, err := h.projectService.GetProject(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ProjectResponse{
			Success: false,
			Error:   "Project not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.ProjectResponse{
		Success: true,
		Project: project,
	})
}

// Generate project files
func (h *Handlers) GenerateProject(c *gin.Context) {
	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Project ID is required",
		})
		return
	}

	project, err := h.projectService.GetProject(projectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Project not found",
		})
		return
	}

	files, err := h.projectService.GenerateProjectFiles(project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to generate project files: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Project files generated successfully",
		"files":   files,
	})
}

// Download project as ZIP
func (h *Handlers) DownloadProject(c *gin.Context) {
	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Project ID is required",
		})
		return
	}

	zipData, filename, err := h.projectService.CreateProjectZIP(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create ZIP: " + err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Length", fmt.Sprintf("%d", len(zipData)))
	c.Data(http.StatusOK, "application/zip", zipData)
}

// Send chat message
func (h *Handlers) ChatMessage(c *gin.Context) {
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ChatResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
		})
		return
	}

	response, err := h.chatService.ProcessMessage(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ChatResponse{
			Success: false,
			Error:   "Failed to process message: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get chat history
func (h *Handlers) GetChatHistory(c *gin.Context) {
	projectID := c.Query("project_id")

	history, err := h.chatService.GetChatHistory(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get chat history: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"history": history,
	})
}
