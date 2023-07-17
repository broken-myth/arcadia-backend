package model

import (
	"gorm.io/gorm"
)

type OwnedPerk struct {
	gorm.Model
	PerkValue uint

	// Relations
	OwnedMiniconID uint         `gorm:"not null"`
	OwnedMinicon   OwnedMinicon `gorm:"foreignKey:OwnedMiniconID;"`
	PerkID         uint         `gorm:"not null;"`
	Perk           Perk         `gorm:"foreignKey:PerkID;"`
}
