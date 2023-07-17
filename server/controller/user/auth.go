package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	userHelper "github.com/delta/arcadia-backend/server/helper/user"
	"github.com/delta/arcadia-backend/utils"

	"github.com/delta/arcadia-backend/server/model"
)

type AuthUserRequest struct {
	Code     string `json:"code" binding:"required"`
	AuthType string `json:"authType" binding:"required"`
}

func AuthUserPOST(c *gin.Context) {
	var req AuthUserRequest

	log := utils.GetControllerLogger("api/user/auth [POST]")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid Request Error:", err)
		helper.SendError(c, http.StatusBadRequest, "Invalid Request")
		return
	}

	if req.Code == "" || (req.AuthType != "LOGIN" && req.AuthType != "SIGNUP") {
		helper.SendError(c, http.StatusBadRequest, "Invalid Request")
		return
	}

	token, err := userHelper.GetOAuth2Token(req.Code)

	if err != nil {
		log.Error("Error in getting token:", err)
		helper.SendError(c, http.StatusInternalServerError, "Some error occurred, Refresh and Try Again")
		return
	}

	user, err := userHelper.GetOAuth2User(token.AccessToken, token.IDToken)

	if err != nil {
		log.Error("Error in getting user:", err)
		helper.SendError(c, http.StatusInternalServerError, "Some error occurred, Refresh and Try Again")
		return
	}

	Name := user.Name
	Email := user.Email

	if len(Name) == 0 || len(Email) == 0 {
		log.Error("Invalid Name or Email")
		helper.SendError(c, http.StatusInternalServerError, "Unable to find User")
		return
	}
	isLogin := false

	if req.AuthType == "LOGIN" {
		isLogin = true
	}

	db := database.GetDB()
	var userDetails model.UserRegistration

	if err = db.Where("Email = ?", Email).First(&userDetails).Error; err != nil {
		//Case : When the Email is not present in the database
		if err == gorm.ErrRecordNotFound {

			// Sign up to proceed
			if isLogin {
				log.Error("User not found")
				helper.SendError(c, http.StatusNotFound, "Please signup to continue")
				return
			}
			// If signup, permission granted
			userReg := model.UserRegistration{
				Email: Email,
				Name:  Name,
			}

			if err := db.Create(&userReg).Error; err != nil {
				log.Error("Failed to create user:", err)
				helper.SendError(c, http.StatusInternalServerError, "Unable to create user, Try Again")
				return
			}

			jwtToken, err := userHelper.GenerateToken(userReg.ID)

			if err != nil {
				log.Error("Token Not generated:", err)
				helper.SendError(c, http.StatusInternalServerError, "Unknown Error occurred, Try Again")
				return
			}

			helper.SendResponse(c, http.StatusOK, jwtToken)
			return
		}

		log.Error("Error in finding User")
		helper.SendError(c, http.StatusInternalServerError, "Unknown Error occurred, Try Again")
		return
	}

	//Case User is present in DB
	if isLogin {
		statusCode := http.StatusOK

		//If form is not filled when state is login
		if !userDetails.FormFilled {
			statusCode = http.StatusForbidden
		}

		//if the form is fully filled, token is generated for login state
		jwtToken, err := userHelper.GenerateToken(userDetails.ID)

		if err != nil {
			log.Error("Token Not generated:", err)
			helper.SendError(c, http.StatusInternalServerError, "Unknown Error occurred, Try Again")
			return
		}

		helper.SendResponse(c, statusCode, jwtToken)
		return
	}
	// if the user has fully filled the form when the state is signup
	if userDetails.FormFilled {
		log.Error("User already registered")
		helper.SendError(c, http.StatusConflict, "User already registered, Login to continue")
		return
	}
	// if the user has not filled the form when the state is signup
	jwtToken, err := userHelper.GenerateToken(userDetails.ID)
	if err != nil {
		log.Error("Token Not generated:", err)
		helper.SendError(c, http.StatusInternalServerError, "Unknown Error occurred, Try Again")
		return
	}
	helper.SendResponse(c, http.StatusOK, jwtToken)
}
