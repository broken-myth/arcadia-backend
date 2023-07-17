package router

import (
	controller "github.com/delta/arcadia-backend/server/controller/admin"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"
)

// Router for the Admin entity
func adminRouter(superRoute *gin.RouterGroup) {
	adminRoutes := superRoute.Group("/admin")

	adminRoutes.POST("/login", controller.AdminLoginPOST)

	adminRoutes.Use(middleware.Auth)
	{
		adminRoutes.GET("/verify", controller.AdminVerifyGET)
		adminRoutes.GET("/constants", controller.GetConstantsGET)
		adminRoutes.PATCH("/constants", controller.UpdateConstantsPATCH)
		adminRoutes.PATCH("/update_leaderboard", controller.UpdateLeaderboardPATCH)
	}

}
