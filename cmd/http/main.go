package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/babyplug/go-clean-arch/internal/adapter/background"
	"github.com/babyplug/go-clean-arch/internal/adapter/config"
	httpAdapter "github.com/babyplug/go-clean-arch/internal/adapter/handler/http"
	"github.com/babyplug/go-clean-arch/internal/adapter/infra/mongo"
	"github.com/babyplug/go-clean-arch/internal/adapter/infra/mongo/repo"
	"github.com/babyplug/go-clean-arch/internal/core/service"
)

func main() {
	cfg := config.Load()
	log.Printf("Config: %+v\n", cfg)

	ctx := context.Background()

	// Init database
	client, err := mongo.New(ctx, cfg.MongoURI)
	defer client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userRepo := repo.NewUserRepo(client)
	userService := service.NewUser(userRepo)
	userHandler := httpAdapter.NewUserHandler(userService)

	// Init router
	router, err := httpAdapter.NewRouter(
		cfg,
		userHandler,
	)

	// Background user logger
	stopCh := make(chan struct{})
	background.StartUserCountLogger(userRepo, stopCh)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router.Handler(),
	}

	// Start server in goroutine
	go func() {
		log.Println("Server started on " + cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	close(stopCh)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown: ", err)
	}

	log.Println("Servers gracefully stopped")

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 10 seconds.")
	log.Println("Server exiting")
}
