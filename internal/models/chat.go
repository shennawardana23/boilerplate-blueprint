package models

import "time"

// ChatMessage represents a chat message
type ChatMessage struct {
	ID        string    `json:"id"`
	Role      string    `json:"role"` // "user" or "assistant"
	Content   string    `json:"content"`
	ProjectID string    `json:"project_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// ChatRequest represents a chat API request
type ChatRequest struct {
	Message   string `json:"message" binding:"required"`
	ProjectID string `json:"project_id,omitempty"`
	Context   string `json:"context,omitempty"`
}

// ChatResponse represents a chat API response
type ChatResponse struct {
	Success     bool                `json:"success"`
	Message     *ChatMessage        `json:"message,omitempty"`
	Error       string              `json:"error,omitempty"`
	Suggestions []ProjectSuggestion `json:"suggestions,omitempty"`
}

// ProjectSuggestion represents AI suggestions for project configuration
type ProjectSuggestion struct {
	Type       string      `json:"type"` // "framework", "database", "feature", etc.
	Value      string      `json:"value"`
	Reason     string      `json:"reason"`
	Confidence float64     `json:"confidence"` // 0-1
	Apply      bool        `json:"apply"`      // Whether to auto-apply
	Options    interface{} `json:"options,omitempty"`
}

// ChatHistory represents a conversation history
type ChatHistory struct {
	ProjectID string        `json:"project_id"`
	Messages  []ChatMessage `json:"messages"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
