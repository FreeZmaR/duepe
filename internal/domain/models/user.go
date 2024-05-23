package models

import (
	"duepe/internal/domain/polices"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Role      polices.UserRole
	CreatedAt time.Time
	UpdatedAt *time.Time
}
