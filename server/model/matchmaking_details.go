package model

import (
	"gorm.io/gorm"
)

type MatchmakingDetails struct {
	gorm.Model
	AttackerLineUp string `gorm:"type:text;not null"`
	DefenderLineUp string `gorm:"type:text;not null"`

	// Relations
	AttackerID uint `gorm:"not null"`
	Attacker   User `gorm:"foreignKey:AttackerID;"`
	DefenderID uint `gorm:"not null"`
	Defender   User `gorm:"foreignKey:DefenderID;"`
}
