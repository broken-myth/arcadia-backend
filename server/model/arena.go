package model

import (
	"gorm.io/gorm"
)

type Arena struct {
	gorm.Model
	Name        string `gorm:"not null;"`
	Description string `gorm:"not null;"`
}
