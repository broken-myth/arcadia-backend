package main

import (
	"fmt"

	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server"
	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/fatih/color"
)

func main() {
	config.InitConfig()

	utils.InitLogger()

	database.ConnectMySQLdb()

	database.ConnectRedisDB()

	model.MigrateDB()

	err := helpers.InitConstants()
	if err != nil {
		fmt.Print(color.RedString("Error Initializing Constants = %v", err))
	}

	err = helpers.UpdateRedis()
	if err != nil {
		fmt.Print(color.RedString("Error Updating Entire leaderboard = %v", err))
	}

	server.Run()

}
