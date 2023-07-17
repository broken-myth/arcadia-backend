package controller

import (
	"net/http"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/gin-gonic/gin"
)

// AdminVerify godoc
//
//	@Summary		Admin Verify
//	@Description	Checks if the admin is logged in
//	@Tags			Admin
//	@Success		200	{string}	string					"Success"
//	@Failure		401	{object}	helper.ErrorResponse	"Unauthorized"
//	@Router			/api/admin/verify [get]
//
//	@Security		ApiKeyAuth
func AdminVerifyGET(c *gin.Context) {
	if !c.MustGet("isAdmin").(bool) {
		helper.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	helper.SendError(c, http.StatusOK, "Admin Verified")
}
