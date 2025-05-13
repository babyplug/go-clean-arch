package service

import (
	"errors"
	"sync"

	"github.com/babyplug/go-clean-arch/internal/core/domain"
	"github.com/babyplug/go-clean-arch/internal/core/port"

	"golang.org/x/crypto/bcrypt"
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

func (s *userServiceImpl) Create(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userServiceImpl) GetByID(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userServiceImpl) GetByEmail(email string) (*domain.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userServiceImpl) List() ([]*domain.User, error) {
	return s.repo.List()
}

func (s *userServiceImpl) Update(user *domain.User) error {
	return s.repo.Update(user)
}

func (s *userServiceImpl) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *userServiceImpl) Count() (int, error) {
	return s.repo.Count()
}

func (s *userServiceImpl) RegisterUser(input *domain.User) (*domain.User, error) {
	existing, _ := s.repo.GetByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	input.Password = string(hashedPassword)
	if err := s.repo.Create(input); err != nil {
		return nil, err
	}
	return input, nil
}

func (s *userServiceImpl) AuthenticateUser(email, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
