package tests

import (
	"fmt"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/fatih/color"
)

func TestTrophyGain() {

	attTrophy, defTrophy := helper.CalculateTrophyGain(100, 107, 2, 0)
	if attTrophy != 28 || defTrophy != -25 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 28, -25 for (100, 107, 2, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helper.CalculateTrophyGain(107, 100, 2, 0)
	if attTrophy != 24 || defTrophy != -21 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 24, -21 for (107, 100, 2, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helper.CalculateTrophyGain(100, 107, 0, 0)
	if attTrophy != 0 || defTrophy != 0 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 0, 0 for (100, 107, 0, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helper.CalculateTrophyGain(90, 107, 1, 0)
	if attTrophy != 29 || defTrophy != -26 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of 29, -26 for (90, 107, 1, 0) \n", attTrophy, defTrophy))
	}

	attTrophy, defTrophy = helper.CalculateTrophyGain(97, 107, 0, 1)
	if attTrophy != -18 || defTrophy != 21 {
		fmt.Print(color.RedString("Got Att = %d, Def = %d instead of -18, 21 for (97, 107, 0, 1) \n", attTrophy, defTrophy))
	}

	fmt.Print(color.GreenString("TestTrophyGain Completed \n \n"))

}
