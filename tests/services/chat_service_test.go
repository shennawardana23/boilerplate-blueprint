package services_test

import (
	"fmt"
	"testing"

	"boilerplate-blueprint/internal/models"
	"boilerplate-blueprint/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewChatService(t *testing.T) {
	service := services.NewChatService()

	assert.NotNil(t, service)
}

func TestChatService_ProcessMessage_GoLanguage(t *testing.T) {
	service := services.NewChatService()

	req := &models.ChatRequest{
		Message:   "I want to build a Go web application",
		ProjectID: "test-project",
		Context:   "project setup",
	}

	response, err := service.ProcessMessage(req)

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)
	assert.NotNil(t, response.Message)
	assert.Equal(t, "assistant", response.Message.Role)
	assert.Contains(t, response.Message.Content, "Go")
	assert.NotEmpty(t, response.Suggestions)

	// Check for Go language suggestion
	var goSuggestion *models.ProjectSuggestion
	for _, suggestion := range response.Suggestions {
		if suggestion.Type == "language" && suggestion.Value == "go" {
			goSuggestion = &suggestion
			break
		}
	}
	assert.NotNil(t, goSuggestion)
	assert.Equal(t, "go", goSuggestion.Value)
	assert.Equal(t, "language", goSuggestion.Type)
	assert.Greater(t, goSuggestion.Confidence, 0.8)
}

func TestChatService_ProcessMessage_PHPLanguage(t *testing.T) {
	service := services.NewChatService()

	req := &models.ChatRequest{
		Message:   "I need a PHP CodeIgniter project",
		ProjectID: "test-project",
		Context:   "project setup",
	}

	response, err := service.ProcessMessage(req)

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)
	assert.NotNil(t, response.Message)
	assert.Equal(t, "assistant", response.Message.Role)
	assert.Contains(t, response.Message.Content, "PHP")
	assert.NotEmpty(t, response.Suggestions)

	// Check for PHP language suggestion
	var phpSuggestion *models.ProjectSuggestion
	for _, suggestion := range response.Suggestions {
		if suggestion.Type == "language" && suggestion.Value == "php" {
			phpSuggestion = &suggestion
			break
		}
	}
	assert.NotNil(t, phpSuggestion)
	assert.Equal(t, "php", phpSuggestion.Value)
	assert.Equal(t, "language", phpSuggestion.Type)
}

func TestChatService_ProcessMessage_DatabaseDetection(t *testing.T) {
	service := services.NewChatService()

	req := &models.ChatRequest{
		Message:   "I want to use PostgreSQL database",
		ProjectID: "test-project",
		Context:   "database setup",
	}

	response, err := service.ProcessMessage(req)

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)

	// Check for database suggestion
	var dbSuggestion *models.ProjectSuggestion
	for _, suggestion := range response.Suggestions {
		if suggestion.Type == "database" && suggestion.Value == "postgresql" {
			dbSuggestion = &suggestion
			break
		}
	}
	assert.NotNil(t, dbSuggestion)
	assert.Equal(t, "postgresql", dbSuggestion.Value)
	assert.Equal(t, "database", dbSuggestion.Type)
	assert.Greater(t, dbSuggestion.Confidence, 0.9)
}

func TestChatService_ProcessMessage_DefaultResponse(t *testing.T) {
	service := services.NewChatService()

	req := &models.ChatRequest{
		Message:   "Hello, how are you?",
		ProjectID: "test-project",
		Context:   "greeting",
	}

	response, err := service.ProcessMessage(req)

	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)
	assert.NotNil(t, response.Message)
	assert.Contains(t, response.Message.Content, "boilerplate")
	assert.Contains(t, response.Message.Content, "Go")
	assert.Contains(t, response.Message.Content, "PHP")
}

func TestChatService_GetChatHistory_ExistingProject(t *testing.T) {
	service := services.NewChatService()

	// Process a message first to create history
	req := &models.ChatRequest{
		Message:   "Test message",
		ProjectID: "test-project",
	}

	_, err := service.ProcessMessage(req)
	require.NoError(t, err)

	// Get chat history
	history, err := service.GetChatHistory("test-project")

	require.NoError(t, err)
	assert.NotNil(t, history)
	assert.Equal(t, "test-project", history.ProjectID)
	assert.Len(t, history.Messages, 2) // User message + assistant response
	assert.False(t, history.CreatedAt.IsZero())
	assert.False(t, history.UpdatedAt.IsZero())

	// Check message roles
	assert.Equal(t, "user", history.Messages[0].Role)
	assert.Equal(t, "assistant", history.Messages[1].Role)
}

func TestChatService_GetChatHistory_NewProject(t *testing.T) {
	service := services.NewChatService()

	// Get chat history for non-existent project
	history, err := service.GetChatHistory("new-project")

	require.NoError(t, err)
	assert.NotNil(t, history)
	assert.Equal(t, "new-project", history.ProjectID)
	assert.Empty(t, history.Messages)
	assert.False(t, history.CreatedAt.IsZero())
}

func TestChatService_GetChatHistory_EmptyProjectID(t *testing.T) {
	service := services.NewChatService()

	// Get chat history with empty project ID
	history, err := service.GetChatHistory("")

	require.NoError(t, err)
	assert.NotNil(t, history)
	assert.Equal(t, "general", history.ProjectID)
	assert.Empty(t, history.Messages)
}

func TestChatService_ConcurrentAccess(t *testing.T) {
	service := services.NewChatService()

	// Test concurrent message processing
	done := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			req := &models.ChatRequest{
				Message:   fmt.Sprintf("Concurrent message %d", i),
				ProjectID: "concurrent-project",
			}
			_, err := service.ProcessMessage(req)
			assert.NoError(t, err)
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 5; i++ {
		<-done
	}

	// Verify all messages were processed
	history, err := service.GetChatHistory("concurrent-project")
	require.NoError(t, err)
	assert.Len(t, history.Messages, 10) // 5 user messages + 5 assistant responses
}

func TestChatService_MessageOrdering(t *testing.T) {
	service := services.NewChatService()

	// Process multiple messages
	req1 := &models.ChatRequest{
		Message:   "First message",
		ProjectID: "order-test",
	}
	req2 := &models.ChatRequest{
		Message:   "Second message",
		ProjectID: "order-test",
	}

	_, err := service.ProcessMessage(req1)
	require.NoError(t, err)
	_, err = service.ProcessMessage(req2)
	require.NoError(t, err)

	history, err := service.GetChatHistory("order-test")
	require.NoError(t, err)

	// Verify message ordering
	assert.Len(t, history.Messages, 4) // 2 user + 2 assistant
	assert.Equal(t, "user", history.Messages[0].Role)
	assert.Equal(t, "assistant", history.Messages[1].Role)
	assert.Equal(t, "user", history.Messages[2].Role)
	assert.Equal(t, "assistant", history.Messages[3].Role)
	assert.Equal(t, "First message", history.Messages[0].Content)
	assert.Equal(t, "Second message", history.Messages[2].Content)
}
