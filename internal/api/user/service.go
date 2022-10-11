package user

import (
	"context"
	"errors"

	"github.com/phungvhbui/go-archi/internal/datastore/model"
	repo "github.com/phungvhbui/go-archi/internal/datastore/repository"
	"github.com/phungvhbui/go-archi/internal/datastore/transaction"
	"github.com/phungvhbui/go-archi/internal/mapper"
	"github.com/phungvhbui/go-archi/internal/model/dto"
	"github.com/phungvhbui/go-archi/internal/stripe"
)

type UserService interface {
	GetAll(context.Context) ([]dto.UserDTO, error)
	Create(context.Context, dto.UserDTO) (dto.UserDTO, error)
}

type userService struct {
	repository repo.UserRepository
	transactor transaction.Transactor
	stripe     *stripe.StripeService
}

func NewUserService(repository repo.UserRepository, transactor transaction.Transactor, stripe *stripe.StripeService) *userService {
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

	dtos, err := mapper.MapList[model.User, dto.UserDTO](entities)
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

	user, err := mapper.MapObject[dto.UserDTO, model.User](request)

	err = s.transactor.Do(func(store transaction.DataStore) error {
		// Create user
		err := store.UserRepository().Save(ctx, &user)
		if err != nil {
			return errCreateUser
		}

		organization := model.Organization{
			Name:   user.AccountUUID.String(),
			IsUser: true,
			Users: []model.User{
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

	userDTO, err := mapper.MapObject[model.User, dto.UserDTO](user)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return userDTO, nil
}
