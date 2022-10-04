package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Organizations []Organization `gorm:"many2many:user_organizations;"`
	AccountUUID   uuid.UUID      `gorm:"<-:create;unique;not null"`
}
