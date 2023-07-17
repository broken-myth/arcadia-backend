package model

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name        string `gorm:"required;unique"`
	Description string `gorm:"required;"`
	ImageURL    string `gorm:"required;"`
}
