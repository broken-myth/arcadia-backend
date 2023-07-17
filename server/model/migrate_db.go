package model

import (
	"github.com/delta/arcadia-backend/database"
)

func MigrateDB() {
	db := database.GetDB()

	for _, model := range []interface{}{
		// Include models here to auto migrate
		Constant{},
		Region{},
		Character{},
		MiniconType{},
		Admin{},
		Minicon{},
		Perk{},
		Target{},
		Lootbox{},
		UserRegistration{},
		User{},
		OwnedMinicon{},
		OwnedPerk{},
		Lineup{},
		GeneratedLootbox{},
		MatchmakingDetails{},
		SimulationDetail{},
		BattleResult{},
	} {
		if err := db.AutoMigrate(&model); err != nil {
			panic(err)
		}
	}
}
