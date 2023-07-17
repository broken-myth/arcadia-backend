package helper

import (
	"fmt"
	"math"

	"github.com/delta/arcadia-backend/utils"
)

func CalculateTrophyGain(attRank uint, defRank uint, attSurvivors int, defSurvivors int) (
	attTrophy int, defTrophy int) {

	params := map[string]interface{}{
		"attRank":      attRank,
		"defRank":      defRank,
		"attSurvivors": attSurvivors,
		"defSurvivors": defSurvivors,
	}

	var log = utils.GetFunctionLoggerWithFields("CalculateTrophyGain", params)

	if attSurvivors == defSurvivors {
		return 0, 0
	}

	miniconsInLineup, err := GetConstant("minicons_in_lineup")
	if err != nil {
		log.Error("Error fetching constant minicons_in_lineup")
		return 0, 0
	}

	rankRangeInt, err := GetConstant("matchmaking_rank_range")
	if err != nil {
		log.Error("Error fetching constant matchmaking_rank_range")
		return 0, 0
	}
	rankRange := float32(rankRangeInt)

	minTrophyGain, err := GetConstant("min_trophy_gain")
	if err != nil {
		log.Error("Error fetching constant min_trophy_gain")
		return 0, 0
	}

	trophyRangeInt, err := GetConstant("trophy_gain_range")
	if err != nil {
		log.Error("Error fetching constant trophy_gain_range")
		return 0, 0
	}
	trophyRange := float32(trophyRangeInt)

	trophyDiff, err := GetConstant("trophy_diff_loser")
	if err != nil {
		log.Error("Error fetching constant trophy_diff_loser")
		return 0, 0
	}

	survivorTrophyRange, err := GetConstant("survivor_trophy_range")
	if err != nil {
		log.Error("Error fetching constant survivor_trophy_range")
		return 0, 0
	}

	// Determine Winners' Trophies
	winnerTrophies := float32(0)
	rankDiff := float32(0)

	if attSurvivors > defSurvivors {
		rankDiff = float32(int(attRank) - int(defRank))
	} else {
		rankDiff = float32(int(defRank) - int(attRank))
	}

	// Main factor = Difference in Ranks
	winnerTrophies = calculateRankDiffTrophy(rankDiff, rankRange, trophyRange)

	// Secondary Factor = proportional to number of surviving minicons
	survivorTrohpies := float32(math.Abs(float64(attSurvivors-defSurvivors))/
		float64(miniconsInLineup)) * float32(survivorTrophyRange)

	winnerTrophies += survivorTrohpies

	if winnerTrophies > trophyRange {
		winnerTrophies = trophyRange
	}
	winnerTrophies += float32(minTrophyGain)

	if attSurvivors > defSurvivors {
		log.Debug(fmt.Sprintf("att Won, att_Troph = %d, def_Trophies= %d",
			int(winnerTrophies), -(int(winnerTrophies) - (trophyDiff))))

		return int(winnerTrophies), -(int(winnerTrophies) - trophyDiff)
	}

	log.Debug(fmt.Sprintf("def Won, att_Troph = %d, def_Trophies= %d",
		-(int(winnerTrophies) - trophyDiff), int(winnerTrophies)))

	return -(int(winnerTrophies) - trophyDiff), int(winnerTrophies)
}

func calculateRankDiffTrophy(rankdiff float32, rankRange float32, trophyRange float32) (initTrophyGain float32) {

	if rankdiff >= -0.3*rankRange && rankdiff <= 0.3*rankRange {
		return 0.5 * trophyRange
	} else if rankdiff >= -0.7*rankRange && rankdiff <= -0.4*rankRange {
		return 0.7 * trophyRange
	} else if rankdiff <= -0.8*rankRange {
		return 0.9 * trophyRange
	} else if rankdiff >= 0.4*rankRange && rankdiff <= 0.7*rankRange {
		return 0.3 * trophyRange
	} else if rankdiff >= 0.8*rankRange {
		return 0.1 * trophyRange
	}

	return 0
}
