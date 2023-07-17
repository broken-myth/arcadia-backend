package model

import (
	"gorm.io/gorm"
)

type BattleResult struct {
	gorm.Model
	Result         int `gorm:"not null;"` // Attacker's PoV = 1,0,-1 (win, draw, loss respectively)
	WinnerTrophies int `gorm:"not null;"`
	LoserTrophies  int `gorm:"not null;"`

	// Relations
	MatchID            uint               `gorm:"not null;"`
	MatchmakingDetails MatchmakingDetails `gorm:"foreignKey:MatchID;"`
}
