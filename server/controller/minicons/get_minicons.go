package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MiniconID struct {
	MiniconID uint `json:"miniconID"`
}

type UnlockedMiniconsResponse struct {
	Name           string `json:"name"`
	ImageLink      string `json:"imageLink"`
	Level          uint   `json:"level"`
	MiniconID      uint   `json:"miniconID"`
	OwnedMiniconID uint   `json:"ownedMiniconID"`
	Type           string `json:"type"`
}

type LockedMiniconsResponse struct {
	Name      string `json:"name"`
	MiniconID uint   `json:"miniconID"`
}

type GetMiniconsResponse struct {
	Lineup   []UnlockedMiniconsResponse `json:"lineup"`
	Unlocked []UnlockedMiniconsResponse `json:"unlocked"`
	Locked   []LockedMiniconsResponse   `json:"locked"`
}

// converts something like [{1},{2}] to [1,2]
func minicontoMiniconID(ObjArr []MiniconID) []uint {
	var idArr []uint

	for i := 0; i < len(ObjArr); i++ {
		idArr = append(idArr, ObjArr[i].MiniconID)
	}

	return idArr
}

// Query for fetching lineup minicons
func fetchLineup(userID uint, db *gorm.DB) ([]UnlockedMiniconsResponse, error) {

	var lineup []UnlockedMiniconsResponse

	err := db.Model(model.Lineup{}).
		Select("minicons.name as name, minicons.image_link as image_link , owned_minicons.level as level",
			"minicons.id as  minicon_id, owned_minicons.id as owned_minicon_id, minicon_types.name as type").
		Joins("JOIN owned_minicons on owned_minicons.id=lineups.owned_minicon_id").
		Joins("JOIN minicons on minicons.id=owned_minicons.minicon_id").
		Joins("JOIN minicon_types on minicon_types.id= minicons.type_id").
		Where("creator_id = ?", userID).
		Find(&lineup).Error

	return lineup, err
}

// Query for fetching all owned minicons except those in the lineup
func fetchUnlocked(userID uint, db *gorm.DB) (unlockedMinicons []UnlockedMiniconsResponse, getLineupMiniconError error,
	getUnlockedMiniconsError error) {

	var lineupMiniconID []MiniconID

	getLineupMiniconError = db.Model(model.Lineup{}).Select("owned_minicons.minicon_id as minicon_id").
		Joins("JOIN owned_minicons on owned_minicons.id=lineups.owned_minicon_id").
		Where("creator_id = ?", userID).
		Find(&lineupMiniconID).Error

	var unlocked []UnlockedMiniconsResponse

	lineupMiniconIDs := minicontoMiniconID(lineupMiniconID)

	if len(lineupMiniconID) != 0 {
		getUnlockedMiniconsError = db.Model(model.OwnedMinicon{}).
			Select("minicons.name as name, minicons.image_link as image_link , owned_minicons.level as level",
				"minicons.id as minicon_id, owned_minicons.id as owned_minicon_id, minicon_types.name as type").
			Joins("JOIN minicons on minicons.id=owned_minicons.minicon_id").
			Joins("JOIN minicon_types on minicon_types.id= minicons.type_id").
			Where(" owner_id = ? AND minicons.id NOT IN ? ", userID, lineupMiniconIDs).
			Find(&unlocked).Error

	} else {
		getUnlockedMiniconsError = db.Model(model.OwnedMinicon{}).
			Select("minicons.name as name, minicons.image_link as image_link , owned_minicons.level as level",
				"minicons.id as minicon_id, owned_minicons.id as owned_minicon_id, minicon_types.name as type").
			Joins("JOIN minicons on minicons.id=owned_minicons.minicon_id").
			Joins("JOIN minicon_types on minicon_types.id= minicons.type_id").
			Where(" owner_id = ?", userID).
			Find(&unlocked).Error
	}

	return unlocked, getLineupMiniconError, getUnlockedMiniconsError
}

// Query for fetching all minicons except those in the owned minicons table
func fetchLocked(userID uint, db *gorm.DB) (lockedMinicons []LockedMiniconsResponse, getAllUnlockedMiniconsError error,
	getAllLockedMiniconsError error) {

	var allUnlockedMiniconID []MiniconID

	getAllUnlockedMiniconsError = db.Model(model.OwnedMinicon{}).
		Select("owned_minicons.minicon_id as minicon_id").
		Where("owner_id = ?", userID).
		Find(&allUnlockedMiniconID).Error

	miniconIDs := minicontoMiniconID(allUnlockedMiniconID)

	if len(miniconIDs) != 0 {
		getAllLockedMiniconsError = db.Model(model.Minicon{}).Select("minicons.name as name , minicons.id as minicon_id").
			Where("minicons.id NOT IN ?", miniconIDs).
			Find(&lockedMinicons).Error
	} else {
		getAllLockedMiniconsError = db.Model(model.Minicon{}).Select("minicons.name as name , minicons.id as minicon_id").
			Find(&lockedMinicons).Error
	}

	return lockedMinicons, getAllUnlockedMiniconsError, getAllLockedMiniconsError
}

// GetMinicons godoc
//
//	@Summary		Get All Minicons
//	@Description	Get All Minicons
//	@Tags			Minicon
//	@Produce		json
//	@Success		200	{object}	controller.GetMiniconsResponse	"Success"
//	@Failure		400	{object}	helper.ErrorResponse			"Bad Request"
//	@Failure		401	{object}	helper.ErrorResponse			"Unauthorized"
//	@Failure		500	{object}	helper.ErrorResponse			"Internal Server Error"
//	@Router			/api/minicon [GET]
//
//	@Security		ApiKeyAuth
func FetchMiniconsGET(c *gin.Context) {

	userID := c.GetUint("userID")

	db := database.GetDB()
	tx := db.Begin()

	log := utils.GetControllerLogger("api/minicon [GET]")

	Lineup, getLineupError := fetchLineup(userID, tx)

	if getLineupError != nil {
		log.Error("An unexpected error occurred while fetching lineup minicons. error = ", getLineupError)
		helper.SendError(c, http.StatusInternalServerError, "An unexpected error occurred. Please try again later") //
		return
	}

	Unlocked, getLineupMiniconError, getUnlockedMiniconsError := fetchUnlocked(userID, tx)

	if getLineupMiniconError != nil || getUnlockedMiniconsError != nil {
		log.Error("An unexpected error occurred while fetching unlocked minicons. errors = ",
			getLineupMiniconError, "\n", getUnlockedMiniconsError)
		helper.SendError(c, http.StatusInternalServerError, "An unexpected error occurred. Please try again later")
		return
	}

	Locked, getAllUnlockedMiniconsError, getAllLockedMiniconsError := fetchLocked(userID, tx)

	if getAllUnlockedMiniconsError != nil || getAllLockedMiniconsError != nil {
		log.Error("An unexpected error occurred while fetching locked minicons. errors = ",
			getAllUnlockedMiniconsError, "\n", getAllLockedMiniconsError) //
		helper.SendError(c, http.StatusInternalServerError, "An unexpected error occurred. Please try again later")
		return
	}

	tx.Commit()

	res := GetMiniconsResponse{
		Lineup:   Lineup,
		Unlocked: Unlocked,
		Locked:   Locked,
	}

	helper.SendResponse(c, http.StatusOK, res)
}
