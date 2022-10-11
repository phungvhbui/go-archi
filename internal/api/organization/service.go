package organization

import (
	"context"

	"github.com/phungvhbui/go-archi/internal/datastore/model"
	repo "github.com/phungvhbui/go-archi/internal/datastore/repository"
	"github.com/phungvhbui/go-archi/internal/mapper"
	"github.com/phungvhbui/go-archi/internal/model/dto"
)

type OrganizationService interface {
	GetAll(context.Context) ([]dto.OrganizationDTO, error)
}

type organizationService struct {
	repository repo.OrganizationRepository
}

func NewOrganizationService(repository repo.OrganizationRepository) *organizationService {
	return &organizationService{
		repository: repository,
	}
}

func (s *organizationService) GetAll(ctx context.Context) ([]dto.OrganizationDTO, error) {
	entities, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos, err := mapper.MapList[model.Organization, dto.OrganizationDTO](entities)
	if err != nil {
		return nil, err
	}

	return dtos, err
}
