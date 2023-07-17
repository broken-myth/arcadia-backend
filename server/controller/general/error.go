package controller

import (
	"net/http"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

type ErrorRequest struct {
	Error string `json:"error" form:"error" binding:"required"`
}

// LogClientSideError godoc
//
//	@Summary		Log Client Side Error
//	@Description	Log Client Side Error
//	@Tags			General
//	@Param			message	formData	string	true	"Client Side Error message"
//	@Produce		json
//	@Success		200	{string}	string					"Success"
//	@Failure		401	{object}	helper.ErrorResponse	"Unauthorized"
//
//	@Failure		400	{object}	helper.ErrorResponse	"Internal Error"
//
//	@Router			/api/error [post]
//
//	@Security		ApiKeyAuth
func LogClientSideErrorPOST(c *gin.Context) {
	var req ErrorRequest

	log := utils.GetControllerLogger("api/error [POST]")

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendResponse(c, http.StatusBadRequest, "Error")
		return
	}

	log.Errorln(req)

	helper.SendResponse(c, http.StatusOK, "Error sent to backened")
}
