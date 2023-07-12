package router

import (
	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"

	"time"

	cors "github.com/itsjamie/gin-cors"
)

var Router *gin.Engine

func Init() {

	config := config.GetConfig()

	if config.AppEnv == "DOCKER" {
		gin.SetMode(gin.ReleaseMode)
		Router = gin.New()

		// Logger
		Router.Use(middleware.LoggerMiddleware)

		// Recovery
		Router.Use(gin.Recovery())
	} else {
		gin.SetMode(gin.DebugMode)
		Router = gin.Default()
		allowedOrigins := config.AllowedOrigins
		Router.Use(cors.Middleware(cors.Config{
			Origins:         allowedOrigins,
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     false,
			ValidateHeaders: false,
		}))
	}

	apiRoutes := Router.Group("/api")

	userRouter(apiRoutes)
	matchRouter(apiRoutes)
	generalRouter(apiRoutes)

}
