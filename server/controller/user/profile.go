package controller

import (
	"net/http"
	"strconv"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
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
	Trophies         int    `json:"trophies"`
	XP               uint   `json:"xp"`
	NumberOfMinicons int64  `json:"numberOfMinicons"`
	CharacterURL     string `json:"characterURL"`
	AvatarURL        string `json:"avatarUrl"`
	Rank             uint   `json:"rank"`
}

type UpdateProfileRequest struct {
	IntendedUpdate string `json:"intendedUpdate" form:"intendedUpdate" binding:"required"`
	NewValue       string `json:"newValue" form:"newValue" binding:"required"`
}

type UpdateProfileResponse struct {
	IntendedUpdate string `json:"intendedUpdate"`
	UpdatedValue   string `json:"newValue"`
}

type IntendedUpdateType string

const (
	Name      IntendedUpdateType = "name"
	College   IntendedUpdateType = "college"
	Contact   IntendedUpdateType = "contact"
	Character IntendedUpdateType = "character"
)

// GetProfile godoc
//
//	@Summary		Get user profile
//	@Description	Gets user profile
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	UpdateProfileResponse	"Success"
//	@Failure		401	{object}	helper.ErrorResponse	"Unauthorized"
//
//	@Failure		403	{object}	helper.ErrorResponse	"User not found"
//
//	@Failure		500	{object}	helper.ErrorResponse	"Internal Server Error"
//	@Router			/user/profile [GET]
//
//	@Security		ApiKeyAuth
func GetProfileGET(c *gin.Context) {
	userID := c.GetUint("userID")

	db := database.GetDB()

	var user model.User

	log := utils.GetControllerLogger("/user/profile [GET]")

	if err := db.Preload("UserRegistration").Preload("Character").First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Error("User not found:", err)
			helper.SendError(c, http.StatusForbidden, "User not found")
		} else {
			log.Error("Error in getting user profile:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in getting user profile")
		}
		return
	}

	var ownedMinicons model.OwnedMinicon
	var numOfMinicons int64

	if err := db.Model(&ownedMinicons).Where("owner_id = ?", userID).Count(&numOfMinicons).Error; err != nil {
		log.Error("Error in getting owned minicons:", err)
		helper.SendError(c, http.StatusInternalServerError, "Error in getting user profile")
		return
	}

	rank, err := helper.GetUserRank(user.ID)

	if err != nil {
		log.Error("Error in getting user rank:", err)
		helper.SendError(c, http.StatusInternalServerError, "Error in getting user profile")
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
		CharacterURL:     user.Character.ImageURL,
		Rank:             rank,
	}

	helper.SendResponse(c, http.StatusOK, res)
}

// UpdateProfile godoc
//
//	@Summary		Update user profile
//	@Description	Updates user profile
//	@Tags			User
//	@Produce		json
//	@Param			intendedUpdate	formData	IntendedUpdateType		true	"Intended update"
//	@Param			newValue		formData	string					true	"New value"
//	@Success		200				{object}	UpdateProfileResponse	"Success"
//	@Failure		400				{object}	helper.ErrorResponse	"Bad Request or Invalid intended update"
//	@Failure		401				{object}	helper.ErrorResponse	"Unauthorized"
//
//	@Failure		403				{object}	helper.ErrorResponse	"User not found"
//
//	@Failure		500				{object}	helper.ErrorResponse	"Internal Server Error"
//	@Router			/user/profile/update [PATCH]
//
//	@Security		ApiKeyAuth
func UpdateUserProfilePATCH(c *gin.Context) {
	userID := c.GetUint("userID")

	var req UpdateProfileRequest

	log := utils.GetControllerLogger("/user/profile/update [PATCH]")

	if err := c.Bind(&req); err != nil {
		log.Error("Binding Error:", err)
		helper.SendError(c, http.StatusBadRequest, "Error")
		return
	}

	db := database.GetDB()

	var user model.User

	if err := db.Preload("UserRegistration").First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Error("User not found:", err)
			helper.SendError(c, http.StatusForbidden, "User not found")
		} else {
			log.Error("Error in updating user profile:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in updating user profile")
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
			log.Error("Invalid intended update:", err)
			helper.SendError(c, http.StatusBadRequest, "Invalid intended update")
			return
		}

		characterID := uint(parsedCharID)

		if err := db.First(&model.Character{}, characterID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Error("Invalid CharacterID:", err)
				helper.SendError(c, http.StatusBadRequest, "Invalid CharacterID")
			} else {
				log.Error("Error in getting character:", err)
				helper.SendError(c, http.StatusInternalServerError, "Error in updating user profile")
			}
		}

		user.CharacterID = characterID
	default:
		helper.SendError(c, http.StatusBadRequest, "Invalid intended update")
		return
	}

	var res UpdateProfileResponse

	if req.IntendedUpdate == "character" {
		// update User Table
		if err := db.Save(&user).Error; err != nil {
			log.Error("Error in updating user profile:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in updating user profile")
			return
		}

		var url string

		if err := db.Model(&user.Character).Select("image_url").
			Where("id = ?", user.CharacterID).Scan(&url).Error; err != nil {
			log.Error("Error in getting avatar_url:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in updating user profile")
			return
		}

		res = UpdateProfileResponse{
			IntendedUpdate: req.IntendedUpdate,
			UpdatedValue:   url,
		}

	} else {
		// update UserRegistration Table
		if err := db.Save(&user.UserRegistration).Error; err != nil {
			log.Error("Error in updating user_registration:", err)
			helper.SendError(c, http.StatusInternalServerError, "Error in updating user profile")
			return
		}

		res = UpdateProfileResponse{
			IntendedUpdate: req.IntendedUpdate,
			UpdatedValue:   req.NewValue,
		}
	}

	helper.SendResponse(c, http.StatusOK, res)
}
