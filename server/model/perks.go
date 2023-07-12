package model

import (
	"gorm.io/gorm"
)

type Perks struct {
	gorm.Model
	Trigger     string `gorm:"not null;"`
	Effect      string `gorm:"not null;"`
	Description string `gorm:"not null;"`
}
