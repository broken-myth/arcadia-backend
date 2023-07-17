package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/simulator"
	"github.com/delta/arcadia-backend/utils"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	matchHelper "github.com/delta/arcadia-backend/server/helper/match"
)

type OpponentDetails struct {
	Username     string `json:"username"`
	Xp           uint   `json:"xp"`
	Trophies     uint   `json:"trophies"`
	CharacterURL string `json:"characterURL"`
}

type StartMatchResponse struct {
	MatchID         uint            `json:"matchId"`
	OpponentDetails OpponentDetails `json:"opponent"`
}

// StartMatch godoc
//
//	@Summary		Start a Match
//	@Description	Start a Match
//	@Tags			Match
//	@Produce		json
//	@Success		200	{object}	controller.StartMatchResponse	"Success"
//	@Failure		401	{object}	helper.ErrorResponse			"Unauthorized"
//
//	@Failure		403	{object}	helper.ErrorResponse			"Daily Attack Limit Reached (OR) Lineup not set (OR) Arena closed"
//
//	@Failure		500	{object}	helper.ErrorResponse			"Internal Server Error"
//	@Router			/api/match/start [get]
//
//	@Security		ApiKeyAuth
func StartMatchGET(c *gin.Context) {
	userID := c.GetUint("userID")

	var log = utils.GetControllerLogger("api/match/start [GET]")

	errorMessage := "Error in finding a match right now. Please try again later."
	db := database.GetDB()

	dailyAttackLimit, err := helper.GetConstant("daily_attack_limit")
	if err != nil {
		log.Error("Error in getting daily attacks limit. Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	var attacksToday int64
	today := fmt.Sprint(time.Now().Format("2006-01-02"), "%")

	if err = db.Model(&model.MatchmakingDetails{}).Where("attacker_id = ? AND created_at LIKE ?", userID, today).
		Count(&attacksToday).Error; err != nil {
		log.Error("Error in getting attacks today. Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	if attacksToday >= int64(dailyAttackLimit) {
		helper.SendError(c, http.StatusForbidden, "You have exhausted your attacks for today. Please try again tomorrow.")
		return
	}

	// i. finding a match
	defenderID, err := matchHelper.MatchMaker(userID)
	if err != nil {
		if err == helper.ErrNoLineup {
			helper.SendError(c, http.StatusForbidden, fmt.Sprint("You don't have a lineup. Create one in the dashboard.",
				" You might have to collect more minicons first"))
			return
		}
		log.Error("Error in finding opponent. Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	tx := db.Begin()

	// i. creatng a matchmaing_details DB entry

	matchID, err := matchHelper.CreateMatchmakingDetails(tx, userID, defenderID)

	if err != nil {
		tx.Rollback()
		log.Error("Error while creating Matchmaking Details DB entry - ", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	attSurvivors, defSurvivors, err := simulator.Start(tx, matchID)
	if err != nil {
		log.Error("Error in simulation - ", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	// ToDo: Check if all below functions are even needed at the end
	// Fetching attacker and defender
	var attacker model.User
	if err := tx.First(&attacker, userID).Error; err != nil {
		log.Error("Error while finding attacker. Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}
	var defender model.User
	if err := db.Preload("UserRegistration").Preload("Character").
		Where("id = ?", defenderID).First(&defender).Error; err != nil {
		tx.Rollback()
		log.Error("Error in fetching user details of defender:", defenderID, ". Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	// ii. creating a battle_results DB entry
	err3 := matchHelper.CreateBattleResult(tx, matchID, attacker, defender,
		attSurvivors, defSurvivors)
	if err3 != nil {
		tx.Rollback()
		log.Error("Error in creating Battle Details DB entry - ", err3)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	if err = helper.UpdateXpUserLineup(tx, userID); err != nil {
		tx.Rollback()
		log.Error("Error in updating XP of user lineup. Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	if err = helper.UpdateXpUserLineup(tx, defenderID); err != nil {
		tx.Rollback()
		log.Error("Error in updating XP of user lineup. Error:", err)
		helper.SendError(c, http.StatusInternalServerError, errorMessage)
		return
	}

	response := StartMatchResponse{
		MatchID: matchID,
		OpponentDetails: OpponentDetails{
			Username:     defender.UserRegistration.Username,
			Xp:           defender.XP,
			Trophies:     uint(defender.Trophies),
			CharacterURL: defender.Character.ImageURL,
		},
	}

	tx.Commit()
	helper.SendResponse(c, http.StatusOK, response)
}
