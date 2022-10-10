package transaction

import (
	"gorm.io/gorm"

	repo "github.com/phungvhbui/go-archi/internal/datastore/repository"
)

type DataStore interface {
	UserRepository() repo.UserRepository
	OrganizationRepository() repo.OrganizationRepository
}

type dataStore struct {
	userRepository         repo.UserRepository
	organizationRepository repo.OrganizationRepository
}

func (s *dataStore) UserRepository() repo.UserRepository {
	return s.userRepository
}

func (s *dataStore) OrganizationRepository() repo.OrganizationRepository {
	return s.organizationRepository
}

type Transaction func(store DataStore) error

type Transactor interface {
	Do(Transaction) error
}

type transactor struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) Transactor {
	return &transactor{db: db}
}

func (t *transactor) Do(fn Transaction) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		dataStore := &dataStore{
			userRepository:         repo.NewUserRepository(tx),
			organizationRepository: repo.NewOrganizationRepository(tx),
		}
		return fn(dataStore)
	})
}
