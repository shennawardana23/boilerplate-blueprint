package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *Handlers) {
	// API group
	api := router.Group("/api")
	{
		// Health check
		api.GET("/health", handlers.Health)

		// Template endpoints
		api.GET("/templates", handlers.GetTemplates)

		// Project endpoints
		api.POST("/projects", handlers.CreateProject)
		api.GET("/projects/:id", handlers.GetProject)
		api.POST("/projects/:id/generate", handlers.GenerateProject)
		api.GET("/projects/:id/download", handlers.DownloadProject)

		// Chat endpoints
		api.POST("/chat/message", handlers.ChatMessage)
		api.GET("/chat/history", handlers.GetChatHistory)
	}
}
