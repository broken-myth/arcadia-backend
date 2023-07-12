package main

import (
	"fmt"

	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/database"
	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/delta/arcadia-backend/server/model"
	tests "github.com/delta/arcadia-backend/tests/test_functions"
	"github.com/delta/arcadia-backend/utils"
)

func main() {

	config.InitConfig()
	utils.InitLogger()
	database.ConnectMySQLdb()
	database.ConnectRedisDB()
	model.MigrateDB()
	err := helpers.InitConstants()
	if err != nil {
		fmt.Printf("Error Initializing Constants = %v", err)
	}

	fmt.Print("Running Tests: \n")

	tests.TestTrophyGain()
	tests.TestLeaderboard()

}
