package models

import (
	"time"
)

type User struct {
	*Model
	Name     string    `gorm:"not null" json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	TagName  string    `gorm:"unique;not null" json:"tag_name"`
	Phone    string    `gorm:"not null" json:"phone"`
	Gender   string    `gorm:"not null" json:"gender"`
	Birthday time.Time `gorm:"not null" json:"birthday"`
	Address  string    `gorm:"not null" json:"address"`
	City     string    `gorm:"not null" json:"city"`
	Province string    `gorm:"not null" json:"province"`
	Country  string    `gorm:"not null" json:"country"`
	ZipCode  string    `gorm:"not null" json:"zip_code"`
}
