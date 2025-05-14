package http

import (
	"strings"

	"github.com/babyplug/go-clean-arch/internal/adapter/config"
	"github.com/babyplug/go-clean-arch/internal/adapter/handler/http/middleware"
	"github.com/babyplug/go-clean-arch/internal/core/port"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	// Disable debug mode in production
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	originsList := strings.Split(config.AllowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

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
			user.POST("/register", userHandler.Register)

			authUser := user.Group("").Use(middleware.AuthMiddleware(token))
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

	return &Router{r}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
