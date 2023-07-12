package model

import (
	"gorm.io/gorm"
)

type OwnedMinicon struct {
	gorm.Model
	Health     uint   `gorm:"not null;"`
	Attack     uint   `gorm:"not null;"`
	Xp         uint   `gorm:"not null;"`
	Perk1Value string `gorm:"not null;"`
	Perk2Value string `gorm:"not null;"`
	Perk3Value string `gorm:"not null;"`

	// Relations
	OwnerID   uint    `gorm:"not null;"`
	Owner     User    `gorm:"foreignKey:OwnerID;"`
	MiniconID uint    `gorm:"not null"`
	Minicon   Minicon `gorm:"foreignKey:MiniconID;"`
}
