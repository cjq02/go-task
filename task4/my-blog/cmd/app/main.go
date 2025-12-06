package main

import (
	"io"
	"log"
	"os"

	"blog-backend/internal/app"
	"blog-backend/internal/config"
	"blog-backend/internal/logger"
	"blog-backend/internal/model"
	"blog-backend/internal/service"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = os.Getenv("APP_ENV")
	}
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		if err := godotenv.Load(".env.local"); err != nil {
			if err := godotenv.Load(); err != nil {
				log.Printf("Note: No .env file found (tried: %s, .env.local, .env), using environment variables or defaults", envFile)
			} else {
				log.Printf("Loaded .env file")
			}
		} else {
			log.Printf("Loaded .env.local file")
		}
	} else {
		log.Printf("Loaded %s file", envFile)
	}

	// Initialize database
	db, err := initDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize logger
	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "./logs"
	}
	appLogger, err := logger.NewLogger(logDir)
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer appLogger.Close()

	// Redirect standard log output
	log.SetOutput(io.MultiWriter(os.Stdout, appLogger.GetWriter()))

	// Initialize configurations
	appConfig := config.NewAppConfig(appLogger)

	// Initialize services
	userService := service.NewUserService(db.DB)
	postService := service.NewPostService(db.DB)
	commentService := service.NewCommentService(db.DB)

	// Setup router
	r := app.SetupRouter(userService, postService, commentService, appConfig, appLogger)

	// Get port from environment or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDatabase() (*config.Database, error) {
	dbConfig := config.NewDatabaseConfig()
	db, err := config.NewDatabase(dbConfig)
	if err != nil {
		return nil, err
	}

	if err := autoMigrate(db.DB); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	log.Println("Running auto-migration...")
	return db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}
