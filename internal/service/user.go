package service

import "github.com/phungvhbui/go-archi/internal/repository"

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Get() string {
	return s.repository.Get()
}
