package router

import (
	controller "github.com/delta/arcadia-backend/server/controller/general"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"
)

func generalRouter(superRoute *gin.RouterGroup) {

	// Fetch all characters
	superRoute.GET("/characters", controller.GetCharactersGET)

	// Fetch Leaderboard
	superRoute.GET("/leaderboard/:page", controller.FetchLeaderboardGET)

	superRoute.Use(middleware.Auth)
	{
		superRoute.GET("/lootbox", controller.LootboxGET)

		superRoute.POST("/lootbox/open", controller.LootboxOpenPOST)

		superRoute.POST("/error", controller.LogClientSideErrorPOST)

	}
}
