package service

import (
	"context"
	"errors"

	"github.com/babyplug/go-clean-arch/internal/core/port"
	"github.com/babyplug/go-clean-arch/internal/core/util"
)

type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

func NewAuth(userRepo port.UserRepository, ts port.TokenService) port.AuthService {
	return &AuthService{
		repo: userRepo,
		ts:   ts,
	}
}

func (a *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = util.ComparePassword(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
		// return "", domain.ErrInvalidCredentials
	}
	

	token, err := a.ts.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
