package main

import (
	"fmt"

	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
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
	err := helper.InitConstants()
	if err != nil {
		fmt.Printf("Error Initializing Constants = %v", err)
	}

	fmt.Print("\n Running Tests: \n")

	tests.TestMatchPipeline()
	tests.TestTrophyGain()
	tests.TestLeaderboard()
	tests.TestUpdateXpLevels()
}
