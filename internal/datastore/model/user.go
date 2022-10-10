package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountUUID   uuid.UUID      `gorm:"<-:create;unique;not null"`
	Organizations []Organization `gorm:"many2many:user_organizations;"`
}
