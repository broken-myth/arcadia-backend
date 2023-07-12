package model

import (
	"gorm.io/gorm"
)

type Lineup struct {
	gorm.Model
	PositionNo uint `gorm:"not null"`

	// Relations
	CreatorID      uint         `gorm:"not null"`
	Creator        User         `gorm:"foreignKey:CreatorID"`
	OwnedMiniconID uint         `gorm:"not null"`
	OwnedMinicon   OwnedMinicon `gorm:"foreignKey:OwnedMiniconID;"`
}
