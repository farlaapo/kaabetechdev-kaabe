package main

import (
	"log"

	"dalabio/internal/framework/driver/db"
	"dalabio/internal/interface_adapter/controller"
	"dalabio/internal/interface_adapter/gateway"
	"dalabio/internal/interface_adapter/routes"
	"dalabio/internal/service"
	"dalabio/pkg/config"

	"github.com/gin-contrib/cors" // Import CORS package
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
		log.Fatal("Error connecting to the database:", err)
	}
	defer database.Close()

	// Create the required tables if they don't exist
	if err := db.CreateTables(database); err != nil {
		log.Fatal("Error creating tables:", err)
	}

	// Initialize the repositories
	userRepository := gateway.NewUserRepository(database)
	tokenRepository := gateway.NewTokenRepository(database)
	courseRepository := gateway.NewCourseRepository(database)
	SpaceRepository := gateway.NewSpaceRepository(database)
	meetingRepository := gateway.NewMeetingRepository(database)
	paymentRepository := gateway.NewPaymentRepository(database)

	// Initialize the services
	userService := service.NewUserService(userRepository, tokenRepository)
	courseService := service.NewCourseService(courseRepository, tokenRepository)
	spaceService := service.NewSpaceService(SpaceRepository, tokenRepository)
	meetingService := service.NewMeetingService(meetingRepository, tokenRepository)
	paymentService := service.NewPaymentService(paymentRepository, tokenRepository)

	// Initialize the controllers
	userController := controller.NewUserController(userService)
	courseController := controller.NewCourseController(courseService)
	spaceController := controller.NewSpaceController(spaceService)
	meetingController := controller.NewMeetingController(meetingService)
	paymentController := controller.NewPaymentController(paymentService)

	// Initialize Gin router
	r := gin.Default()

	// Apply CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register user-related routes with token repository for middleware
	routes.RegisterUserRoutes(r, userController, tokenRepository)
	routes.RegisterCoursesRoutes(r, courseController, tokenRepository)
	routes.RegisterSpacesRoutes(r, spaceController, tokenRepository)
	routes.RegisterMeetingRoutes(r, meetingController, tokenRepository)
	routes.RegisterPaymentRoutes(r, paymentController, tokenRepository)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
