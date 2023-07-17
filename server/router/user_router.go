package router

import (
	controller "github.com/delta/arcadia-backend/server/controller/user"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"
)

// Router for the User entity
func userRouter(superRoute *gin.RouterGroup) {
	userRoutes := superRoute.Group("/user")

	userRoutes.POST("/auth", controller.AuthUserPOST)

	userRoutes.Use(middleware.Auth)
	{
		userRoutes.POST("/signup/complete", controller.SignupUserPOST)

		userRoutes.GET("/profile", controller.GetProfileGET)

		userRoutes.PATCH("/profile", controller.UpdateUserProfilePATCH)
	}
}
