package entity

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name     string `gorm:"not null"`
	IsUser   bool   `gorm:"<-:create;index;not null"`
	StripeId string `gorm:"not null"`
	Users    []User `gorm:"many2many:user_organizations;"`
}

type UserOrganization struct {
	UserId         uint `gorm:"primaryKey"`
	OrganizationId uint `gorm:"primaryKey"`
}
