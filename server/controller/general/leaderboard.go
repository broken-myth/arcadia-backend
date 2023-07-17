package controller

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/delta/arcadia-backend/database"
	generalHelper "github.com/delta/arcadia-backend/server/helper/general"
	userHelper "github.com/delta/arcadia-backend/server/helper/user"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/gin-gonic/gin"
)

type LeaderboardRow struct {
	UserID    uint   `json:"userId"`
	Username  string `json:"username"`
	Trophies  uint   `json:"trophies"`
	Rank      uint   `json:"rank"`
	XP        uint64 `json:"xp"`
	AvatarURL string `json:"avatarUrl"`
}

type GetLeaderboardResponse struct {
	Leaderboard []LeaderboardRow `json:"leaderboard"`
	Pages       int              `json:"pages"`
}

// To remove object in given index of a slice
func removeLeaderboardEntityFromSlice(slice []LeaderboardRow, index int) []LeaderboardRow {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

// FetchLeaderboard godoc
//
//	@Summary		Get Leaderboard
//	@Description	Get Leaderboard
//	@Tags			General
//	@Accept			json
//	@Param			page	path	uint	true	"Page number"
//	@Produce		json
//	@Success		200	{object}	controller.GetLeaderboardResponse	"Success"
//	@Failure		401	{object}	generalHelper.ErrorResponse			"Unauthorized"
//
//	@Failure		500	{object}	generalHelper.ErrorResponse			"Internal Error"
//
//	@Router			/api/leaderboard/{page} [get]
func FetchLeaderboardGET(c *gin.Context) {
	param := c.Param("page")

	log := utils.GetControllerLogger("/api/leaderboard/:page [GET]")

	page, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		generalHelper.SendError(c, http.StatusInternalServerError, "Unknown error. Please try again later")
		return
	}

	if page < 1 {
		page = 1
	}

	leaderboard, err := GetEntireLeaderboard()
	if err != nil {
		log.Error(err)
		generalHelper.SendError(c, http.StatusInternalServerError, "Unknown error. Please try again later")
		return
	}

	authHeader := c.Request.Header.Get("Authorization")

	if len(strings.Split(authHeader, " ")) == 1 {
		res := paginateLeaderboard(strings.Compare(param, "") != 0, page, 10, leaderboard)
		generalHelper.SendResponse(c, http.StatusOK, res)
		return
	}

	userID, err := userHelper.ValidateToken(authHeader)
	if err != nil {
		generalHelper.SendResponse(c, http.StatusUnauthorized, err)
		return
	}

	var userRankDetails []LeaderboardRow
	// if user is logged in

	userRank, _ := generalHelper.GetUserRank(userID)

	if userRank >= 1 {
		userRankDetails, _ = GetLeaderboardRange(userRank, 1)
	}

	// first value is user's leaderboard details (followed by the actual entire leaderboard)
	leaderboard = append(userRankDetails, leaderboard...)
	res := paginateLeaderboard(strings.Compare(param, "") != 0, page, 10, leaderboard)
	generalHelper.SendResponse(c, http.StatusOK, res)

}

func paginateLeaderboard(
	paginateWithPage bool,
	page int,
	windowSize int,
	leaderboard []LeaderboardRow) GetLeaderboardResponse {
	var sliceStartIndex int
	var sliceEndIndex int
	leaderboardSlice := leaderboard
	noOfPages := 1 // Total pages in the leaderboard (init with 1)

	if len(leaderboardSlice) <= windowSize {
		// If the total number of users is less than the window size
		// then we don't need to paginate
		return GetLeaderboardResponse{
			Leaderboard: leaderboardSlice,
			Pages:       noOfPages,
		}
	}

	// If the first entry in the array is not of rank 1 or first 2 entries are rank 1 it means a user is authenticated
	userIsAuthenticated := leaderboardSlice[0].Rank != 1 ||
		(leaderboardSlice[0].Rank == 1 && leaderboardSlice[1].Rank == 1)

	if userIsAuthenticated {
		// Remove the second entry for the authenticatedUser in the rest of the leaderboard
		leaderboardSlice = removeLeaderboardEntityFromSlice(leaderboardSlice, int(leaderboardSlice[0].Rank))

		if !paginateWithPage || page <= 1 {
			// Send 1,2,3 and the user's rank and then [windowsize] other users
			page = 1
			sliceStartIndex = 0
		} else {
			sliceStartIndex = windowSize + 4 + (page-2)*windowSize
		}

		sliceEndIndex = windowSize + 4 + (page-1)*windowSize

		// Handling total number of pages present
		if len(leaderboardSlice) > windowSize+4 {
			noOfPages += int(math.Ceil(float64(len(leaderboardSlice)-(windowSize+4)) / float64(windowSize)))
		}
	} else {
		if !paginateWithPage || page <= 1 {
			//Send the data for page 1 (ranks 1,2,3 and [windowSize] users after that)
			page = 1
			sliceStartIndex = 0
		} else {
			sliceStartIndex = windowSize + 3 + (page-2)*windowSize
		}
		sliceEndIndex = windowSize + 3 + (page-1)*windowSize
		// Handling total number of pages present
		if len(leaderboardSlice) > windowSize+3 {
			noOfPages += int(math.Ceil(float64(len(leaderboardSlice)-(windowSize+3)) / float64(windowSize)))
		}
	}

	// Placing the end index to the last element of the array in case the index goes out of range
	if sliceEndIndex >= len(leaderboardSlice) {
		sliceEndIndex = len(leaderboardSlice)
	}

	if len(leaderboardSlice) >= windowSize {
		leaderboardSlice = leaderboardSlice[sliceStartIndex:sliceEndIndex]
	}

	return GetLeaderboardResponse{
		Leaderboard: leaderboardSlice,
		Pages:       noOfPages,
	}
}

func GetEntireLeaderboard() (ranks []LeaderboardRow, err error) {
	var log = utils.GetFunctionLogger("GetEntireLeaderboard")

	leaderboard, _ := GetLeaderboardRange(1, 0)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return leaderboard, nil
}

func GetLeaderboardRange(startRank uint, count uint) (ranks []LeaderboardRow, err error) {
	// set count = 0 to get all entries after a certain rank

	var log = utils.GetFunctionLogger("GetLeaderboardRange")

	results, err := generalHelper.GetUsersWithRankInRange(startRank, count)
	if err != nil {
		log.Error(err)
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
			AvatarURL: RankedUser.Character.ImageURL,
		}

		leaderboard = append(leaderboard, rowLeaderboard)
		rank = rank + 1
	}

	return leaderboard, nil
}
