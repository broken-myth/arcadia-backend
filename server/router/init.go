package router

import (
	"github.com/delta/arcadia-backend/config"
	controller "github.com/delta/arcadia-backend/server/controller/general"
	"github.com/delta/arcadia-backend/server/middleware"
	"github.com/gin-gonic/gin"

	"time"

	docs "github.com/delta/arcadia-backend/docs"
	cors "github.com/itsjamie/gin-cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Router *gin.Engine

func Init() {
	config := config.GetConfig()

	if config.AppEnv == "DOCKER" {
		gin.SetMode(gin.ReleaseMode)
		Router = gin.New()

		// Recovery
		Router.Use(gin.Recovery())
	} else {
		gin.SetMode(gin.DebugMode)
		Router = gin.Default()

	}

	// Logger
	Router.Use(middleware.LoggerMiddleware)

	allowedOrigins := config.AllowedOrigins
	Router.Use(cors.Middleware(cors.Config{
		Origins:         allowedOrigins,
		Methods:         "GET, POST, PATCH, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	Router.Use(middleware.RateLimiter())
	{

		Router.GET("/", controller.Ping) // As a test controller
		docs.SwaggerInfo.BasePath = ""
		Router.GET("/"+config.SwaggerURL+"/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		apiRoutes := Router.Group("/api")

		userRouter(apiRoutes)
		miniconsRouter(apiRoutes)
		matchRouter(apiRoutes)
		adminRouter(apiRoutes)
		generalRouter(apiRoutes)

	}

}
