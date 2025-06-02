package services

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type UserService struct {
	usersRepository port.UsersRepository
}

func NewUsersService(usersRepository port.UsersRepository) *UserService {
	return &UserService{
		usersRepository: usersRepository,
	}
}

func (s *UserService) Register(ctx context.Context, user *domain.Users) (*domain.Users, error) {
	return s.usersRepository.CreateUser(ctx, user)
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*domain.Users, error) {
	return s.usersRepository.GetUserById(ctx, id)
}

func (s *UserService) GetUsersByRole(ctx context.Context, role string) ([]domain.Users, error) {
	return s.usersRepository.GetUsersByRole(ctx, role)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, user *domain.Users) (*domain.Users, error) {
	return s.usersRepository.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.usersRepository.DeleteUser(ctx, id)
}
