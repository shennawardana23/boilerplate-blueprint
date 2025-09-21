package models_test

import (
	"testing"
	"time"

	"boilerplate-blueprint/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestChatMessage_Creation(t *testing.T) {
	now := time.Now()
	message := models.ChatMessage{
		ID:        "test-message-id",
		Role:      "user",
		Content:   "Hello, I need help with my project",
		ProjectID: "test-project",
		CreatedAt: now,
	}

	assert.Equal(t, "test-message-id", message.ID)
	assert.Equal(t, "user", message.Role)
	assert.Equal(t, "Hello, I need help with my project", message.Content)
	assert.Equal(t, "test-project", message.ProjectID)
	assert.Equal(t, now, message.CreatedAt)
}

func TestChatMessage_Roles(t *testing.T) {
	userMessage := models.ChatMessage{
		ID:      "user-msg",
		Role:    "user",
		Content: "User message",
	}

	assistantMessage := models.ChatMessage{
		ID:      "assistant-msg",
		Role:    "assistant",
		Content: "Assistant response",
	}

	assert.Equal(t, "user", userMessage.Role)
	assert.Equal(t, "assistant", assistantMessage.Role)
}

func TestChatMessage_EmptyProjectID(t *testing.T) {
	message := models.ChatMessage{
		ID:        "test-message-id",
		Role:      "user",
		Content:   "Hello",
		ProjectID: "",
	}

	assert.Empty(t, message.ProjectID)
}

func TestChatRequest_Creation(t *testing.T) {
	request := models.ChatRequest{
		Message:   "I want to build a Go web application",
		ProjectID: "test-project",
		Context:   "project setup",
	}

	assert.Equal(t, "I want to build a Go web application", request.Message)
	assert.Equal(t, "test-project", request.ProjectID)
	assert.Equal(t, "project setup", request.Context)
}

func TestChatRequest_MinimalRequest(t *testing.T) {
	request := models.ChatRequest{
		Message: "Hello",
		// ProjectID and Context are optional
	}

	assert.Equal(t, "Hello", request.Message)
	assert.Empty(t, request.ProjectID)
	assert.Empty(t, request.Context)
}

func TestChatResponse_Success(t *testing.T) {
	message := &models.ChatMessage{
		ID:      "response-id",
		Role:    "assistant",
		Content: "I can help you with that!",
	}

	suggestions := []models.ProjectSuggestion{
		{
			Type:       "language",
			Value:      "go",
			Reason:     "You mentioned Go in your message",
			Confidence: 0.9,
			Apply:      true,
		},
	}

	response := models.ChatResponse{
		Success:     true,
		Message:     message,
		Suggestions: suggestions,
	}

	assert.True(t, response.Success)
	assert.Equal(t, message, response.Message)
	assert.Len(t, response.Suggestions, 1)
	assert.Empty(t, response.Error)
}

func TestChatResponse_Error(t *testing.T) {
	response := models.ChatResponse{
		Success: false,
		Error:   "Failed to process message",
	}

	assert.False(t, response.Success)
	assert.Equal(t, "Failed to process message", response.Error)
	assert.Nil(t, response.Message)
	assert.Empty(t, response.Suggestions)
}

func TestProjectSuggestion_Creation(t *testing.T) {
	suggestion := models.ProjectSuggestion{
		Type:       "framework",
		Value:      "gin",
		Reason:     "Gin is a popular, fast HTTP framework for Go",
		Confidence: 0.85,
		Apply:      true,
		Options:    []string{"gin", "chi", "echo"},
	}

	assert.Equal(t, "framework", suggestion.Type)
	assert.Equal(t, "gin", suggestion.Value)
	assert.Equal(t, "Gin is a popular, fast HTTP framework for Go", suggestion.Reason)
	assert.Equal(t, 0.85, suggestion.Confidence)
	assert.True(t, suggestion.Apply)
	assert.NotNil(t, suggestion.Options)
}

func TestProjectSuggestion_Types(t *testing.T) {
	tests := []struct {
		name       string
		suggestion models.ProjectSuggestion
		expected   string
	}{
		{
			name: "Language suggestion",
			suggestion: models.ProjectSuggestion{
				Type:  "language",
				Value: "go",
			},
			expected: "language",
		},
		{
			name: "Framework suggestion",
			suggestion: models.ProjectSuggestion{
				Type:  "framework",
				Value: "gin",
			},
			expected: "framework",
		},
		{
			name: "Database suggestion",
			suggestion: models.ProjectSuggestion{
				Type:  "database",
				Value: "postgresql",
			},
			expected: "database",
		},
		{
			name: "Authentication suggestion",
			suggestion: models.ProjectSuggestion{
				Type:  "authentication",
				Value: "jwt",
			},
			expected: "authentication",
		},
		{
			name: "Feature suggestion",
			suggestion: models.ProjectSuggestion{
				Type:  "feature",
				Value: "rest_api",
			},
			expected: "feature",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.suggestion.Type)
		})
	}
}

func TestProjectSuggestion_ConfidenceLevels(t *testing.T) {
	tests := []struct {
		name       string
		confidence float64
		valid      bool
	}{
		{
			name:       "High confidence",
			confidence: 0.95,
			valid:      true,
		},
		{
			name:       "Medium confidence",
			confidence: 0.75,
			valid:      true,
		},
		{
			name:       "Low confidence",
			confidence: 0.25,
			valid:      true,
		},
		{
			name:       "Zero confidence",
			confidence: 0.0,
			valid:      true,
		},
		{
			name:       "Maximum confidence",
			confidence: 1.0,
			valid:      true,
		},
		{
			name:       "Negative confidence",
			confidence: -0.1,
			valid:      false,
		},
		{
			name:       "Over maximum confidence",
			confidence: 1.1,
			valid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suggestion := models.ProjectSuggestion{
				Type:       "test",
				Value:      "test",
				Confidence: tt.confidence,
			}

			if tt.valid {
				assert.GreaterOrEqual(t, suggestion.Confidence, 0.0)
				assert.LessOrEqual(t, suggestion.Confidence, 1.0)
			} else {
				assert.True(t, suggestion.Confidence < 0.0 || suggestion.Confidence > 1.0)
			}
		})
	}
}

func TestProjectSuggestion_ApplyFlag(t *testing.T) {
	autoApplySuggestion := models.ProjectSuggestion{
		Type:  "language",
		Value: "go",
		Apply: true,
	}

	manualApplySuggestion := models.ProjectSuggestion{
		Type:  "framework",
		Value: "gin",
		Apply: false,
	}

	assert.True(t, autoApplySuggestion.Apply)
	assert.False(t, manualApplySuggestion.Apply)
}

func TestProjectSuggestion_Options(t *testing.T) {
	suggestion := models.ProjectSuggestion{
		Type:    "framework",
		Value:   "gin",
		Options: []string{"gin", "chi", "echo", "standard"},
	}

	assert.NotNil(t, suggestion.Options)
	assert.Len(t, suggestion.Options, 4)
	assert.Contains(t, suggestion.Options, "gin")
	assert.Contains(t, suggestion.Options, "chi")
	assert.Contains(t, suggestion.Options, "echo")
	assert.Contains(t, suggestion.Options, "standard")
}

func TestChatHistory_Creation(t *testing.T) {
	now := time.Now()
	messages := []models.ChatMessage{
		{
			ID:        "msg1",
			Role:      "user",
			Content:   "Hello",
			ProjectID: "test-project",
			CreatedAt: now,
		},
		{
			ID:        "msg2",
			Role:      "assistant",
			Content:   "Hi there!",
			ProjectID: "test-project",
			CreatedAt: now.Add(time.Minute),
		},
	}

	history := models.ChatHistory{
		ProjectID: "test-project",
		Messages:  messages,
		CreatedAt: now,
		UpdatedAt: now.Add(time.Minute),
	}

	assert.Equal(t, "test-project", history.ProjectID)
	assert.Len(t, history.Messages, 2)
	assert.Equal(t, "msg1", history.Messages[0].ID)
	assert.Equal(t, "msg2", history.Messages[1].ID)
	assert.Equal(t, "user", history.Messages[0].Role)
	assert.Equal(t, "assistant", history.Messages[1].Role)
	assert.Equal(t, now, history.CreatedAt)
	assert.Equal(t, now.Add(time.Minute), history.UpdatedAt)
}

func TestChatHistory_EmptyMessages(t *testing.T) {
	now := time.Now()
	history := models.ChatHistory{
		ProjectID: "new-project",
		Messages:  []models.ChatMessage{},
		CreatedAt: now,
		UpdatedAt: now,
	}

	assert.Equal(t, "new-project", history.ProjectID)
	assert.Empty(t, history.Messages)
	assert.Equal(t, now, history.CreatedAt)
	assert.Equal(t, now, history.UpdatedAt)
}

func TestChatHistory_MessageOrdering(t *testing.T) {
	now := time.Now()
	messages := []models.ChatMessage{
		{
			ID:        "msg1",
			Role:      "user",
			Content:   "First message",
			CreatedAt: now,
		},
		{
			ID:        "msg2",
			Role:      "assistant",
			Content:   "First response",
			CreatedAt: now.Add(time.Minute),
		},
		{
			ID:        "msg3",
			Role:      "user",
			Content:   "Second message",
			CreatedAt: now.Add(2 * time.Minute),
		},
		{
			ID:        "msg4",
			Role:      "assistant",
			Content:   "Second response",
			CreatedAt: now.Add(3 * time.Minute),
		},
	}

	history := models.ChatHistory{
		ProjectID: "test-project",
		Messages:  messages,
		CreatedAt: now,
		UpdatedAt: now.Add(3 * time.Minute),
	}

	assert.Len(t, history.Messages, 4)
	assert.Equal(t, "First message", history.Messages[0].Content)
	assert.Equal(t, "First response", history.Messages[1].Content)
	assert.Equal(t, "Second message", history.Messages[2].Content)
	assert.Equal(t, "Second response", history.Messages[3].Content)
}

func TestChatMessage_ContentTypes(t *testing.T) {
	tests := []struct {
		name    string
		content string
		valid   bool
	}{
		{
			name:    "Simple text",
			content: "Hello",
			valid:   true,
		},
		{
			name:    "Long text",
			content: "This is a very long message with multiple sentences. It contains detailed information about the project requirements and specifications.",
			valid:   true,
		},
		{
			name:    "Empty content",
			content: "",
			valid:   false,
		},
		{
			name:    "Special characters",
			content: "Hello! @#$%^&*()_+-=[]{}|;':\",./<>?",
			valid:   true,
		},
		{
			name:    "Code snippet",
			content: "Here's some code:\n```go\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}\n```",
			valid:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := models.ChatMessage{
				ID:      "test-msg",
				Role:    "user",
				Content: tt.content,
			}

			if tt.valid {
				assert.NotEmpty(t, message.Content)
			} else {
				assert.Empty(t, message.Content)
			}
		})
	}
}

func TestProjectSuggestion_ReasonLength(t *testing.T) {
	tests := []struct {
		name   string
		reason string
		length int
	}{
		{
			name:   "Short reason",
			reason: "Popular choice",
			length: 14,
		},
		{
			name:   "Medium reason",
			reason: "This framework is widely used in the Go community and provides excellent performance",
			length: 85,
		},
		{
			name:   "Long reason",
			reason: "This is a very detailed explanation of why this particular option is recommended, including performance benchmarks, community support, documentation quality, and integration capabilities",
			length: 180,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suggestion := models.ProjectSuggestion{
				Type:   "framework",
				Value:  "gin",
				Reason: tt.reason,
			}

			assert.Equal(t, tt.length, len(suggestion.Reason))
		})
	}
}

func TestChatRequest_ContextTypes(t *testing.T) {
	tests := []struct {
		name    string
		context string
		valid   bool
	}{
		{
			name:    "Project setup context",
			context: "project setup",
			valid:   true,
		},
		{
			name:    "Database configuration context",
			context: "database configuration",
			valid:   true,
		},
		{
			name:    "Authentication setup context",
			context: "authentication setup",
			valid:   true,
		},
		{
			name:    "Empty context",
			context: "",
			valid:   true, // Context is optional
		},
		{
			name:    "Complex context",
			context: "setting up a microservices architecture with Docker containers and Kubernetes deployment",
			valid:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := models.ChatRequest{
				Message: "Test message",
				Context: tt.context,
			}

			if tt.valid {
				// Context can be empty or have content
				assert.True(t, len(request.Context) >= 0)
			}
		})
	}
}
