package main

import (
	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/simulator"
	"github.com/delta/arcadia-backend/utils"
)

//	@title			Arcadia API
//	@version		1.0
//	@description	This is the API documentation for Arcadia Backend

//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Authorization token

func main() {
	config.InitConfig()

	utils.InitLogger()

	database.ConnectMySQLdb()

	database.ConnectRedisDB()

	model.MigrateDB()

	err := helper.InitConstants()
	if err != nil {
		utils.Logger.Errorln("Error while initializing constants. Error: ", err)
	} else {
		utils.Logger.Infoln("Constants initialized successfully")
	}

	err = helper.UpdateRedis()
	if err != nil {
		utils.Logger.Errorln("Error while updating redis. Error: ", err)
	} else {
		utils.Logger.Infoln("Redis updated successfully")
	}

	err = simulator.Init()
	if err != nil {
		utils.Logger.Errorln("Error in Initialising Simulator. Error: ", err)
	}

	server.Run()
}
