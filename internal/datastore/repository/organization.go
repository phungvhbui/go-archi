package repository

import (
	"context"

	"github.com/phungvhbui/go-archi/internal/datastore/model"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Repository[model.Organization]
	UpdateStripeId(context.Context, *model.Organization, string) error
}

type organizationRepository struct {
	GormRepository[model.Organization]
}

func NewOrganizationRepository(db *gorm.DB) *organizationRepository {
	return &organizationRepository{
		GormRepository: GormRepository[model.Organization]{
			DB: db,
		},
	}
}

func (r *organizationRepository) UpdateStripeId(ctx context.Context, organization *model.Organization, stripeId string) error {
	organization.StripeId = stripeId
	if err := r.DB.WithContext(ctx).Select("StripeId").Updates(organization).Error; err != nil {
		return err
	}
	return nil
}
