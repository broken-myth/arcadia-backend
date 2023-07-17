package model

import (
	//	"gorm.io/datatypes"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SimulationDetail struct {
	gorm.Model
	SimulationLog     datatypes.JSON `gorm:"not null;"`
	AttackerSurvivors uint           `gorm:"not null;"` // For Testing Purposes
	DefenderSurvivors uint           `gorm:"not null;"` // For Testing Purposes
	// Relations
	MatchID            uint               `gorm:"not null;"`
	MatchmakingDetails MatchmakingDetails `gorm:"foreignKey:MatchID;"`
}
