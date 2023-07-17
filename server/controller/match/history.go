package controller

import (
	"errors"
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

type MatchDetailsResponse struct {
	MatchID      uint   `json:"matchID"`
	MatchType    string `json:"matchType"`
	TrophyChange int    `json:"trophyChange"`

	// Attacker Details
	OpponentUsername  string `json:"opponentUsername"`
	OpponentAvatarURL string `json:"opponentAvatarURL"`
}

type MatchResultResponse struct {
	MatchID uint
	Result  int

	// Battle Result Details
	WinnerTrophies int `gorm:"column:winnerTrophies"`
	LoserTrophies  int `gorm:"column:loserTrophies"`

	// Attacker Details
	AttackerID        uint
	AttackerUsername  string
	AttackerAvatarURL string

	// Defender Details
	DefenderID        uint
	DefenderUsername  string
	DefenderAvatarURL string
}

func calcMatchDetails(userID uint, battleResponse MatchResultResponse,
	res *MatchDetailsResponse) (err error) {
	res.MatchID = battleResponse.MatchID
	res.TrophyChange = 0

	if battleResponse.AttackerID == userID {
		res.MatchType = "Attack"
		res.OpponentUsername = battleResponse.DefenderUsername
		res.OpponentAvatarURL = battleResponse.DefenderAvatarURL

		if battleResponse.Result == 1 {
			res.TrophyChange = battleResponse.WinnerTrophies
		} else if battleResponse.Result == -1 {
			res.TrophyChange = battleResponse.LoserTrophies
		}
	} else if battleResponse.DefenderID == userID {
		res.MatchType = "Defence"
		res.OpponentUsername = battleResponse.AttackerUsername
		res.OpponentAvatarURL = battleResponse.AttackerAvatarURL

		if battleResponse.Result == 1 {
			res.TrophyChange = battleResponse.LoserTrophies
		} else if battleResponse.Result == -1 {
			res.TrophyChange = battleResponse.WinnerTrophies
		}
	} else {
		return errors.New("User not found in match")
	}

	return nil
}

type MatchHistoryResponse []MatchDetailsResponse

func getMatchHistory(userID uint) (matchHistory MatchHistoryResponse, err error) {
	var res MatchHistoryResponse

	// MatchResultResponse is used to store the result of the query
	var battleResponse []MatchResultResponse

	db := database.GetDB()

	if err := db.Model(&model.MatchmakingDetails{}).
		Select("matchmaking_details.id as match_id, matchmaking_details.attacker_id as attacker_id",
			"attacker_user.username as attacker_username, defender_user.username as defender_username",
			"matchmaking_details.defender_id as defender_id, battle_results.result as result",
			"battle_results.winner_trophies as winnerTrophies, battle_results.loser_trophies as loserTrophies",
			"attacker_character.image_url as attacker_avatar_url, defender_character.image_url as defender_avatar_url").
		// Join user_registrations table to get the user id of the user
		Joins("JOIN user_registrations as attacker_user on attacker_user.id = matchmaking_details.attacker_id").
		Joins("JOIN user_registrations as defender_user on defender_user.id = matchmaking_details.defender_id").
		// Join users table to get the character of the user
		Joins("JOIN users as attacker on attacker.user_registration_id = attacker_user.id").
		Joins("JOIN users as defender on defender.user_registration_id = defender_user.id").
		// Join characters table to get the avatar of the user
		Joins("JOIN characters as attacker_character on attacker_character.id = attacker.character_id").
		Joins("JOIN characters as defender_character on defender_character.id = defender.character_id").
		// Join battle_results table to get the result of the match
		Joins("JOIN battle_results on battle_results.match_id = matchmaking_details.id").
		Where("matchmaking_details.attacker_id = ?", userID).
		Or("matchmaking_details.defender_id = ?", userID).
		Order("matchmaking_details.updated_at desc").
		Find(&battleResponse).
		Error; err != nil {
		return res, err
	}

	for i := 0; i < len(battleResponse); i++ {
		var matchDetails MatchDetailsResponse
		err = calcMatchDetails(userID, battleResponse[i], &matchDetails)

		if err != nil {
			return res, err
		}
		res = append(res, matchDetails)
	}

	return res, nil
}

// GetMatchHistory godoc
//
//	@Summary		Get History of Matches
//	@Description	Get History of Matches
//	@Tags			Match
//	@Produce		json
//	@Success		200	{object}	controller.MatchHistoryResponse	"Success"
//	@Failure		401	{object}	helper.ErrorResponse			"Unauthorized"
//	@Failure		404	{object}	helper.ErrorResponse			"Not Found"
//	@Router			/api/match/history [get]
//
//	@Security		ApiKeyAuth
func GetMatchHistoryGET(c *gin.Context) {
	userID := c.GetUint("userID")

	log := utils.GetControllerLogger("/api/match/history [GET]")

	response, err := getMatchHistory(userID)

	if err != nil {
		log.Error(err)
		helper.SendResponse(c, http.StatusNotFound, "Unable to fetch match history")
		return
	}

	helper.SendResponse(c, http.StatusOK, response)
}
