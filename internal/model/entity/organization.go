package entity

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	IsUser   bool   `gorm:"<-:create;index;not null"`
	ClientId string `gorm:"not null;unique"`
	Users    []User `gorm:"many2many:user_organizations;"`
	StripeId string `gorm:"not null;unique"`
}

type UserOrganization struct {
	UserId         uint `gorm:"primaryKey"`
	OrganizationId uint `gorm:"primaryKey"`
}
