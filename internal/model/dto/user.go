package dto

import "github.com/google/uuid"

type UserDTO struct {
	ID          int64
	AccountUUID uuid.UUID
}
