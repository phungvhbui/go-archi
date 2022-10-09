package repository

import (
	"context"
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Repository[entity.Organization]
	UpdateStripeId(context.Context, *entity.Organization, string) error
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

func (r *organizationRepository) UpdateStripeId(ctx context.Context, organization *entity.Organization, stripeId string) error {
	organization.StripeId = stripeId
	if err := r.DB.WithContext(ctx).Select("StripeId").Updates(organization).Error; err != nil {
		return err
	}
	return nil
}
