package helper

import (
	"errors"
	"fmt"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
)

var miniconsInLineup int
var matchmakingRankRange int
var minTrophyGain int
var trophyGainRange int
var trophyDiffLoser int
var survivorTrophyRange int
var defaultTrophyCount int
var successiveDuplicateMatchLimit int
var dailyAttackLimit int
var incrXpMinicon int
var incrXpUser int
var xpLevelMultiplier int
var xpBaseCount int
var levelUpStatMultiplierNumerator int
var levelUpStatMultiplierDenominator int
var isArenaOpen int
var maxMiniconLevel int
var maxUnlockedMinicons int
var typeMultiplierNumerator int
var typeMultiplierDenominator int

func InitConstants() (err error) {

	var log = utils.GetFunctionLogger("InitConstants")

	db := database.GetDB()

	var allConstants []model.Constant

	if err := db.Find(&allConstants).Error; err != nil {
		log.Error(err)
		return errors.New("Error Fetching all Constants")
	}

	for _, constant := range allConstants {
		switch constant.Name {
		case model.MiniconsInLineup:
			miniconsInLineup = constant.Value
		case model.MatchmakingRankRange:
			matchmakingRankRange = constant.Value
		case model.MinTrophyGain:
			minTrophyGain = constant.Value
		case model.TrophyGainRange:
			trophyGainRange = constant.Value
		case model.TrophyDiffLoser:
			trophyDiffLoser = constant.Value
		case model.SurvivorTrophyRange:
			survivorTrophyRange = constant.Value
		case model.DefaultTrophyCount:
			defaultTrophyCount = constant.Value
		case model.SuccessiveDuplicateMatchLimit:
			successiveDuplicateMatchLimit = constant.Value
		case model.DailyAttackLimit:
			dailyAttackLimit = constant.Value
		case model.IncrXpMinicon:
			incrXpMinicon = constant.Value
		case model.IncrXpUser:
			incrXpUser = constant.Value
		case model.XpLevelMultiplier:
			xpLevelMultiplier = constant.Value
		case model.XpBaseCount:
			xpBaseCount = constant.Value
		case model.LevelUpStatMultiplierNumerator:
			levelUpStatMultiplierNumerator = constant.Value
		case model.LevelUpStatMultiplierDenominator:
			levelUpStatMultiplierDenominator = constant.Value
		case model.IsArenaOpen:
			isArenaOpen = constant.Value
		case model.MaxMiniconLevel:
			maxMiniconLevel = constant.Value
		case model.MaxUnlockedMinicons:
			maxUnlockedMinicons = constant.Value
		case model.TypeMultiplierNumerator:
			typeMultiplierNumerator = constant.Value
		case model.TypeMultiplierDenominator:
			typeMultiplierDenominator = constant.Value
		default:
			// fmt.Print(color.RedString("Unassigned Constant %s : %d \n", constant.Name, constant.Value))
			log.Warn(fmt.Sprintf("Unassigned Constant %s : %d", constant.Name, constant.Value))
		}
	}

	log.Info("Constants Fetched!")
	return nil
}

func GetConstant(constName model.ConstantType) (constValue int, err error) {

	var log = utils.GetFunctionLogger("GetConstant")

	switch constName {

	case model.MiniconsInLineup:
		return miniconsInLineup, nil
	case model.MatchmakingRankRange:
		return matchmakingRankRange, nil
	case model.MinTrophyGain:
		return minTrophyGain, nil
	case model.TrophyGainRange:
		return trophyGainRange, nil
	case model.TrophyDiffLoser:
		return trophyDiffLoser, nil
	case model.SurvivorTrophyRange:
		return survivorTrophyRange, nil
	case model.DefaultTrophyCount:
		return defaultTrophyCount, nil
	case model.SuccessiveDuplicateMatchLimit:
		return successiveDuplicateMatchLimit, nil
	case model.DailyAttackLimit:
		return dailyAttackLimit, nil
	case model.IncrXpMinicon:
		return incrXpMinicon, nil
	case model.IncrXpUser:
		return incrXpUser, nil
	case model.XpLevelMultiplier:
		return xpLevelMultiplier, nil
	case model.XpBaseCount:
		return xpBaseCount, nil
	case model.LevelUpStatMultiplierNumerator:
		return levelUpStatMultiplierNumerator, nil
	case model.LevelUpStatMultiplierDenominator:
		return levelUpStatMultiplierDenominator, nil
	case model.IsArenaOpen:
		return isArenaOpen, nil
	case model.MaxMiniconLevel:
		return maxMiniconLevel, nil
	case model.MaxUnlockedMinicons:
		return maxUnlockedMinicons, nil
	case model.TypeMultiplierNumerator:
		return typeMultiplierNumerator, nil
	case model.TypeMultiplierDenominator:
		return typeMultiplierDenominator, nil
	}

	log.Error("Attempted to fetch invalid constant")
	errMessage := fmt.Sprintf("Attempted to fetch invalid constant %s", constName)
	return -1, errors.New(errMessage)
}
