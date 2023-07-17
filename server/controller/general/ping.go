package controller

import (
	"net/http"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/gin-gonic/gin"
)

// Ping godoc
//
//	@Summary		Ping
//	@Description	Checks if the server is up and running
//	@Tags			General
//	@Success		200	{object}	controller.LootboxOpenPOSTResponse	"Success"
//	@Router			/ [get]
func Ping(c *gin.Context) {
	helper.SendResponse(c, http.StatusOK, "pong")
}
