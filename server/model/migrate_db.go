package model

import (
	"github.com/delta/arcadia-backend/database"
)

func MigrateDB() {
	db := database.GetDB()

	for _, model := range []interface{}{
		// Include models here to auto migrate
		UserRegistration{},
		User{},
		Perks{},
		Minicon{},
		Arena{},
		OwnedMinicon{},
		Lineup{},
		Constants{},
		MatchmakingDetails{},
		BattleResult{},
		Lootbox{},
	} {
		if err := db.AutoMigrate(&model); err != nil {
			panic(err)
		}
	}
}
