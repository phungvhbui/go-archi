package service

import (
	"context"
	"github.com/phungvhbui/go-archi/internal/mapper"
	"github.com/phungvhbui/go-archi/internal/repository"

	"github.com/phungvhbui/go-archi/internal/model/dto"
	"github.com/phungvhbui/go-archi/internal/model/entity"
)

type UserService interface {
	GetAll(context.Context) ([]dto.UserDTO, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) GetAll(ctx context.Context) ([]dto.UserDTO, error) {
	entities, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos, err := mapper.MapList[entity.User, dto.UserDTO](entities)
	if err != nil {
		return nil, err
	}

	return dtos, err
}
