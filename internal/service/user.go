package service

import (
	"context"
	"errors"
	"github.com/phungvhbui/go-archi/internal/externalservice"
	"github.com/phungvhbui/go-archi/internal/mapper"
	"github.com/phungvhbui/go-archi/internal/model/dto"
	"github.com/phungvhbui/go-archi/internal/model/entity"
	"github.com/phungvhbui/go-archi/internal/repository"
)

type UserService interface {
	GetAll(context.Context) ([]dto.UserDTO, error)
	Create(context.Context, dto.UserDTO) (dto.UserDTO, error)
}

type userService struct {
	repository repository.UserRepository
	transactor repository.Transactor
	stripe     *externalservice.StripeService
}

func NewUserService(repository repository.UserRepository, transactor repository.Transactor, stripe *externalservice.StripeService) *userService {
	return &userService{
		repository: repository,
		transactor: transactor,
		stripe:     stripe,
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

func (s *userService) Create(ctx context.Context, request dto.UserDTO) (dto.UserDTO, error) {
	var (
		errCreateUser   = errors.New("errCreateUser")
		errCreateOrg    = errors.New("errCreateOrg")
		errCreateStripe = errors.New("errCreateStripe")
		errUpdateOrg    = errors.New("errUpdateOrg")
	)

	user, err := mapper.MapObject[dto.UserDTO, entity.User](request)

	err = s.transactor.Do(func(store repository.DataStore) error {
		// Create user
		err := store.UserRepository().Save(ctx, &user)
		if err != nil {
			return errCreateUser
		}

		organization := entity.Organization{
			Name:   user.AccountUUID.String(),
			IsUser: true,
			Users: []entity.User{
				user,
			},
		}

		// Create organization
		err = store.OrganizationRepository().Save(ctx, &organization)
		if err != nil {
			return errCreateOrg
		}

		// Create stripe id
		id, err := s.stripe.GenerateStripeId()
		if err != nil {
			return errCreateStripe
		}

		// Update organization
		err = store.OrganizationRepository().UpdateStripeId(ctx, &organization, id)
		if err != nil {
			return errUpdateOrg
		}

		return nil
	})

	if errors.As(err, &errUpdateOrg) {
		err = s.stripe.RollbackStripe()
	}

	if err != nil {
		return dto.UserDTO{}, err
	}

	userDTO, err := mapper.MapObject[entity.User, dto.UserDTO](user)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return userDTO, nil
}
