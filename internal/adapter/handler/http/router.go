package http

import (
	"strings"

	"github.com/babyplug/go-clean-arch/internal/adapter/config"
	"github.com/babyplug/go-clean-arch/internal/adapter/middleware"
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
	userHandler *UserHandler, // UserHandler is a struct that handles user-related HTTP requests
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
		user := v1.Group("/user")
		{
			// user.POST("/register", userHandler)
			user.GET("", userHandler.List)
		}
	}

	return &Router{r}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

func (r *Router) Shutdown() {
	r.Shutdown()
}
