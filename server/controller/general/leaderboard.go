package controller

import (
	"net/http"

	"github.com/delta/arcadia-backend/database"
	helpers "github.com/delta/arcadia-backend/server/helpers/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LeaderboardRow struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Trophies  uint   `json:"trophies"`
	Rank      uint   `json:"rank"`
	XP        uint64 `json:"xp"`
	AvatarURL string `json:"avatar_url"`
}

// called by frontend client. Returns response of array of LeaderboardRows
func FetchLeaderboardGET(c *gin.Context) {
	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "FetchLeaderboardGET",
	})

	leaderboard, err := GetEntireLeaderboard()
	if err != nil {
		l.Error(err)
		utils.SendResponse(c, http.StatusInternalServerError, "Error fetching leaderboard. Please try again later")
		return
	}

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		utils.SendResponse(c, http.StatusOK, leaderboard)
		return
	}

	userID, err := utils.ValidateToken(authHeader)
	if err != nil {
		utils.SendResponse(c, http.StatusOK, leaderboard)
		return
	}

	var userRankDetails []LeaderboardRow
	// if user is logged in

	for { // to ensure user's rank doesnt change before we fetch it
		userRank, _ := helpers.GetUserRank(userID)
		userRankDetails, _ = GetLeaderboardRange(userRank, 1)
		if userRankDetails[0].UserID == userID {
			break
		}
	}

	// first value is user's leaderboard details (followed by the actual entire leaderboard)
	leaderboard = append(userRankDetails, leaderboard...)

	utils.SendResponse(c, http.StatusOK, leaderboard)
}

func GetEntireLeaderboard() (ranks []LeaderboardRow, err error) {
	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "GetEntireLeaderboard",
	})

	leaderboard, _ := GetLeaderboardRange(1, 0)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return leaderboard, nil
}

func GetLeaderboardRange(startRank uint, count uint) (ranks []LeaderboardRow, err error) {
	// set count = 0 to get all entries after a certain rank

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":      "GetLeaderboardRange",
		"param_start": startRank,
		"param_count": count,
	})

	results, err := helpers.GetUsersWithRankInRange(startRank, count)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	var leaderboard []LeaderboardRow
	db := database.GetDB()
	rank := startRank

	for _, result := range results {
		RankedUser := &model.User{}
		db.Preload("UserRegistration").Preload("Character").First(RankedUser, result.Member.(string))

		rowLeaderboard := LeaderboardRow{
			UserID:    RankedUser.ID,
			Username:  RankedUser.UserRegistration.Username,
			Trophies:  uint(result.Score),
			Rank:      rank,
			XP:        uint64(RankedUser.XP),
			AvatarURL: RankedUser.Character.AvatarURL,
		}

		leaderboard = append(leaderboard, rowLeaderboard)
		rank = rank + 1
	}

	return leaderboard, nil
}
