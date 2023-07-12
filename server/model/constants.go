package model

import (
	"gorm.io/gorm"
)

type Constants struct {
	gorm.Model
	Key   string `gorm:"not null;"`
	Value int    `gorm:"not null;"`
}
