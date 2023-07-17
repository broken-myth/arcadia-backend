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

type PerkData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PerkValue   uint   `json:"perkValue"`
}

type GetMiniconResponse struct {
	Name        string     `json:"name"`
	Health      uint       `json:"health"`
	Attack      uint       `json:"attack"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	XP          uint       `json:"xp"`
	Type        string     `json:"type"`
	Perks       []PerkData `json:"perks"`
}

// GetMiniconDetails godoc
//
//	@Summary		Get Minicon Details
//	@Description	Get details of a minicon
//	@Tags			Minicon
//	@Produce		json
//	@Param			id	path		uint							true	"Minicon Id"
//	@Success		200	{object}	controller.GetMiniconResponse	"Success"
//	@Failure		401	{object}	helper.ErrorResponse			"Unauthorized"
//
//	@Failure		403	{object}	helper.ErrorResponse			"Not Owned or Doesn't Exist"
//
//	@Failure		500	{object}	helper.ErrorResponse			"Internal Server Error"
//	@Router			/api/minicon/:id [get]
//
//	@Security		ApiKeyAuth
func GetMiniconDetailsGET(c *gin.Context) {
	miniconID := c.Param("id")

	userID := c.GetUint("userID")

	log := utils.GetControllerLogger("api/minicon/:id [GET]")

	db := database.GetDB()

	var ownedMinicon model.OwnedMinicon

	if err := db.Preload("Minicon").Preload("Minicon.Type").
		Where("minicon_id = ? AND owner_id = ?", miniconID, userID).
		First(&ownedMinicon).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.SendError(c, http.StatusForbidden, "You do not own this minicon")
		} else {
			log.Error("Error while fetching owned_minicon. error = ", err)
			helper.SendError(c, http.StatusInternalServerError, "Unable to fetch minicon details. Please refresh.")
		}
		return
	}
	var ownedPerks []model.OwnedPerk

	if err := db.Preload("Perk").Where("owned_minicon_id = ? ", ownedMinicon.ID).Find(&ownedPerks).Error; err != nil {
		log.Error(err)
		helper.SendError(c, http.StatusInternalServerError, "Some Error occurred")
	}

	var miniconPerk []PerkData

	for _, ownedPerk := range ownedPerks {
		miniconPerk = append(miniconPerk, PerkData{Name: ownedPerk.Perk.PerkName,
			Description: ownedPerk.Perk.Description,
			PerkValue:   ownedPerk.PerkValue})
	}

	res := GetMiniconResponse{
		Name:        ownedMinicon.Minicon.Name,
		Health:      ownedMinicon.Minicon.BaseHealth,
		Attack:      ownedMinicon.Attack,
		Description: ownedMinicon.Minicon.Description,
		Image:       ownedMinicon.Minicon.ImageLink,
		Type:        ownedMinicon.Minicon.Type.Name,
		XP:          ownedMinicon.XP,
		Perks:       miniconPerk,
	}

	helper.SendResponse(c, http.StatusOK, res)
}
