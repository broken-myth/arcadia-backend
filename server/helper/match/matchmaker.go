package helper

import (
	"errors"
	"math/rand"
	"time"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
)

func shuffleArray(arr *[]uint) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(*arr); n > 0; n-- {
		randIndex := r.Intn(n)
		(*arr)[n-1], (*arr)[randIndex] = (*arr)[randIndex], (*arr)[n-1]
	}
}

func checkOpponent(opponentsPlayed []uint, defenderID uint) bool {
	for _, opponent := range opponentsPlayed {
		if opponent == defenderID {
			return false
		}
	}

	return true
}

func MatchMaker(attackerID uint) (uint, error) {
	var log = utils.GetFunctionLogger("MatchMaker")

	defenderID := uint(0)

	// get an array of nearby userIDs on leaderboard using redis
	opponentUserIDs, err := helper.FindSuitors(attackerID)
	if err != nil {
		if err == helper.ErrNoLineup {
			return 0, err
		}
		log.Error("Error while calling FindSuitors. Error:", err)
		return defenderID, err
	}

	// shuffles this array
	shuffleArray(&opponentUserIDs)

	successiveDuplicateMatchLimit, err := helper.GetConstant("successive_duplicate_match_limit")
	if err != nil {
		log.Error("Error fetching constant successive_duplicate_match_limit")
		return defenderID, err
	}

	// fetching last played match with
	db := database.GetDB()
	opponentsPlayed := []model.MatchmakingDetails{}

	if err := db.Order("created_at desc").Where("attacker_id = ?", attackerID).Or("defender_id = ?", attackerID).
		Limit(successiveDuplicateMatchLimit).Find(&opponentsPlayed).Error; err != nil {
		log.Error("Error fetching last few matches. Error:", err)
		return defenderID, err
	}

	var opponentsPlayedID []uint
	for _, opponent := range opponentsPlayed {
		var id uint

		if opponent.AttackerID == attackerID {
			id = opponent.DefenderID
		} else {
			id = opponent.AttackerID
		}

		opponentsPlayedID = append(opponentsPlayedID, id)
	}

	// check which opponent the current user haven't played with for last x matches
	for _, opponentID := range opponentUserIDs {
		if checkOpponent(opponentsPlayedID, opponentID) {
			defenderID = opponentID
			break
		}
	}

	if defenderID == 0 {
		errMessage := "Unable to find an opponent"
		log.Error(errMessage)
		return defenderID, errors.New(errMessage)
	}

	return defenderID, nil
}
