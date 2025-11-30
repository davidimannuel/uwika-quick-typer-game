package main

import (
	"log"
	"os"

	"uwika_quick_typer_game/internal/application/services"
	"uwika_quick_typer_game/internal/infrastructure/database"
	"uwika_quick_typer_game/internal/infrastructure/http/router"
	"uwika_quick_typer_game/internal/infrastructure/persistence/postgres"
)

func main() {
	// Database configuration
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "s3cret"),
		DBName:   getEnv("DB_NAME", "quick_typer"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Connect to database
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	tokenRepo := postgres.NewTokenRepository(db)
	themeRepo := postgres.NewThemeRepository(db)
	stageRepo := postgres.NewStageRepository(db)
	phraseRepo := postgres.NewPhraseRepository(db)
	scoreRepo := postgres.NewScoreRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, tokenRepo)
	gameService := services.NewGameService(stageRepo, phraseRepo, scoreRepo)
	adminService := services.NewAdminService(stageRepo, phraseRepo, userRepo, themeRepo)

	// Setup router
	r := router.SetupRouter(authService, gameService, adminService, userRepo)

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

