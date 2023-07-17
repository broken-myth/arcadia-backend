package controller

import (
	"net/http"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

// UpdateLeaderboardPOST godoc
//
//	@Summary		Update Leaderboard
//	@Description	Update Leaderboard
//	@Tags			Admin
//	@Produce		json
//	@Success		200	{string}	string					"Updated Successfully"
//	@Failure		400	{object}	helper.ErrorResponse	"Error in updating constants"
//	@Failure		401	{object}	helper.ErrorResponse	"Unauthorized"
//	@Router			/api/admin/update_leaderboard [patch]
//
//	@Security		ApiKeyAuth
func UpdateLeaderboardPATCH(c *gin.Context) {
	var log = utils.GetControllerLogger("api/update_leaderboard [POST]")

	if !c.MustGet("isAdmin").(bool) {
		log.Error("Unauthorized")
		helper.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := helper.UpdateRedis()

	if err != nil {
		log.Error("Error in Updating the Redis Leaderboard. Error: ", err)
		// !! Note: only send the err.Error() response for admin routes
		helper.SendError(c, http.StatusOK, "Error in Updating the Redis Leaderboard. Error: "+err.Error())
		return
	}

	helper.SendResponse(c, http.StatusOK, "Redis Leaderboard Updated")
}
