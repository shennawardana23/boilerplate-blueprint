package main

import (
	"log"
	"os"
	"runtime"

	"boilerplate-blueprint/internal/api"
	"boilerplate-blueprint/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Check if running in Lambda environment
	if isLambdaEnvironment() {
		log.Println("🚀 Starting Boilerplate Blueprint in AWS Lambda mode...")

		// Initialize services
		templateService := services.NewTemplateService()
		projectService := services.NewProjectService(templateService)
		chatService := services.NewChatService()

		// Initialize handlers
		handlers := api.NewHandlers(projectService, templateService, chatService)

		// Start Lambda handler
		api.StartLambda(handlers)
		return
	}

	// Regular server mode
	startServer()
}

// isLambdaEnvironment checks if we're running in AWS Lambda
func isLambdaEnvironment() bool {
	// Check for Lambda environment variables
	lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT")
	lambdaRuntimeAPI := os.Getenv("AWS_LAMBDA_RUNTIME_API")

	return lambdaTaskRoot != "" || lambdaRuntimeAPI != ""
}

// startServer starts the regular HTTP server
func startServer() {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}

	// Log startup information
	log.Printf("🚀 Starting Boilerplate Blueprint server...")
	log.Printf("📊 Go Version: %s", runtime.Version())
	log.Printf("🖥️  OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)

	// Initialize services
	templateService := services.NewTemplateService()
	projectService := services.NewProjectService(templateService)
	chatService := services.NewChatService()

	// Initialize handlers
	handlers := api.NewHandlers(projectService, templateService, chatService)

	// Create Gin router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://localhost:5173",
		"https://localhost:3000",
		"https://localhost:5173",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	// Setup routes
	api.SetupRoutes(router, handlers)

	// Serve static files (Vue.js build)
	router.Static("/static", "./web/dist")
	router.StaticFile("/", "./web/dist/index.html")

	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("📡 Server will be available at: http://localhost:%s", port)
	log.Printf("🌐 Frontend will be available at: http://localhost:%s", port)

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
