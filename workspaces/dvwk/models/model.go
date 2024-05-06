package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Id        uuid.UUID           `gorm:"type:varchar(37);primary_key" strfmt:"limit:37" json:"id"`
	CreatedAt time.Time           `gorm:"type:datetime;not null" json:"created_at"`
	UpdateAt  time.Time           `gorm:"type:datetime;not null" json:"update_at"`
	DeletedAt sql.Null[time.Time] `gorm:"type:datetime" json:"deleted_at"`
}
