//go:build wireinject
// +build wireinject

//go:generate wire
package app

import (
	"context"

	"go-hexagonal-architecture/internal/adapter/auth/jwt"
	"go-hexagonal-architecture/internal/adapter/config"
	"go-hexagonal-architecture/internal/adapter/handler/http"
	"go-hexagonal-architecture/internal/adapter/infra/mongo"
	repo "go-hexagonal-architecture/internal/adapter/infra/mongo/repo"
	"go-hexagonal-architecture/internal/core/service"

	"github.com/google/wire"
)

// InitializeApplication wires up all dependencies and returns an *Application.
func InitializeApplication(ctx context.Context) (*Application, error) {
	wire.Build(
		config.Load, // *config.Config
		mongo.New,
		repo.NewUserRepo, // port.UserRepository
		jwt.New,          // port.TokenService
		service.NewUser,
		service.NewAuth,
		http.NewUserHandler,
		http.NewAuthHandler,
		New, // app.New (constructor for *Application)
	)
	return nil, nil
}
