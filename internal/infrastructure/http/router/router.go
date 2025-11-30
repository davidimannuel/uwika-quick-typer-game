package router

import (
	"uwika_quick_typer_game/internal/application/services"
	"uwika_quick_typer_game/internal/domain/repositories"
	"uwika_quick_typer_game/internal/infrastructure/http/handlers"
	"uwika_quick_typer_game/internal/infrastructure/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authService *services.AuthService,
	gameService *services.GameService,
	adminService *services.AdminService,
	userRepo repositories.UserRepository,
) *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	gameHandler := handlers.NewGameHandler(gameService, userRepo)
	adminHandler := handlers.NewAdminHandler(adminService)

	// Public routes
	api := r.Group("/api")
	{
		// Auth endpoints
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/profile", middleware.AuthMiddleware(authService), authHandler.Profile)
		}

		// Game endpoints (require authentication)
		game := api.Group("")
		game.Use(middleware.AuthMiddleware(authService))
		{
			game.GET("/stages", gameHandler.GetStages)
			game.GET("/stage/:id", gameHandler.GetStageDetail)
			game.POST("/score/submit", gameHandler.SubmitScore)
			game.GET("/leaderboard", gameHandler.GetLeaderboard)
		}
	}

	// Admin routes (require admin authentication)
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(authService))
	admin.Use(middleware.AdminMiddleware())
	{
		// Theme endpoints (read-only for admin)
		admin.GET("/themes", adminHandler.GetAllThemes)

		// Stage management
		admin.POST("/stage", adminHandler.CreateStage)
		admin.PUT("/stage/:id", adminHandler.UpdateStage)
		admin.DELETE("/stage/:id", adminHandler.DeleteStage)
		admin.GET("/stages", adminHandler.GetAllStages)

		// Phrase management
		admin.POST("/phrase", adminHandler.CreatePhrase)
		admin.PUT("/phrase/:id", adminHandler.UpdatePhrase)
		admin.DELETE("/phrase/:id", adminHandler.DeletePhrase)
		admin.GET("/phrases", adminHandler.GetPhrasesByStage)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
