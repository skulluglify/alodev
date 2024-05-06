package models

import "github.com/google/uuid"

type Categories struct {
	*Model
	Name   string    `gorm:"type:varchar(51);not null" strfmt:"limit:51" json:"name"`
	Kind   string    `gorm:"type:varchar(51);not null" strfmt:"limit:51" json:"kind"`
	UserId uuid.UUID `gorm:"type:varchar(37);not null" strfmt:"limit:37" json:"user_id"` // User
}
