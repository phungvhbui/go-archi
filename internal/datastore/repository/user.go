package repository

import (
	"github.com/phungvhbui/go-archi/internal/datastore/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[model.User]
}

type userRepository struct {
	GormRepository[model.User]
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		GormRepository: GormRepository[model.User]{
			DB: db,
		},
	}
}
