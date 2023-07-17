package tests

import (
	"fmt"

	"github.com/delta/arcadia-backend/database"
	controller "github.com/delta/arcadia-backend/server/controller/general"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/fatih/color"
)

func TestLeaderboard() {

	redisDB := database.GetRedisDB()
	redisDB.FlushAll()

	err := helper.InsertNewUserRedis(1)
	if err != nil {
		fmt.Print(color.RedString("Error inserting user"))
	}
	_ = helper.InsertNewUserRedis(2)
	_ = helper.InsertNewUserRedis(3)
	_ = helper.InsertNewUserRedis(4)

	//
	//

	ldb, err := controller.GetEntireLeaderboard()
	fmt.Println(ldb)
	// fmt.Print("\n")

	if err != nil {
		fmt.Print(color.RedString("Error getting leaderboard"))
	}

	err = helper.UpdateUserTrophies(1, 1200)
	if err != nil {
		fmt.Print(color.RedString("Error updating trophies"))
	}

	_ = helper.UpdateUserTrophies(3, 700)
	_ = helper.UpdateUserTrophies(2, 1700)
	_ = helper.UpdateUserTrophies(4, 1100)

	ldb, err = controller.GetEntireLeaderboard()
	fmt.Print(ldb)
	fmt.Print("\n")

	if err != nil {
		fmt.Print(color.RedString("Error getting leaderboard"))
	}

	ldb, err = controller.GetLeaderboardRange(1, 0)
	if err != nil {
		fmt.Print(color.RedString("Error getting leaderboard"))
	}
	fmt.Print(ldb)
	fmt.Print("\n\n\n\n")

	_ = helper.UpdateRedis()

}
