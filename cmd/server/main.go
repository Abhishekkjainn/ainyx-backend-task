package main

import (
	"ainyx-backend/config"
	"ainyx-backend/db/sqlc"
	"ainyx-backend/internal/handler"
	"ainyx-backend/internal/logger"
	"ainyx-backend/internal/middleware"
	"ainyx-backend/internal/repository"
	"ainyx-backend/internal/routes"
	"ainyx-backend/internal/service"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	logger.InitLogger()
	defer logger.Log.Sync()

	dbConn := config.ConnectDB()
	defer dbConn.Close()

	// Wiring Layers: DB -> Repo -> Service -> Handler
	queries := sqlc.New(dbConn)
	userRepo := repository.NewUserRepository(queries)     // Added Repo Layer
	userService := service.NewUserService(userRepo)       // Injected Repo
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New(fiber.Config{
		AppName: "Ainyx User API",
		// Add custom error handler if desired
	})

	// Middleware (Adopted from Kedar + Abhishek)
	app.Use(recover.New())          // Safety: Catch panics
	app.Use(cors.New())             // CORS support
	app.Use(middleware.RequestLogger) // Abhishek's Logger

	routes.SetupRoutes(app, userHandler)

	// Graceful Shutdown (Kedar's implementation)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start server in goroutine
	go func() {
		logger.Log.Info("Server starting on port 8080")
		if err := app.Listen(":8080"); err != nil {
			logger.Log.Fatal(err.Error())
		}
	}()

	// Wait for interrupt
	<-ctx.Done()
	logger.Log.Info("Shutting down server...")

	// Cleanup
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		logger.Log.Error(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	logger.Log.Info("Server exited gracefully")
}