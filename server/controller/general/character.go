package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

type Character struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl"`
	AvatarURL   string `json:"avatarUrl"`
}

type GetCharactersResponse struct {
	Characters []Character `json:"characters"`
}

// GetCharacters godoc
//
//	@Summary		Get all Characters
//	@Description	Get all Characters
//	@Tags			General
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	controller.GetCharactersResponse	"Success"
//	@Failure		500	{object}	helper.ErrorResponse				"Internal Server Error"
//	@Router			/api/characters [get]
func GetCharactersGET(c *gin.Context) {
	var characters []Character

	db := database.GetDB()

	log := utils.GetControllerLogger("/api/characters [GET]")

	if err := db.Model(model.Character{}).Omit("created_at", "updated_at", "deleted_at").
		Find(&characters).Error; err != nil {
		log.Errorln(err)
		helper.SendError(c, http.StatusInternalServerError, "Unable to get characters, Please try again later")
		return
	}

	res := GetCharactersResponse{
		Characters: characters,
	}

	helper.SendResponse(c, http.StatusOK, res)
}
