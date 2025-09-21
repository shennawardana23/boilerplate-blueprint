package services

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"boilerplate-blueprint/internal/models"

	"github.com/google/uuid"
)

type ChatService struct {
	conversations map[string]*models.ChatHistory
	mu            sync.RWMutex
}

func NewChatService() *ChatService {
	return &ChatService{
		conversations: make(map[string]*models.ChatHistory),
	}
}

func (s *ChatService) ProcessMessage(req *models.ChatRequest) (*models.ChatResponse, error) {
	// Create user message
	userMessage := &models.ChatMessage{
		ID:        uuid.New().String(),
		Role:      "user",
		Content:   req.Message,
		ProjectID: req.ProjectID,
		CreatedAt: time.Now(),
	}

	// Store user message
	s.storeMessage(req.ProjectID, userMessage)

	// Process the message and generate AI response
	assistantMessage, suggestions, err := s.generateAIResponse(req, userMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to generate AI response: %w", err)
	}

	// Store assistant message
	s.storeMessage(req.ProjectID, assistantMessage)

	return &models.ChatResponse{
		Success:     true,
		Message:     assistantMessage,
		Suggestions: suggestions,
	}, nil
}

func (s *ChatService) GetChatHistory(projectID string) (*models.ChatHistory, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if projectID == "" {
		// Return all conversations or create a general one
		projectID = "general"
	}

	history, exists := s.conversations[projectID]
	if !exists {
		// Create new conversation
		history = &models.ChatHistory{
			ProjectID: projectID,
			Messages:  []models.ChatMessage{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	return history, nil
}

func (s *ChatService) storeMessage(projectID string, message *models.ChatMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if projectID == "" {
		projectID = "general"
	}

	history, exists := s.conversations[projectID]
	if !exists {
		history = &models.ChatHistory{
			ProjectID: projectID,
			Messages:  []models.ChatMessage{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		s.conversations[projectID] = history
	}

	history.Messages = append(history.Messages, *message)
	history.UpdatedAt = time.Now()
}

func (s *ChatService) generateAIResponse(req *models.ChatRequest, userMessage *models.ChatMessage) (*models.ChatMessage, []models.ProjectSuggestion, error) {
	// For now, we'll create a simple rule-based response system
	// In a real implementation, this would integrate with OpenAI API

	response, suggestions := s.generateRuleBasedResponse(req.Message, req.Context)

	assistantMessage := &models.ChatMessage{
		ID:        uuid.New().String(),
		Role:      "assistant",
		Content:   response,
		ProjectID: req.ProjectID,
		CreatedAt: time.Now(),
	}

	return assistantMessage, suggestions, nil
}

func (s *ChatService) generateRuleBasedResponse(message, context string) (string, []models.ProjectSuggestion) {
	message = strings.ToLower(message)
	var suggestions []models.ProjectSuggestion

	// Detect language preference
	if strings.Contains(message, "go") || strings.Contains(message, "golang") {
		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "language",
			Value:      "go",
			Reason:     "You mentioned Go/Golang in your message",
			Confidence: 0.9,
			Apply:      true,
		})

		// Suggest Go-specific frameworks
		if strings.Contains(message, "web") || strings.Contains(message, "api") || strings.Contains(message, "server") {
			suggestions = append(suggestions, models.ProjectSuggestion{
				Type:       "framework",
				Value:      "gin",
				Reason:     "Gin is a popular, fast HTTP framework for Go web applications",
				Confidence: 0.8,
				Apply:      true,
			})
		}

		return "Great choice! Go is excellent for building high-performance applications. I can help you set up a Go project with Clean Architecture, including all 17 utility packages for enterprise-grade development. What type of application are you building?", suggestions
	}

	if strings.Contains(message, "php") || strings.Contains(message, "codeigniter") {
		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "language",
			Value:      "php",
			Reason:     "You mentioned PHP or CodeIgniter in your message",
			Confidence: 0.9,
			Apply:      true,
		})

		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "ci_version",
			Value:      "3",
			Reason:     "CodeIgniter 3 is stable and widely used for enterprise applications",
			Confidence: 0.7,
			Apply:      false,
		})

		return "PHP with CodeIgniter is a solid choice for rapid web development! I can help you create a complete MVC application with security features, authentication, and a clean admin panel. Would you prefer CodeIgniter 3 or 4?", suggestions
	}

	// Database detection
	if strings.Contains(message, "postgres") || strings.Contains(message, "postgresql") {
		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "database",
			Value:      "postgresql",
			Reason:     "PostgreSQL is a robust, feature-rich database perfect for enterprise applications",
			Confidence: 0.95,
			Apply:      true,
		})
	} else if strings.Contains(message, "mysql") {
		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "database",
			Value:      "mysql",
			Reason:     "MySQL is widely supported and great for web applications",
			Confidence: 0.9,
			Apply:      true,
		})
	}

	// Authentication detection
	if strings.Contains(message, "auth") || strings.Contains(message, "login") || strings.Contains(message, "user") {
		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "authentication",
			Value:      "jwt",
			Reason:     "JWT is modern, stateless, and perfect for API authentication",
			Confidence: 0.8,
			Apply:      false,
		})
	}

	// Feature detection
	if strings.Contains(message, "api") || strings.Contains(message, "rest") {
		suggestions = append(suggestions, models.ProjectSuggestion{
			Type:       "feature",
			Value:      "rest_api",
			Reason:     "RESTful API structure detected in your requirements",
			Confidence: 0.85,
			Apply:      true,
		})
	}

	// Default response
	if len(suggestions) == 0 {
		return "Hello! I'm here to help you create amazing boilerplate projects. I can generate:\n\n🐹 **Go projects** with Clean Architecture, Gin framework, and 17 utility packages\n🐘 **PHP CodeIgniter projects** with MVC structure and security features\n\nWhat kind of project would you like to build today?", suggestions
	}

	// Generate response based on suggestions
	responseText := "Based on your message, I have some suggestions for your project configuration. "
	if len(suggestions) > 0 {
		responseText += "I've automatically detected some preferences - you can review and modify them in the project configuration form. "
	}
	responseText += "What other requirements do you have for this project?"

	return responseText, suggestions
}

// Mock function for future OpenAI integration
func (s *ChatService) callOpenAI(messages []models.ChatMessage, context string) (string, error) {
	// This would integrate with OpenAI API
	// For now, return a placeholder
	return "OpenAI integration not yet implemented", nil
}

// Helper function to format chat context for AI
func (s *ChatService) formatContextForAI(projectID string) string {
	history, _ := s.GetChatHistory(projectID)
	if history == nil || len(history.Messages) == 0 {
		return "New conversation"
	}

	// Format recent messages for context
	context := "Recent conversation:\n"
	start := len(history.Messages) - 5 // Last 5 messages
	if start < 0 {
		start = 0
	}

	for _, msg := range history.Messages[start:] {
		context += fmt.Sprintf("%s: %s\n", msg.Role, msg.Content)
	}

	return context
}
