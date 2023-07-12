package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
)

type AuthUserRequest struct {
	Code     string `json:"code" binding:"required"`
	AuthType string `json:"authType" binding:"required"`
}

func AuthUserPOST(c *gin.Context) {
	var req AuthUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
		return
	}

	if req.Code == "" || (req.AuthType != "LOGIN" && req.AuthType != "SIGNUP") {
		utils.SendResponse(c, http.StatusBadRequest, "Invalid Request")
		return
	}

	token, err := utils.GetOAuth2Token(req.Code)

	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Error in Authenticating User")
		return
	}

	user, err := utils.GetOAuth2User(token.AccessToken, token.IDToken)

	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Error in Authenticating User")
		return
	}

	Name := user.Name
	Email := user.Email

	if len(Name) == 0 || len(Email) == 0 {
		utils.SendResponse(c, http.StatusInternalServerError, "User not found")
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

			//Sign up to proceed
			if isLogin {
				utils.SendResponse(c, http.StatusNotFound, "User not found. Please signup.")
				return
			}
			//if signup, permission granted
			userReg := model.UserRegistration{
				Email: Email,
				Name:  Name,
			}
			if err := db.Create(&userReg).Error; err != nil {
				utils.SendResponse(c, http.StatusInternalServerError, "Failed to create user")
				return
			}
			jwtToken, err := utils.GenerateToken(userReg.ID)
			if err != nil {
				utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
				return
			}
			utils.SendResponse(c, http.StatusOK, jwtToken)
			return
		}
		utils.SendResponse(c, http.StatusInternalServerError, "Error in finding User")
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
		jwtToken, err := utils.GenerateToken(userDetails.ID)
		if err != nil {
			utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
			return
		}
		utils.SendResponse(c, statusCode, jwtToken)
		return
	}
	// if the user has fully filled the form when the state is signup
	if userDetails.FormFilled {
		utils.SendResponse(c, http.StatusConflict, "User already registered, Login to continue")
		return
	}
	// if the user has not filled the form when the state is signup
	jwtToken, err := utils.GenerateToken(userDetails.ID)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Token Not generated")
		return
	}
	utils.SendResponse(c, http.StatusOK, jwtToken)
}
