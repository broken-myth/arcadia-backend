package router

import (
	controller "github.com/delta/arcadia-backend/server/controller/match"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"
)

// Router for the Match entity
func matchRouter(superRoute *gin.RouterGroup) {
	matchRoutes := superRoute.Group("/match")

	matchRoutes.Use(middleware.Auth)
	{
		matchRoutes.GET("/history", controller.GetMatchHistoryGET)
		matchRoutes.GET("/view/:id", controller.ViewSimulationLogGET)
		matchRoutes.Use(middleware.Arena)
		{
			matchRoutes.GET("/start", controller.StartMatchGET)
		}
	}
}
