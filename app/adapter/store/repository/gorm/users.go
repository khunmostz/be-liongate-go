package gorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	base *BaseGormRepository
}

func NewGormUsersRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		base: NewBaseGormRepository(db),
	}
}

func (r *GormUserRepository) CreateUser(ctx context.Context, user *domain.Users) (*domain.Users, error) {
	// Generate UUID for new user
	user.Id = uuid.New().String()

	if err := r.base.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *GormUserRepository) GetUserById(ctx context.Context, id string) (*domain.Users, error) {
	var user domain.Users
	if err := r.base.db.WithContext(ctx).Preload("Bookings").Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetUsersByRole(ctx context.Context, role string) ([]domain.Users, error) {
	var users []domain.Users
	if err := r.base.db.WithContext(ctx).Preload("Bookings").Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.Users, error) {
	var user domain.Users
	if err := r.base.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) UpdateUser(ctx context.Context, id string, user *domain.Users) (*domain.Users, error) {
	existingUser, err := r.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.base.db.WithContext(ctx).Model(existingUser).Updates(user).Error; err != nil {
		return nil, err
	}

	return r.GetUserById(ctx, id)
}

func (r *GormUserRepository) DeleteUser(ctx context.Context, id string) error {
	return r.base.db.WithContext(ctx).Where("user_id = ?", id).Delete(&domain.Users{}).Error
}
