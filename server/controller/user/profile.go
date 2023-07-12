package controller

import (
	"net/http"
	"strconv"

	"github.com/delta/arcadia-backend/database"
	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetProfileResponse struct {
	Username         string `json:"username"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	College          string `json:"college"`
	Contact          string `json:"contact"`
	Trophies         uint   `json:"trophies"`
	XP               uint   `json:"xp"`
	NumberOfMinicons int64  `json:"numberOfMinicons"`
	CharacterID      uint   `json:"characterId"`
	AvatarURL        string `json:"avatarUrl"`
	Rank             uint   `json:"rank"`
}

type UpdateProfileRequest struct {
	IntendedUpdate string `json:"intendedUpdate" binding:"required"`
	NewValue       string `json:"newValue" binding:"required"`
}

type UpdateProfileResponse struct {
	IntendedUpdate string `json:"intendedUpdate"`
	NewValue       string `json:"newValue"`
	AvatarURL      string `json:"avatarUrl"` // Empty if not Intended update is not character
}

func GetProfileGET(c *gin.Context) {
	userID := c.GetUint("userID")

	db := database.GetDB()

	var user model.User

	if err := db.Preload("UserRegistration").Preload("Character").First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendResponse(c, http.StatusBadRequest, "User not found")
		} else {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in getting user profile")
		}
	}

	var ownedMinicons model.OwnedMinicon
	var numOfMinicons int64

	if err := db.Model(&ownedMinicons).Where("owner_id = ?", userID).Count(&numOfMinicons).Error; err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Error in getting user profile")
		return
	}

	rank, err := helpers.GetUserRank(user.ID)

	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Error in getting user profile")
		return
	}

	res := GetProfileResponse{
		Username:         user.UserRegistration.Username,
		Name:             user.UserRegistration.Name,
		Email:            user.UserRegistration.Email,
		College:          user.UserRegistration.College,
		Contact:          user.UserRegistration.Contact,
		Trophies:         user.Trophies,
		XP:               user.XP,
		NumberOfMinicons: numOfMinicons,
		CharacterID:      user.Character.ID,
		AvatarURL:        user.Character.AvatarURL,
		Rank:             rank,
	}

	utils.SendResponse(c, http.StatusOK, res)
}

func UpdateUserProfilePOST(c *gin.Context) {
	userID := c.GetUint("userID")

	var req UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Error")
		return
	}

	db := database.GetDB()

	var user model.User

	if err := db.Preload("UserRegistration").First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.SendResponse(c, http.StatusBadRequest, "User not found")
		} else {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in updating user profile")
		}
	}

	switch req.IntendedUpdate {
	case "name":
		user.UserRegistration.Name = req.NewValue
	case "college":
		user.UserRegistration.College = req.NewValue
	case "contact":
		user.UserRegistration.Contact = req.NewValue
	case "character":
		parsedCharID, err := strconv.ParseUint(req.NewValue, 10, 32)
		if err != nil {
			utils.SendResponse(c, http.StatusBadRequest, "Invalid intended update")
			return
		}

		characterID := uint(parsedCharID)

		if err := db.First(&model.Character{}, characterID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				utils.SendResponse(c, http.StatusBadRequest, "Invalid CharacterID")
			} else {
				utils.SendResponse(c, http.StatusInternalServerError, "Error in updating user profile")
			}
		}

		user.CharacterID = characterID
	default:
		utils.SendResponse(c, http.StatusBadRequest, "Invalid intended update")
		return
	}

	var res UpdateProfileResponse

	if req.IntendedUpdate == "character" {
		// update User Table
		if err := db.Save(&user).Error; err != nil {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in updating user profile")
			return
		}

		var url string

		if err := db.Model(&user.Character).Select("avatar_url").
			Where("id = ?", user.CharacterID).Scan(&url).Error; err != nil {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in updating user profile")
			return
		}

		res = UpdateProfileResponse{
			IntendedUpdate: req.IntendedUpdate,
			NewValue:       req.NewValue,
			AvatarURL:      url,
		}

	} else {
		// update UserRegistration Table
		if err := db.Save(&user.UserRegistration).Error; err != nil {
			utils.SendResponse(c, http.StatusInternalServerError, "Error in updating user profile")
			return
		}

		res = UpdateProfileResponse{
			IntendedUpdate: req.IntendedUpdate,
			NewValue:       req.NewValue,
		}
	}

	utils.SendResponse(c, http.StatusOK, res)
}
