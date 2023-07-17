package tests

import (
	"fmt"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/fatih/color"
)

func TestUpdateXpLevels() {
	db := database.GetDB()
	tx := db.Begin()

	var minicon model.OwnedMinicon
	var ownedPerk model.OwnedPerk
	tx.First(&minicon)
	tx.Where("owned_minicon_id = ?", minicon.ID).First(&ownedPerk)
	ownedPerk.PerkValue = 160
	minicon.Level = 1
	minicon.XP = 1000
	tx.Save(&ownedPerk)
	tx.Save(&minicon)

	for i := 0; i < 30; i++ {
		err := helper.UpdateXpUserLineup(tx, minicon.OwnerID)
		if err != nil {
			fmt.Print(color.RedString("Error while updating xp of minicon. Error = %v", err))
			return
		}
	}

	tx.First(&minicon)
	tx.Where("owned_minicon_id = ?", minicon.ID).First(&ownedPerk)

	if minicon.Level != 2 || minicon.XP != 2140 || ownedPerk.PerkValue != 176 {
		fmt.Print(color.RedString(fmt.Sprint("Got level = ", minicon.Level, " instead of 2 or xp = ",
			minicon.XP, " instead of 2140, or perk value = ", ownedPerk.PerkValue,
			" instead of 176 in UpdateLevelMinicon & UpdateXpMinicon Testing \n\n",
		)))
	}

	tx.Rollback()
	fmt.Print(color.GreenString("TestUpdateXpLevels passed\n"))
}
