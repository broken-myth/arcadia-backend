package controller

import (
	"fmt"
	"net/http"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MatchHistoryResponse struct {
	MatchID          uint   `json:"matchID"`
	OpponentID       uint   `json:"opponentID"`
	OpponentUsername string `json:"opponentUsername"`
	MatchType        string `json:"matchType"`
	TrophyChange     int    `json:"trophyChange"`
}

type Response struct {
	UserID       uint                   `json:"userID"`
	Username     string                 `json:"username"`
	MatchHistory []MatchHistoryResponse `json:"matchHistory"`
}

func GetMatchHistoryGET(c *gin.Context) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "GetMatchHistory",
	})

	userID := c.GetUint("userID")

	db := database.GetDB()

	var details []model.BattleResult

	condition1 := "MatchmakingDetails.Attacker.UserRegistration"
	condition2 := "MatchmakingDetails.Defender.UserRegistration"

	if err := db.Preload(condition1).Preload(condition2).Find(&details).Error; err != nil {
		l.Error("Error in finding Match History Details")
		fmt.Println(err)
		utils.SendResponse(c, http.StatusBadRequest, "Details not found")
		return
	}

	var response Response

	response.UserID = userID

	for i := 0; i < len(details); i++ {
		if details[i].MatchmakingDetails.AttackerID == userID {

			response.Username = details[i].MatchmakingDetails.Attacker.UserRegistration.Username
			obj := MatchHistoryResponse{
				MatchID:          details[i].ID,
				OpponentID:       details[i].MatchmakingDetails.DefenderID,
				OpponentUsername: details[i].MatchmakingDetails.Defender.UserRegistration.Username,
				MatchType:        "Attacker",
			}
			if details[i].Result == 1 {
				obj.TrophyChange = int(details[i].TrophyGain)
			} else if details[i].Result == -1 {
				obj.TrophyChange = -int(details[i].TrophyLost)
			} else {
				obj.TrophyChange = 0
			}
			response.MatchHistory = append(response.MatchHistory, obj)

		} else if details[i].MatchmakingDetails.DefenderID == userID {

			response.Username = details[i].MatchmakingDetails.Defender.UserRegistration.Username
			obj := MatchHistoryResponse{
				MatchID:          details[i].ID,
				OpponentID:       details[i].MatchmakingDetails.AttackerID,
				OpponentUsername: details[i].MatchmakingDetails.Attacker.UserRegistration.Username,
				MatchType:        "Defender",
			}
			if details[i].Result == 1 {
				obj.TrophyChange = -int(details[i].TrophyLost)
			} else if details[i].Result == -1 {
				obj.TrophyChange = int(details[i].TrophyGain)
			} else {
				obj.TrophyChange = 0
			}
			response.MatchHistory = append(response.MatchHistory, obj)
		}
	}

	utils.SendResponse(c, http.StatusOK, response)
}
