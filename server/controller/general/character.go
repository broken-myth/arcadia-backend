package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
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

func GetCharactersGET(c *gin.Context) {
	var characters []model.Character

	db := database.GetDB()

	if err := db.Find(&characters).Error; err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, "Error while fetching characters")
	}

	var res GetCharactersResponse

	for _, character := range characters {
		res.Characters = append(res.Characters, Character{
			ID:          character.ID,
			Name:        character.Name,
			Description: character.Description,
			ImageURL:    character.ImageURL,
			AvatarURL:   character.AvatarURL,
		})
	}

	utils.SendResponse(c, http.StatusOK, res)
}
