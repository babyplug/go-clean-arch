package http

import (
	"strings"

	"go-hexagonal-architecture/internal/adapter/config"
	"go-hexagonal-architecture/internal/adapter/handler/http/middleware"
	"go-hexagonal-architecture/internal/core/port"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "go-hexagonal-architecture/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.Config, // config.Config is a struct that holds configuration values
	token port.TokenService,
	userHandler *UserHandler, // UserHandler is a struct that handles user-related HTTP requests
	authHandler *AuthHandler, // AuthHandler is a struct that handles authentication-related HTTP requests
) (*Router, error) {
	gin.SetMode(gin.DebugMode)

	// Disable debug mode in production
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = strings.Split(config.AllowedOrigins, ",")
	ginConfig.AllowMethods = strings.Split(config.AllowedMethods, ",")
	ginConfig.AllowHeaders = strings.Split(config.AllowedHeaders, ",")

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), cors.New(ginConfig), middleware.LoggingMiddleware())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := r.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("", userHandler.Register)

			authUser := user.Group("")
			// .Use(middleware.AuthMiddleware(token))
			{
				authUser.GET("", userHandler.List)
				authUser.GET("/:id", userHandler.GetByID)
				authUser.PUT("/:id", userHandler.Update)
				authUser.DELETE("/:id", userHandler.Delete)
			}
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{r}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
