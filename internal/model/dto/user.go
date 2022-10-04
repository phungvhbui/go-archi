package dto

import "github.com/google/uuid"

type User struct {
	ID          int64
	AccountUUID uuid.UUID
}
