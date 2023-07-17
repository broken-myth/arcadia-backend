package model

import (
	"gorm.io/gorm"
)

// User Registration Database
type UserRegistration struct {
	gorm.Model
	Email      string `gorm:"required;unique"`
	Name       string `gorm:"required;"`
	FormFilled bool   `gorm:"default:false;"`
	Username   string `gorm:"default:null;unique"`
	College    string `gorm:"default:null;"`
	Contact    string `gorm:"default:null;"`
}
