package controller

import (
	"fmt"
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignupRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	College  string `json:"college" form:"college" binding:"required"`
	Contact  string `json:"contact" form:"contact" binding:"required"`
}

const defaultXP = 0

func SignupUserPOST(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, "Error")
		return
	}

	id := c.GetUint("userID")

	params := map[string]interface{}{
		"userID": id,
	}

	log := utils.GetControllerLoggerWithFields("/api/signup/complete [POST]", params)

	db := database.GetDB()
	tx := db.Begin()

	var userDetails model.UserRegistration
	var usernameAlreadyTaken = true

	if err := db.Where("username = ?", req.Username).First(&userDetails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			usernameAlreadyTaken = false
		} else {
			log.Error("Error in fetching User:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
			return
		}
	}

	if usernameAlreadyTaken {
		helper.SendError(c, http.StatusBadRequest, "Username already taken. Please choose another username.")
		return
	}

	if err := tx.First(&userDetails, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.SendError(c, http.StatusBadRequest, "User has not Registered")
		} else {
			log.Error("Error in fetching UserReg:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
		}
		return
	}

	if userDetails.FormFilled {
		helper.SendError(c, http.StatusBadRequest, "User has already Filled the Form")
		return
	}

	userDetails.Username = req.Username
	userDetails.College = req.College
	userDetails.Contact = req.Contact
	userDetails.FormFilled = true

	if err := tx.Save(&userDetails).Error; err != nil {
		log.Error("Error in saving UserReg:", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	defaultTrophies, err := helper.GetConstant("default_trophy_count")
	if err != nil || defaultTrophies == 0 {
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, fmt.Sprint("Error in Signing Up User.",
			" Constants may not have been initted properly."))
		return
	}

	user := model.User{
		ID:                 userDetails.ID,
		Username:           userDetails.Username,
		UserRegistrationID: userDetails.ID,
		Trophies:           defaultTrophies,
		XP:                 defaultXP,
		CharacterID:        1,
	}

	if err := tx.Create(&user).Error; err != nil {
		log.Error("Error in creating user:", err)
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	var defaultLootboxes []model.Lootbox

	if err := tx.Find(&defaultLootboxes).Error; err != nil {
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	if len(defaultLootboxes) == 0 {
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	var generatedLootboxes []model.GeneratedLootbox

	for _, lootbox := range defaultLootboxes {
		generatedLootbox := model.GeneratedLootbox{
			UserID:    user.ID,
			LootboxID: lootbox.ID,
		}
		generatedLootboxes = append(generatedLootboxes, generatedLootbox)
	}

	if err := tx.Create(&generatedLootboxes).Error; err != nil {
		tx.Rollback()
		helper.SendError(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	tx.Commit()

	helper.SendResponse(c, http.StatusOK, "All Ready! Login to Begin")
}
