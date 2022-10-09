package repository

import (
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Repository[entity.Organization]
}

type organizationRepository struct {
	GormRepository[entity.Organization]
}

func NewOrganizationRepository(db *gorm.DB) *organizationRepository {
	return &organizationRepository{
		GormRepository: GormRepository[entity.Organization]{
			DB: db,
		},
	}
}
