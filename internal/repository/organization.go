package repository

import (
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	gormRepository[entity.Organization]
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{
		gormRepository: gormRepository[entity.Organization]{
			db: db,
		},
	}
}
