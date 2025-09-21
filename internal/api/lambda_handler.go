package api

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

// StartLambda starts the Lambda handler
func StartLambda(handlers *Handlers) {
	// Set Gin to release mode for Lambda
	gin.SetMode(gin.ReleaseMode)

	// Create router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configure CORS for Lambda
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Setup routes
	SetupRoutes(router, handlers)

	// Start Lambda with Gin adapter
	lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		// Create a simple response for now
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusOK,
			Body:       "Lambda integration coming soon. Use local deployment for now.",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}, nil
	})
}
