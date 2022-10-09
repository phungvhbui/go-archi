package repository

import (
	"context"
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[entity.User]
	GetUserByEmail(context.Context, string) (entity.User, error)
}

type userRepository struct {
	GormRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		GormRepository: GormRepository[entity.User]{
			DB: db,
		},
	}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	if err := r.DB.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return *new(entity.User), err
	}

	return user, nil
}
