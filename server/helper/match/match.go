package helper

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"

	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
)

func CreateMatchmakingDetails(tx *gorm.DB, attackerID uint, defenderID uint) (uint, error) {
	//Fetching Player Lineups
	AttackerLineupJSON, err := helper.ReturnPlayerMiniconLineup(tx, attackerID)
	if err != nil {
		return 0, err
	}
	DefenderLineupJSON, err := helper.ReturnPlayerMiniconLineup(tx, defenderID)
	if err != nil {
		return 0, err
	}

	//Making MatchmakingDetails object in DB
	matchUp := model.MatchmakingDetails{
		AttackerID:     attackerID,
		DefenderID:     defenderID,
		AttackerLineUp: datatypes.JSON(AttackerLineupJSON),
		DefenderLineUp: datatypes.JSON(DefenderLineupJSON),
	}

	if err := tx.Create(&matchUp).Error; err != nil {
		return 0, err
	}

	return matchUp.ID, nil
}

func CreateBattleResult(tx *gorm.DB, matchID uint, attacker model.User, defender model.User,
	attSurvivors uint, defSurvivors uint) error {

	var log = utils.GetFunctionLogger("CreateBattleResult")

	// get from redis leaderboard
	attRank, err := helper.GetUserRank(attacker.ID)
	if err != nil {
		log.Error("Error while getting attacker rank. Error:", err)
		return err
	}
	defRank, err := helper.GetUserRank(defender.ID)
	if err != nil {
		log.Error("Error while getting defender rank. Error:", err)
		return err
	}

	// calculate trophies
	attackerTrophy, defenderTrophy := helper.CalculateTrophyGain(attRank, defRank,
		int(attSurvivors), int(defSurvivors))

	// update trophies
	result, winnerTrophies, loserTrophies := EvaluteBattleResult(attRank, defRank, attSurvivors, defSurvivors)

	attacker.Trophies += attackerTrophy
	if attacker.Trophies < 0 {
		attacker.Trophies = 0
	}
	defender.Trophies += defenderTrophy
	if defender.Trophies < 0 {
		defender.Trophies = 0
	}

	battleResult := model.BattleResult{
		MatchID:        matchID,
		Result:         result,
		WinnerTrophies: winnerTrophies,
		LoserTrophies:  loserTrophies,
	}

	if err := tx.Create(&battleResult).Error; err != nil {
		log.Error("Error creating battle result: ", err)
		return err
	}

	if err = tx.Save(&attacker).Error; err != nil {
		log.Error("Error saving attacker: ", err)
		return err
	}
	if err = tx.Save(&defender).Error; err != nil {
		log.Error("Error saving defender: ", err)
		return err
	}

	// Redis updates - so not an issue even if it fails
	err = helper.UpdateUserTrophies(attacker.ID, attacker.Trophies)
	if err != nil {
		log.Error(err)
	}
	err = helper.UpdateUserTrophies(defender.ID, defender.Trophies)
	if err != nil {
		log.Error(err)
	}

	return nil
}

// EvaluteBattleResult returns the result of the match, and the trophies gained by the winner and loser
func EvaluteBattleResult(attRank uint, defRank uint, attSurvivors uint, defSurvivors uint) (int, int, int) {

	var result int
	var winnerTrophies int
	var loserTrophies int

	attackerTrophy, defenderTrophy := helper.CalculateTrophyGain(attRank, defRank,
		int(attSurvivors), int(defSurvivors))

	if attSurvivors > defSurvivors {
		result = 1
		winnerTrophies = attackerTrophy
		loserTrophies = defenderTrophy
		return result, winnerTrophies, loserTrophies
	} else if attSurvivors < defSurvivors {
		result = -1
		winnerTrophies = defenderTrophy
		loserTrophies = attackerTrophy
		return result, winnerTrophies, loserTrophies
	} else {
		result = 0
		winnerTrophies = attackerTrophy
		loserTrophies = defenderTrophy
		return result, winnerTrophies, loserTrophies
	}
}
