package helpers

import (
	"fmt"
	"math"

	"github.com/delta/arcadia-backend/utils"
	"github.com/sirupsen/logrus"
)

func CalculateTrophyGain(attRank uint, defRank uint, attSurvivors int, defSurvivors int) (
	attTrophy uint, defTrophy uint) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":         "CalculateTrophyGain",
		"param_att_rank": attRank,
		"param_def_rank": defRank,
		"param_att_surv": attSurvivors,
		"param_def_surv": defSurvivors,
	})

	if attSurvivors == defSurvivors {
		return 0, 0
	}

	miniconsInLineup, err := GetConstant("minicons_in_lineup")
	if err != nil {
		l.Error("Error fetching constant minicons_in_lineup")
		return 0, 0
	}

	rankRangeInt, err := GetConstant("matchmaking_rank_range")
	if err != nil {
		l.Error("Error fetching constant matchmaking_rank_range")
		return 0, 0
	}
	rankRange := float32(rankRangeInt)

	minTrophyGain, err := GetConstant("min_trophy_gain")
	if err != nil {
		l.Error("Error fetching constant min_trophy_gain")
		return 0, 0
	}

	trophyRangeInt, err := GetConstant("trophy_gain_range")
	if err != nil {
		l.Error("Error fetching constant trophy_gain_range")
		return 0, 0
	}
	trophyRange := float32(trophyRangeInt)

	trophyDiff, err := GetConstant("trophy_diff_loser")
	if err != nil {
		l.Error("Error fetching constant trophy_diff_loser")
		return 0, 0
	}

	survivorTrophyRange, err := GetConstant("survivor_trophy_range")
	if err != nil {
		l.Error("Error fetching constant survivor_trophy_range")
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
	if rankDiff >= -0.3*rankRange && rankDiff <= 0.3*rankRange {
		winnerTrophies = 0.5 * trophyRange
	} else if rankDiff >= -0.7*rankRange && rankDiff <= -0.4*rankRange {
		winnerTrophies = 0.7 * trophyRange
	} else if rankDiff <= -0.8*rankRange {
		winnerTrophies = 0.9 * trophyRange
	} else if rankDiff >= 0.4*rankRange && rankDiff <= 0.7*rankRange {
		winnerTrophies = 0.3 * trophyRange
	} else if rankDiff >= 0.8*rankRange {
		winnerTrophies = 0.1 * trophyRange
	}

	// Secondary Factor = proportional to number of surviving minicons
	survivorTrohpies := float32(math.Abs(float64(attSurvivors-defSurvivors))/
		float64(miniconsInLineup)) * float32(survivorTrophyRange)

	winnerTrophies += survivorTrohpies

	if winnerTrophies > trophyRange {
		winnerTrophies = trophyRange
	}
	winnerTrophies += float32(minTrophyGain)

	if attSurvivors > defSurvivors {
		l.Debug(fmt.Sprintf("att Won, att_Troph = %d, def_Trophies= %d",
			int(winnerTrophies), (int(winnerTrophies) - (trophyDiff))))

		return uint(winnerTrophies), (uint(winnerTrophies) - uint(trophyDiff))
	}

	l.Debug(fmt.Sprintf("def Won, att_Troph = %d, def_Trophies= %d",
		(int(winnerTrophies) - (trophyDiff)), int(winnerTrophies)))

	return (uint(winnerTrophies) - uint(trophyDiff)), uint(winnerTrophies)
}
