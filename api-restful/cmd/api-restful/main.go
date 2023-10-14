package main

import (
	"context"
	"getting-started/api-restful/internal/adapter/handler"
	"getting-started/api-restful/internal/adapter/repository/postgres"
	"getting-started/api-restful/internal/core/service"
	"log/slog"
	"os"

	_ "getting-started/api-restful/docs"

	"github.com/joho/godotenv"
)

//	@title			Marvel API
//	@description	API about the Marvel Cinematic Universe (MCU)
//	@version		1.0.0
//
//	@contact.name	Kyrex
//	@contact.url	github.com/kyrex23
//
//	@host			localhost:8080
//	@basePath		/api/v1
func main() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("Error ocurred loading env file", "error", err)
		os.Exit(1)
	}

	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConnection := os.Getenv("DB_CONNECTION")
	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")
	httpUrl := httpHost + ":" + httpPort

	slog.Info("Starting the application", "app", appName, "env", env)

	// Init database
	ctx := context.Background()
	db, err := postgres.NewDB(ctx)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("Successfully connected to the database", "db", dbConnection)

	// ==================== DEPENDENCY INJECTION ====================
	heroRepository := postgres.NewHeroRepository(db)
	heroService := service.NewHeroService(heroRepository)
	heroHandler := handler.NewHeroHandler(heroService)

	router, err := handler.NewRouter(*heroHandler)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start HTTP Server
	slog.Info("Starting the HTTP server", "listen_address", httpUrl)
	if err = router.Serve(httpUrl); err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
