package externalservice

import (
	"errors"
	"github.com/rs/zerolog/log"
)

type StripeService struct{}

func NewStripe() *StripeService {
	return &StripeService{}
}

func (s *StripeService) GenerateStripeId() (string, error) {
	log.Print("create stripe not ok")
	return "success", errors.New("not ok")
}

func (s *StripeService) RollbackStripe() error {
	log.Print("rollback ok")
	return nil
}
