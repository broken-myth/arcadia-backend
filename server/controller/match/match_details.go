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

type MatchDetailsResponse struct {
	MatchID          uint   `json:"matchID"`
	OpponentID       uint   `json:"opponentID"`
	OpponentUsername string `json:"opponentUsername"`
	MatchType        string `json:"matchType"`
	TrophyChange     int    `json:"trophyChange"`
}

type MatchResponse struct {
	UserID       uint                 `json:"userID"`
	Username     string               `json:"username"`
	MatchDetails MatchDetailsResponse `json:"matchDetails"`
}

func GetMatchDetailsGET(c *gin.Context) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "GetMatchDetails",
	})

	userID := c.GetUint("userID")

	matchID := c.Param("id")

	db := database.GetDB()

	var matchDetails model.BattleResult

	condition1 := "MatchmakingDetails.Attacker.UserRegistration"
	condition2 := "MatchmakingDetails.Defender.UserRegistration"

	if err := db.Preload(condition1).Preload(condition2).Where("id = ?", matchID).First(&matchDetails).Error; err != nil {
		l.Error("Error in finding Match Details")
		fmt.Println(err)
		utils.SendResponse(c, http.StatusBadRequest, "Details not found")
		return
	}

	if userID != matchDetails.MatchmakingDetails.AttackerID && userID != matchDetails.MatchmakingDetails.DefenderID {
		utils.SendResponse(c, http.StatusUnauthorized, "Access Denied")
		return
	}

	var response MatchResponse

	if matchDetails.MatchmakingDetails.AttackerID == userID {

		response.UserID = userID
		response.Username = matchDetails.MatchmakingDetails.Attacker.UserRegistration.Username
		response.MatchDetails.MatchID = matchDetails.ID
		response.MatchDetails.OpponentID = matchDetails.MatchmakingDetails.DefenderID
		response.MatchDetails.OpponentUsername = matchDetails.MatchmakingDetails.Defender.UserRegistration.Username
		response.MatchDetails.MatchType = "Attacker"

		if matchDetails.Result == 1 {
			response.MatchDetails.TrophyChange = int(matchDetails.TrophyGain)
		} else if matchDetails.Result == -1 {
			response.MatchDetails.TrophyChange = -int(matchDetails.TrophyLost)
		} else {
			response.MatchDetails.TrophyChange = 0
		}

	} else if matchDetails.MatchmakingDetails.DefenderID == userID {

		response.UserID = userID
		response.Username = matchDetails.MatchmakingDetails.Defender.UserRegistration.Username
		response.MatchDetails.MatchID = matchDetails.ID
		response.MatchDetails.OpponentID = matchDetails.MatchmakingDetails.AttackerID
		response.MatchDetails.OpponentUsername = matchDetails.MatchmakingDetails.Attacker.UserRegistration.Username
		response.MatchDetails.MatchType = "Defender"

		if matchDetails.Result == 1 {
			response.MatchDetails.TrophyChange = -int(matchDetails.TrophyLost)
		} else if matchDetails.Result == -1 {
			response.MatchDetails.TrophyChange = int(matchDetails.TrophyGain)
		} else {
			response.MatchDetails.TrophyChange = 0
		}
	}

	utils.SendResponse(c, http.StatusOK, response)
}
