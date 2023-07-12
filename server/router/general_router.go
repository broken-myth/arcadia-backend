package router

import (
	controller "github.com/delta/arcadia-backend/server/controller/general"
	"github.com/gin-gonic/gin"
)

func generalRouter(superRoute *gin.RouterGroup) {

	// Fetch all characters
	superRoute.GET("/characters", controller.GetCharactersGET)

	// Fetch Leaderboard
	superRoute.GET("/leaderboard", controller.FetchLeaderboardGET)

}
