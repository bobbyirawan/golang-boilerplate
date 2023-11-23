package utils

import (
	"github.com/google/uuid"
)

type UUID struct {
	id uuid.UUID
}

func NewIdGenertor() UUID {
	return UUID{uuid.New()}
}

func (u UUID) UserID() string {
	return u.id.String()
}

// GenerateUUID generates a new UUID using the google/uuid package.
func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

// func GenerateRequestID() string {

// }
