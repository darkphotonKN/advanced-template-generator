package config

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/darkphotonKN/go-template-generator/internal/item"
	"github.com/darkphotonKN/go-template-generator/internal/middleware"
)

func SetupRoutes(db *sqlx.DB, logger *slog.Logger) *gin.Engine {
	// Set Gin to release mode in production
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Add middleware
	router.Use(middleware.RequestID())
	router.Use(middleware.RequestLogger(logger))
	router.Use(middleware.StructuredLogger(logger))
	router.Use(corsMiddleware())

	// Initialize services
	repo := item.NewRepository(db)
	service := item.NewService(repo, logger)
	handler := item.NewHandler(service, logger)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes
	api := router.Group("/api")
	{
		// Item endpoints
		items := api.Group("/items")
		{
			items.POST("", handler.CreateItem)
			items.GET("", handler.ListItems)
			items.GET("/:id", handler.GetItem)
			items.PUT("/:id", handler.UpdateItem)
			items.DELETE("/:id", handler.DeleteItem)
		}
	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	return cors.New(config)
}