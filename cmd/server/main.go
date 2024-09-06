package main

import (
	"log"

	"dalabio/internal/framework/driver/db"
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/interface_adapter/gateway"
	"dalabio/internal/interface_adapter/routes"
	"dalabio/internal/service"
	"dalabio/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file (it will look for .env in the root directory)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using environment variables.")
	}
}

func main() {
	// Load the database configuration from environment variables or .env
	dbConfig := config.LoadDBConfig()

	// Debug: Print the loaded database configuration
	log.Printf("DB Config: Host=%s, Port=%s, User=%s, Password=%s, DBName=%s, SSLMode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)

	// Connect to the database
	database, err := db.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Create the required tables if they don't exist
	if err := db.CreateTables(database); err != nil {
		log.Fatal(err)
	}

	// Initialize the repositories
	userRepository := gateway.NewUserRepository(database)
	tokenRepository := gateway.NewTokenRepository(database)

	// Initialize the services
	userService := service.NewUserService(userRepository, tokenRepository)

	// Initialize the controllers
	userController := controller.NewUserController(userService)

	// Initialize Gin router
	r := gin.Default()

	// Register user-related routes with token repository for middleware
	routes.RegisterUserRoutes(r, userController, tokenRepository)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
