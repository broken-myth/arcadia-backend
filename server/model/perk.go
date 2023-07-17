package model

import (
	"gorm.io/gorm"
)

type Perk struct {
	gorm.Model
	PerkTrigger string `gorm:"not null;"`
	PerkName    string `gorm:"not null;"`
	Effect      string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	BaseValue   uint   `gorm:"not null;"`

	// Relation
	MiniconID uint    `gorm:"not null;"`
	Minicon   Minicon `gorm:"foreignKey:MiniconID;"`
}
