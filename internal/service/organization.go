package service

import (
	"context"
	"github.com/phungvhbui/go-archi/internal/mapper"
	"github.com/phungvhbui/go-archi/internal/model/dto"
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"github.com/phungvhbui/go-archi/internal/repository"
)

type OrganizationService interface {
	GetAll(context.Context) ([]dto.OrganizationDTO, error)
}

type organizationService struct {
	repository repository.OrganizationRepository
}

func NewOrganizationService(repository repository.OrganizationRepository) *organizationService {
	return &organizationService{
		repository: repository,
	}
}

func (s *organizationService) GetAll(ctx context.Context) ([]dto.OrganizationDTO, error) {
	entities, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos, err := mapper.MapList[entity.Organization, dto.OrganizationDTO](entities)
	if err != nil {
		return nil, err
	}

	return dtos, err
}
