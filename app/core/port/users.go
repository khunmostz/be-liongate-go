package port

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

type UsersRepository interface {
	CreateUser(context context.Context, user *domain.Users) (*domain.Users, error)
	GetUserById(context context.Context, id string) (*domain.Users, error)
	GetUsersByRole(context context.Context, role string) ([]domain.Users, error)
	GetUserByUsername(context context.Context, username string) (*domain.Users, error)
	UpdateUser(context context.Context, id string, user *domain.Users) (*domain.Users, error)
	DeleteUser(context context.Context, id string) error
}

type UsersService interface {
	Register(context context.Context, user *domain.Users) (*domain.Users, error)
	GetUserById(context context.Context, id string) (*domain.Users, error)
	GetUsersByRole(context context.Context, role string) ([]domain.Users, error)
	UpdateUser(context context.Context, id string, user *domain.Users) (*domain.Users, error)
	DeleteUser(context context.Context, id string) error
}
