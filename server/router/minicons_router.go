package router

import (
	controller "github.com/delta/arcadia-backend/server/controller/minicons"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"
)

func miniconsRouter(superRoute *gin.RouterGroup) {
	miniconRoutes := superRoute.Group("/minicon")

	miniconRoutes.Use(middleware.Auth)
	{
		miniconRoutes.GET("/", controller.FetchMiniconsGET)
		miniconRoutes.GET("/:id", controller.GetMiniconDetailsGET)
		miniconRoutes.PATCH("/updateLineup", controller.UpdateLineupPATCH)
	}
}
