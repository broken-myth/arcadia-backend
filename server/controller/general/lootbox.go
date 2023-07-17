package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Lootbox struct {
	X         string `json:"x"`      // Encrypted X coordinate of the lootbox
	Y         string `json:"y"`      // Encrypted Y coordinate of the lootbox
	IsOpen    bool   `json:"isOpen"` // Is the lootbox open?
	LootboxID uint   `json:"lootboxID"`
	Region    string `json:"region"`
}

type LootboxGETResponse struct {
	Lootboxes []Lootbox `json:"lootboxes"`
}

type LootboxOpenPOSTRequest struct {
	X         int  `json:"x" form:"x" binding:"required"` // Player's Tile X coordinate
	Y         int  `json:"y" form:"y" binding:"required"` // Player's Tile Y coordinate
	LootboxID uint `json:"lootboxID" form:"lootboxID" binding:"required"`
}

type Unlocked struct {
	MiniconName        string `json:"name"`
	MiniconDescription string `json:"description"`
	MiniconImage       string `json:"image"`
}

type LootboxOpenPOSTResponse struct {
	Unlocked  Unlocked `json:"unlocked"`
	LootboxID uint     `json:"lootboxID"`
}

// GetLootbox godoc
//
//	@Summary		Get Lootboxes
//	@Description	Get all lootboxes of the user
//	@Tags			General
//	@Produce		json
//	@Success		200	{object}	controller.LootboxGETResponse	"Success"
//	@Failure		401	{object}	helper.ErrorResponse			"Unauthorized"
//
//	@Failure		500	{object}	helper.ErrorResponse			"Internal Error"
//
//	@Router			/api/lootbox [get]
//
//	@Security		ApiKeyAuth
func LootboxGET(c *gin.Context) {
	userID := c.GetUint("userID")

	var lootboxes []model.GeneratedLootbox

	log := utils.GetControllerLogger("/api/lootbox [GET]")

	db := database.GetDB()

	if err := db.Preload("Lootbox.Region").Where("user_id = ?", userID).
		Find(&lootboxes).Error; err != nil {
		log.Errorln("Error while getting lootboxes: ", err)
		helper.SendError(c, http.StatusInternalServerError, "Unable to get lootboxes, Please try again later")
		return
	}

	tokenArray := strings.Split(c.Request.Header.Get("Authorization"), " ")

	if len(tokenArray) != 2 {
		log.Errorln("Token not found")
		helper.SendError(c, http.StatusUnauthorized, "You are not authorized to perform this action")
		return
	}

	token := tokenArray[1]

	if token == "" {
		log.Errorln("Token not found")
		helper.SendError(c, http.StatusUnauthorized, "You are not authorized to perform this action")
		return
	}

	var lootboxResponse []Lootbox
	var encryptedX string
	var encryptedY string
	var err error

	key := utils.GenerateKey(token)

	for _, lootbox := range lootboxes {
		stringX := key + strconv.Itoa(lootbox.Lootbox.X)
		stringY := key + strconv.Itoa(lootbox.Lootbox.Y)

		if encryptedX, err = utils.Encrypt(stringX, key); err != nil {
			log.Errorln("Error while encrypting lootbox X: ", err)
			helper.SendError(c, http.StatusInternalServerError, "Unable to get lootboxes, Please try again later")
			return
		}

		if encryptedY, err = utils.Encrypt(stringY, key); err != nil {
			log.Errorln("Error while encrypting lootbox Y: ", err)
			helper.SendError(c, http.StatusInternalServerError, "Unable to get lootboxes, Please try again later")
			return
		}

		lootboxResponse = append(lootboxResponse, Lootbox{
			X:         encryptedX,
			Y:         encryptedY,
			IsOpen:    lootbox.IsOpen,
			LootboxID: lootbox.ID,
			Region:    lootbox.Lootbox.Region.Name,
		})
	}

	var res LootboxGETResponse

	res.Lootboxes = lootboxResponse

	helper.SendResponse(c, http.StatusOK, res)
}

// LootboxOpen godoc
//
//	@Summary		Open Lootbox
//	@Description	Open a lootbox
//	@Tags			General
//	@Produce		json
//	@Param			x			formData	uint								true	"Player's Tile X coordinate"
//	@Param			y			formData	uint								true	"Player's Tile Y coordinate"
//	@Param			lootboxID	formData	uint								true	"Lootbox ID"
//	@Success		200			{object}	controller.LootboxOpenPOSTResponse	"Success"
//	@Failure		401			{object}	helper.ErrorResponse				"Unauthorized"
//	@Failure		403			{object}	helper.ErrorResponse				"Minicon Limit Reached"
//	@Failure		409			{object}	helper.ErrorResponse				"Lootbox already opened"
//	@Failure		418			{object}	helper.ErrorResponse				"LOL, Nice Try"
//	@Failure		500			{object}	helper.ErrorResponse				"Internal Error"
//	@Router			/api/lootbox/open [post]
//
//	@Security		ApiKeyAuth
func LootboxOpenPOST(c *gin.Context) {
	var req LootboxOpenPOSTRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		helper.SendError(c, http.StatusBadRequest, "Bad request")
		return
	}

	params := map[string]interface{}{
		"x":         req.X,
		"y":         req.Y,
		"lootboxID": req.LootboxID,
	}

	userID := c.GetUint("userID")

	log := utils.GetControllerLoggerWithFields("/api/lootbox/open [POST]", params)

	var lootbox model.GeneratedLootbox

	db := database.GetDB()
	tx := db.Begin()

	// get owned minicons count
	// if gte maxunlocked, return error

	maxMiniconCount, err := helper.GetConstant("max_unlocked_minicons")
	if err != nil {
		log.Error("Error while fetching constant max_unlocked_minicons. Error = ", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Please try again later")
		return
	}

	var ownedMiniconsCount int64
	if err := db.Where("owner_id = ?", userID).Find(&model.OwnedMinicon{}).Count(&ownedMiniconsCount).Error; err != nil {
		tx.Rollback()
		log.Error("Error in fetching count of owned minicons of user:", userID, ". Error:", err)
		helper.SendError(c, http.StatusInternalServerError, "Please try again later")
		return
	}

	if int(ownedMiniconsCount) >= maxMiniconCount {
		tx.Rollback()
		// log.Warn("User has reached the maximum number of minicons he can unlock today. User ID = ", userID,
		// 	". Max unlocked minicons = ", maxMiniconCount, ". Owned minicons count = ", ownedMiniconsCount)
		helper.SendError(c, http.StatusForbidden, fmt.Sprint("You have reached the maximum number of minicons",
			" you can unlock today. Please try again tomorrow"))
		return
	}

	if err := tx.Preload("Lootbox").First(&lootbox, "id = ? AND user_id = ?", req.LootboxID, userID).Error; err != nil {
		log.Errorln(err)
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			helper.SendError(c, http.StatusNotFound, "Could not find lootboxes, Try again")
			return
		}
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	if lootbox.IsOpen {
		log.Errorln("Lootbox was already opened")
		tx.Rollback()
		helper.SendError(c, http.StatusConflict, "Lootbox already opened!")
		return
	}

	if math.Abs(float64(lootbox.Lootbox.X-req.X)) >= 3 && math.Abs(float64(lootbox.Lootbox.Y-req.Y)) >= 3 {
		log.Errorln("Player is too far away from the lootbox")
		tx.Rollback()
		helper.SendError(c, http.StatusTeapot, "Nice try bro") // Change this if you dare to
		return
	}

	lootbox.IsOpen = true
	var minicon model.Minicon

	if err := tx.First(&minicon, "id = ?", lootbox.Lootbox.UnlocksID).Error; err != nil {
		log.Errorln("Error finding minicon: ", err)
		tx.Rollback()

		if err == gorm.ErrRecordNotFound {
			helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
			return
		}

		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	defaultXP, err := helper.GetConstant("xp_base_count")
	if err != nil {
		log.Error("Error while fetching constant xp_base_count. Error = ", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Please try again later")
		return
	}

	var ownedMinicon model.OwnedMinicon

	if err := tx.First(&ownedMinicon, "owner_id = ? AND minicon_id = ?", userID, minicon.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ownedMinicon = model.OwnedMinicon{
				MiniconID: minicon.ID,
				OwnerID:   userID,
				XP:        uint(defaultXP),

				// Set base stats
				Attack: minicon.BaseAttack,
				Health: minicon.BaseHealth,
			}
			if err := tx.Create(&ownedMinicon).Error; err != nil {

				log.Errorln("Error creating owned minicon: ", err)
				tx.Rollback()
				helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
				return

			}
		} else {
			log.Errorln("Error retrieving owned minicon: ", err)
			tx.Rollback()
			helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
			return
		}
	} else {
		log.Errorln("Found owned minicon of unopened lootbox: ", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	var perks []model.Perk

	if err := tx.Find(&perks, "minicon_id = ?", minicon.ID).Error; err != nil {
		log.Errorln("Error fetching perks: ", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	for _, perk := range perks {
		var ownedPerk model.OwnedPerk
		if err := tx.First(&ownedPerk, "owned_minicon_id = ? AND perk_id = ?", ownedMinicon.ID, perk.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ownedPerk = model.OwnedPerk{
					PerkValue:      perk.BaseValue,
					OwnedMiniconID: ownedMinicon.ID,
					PerkID:         perk.ID,
				}
				if err := tx.Create(&ownedPerk).Error; err != nil {
					log.Errorln("Error creating owned perks: ", err)
					tx.Rollback()
					helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
					return
				}
			} else {
				log.Errorln("Error retrieving owned perk: ", err)
				tx.Rollback()
				helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
				return
			}
		} else {
			log.Errorln("Found existing owned perk for a new owned minicon: ", err)
			tx.Rollback()
			helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
			return
		}
	}

	if err := tx.Save(&lootbox).Error; err != nil {
		log.Errorln("Error saving opened lootbox. Error: ", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorln("Error committing transaction: ", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	helper.SendResponse(c, http.StatusOK, LootboxOpenPOSTResponse{
		Unlocked: Unlocked{
			MiniconName:        minicon.Name,
			MiniconDescription: minicon.Description,
			MiniconImage:       minicon.ImageLink,
		},
		LootboxID: req.LootboxID,
	})
}
