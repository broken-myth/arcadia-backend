package tests

import (
	"fmt"

	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/fatih/color"
)

func TestTrophyGain() {

	attTrophy, defTrophy := helpers.CalculateTrophyGain(100, 107, 2, 0)
	if attTrophy != 28 || defTrophy != 23 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 28, 23 for (100, 107, 2, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helpers.CalculateTrophyGain(107, 100, 2, 0)
	if attTrophy != 24 || defTrophy != 19 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 24, 19 for (107, 100, 2, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helpers.CalculateTrophyGain(100, 107, 0, 0)
	if attTrophy != 0 || defTrophy != 0 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 0, 0 for (100, 107, 0, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helpers.CalculateTrophyGain(90, 107, 1, 0)
	if attTrophy != 29 || defTrophy != 24 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 29, 24 for (90, 107, 1, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helpers.CalculateTrophyGain(97, 107, 0, 1)
	if attTrophy != 16 || defTrophy != 21 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 16, 21 for (97, 107, 0, 1) \n", attTrophy, defTrophy))
	}

	fmt.Print(color.GreenString("TestTrophyGain Completed \n \n"))

}
