package tests

import (
	"fmt"

	"github.com/delta/arcadia-backend/database"
	controller "github.com/delta/arcadia-backend/server/controller/general"
	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/fatih/color"
)

func TestLeaderboard() {

	redisDB := database.GetRedisDB()
	redisDB.FlushAll()

	err := helpers.InsertNewUserRedis(1)
	if err != nil {
		fmt.Print(color.RedString("Error inserting user"))
	}
	_ = helpers.InsertNewUserRedis(2)
	_ = helpers.InsertNewUserRedis(3)
	_ = helpers.InsertNewUserRedis(4)

	//
	//

	ldb, err := controller.GetEntireLeaderboard()
	fmt.Println(ldb)
	// fmt.Print("\n")

	if err != nil {
		fmt.Print(color.RedString("Error getting leaderboard"))
	}

	err = helpers.UpdateUserTrophies(1, 1200)
	if err != nil {
		fmt.Print(color.RedString("Error updating trophies"))
	}

	_ = helpers.UpdateUserTrophies(3, 700)
	_ = helpers.UpdateUserTrophies(2, 1700)
	_ = helpers.UpdateUserTrophies(4, 1100)

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
	fmt.Print("\n")

	_ = helpers.UpdateRedis()

}
