package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

type ConstantsResponse struct {
	Name  model.ConstantType `json:"name"`
	Value int                `json:"value"`
}

type UpdateConstantsRequest struct {
	Name     model.ConstantType `json:"name" form:"name" binding:"required"`
	NewValue int                `json:"newValue" form:"newValue" binding:"required"`
}

// GetConstants godoc
//
//	@Summary		Get all Constants
//	@Description	Get all Constants
//	@Tags			Admin
//	@Produce		json
//	@Success		200	{object}	[]ConstantsResponse
//	@Failure		400	{object}	helper.ErrorResponse
//	@Router			/api/admin/constants [get]
//
//	@Security		ApiKeyAuth
func GetConstantsGET(c *gin.Context) {
	var log = utils.GetControllerLogger("api/constants [GET]")
	db := database.GetDB()

	var response []ConstantsResponse

	if err := db.Model(&model.Constant{}).
		Select("name, value").
		Find(&response).
		Error; err != nil {
		log.Error("Unable to fetch constants. Error: ", err)
		// !! Note: only send the err.Error() response for admin routes
		helper.SendError(c, http.StatusBadRequest, "Unable to fetch constants. Error: "+err.Error())
		return
	}

	helper.SendResponse(c, http.StatusOK, response)
}

// UpdateConstant godoc
//
//	@Summary		Update a constant
//	@Description	Update a constant
//	@Tags			Admin
//	@Produce		json
//	@Param			name		formData	model.ConstantType		true	"Name of the constant"
//	@Param			newValue	formData	int						true	"New value of the constant"
//	@Success		200			{string}	string					"Success"
//	@Failure		400			{object}	helper.ErrorResponse	"Error in updating constants"
//	@Failure		401			{object}	helper.ErrorResponse	"Unauthorized"
//	@Router			/api/admin/constants [patch]
//
//	@Security		ApiKeyAuth
func UpdateConstantsPATCH(c *gin.Context) {
	var log = utils.GetControllerLogger("api/constants [POST]")

	if !c.MustGet("isAdmin").(bool) {
		log.Error("Unauthorized")
		helper.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req UpdateConstantsRequest

	if err := c.Bind(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, "Invalid Request")
		return
	}

	db := database.GetDB()

	if err := db.Model(&model.Constant{}).Where("Name = ?", req.Name).Update("value", req.NewValue).Error; err != nil {
		log.Error("Error in updating constants. Error: ", err)
		// !! Note: only send the err.Error() response for admin routes
		helper.SendError(c, http.StatusBadRequest, "Error in updating constants. Error: "+err.Error())
		return
	}

	err := helper.InitConstants()
	if err != nil {
		log.Error("Error in initiating constants. Error: ", err)
		// !! Note: only send the err.Error() response for admin routes
		helper.SendError(c, http.StatusBadRequest, "Error in initiating constants. Error: "+err.Error())
		return
	}

	helper.SendResponse(c, http.StatusOK, "Updated Successfully")
}
