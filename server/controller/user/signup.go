package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	College  string `json:"college" binding:"required"`
	Contact  string `json:"contact" binding:"required"`
}

const defaultXP = 0

func SignupUserPOST(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Error")
		return
	}

	id := c.GetUint("userID")

	db := database.GetDB()

	var userDetails model.UserRegistration
	var usernameAlreadyTaken = true

	if err := db.Where("username = ?", req.Username).First(&userDetails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			usernameAlreadyTaken = false
		} else {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in Signing Up User")
			return
		}
	}

	if usernameAlreadyTaken {
		utils.SendResponse(c, http.StatusBadRequest, "Username already taken")
		return
	}

	if err := db.First(&userDetails, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendResponse(c, http.StatusBadRequest, "User has not Registered")
		} else {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in Signing Up User")
		}
		return
	}

	if userDetails.FormFilled {
		utils.SendResponse(c, http.StatusBadRequest, "User has already Filled the Form")
		return
	}

	userDetails.Username = req.Username
	userDetails.College = req.College
	userDetails.Contact = req.Contact
	userDetails.FormFilled = true

	tx := db.Begin()

	if err := tx.Save(&userDetails).Error; err != nil {
		tx.Rollback()
		utils.SendResponse(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	defaultTrophies, err := helpers.GetConstant("default_trophy_count")
	if err != nil {
		tx.Rollback()
		utils.SendResponse(c, http.StatusInternalServerError, "Error in Signing Up User4")
		return
	}

	user := model.User{
		ID:                 userDetails.ID,
		UserRegistrationID: userDetails.ID,
		Trophies:           uint(defaultTrophies),
		XP:                 defaultXP,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		utils.SendResponse(c, http.StatusInternalServerError, "Error in Signing Up User")
		return
	}

	err = helpers.InsertNewUserRedis(userDetails.ID)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Failed to initialise user")
		return
	}

	tx.Commit()

	utils.SendResponse(c, http.StatusOK, "All Ready! Login to Begin")
}
