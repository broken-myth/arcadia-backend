package model

import (
	"gorm.io/gorm"
)

// Admin Database
type Admin struct {
	gorm.Model
	Username string `gorm:"required;unique"`
	Password string `gorm:"required;"`
}
