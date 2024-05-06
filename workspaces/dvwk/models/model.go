package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Id        uuid.UUID           `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time           `gorm:"type:timestamp;not null" json:"created_at"`
	UpdateAt  time.Time           `gorm:"type:timestamp;not null" json:"update_at"`
	DeletedAt sql.Null[time.Time] `gorm:"type:timestamp" json:"deleted_at"`
}
