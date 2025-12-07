package app

import (
	"blog-backend/internal/config"
	"blog-backend/internal/handler"
	"blog-backend/internal/logger"
	"blog-backend/internal/middleware"
	"blog-backend/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userService *service.UserService,
	postService *service.PostService,
	commentService *service.CommentService,
	appConfig *config.AppConfig,
	appLogger *logger.Logger,
) *gin.Engine {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "your-secret-key"
	}
	jwtService := service.NewJWTService(secretKey)
	authMiddleware := middleware.NewAuthMiddleware(secretKey)

	userHandler := handler.NewUserHandler(userService, jwtService, appLogger)
	postHandler := handler.NewPostHandler(postService, appLogger)
	commentHandler := handler.NewCommentHandler(commentService, appLogger)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/logout", authMiddleware.JWTAuth(), userHandler.Logout)
		}

		users := api.Group("/users")
		users.Use(authMiddleware.JWTAuth())
		{
			users.GET("/me", userHandler.GetProfile)
		}

		posts := api.Group("/posts")
		{
			posts.GET("", postHandler.List)
			posts.GET("/:id", postHandler.GetByID)
			posts.POST("", authMiddleware.JWTAuth(), postHandler.Create)
			posts.PUT("/:id", authMiddleware.JWTAuth(), postHandler.Update)
			posts.DELETE("/:id", authMiddleware.JWTAuth(), postHandler.Delete)
		}

		comments := api.Group("/comments")
		{
			comments.GET("/post/:postId", commentHandler.ListByPostID)
			comments.GET("/:id", commentHandler.GetByID)
			comments.POST("", authMiddleware.JWTAuth(), commentHandler.Create)
			comments.PUT("/:id", authMiddleware.JWTAuth(), commentHandler.Update)
			comments.DELETE("/:id", authMiddleware.JWTAuth(), commentHandler.Delete)
		}
	}

	return r
}
