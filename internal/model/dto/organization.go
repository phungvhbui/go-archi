package dto

type OrganizationDTO struct {
	ID       int64
	IsUser   bool
	ClientId string
	StripeId string
}
