package models

import "github.com/google/uuid"

type Categories struct {
	*Model
	Name   string    `gorm:"not null" json:"name"`
	Kind   string    `gorm:"not null" json:"kind"`
	UserId uuid.UUID `gorm:"not null" json:"user_id"` // User
}
