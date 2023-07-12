package model

import (
	"gorm.io/gorm"
)

type BattleResult struct {
	gorm.Model
	Result     int  `gorm:"not null;"` // Attacker's PoV = 1,0,-1 (win, draw, loss respectively)
	TrophyGain uint `gorm:"not null;"` // Must be +ve
	TrophyLost uint `gorm:"not null;"` // Must be +ve

	// Relations
	MatchID            uint               `gorm:"not null;"`
	MatchmakingDetails MatchmakingDetails `gorm:"foreignKey:MatchID;"`
}
