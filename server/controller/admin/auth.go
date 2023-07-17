package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
	generalHelper "github.com/delta/arcadia-backend/server/helper/general"
	userHelper "github.com/delta/arcadia-backend/server/helper/user"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// passwords hashed with 14 rounds
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AdminLogin godoc
//
//	@Summary		Admin Login
//	@Description	Admin Login
//	@Tags			Admin
//	@Produce		json
//	@Param			username	formData	string					true	"Username of the admin"
//	@Param			password	formData	string					true	"Password of the admin"
//	@Success		200			{string}	string					"Success"
//	@Failure		400			{object}	helper.ErrorResponse	"Error in updating constants"
//	@Router			/api/admin/login [post]
func AdminLoginPOST(c *gin.Context) {
	var req AdminLoginRequest

	if err := c.Bind(&req); err != nil {
		generalHelper.SendError(c, http.StatusBadRequest, "Invalid Request")
		return
	}
	params := map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	}

	var log = utils.GetControllerLoggerWithFields("/api/admin/login [POST]", params)

	db := database.GetDB()

	var adminDetails model.Admin

	if err := db.Where("username = ?", req.Username).First(&adminDetails).Error; err != nil {
		log.Error("Username Not Found. Error: ", err)
		generalHelper.SendError(c, http.StatusBadRequest, "Incorrect Credentials")
		return
	}

	if checkPasswordHash(req.Password, adminDetails.Password) {
		// the userID 0 denotes to admin
		jwtToken, err := userHelper.GenerateToken(0)
		if err != nil {
			log.Error("Token Not generated. Error: ", err)
			generalHelper.SendError(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		generalHelper.SendResponse(c, http.StatusOK, jwtToken)
		return
	}

	generalHelper.SendError(c, http.StatusBadRequest, "Incorrect Credentials")
}
