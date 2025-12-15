package main

import (
	"ainyx-backend/config"
	"ainyx-backend/db/sqlc"
	"ainyx-backend/internal/handler"
	"ainyx-backend/internal/logger"
	"ainyx-backend/internal/middleware"
	"ainyx-backend/internal/routes"
	"ainyx-backend/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
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

	
	queries := sqlc.New(dbConn)

	
	userService := service.NewUserService(queries)
	userHandler := handler.NewUserHandler(userService)

	
	app := fiber.New(fiber.Config{
		AppName: "Ainyx User API",
	})

	
	app.Use(middleware.RequestLogger)

	
	routes.SetupRoutes(app, userHandler)

	
	logger.Log.Info("Server starting on port 8080")
	log.Fatal(app.Listen(":8080"))
}