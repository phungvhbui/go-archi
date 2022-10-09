package repository

import (
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[entity.User]
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
