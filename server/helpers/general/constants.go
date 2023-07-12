package helpers

import (
	"errors"
	"fmt"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var miniconsInLineup int
var matchmakingRankRange int
var minTrophyGain int
var trophyGainRange int
var trophyDiffLoser int
var survivorTrophyRange int
var defaultTrophyCount int

func InitConstants() (err error) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "InitConstants",
	})

	db := database.GetDB()

	var allConstants []model.Constants

	if err := db.Find(&allConstants).Error; err != nil {
		l.Error(err)
		return errors.New("Error Fetching Constants")
	}

	for _, constant := range allConstants {
		switch constant.Key {
		case "minicons_in_lineup":
			miniconsInLineup = constant.Value
		case "matchmaking_rank_range":
			matchmakingRankRange = constant.Value
		case "min_trophy_gain":
			minTrophyGain = constant.Value
		case "trophy_gain_range":
			trophyGainRange = constant.Value
		case "trophy_diff_loser":
			trophyDiffLoser = constant.Value
		case "survivor_trophy_range":
			survivorTrophyRange = constant.Value
		case "default_trophy_count":
			defaultTrophyCount = constant.Value
		default:
			fmt.Print(color.RedString("Unassigned Constant %s : %d \n", constant.Key, constant.Value))
			l.Warn(fmt.Sprintf("Unassigned Constant %s : %d", constant.Key, constant.Value))
		}
	}

	l.Info("Constants Fetched!")
	return nil
}

func GetConstant(constKey string) (constValue int, err error) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":          "GetConstant",
		"param_const_key": constKey,
	})

	switch constKey {

	case "minicons_in_lineup":
		return miniconsInLineup, nil
	case "matchmaking_rank_range":
		return matchmakingRankRange, nil
	case "min_trophy_gain":
		return minTrophyGain, nil
	case "trophy_gain_range":
		return trophyGainRange, nil
	case "trophy_diff_loser":
		return trophyDiffLoser, nil
	case "survivor_trophy_range":
		return survivorTrophyRange, nil
	case "default_trophy_count":
		return defaultTrophyCount, nil
	}

	l.Error("Attempted to fetch invalid constant")
	errMessage := fmt.Sprintf("Attempted to fetch invalid constant %s", constKey)
	return -1, errors.New(errMessage)
}
