package controller

import (
	"fmt"
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	miniconHelper "github.com/delta/arcadia-backend/server/helper/minicon"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateLineupRequest struct {
	LineupIDArr []uint `json:"lineupIDArr" binding:"required"`
}

type UpdateLineupResponse struct {
	MiniconID uint `json:"miniconID"`
}

// UpdateLineup godoc
//
//	@Summary		Update Minicon Lineup
//	@Description	Update Minicon Lineup
//	@Tags			Minicon
//	@Produce		json
//
//	@Param			json	body		UpdateLineupRequest				true	"Lineup ID Array"
//
//	@Success		200		{object}	controller.UpdateLineupResponse	"Success"
//	@Failure		400		{object}	helper.ErrorResponse			"Bad Request"
//	@Failure		401		{object}	helper.ErrorResponse			"Unauthorized"
//	@Failure		500		{object}	helper.ErrorResponse			"Internal Server Error"
//	@Router			/api/minicon/updateLineup [patch]
//
//	@Security		ApiKeyAuth
func UpdateLineupPATCH(c *gin.Context) {

	userID := c.GetUint("userID")

	var req UpdateLineupRequest

	if err := c.Bind(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, "Invalid request, please check the request body")
		return
	}

	params := map[string]interface{}{
		"lineupIDArr": req.LineupIDArr,
		"userID":      userID,
	}

	var log = utils.GetControllerLoggerWithFields("/api/minicons/update_lineup [POST]", params)

	db := database.GetDB()
	tx := db.Begin()

	MiniconsInLineup, err := helper.GetConstant("minicons_in_lineup")
	if err != nil {
		tx.Rollback()
		log.Error("Error fetching constant minicons_in_lineup")
		return
	}

	if len(req.LineupIDArr) < MiniconsInLineup {
		tx.Rollback()
		helper.SendError(c, http.StatusBadRequest, fmt.Sprint("Not enough minicons in the lineup.",
			" There should be ", MiniconsInLineup, " minicons in a lineup."))
		return
	} else if len(req.LineupIDArr) > MiniconsInLineup {
		tx.Rollback()
		log.Error("Excess number of minicons in the lineup")
		helper.SendError(c, http.StatusBadRequest, fmt.Sprint("Excess number of minicons in the lineup.",
			" There should be ", MiniconsInLineup, " minicons in a lineup."))
		return
	}

	if miniconHelper.DuplicateInArray(req.LineupIDArr) {
		tx.Rollback()
		helper.SendError(c, http.StatusBadRequest,
			"Duplicates exist in the minicon lineup, please create a new lineup and try again.")
		return
	}

	var ownedMinicon []model.OwnedMinicon

	for i := 0; i < len(req.LineupIDArr); i++ {
		if err := tx.Where("owner_id = ? AND id = ?", userID, req.LineupIDArr[i]).
			First(&ownedMinicon).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				log.Error("Unowned minicon in the lineup")
				helper.SendError(c, http.StatusBadRequest, "There are some minicons in the lineup that you don't own.")
			} else {
				tx.Rollback()
				log.Error("Error fetching Lineup. Error: ", err)
				helper.SendError(c, http.StatusInternalServerError, "An unexpected error occurred. Please try again later")
			}
			return
		}
	}

	//to check whether user already had a lineup
	doesLineUpExist := true
	var existingLineup model.Lineup

	//check if the user has a lineup
	if err := tx.Where("creator_id = ?", userID).First(&existingLineup).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			doesLineUpExist = false
		} else {
			tx.Rollback()
			log.Error("Error fetching Lineup. Error: ", err)
			helper.SendError(c, http.StatusBadRequest, "An unexpected error occurred. Please try again later")
			return
		}
	}

	var updateLineup []model.Lineup

	if doesLineUpExist {
		for i := 0; i < len(req.LineupIDArr); i++ {
			if err := tx.Model(&updateLineup).Where("creator_id = ? AND position_no = ?", userID, uint(i+1)).
				Update("owned_minicon_id", req.LineupIDArr[i]).Error; err != nil {
				log.Error("Error updating Lineup. Error: ", err)
				tx.Rollback()
				helper.SendError(c, http.StatusBadRequest, "An unexpected error occurred. Please try again later")
				return
			}
		}
		tx.Commit()
		helper.SendResponse(c, http.StatusOK, "Lineup updated successfully")
		return
	}

	for i := 0; i < len(req.LineupIDArr); i++ {

		lineupRow := model.Lineup{CreatorID: userID, PositionNo: uint(i + 1), OwnedMiniconID: req.LineupIDArr[i]}

		if err := tx.Save(&lineupRow).Error; err != nil {
			tx.Rollback()
			log.Error("Error creating Lineup. Error: ", err)
			helper.SendError(c, http.StatusBadRequest, "An unexpected error occurred. Please try again later")
			return
		}
	}

	// add to leaderboard
	err = helper.InsertNewUserRedis(userID)
	if err != nil {
		tx.Rollback()
		log.Error("Error adding user to redis when creating a lineup. Error: ", err)
		return
	}

	tx.Commit()

	helper.SendResponse(c, http.StatusOK, "Lineup created successfully")
}
