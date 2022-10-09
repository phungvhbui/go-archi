package repository

import (
	"gorm.io/gorm"
)

type DataStore interface {
	UserRepository() UserRepository
	OrganizationRepository() OrganizationRepository
}

type dataStore struct {
	userRepository         UserRepository
	organizationRepository OrganizationRepository
}

func (s *dataStore) UserRepository() UserRepository {
	return s.userRepository
}

func (s *dataStore) OrganizationRepository() OrganizationRepository {
	return s.organizationRepository
}

type Transaction func(store DataStore) error

type Transactor interface {
	Do(Transaction) error
}

type transactor struct {
	db *gorm.DB
}

func New(db *gorm.DB) Transactor {
	return &transactor{db: db}
}

func (t *transactor) Do(fn Transaction) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		dataStore := &dataStore{
			userRepository:         NewUserRepository(tx),
			organizationRepository: NewOrganizationRepository(tx),
		}
		return fn(dataStore)
	})
}
