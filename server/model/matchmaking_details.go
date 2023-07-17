package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MatchmakingDetails struct {
	gorm.Model
	AttackerLineUp datatypes.JSON `gorm:"not null;"`
	DefenderLineUp datatypes.JSON `gorm:"not null;"`

	// Relations
	AttackerID uint `gorm:"not null"`
	Attacker   User `gorm:"foreignKey:AttackerID;"`
	DefenderID uint `gorm:"not null"`
	Defender   User `gorm:"foreignKey:DefenderID;"`
}
