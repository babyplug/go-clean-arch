package service

import (
	"context"
	"errors"
	"sync"

	"github.com/babyplug/go-clean-arch/internal/core/domain"
	"github.com/babyplug/go-clean-arch/internal/core/port"
	"github.com/babyplug/go-clean-arch/internal/core/util"
)

var (
	userServiceOnce sync.Once
	userService     *userServiceImpl
)

type userServiceImpl struct {
	repo port.UserRepository
}

func NewUser(repo port.UserRepository) port.UserService {
	userServiceOnce.Do(func() {
		userService = &userServiceImpl{repo: repo}
	})

	return userService
}

func (s *userServiceImpl) Create(ctx context.Context, user *domain.User) error {
	existing, _ := s.repo.GetByEmail(ctx, user.Email)
	if existing != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userServiceImpl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *userServiceImpl) List(ctx context.Context) ([]*domain.User, error) {
	return s.repo.List(ctx)
}

func (s *userServiceImpl) Update(ctx context.Context, user *domain.User) error {
	existingUser, err := s.repo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}

	if existingUser.Email != user.Email {
		existing, _ := s.repo.GetByEmail(ctx, user.Email)
		if existing != nil {
			return errors.New("email already exists")
		}
	}

	return s.repo.Update(ctx, user)
}

func (s *userServiceImpl) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *userServiceImpl) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}
