package port

import "github.com/babyplug/go-clean-arch/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
	Count() (int, error)
}

type UserService interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	List() ([]*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
	Count() (int, error)
}
