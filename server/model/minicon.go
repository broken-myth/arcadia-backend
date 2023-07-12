package model

import (
	"gorm.io/gorm"
)

type Minicon struct {
	gorm.Model
	Name           string `gorm:"not null;"`
	BaseHealth     uint   `gorm:"not null;"`
	BaseAttack     uint   `gorm:"not null;"`
	Description    string `gorm:"not null;"`
	ImageLink      string `gorm:"not null;"`
	Perk1ID        uint   `gorm:"not null;"`
	Perk2ID        uint   `gorm:"not null;"`
	Perk3ID        uint   `gorm:"not null;"`
	BasePerk1Value string `gorm:"not null;"`
	BasePerk2Value string `gorm:"not null;"`
	BasePerk3Value string `gorm:"not null;"`
}
