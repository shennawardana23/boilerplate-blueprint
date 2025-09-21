package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"boilerplate-blueprint/internal/api"
	"boilerplate-blueprint/internal/models"
	"boilerplate-blueprint/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestHandlers() *api.Handlers {
	templateService := services.NewTemplateService()
	projectService := services.NewProjectService(templateService)
	chatService := services.NewChatService()
	return api.NewHandlers(projectService, templateService, chatService)
}

func TestNewHandlers(t *testing.T) {
	templateService := services.NewTemplateService()
	projectService := services.NewProjectService(templateService)
	chatService := services.NewChatService()

	handlers := api.NewHandlers(projectService, templateService, chatService)

	assert.NotNil(t, handlers)
}

func TestHandlers_Health(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handlers.Health)

	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "healthy", response["status"])
	assert.Equal(t, "boilerplate-blueprint", response["service"])
	assert.Equal(t, "1.0.0", response["version"])
}

func TestHandlers_GetTemplates(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/templates", handlers.GetTemplates)

	req, err := http.NewRequest("GET", "/templates", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["templates"])

	templates := response["templates"].([]interface{})
	assert.Len(t, templates, 2)

	// Check Go template
	goTemplate := templates[0].(map[string]interface{})
	assert.Equal(t, "go", goTemplate["language"])
	assert.Equal(t, "Go Clean Architecture", goTemplate["name"])

	// Check PHP template
	phpTemplate := templates[1].(map[string]interface{})
	assert.Equal(t, "php", phpTemplate["language"])
	assert.Equal(t, "PHP CodeIgniter MVC", phpTemplate["name"])
}

func TestHandlers_CreateProject_ValidGoProject(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/projects", handlers.CreateProject)

	projectReq := models.ProjectRequest{
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

	jsonData, err := json.Marshal(projectReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonData))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.ProjectResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Project)
	assert.Equal(t, projectReq.Name, response.Project.Name)
	assert.Equal(t, projectReq.Language, response.Project.Language)
	assert.Equal(t, projectReq.Description, response.Project.Description)
	assert.NotEmpty(t, response.Project.ID)
}

func TestHandlers_CreateProject_ValidPHPProject(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/projects", handlers.CreateProject)

	projectReq := models.ProjectRequest{
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

	jsonData, err := json.Marshal(projectReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonData))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.ProjectResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Project)
	assert.Equal(t, projectReq.Name, response.Project.Name)
	assert.Equal(t, projectReq.Language, response.Project.Language)
	assert.Equal(t, projectReq.Description, response.Project.Description)
	assert.NotEmpty(t, response.Project.ID)
}

func TestHandlers_CreateProject_InvalidJSON(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/projects", handlers.CreateProject)

	req, err := http.NewRequest("POST", "/projects", bytes.NewBufferString("invalid json"))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.ProjectResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.NotEmpty(t, response.Error)
}

func TestHandlers_GetProject_ValidID(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/projects", handlers.CreateProject)
	router.GET("/projects/:id", handlers.GetProject)

	// First create a project
	projectReq := models.ProjectRequest{
		Name:        "test-project",
		Language:    models.LanguageGo,
		Description: "A test project",
	}

	jsonData, err := json.Marshal(projectReq)
	require.NoError(t, err)

	createReq, err := http.NewRequest("POST", "/projects", bytes.NewBuffer(jsonData))
	require.NoError(t, err)
	createReq.Header.Set("Content-Type", "application/json")

	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)

	assert.Equal(t, http.StatusOK, createW.Code)

	var createResponse models.ProjectResponse
	err = json.Unmarshal(createW.Body.Bytes(), &createResponse)
	require.NoError(t, err)

	projectID := createResponse.Project.ID

	// Now get the project
	getReq, err := http.NewRequest("GET", fmt.Sprintf("/projects/%s", projectID), nil)
	require.NoError(t, err)

	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusOK, getW.Code)

	var getResponse models.ProjectResponse
	err = json.Unmarshal(getW.Body.Bytes(), &getResponse)
	require.NoError(t, err)

	assert.True(t, getResponse.Success)
	assert.NotNil(t, getResponse.Project)
	assert.Equal(t, projectID, getResponse.Project.ID)
	assert.Equal(t, projectReq.Name, getResponse.Project.Name)
}

func TestHandlers_GetProject_InvalidID(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/projects/:id", handlers.GetProject)

	req, err := http.NewRequest("GET", "/projects/non-existent-id", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response models.ProjectResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.NotEmpty(t, response.Error)
	assert.Contains(t, response.Error, "project not found")
}

func TestHandlers_ChatMessage_ValidMessage(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/chat/message", handlers.ChatMessage)

	chatReq := models.ChatRequest{
		Message:   "I want to build a Go web application",
		ProjectID: "test-project",
		Context:   "project setup",
	}

	jsonData, err := json.Marshal(chatReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/chat/message", bytes.NewBuffer(jsonData))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.ChatResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response.Success)
	assert.NotNil(t, response.Message)
	assert.Equal(t, "assistant", response.Message.Role)
	assert.NotEmpty(t, response.Message.Content)
	assert.NotEmpty(t, response.Suggestions)
}

func TestHandlers_ChatMessage_InvalidJSON(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/chat/message", handlers.ChatMessage)

	req, err := http.NewRequest("POST", "/chat/message", bytes.NewBufferString("invalid json"))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.ChatResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response.Success)
	assert.NotEmpty(t, response.Error)
}

func TestHandlers_GetChatHistory_ValidProjectID(t *testing.T) {
	handlers := setupTestHandlers()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/chat/message", handlers.ChatMessage)
	router.GET("/chat/history", handlers.GetChatHistory)

	// First send a message to create history
	chatReq := models.ChatRequest{
		Message:   "Hello, I need help with my project",
		ProjectID: "test-project",
	}

	jsonData, err := json.Marshal(chatReq)
	require.NoError(t, err)

	messageReq, err := http.NewRequest("POST", "/chat/message", bytes.NewBuffer(jsonData))
	require.NoError(t, err)
	messageReq.Header.Set("Content-Type", "application/json")

	messageW := httptest.NewRecorder()
	router.ServeHTTP(messageW, messageReq)

	assert.Equal(t, http.StatusOK, messageW.Code)

	// Now get chat history
	historyReq, err := http.NewRequest("GET", "/chat/history?project_id=test-project", nil)
	require.NoError(t, err)

	historyW := httptest.NewRecorder()
	router.ServeHTTP(historyW, historyReq)

	assert.Equal(t, http.StatusOK, historyW.Code)

	var response map[string]interface{}
	err = json.Unmarshal(historyW.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["history"])

	history := response["history"].(map[string]interface{})
	assert.Equal(t, "test-project", history["project_id"])
	assert.NotNil(t, history["messages"])
}
