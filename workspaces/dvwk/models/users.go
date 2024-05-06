package models

import (
	"time"
)

type User struct {
	*Model
	Name     string    `gorm:"type:varchar(101);not null" strfmt:"limit:101" json:"name"`
	Email    string    `gorm:"type:varchar(101);unique;not null" strfmt:"limit:101" json:"email"`
	TagName  string    `gorm:"type:varchar(101);unique;not null" strfmt:"limit:101" json:"tag_name"`
	Phone    string    `gorm:"type:varchar(21);not null" strfmt:"limit:21" json:"phone"`
	Gender   string    `gorm:"type:varchar(101);not null" strfmt:"limit:101" json:"gender"`
	Birthday time.Time `gorm:"type:datetime;not null" json:"birthday"`
	Address  string    `gorm:"type:varchar(201);not null" strfmt:"limit:201" json:"address"`
	City     string    `gorm:"type:varchar(51);not null" strfmt:"limit:51" json:"city"`
	Province string    `gorm:"type:varchar(51);not null" strfmt:"limit:51" json:"province"`
	Country  string    `gorm:"type:varchar(51);not null" strfmt:"limit:51" json:"country"`
	ZipCode  string    `gorm:"type:varchar(21);not null" strfmt:"limit:21" json:"zip_code"`
}
