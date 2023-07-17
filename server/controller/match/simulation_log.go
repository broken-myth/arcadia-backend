package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	model "github.com/delta/arcadia-backend/server/model"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Result of each individual round
type SimulationLog struct {
	Headers         []string
	SimulationLogDB datatypes.JSON
}

// called at the start of match to initialise the headers of BattleResult Object
func MakeSimulationLog(SimulationDetails datatypes.JSON) SimulationLog {
	BattleResultObject := SimulationLog{}
	BattleResultObject.Headers = append(BattleResultObject.Headers,
		"Event",
		"Type",
		"Trigger",
		"Sender",
		"Receiver",
		"DeltaValue")
	BattleResultObject.SimulationLogDB = SimulationDetails
	return BattleResultObject
}

type Debug struct {
	AttackerSurvivors uint
	DefenderSurvivors uint
}

type SimulationLogResponse struct {
	AttackerLineUp  datatypes.JSON
	DefenderLineUp  datatypes.JSON
	SimulationLog   SimulationLog
	Debug           Debug
	WinnerTrophies  int
	LoserTrophies   int
	Result          int
	OpponentName    string
	OpponentCharURL string
	IsAttacker      bool
}

func ViewSimulationLogGET(c *gin.Context) {
	var MatchmakingDetails model.MatchmakingDetails
	var SimulationLogDetails model.SimulationDetail
	var BattleResult model.BattleResult
	userID := c.GetUint("userID")
	db := database.GetDB()
	matchParamID := c.Param("id")

	matchID, err := strconv.ParseUint(matchParamID, 10, 32)
	if err != nil {
		helper.SendError(c, http.StatusBadRequest, "Invalid Match Id")
		return
	}
	// Fetching Matchmaking Details from DB
	err = db.Select("attacker_line_up,defender_line_up,attacker_id,defender_id").
		Where("id = ?", matchID).First(&MatchmakingDetails).Error

	if err != nil {
		// When Match does not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendError(c, http.StatusNotFound, "Match does not exist")
			return
		}
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	// When the user accessing the match is neither the attacker or defender
	if MatchmakingDetails.AttackerID != userID && MatchmakingDetails.DefenderID != userID {
		helper.SendError(c, http.StatusForbidden, "Not Authorized to view match")
		return
	}

	// Fetching Simulation Log from DB
	err = db.Select("simulation_log , attacker_survivors , defender_survivors").Where(" match_id = ? ", matchID).
		First(&SimulationLogDetails).Error

	if err != nil {
		// When Log does not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendError(c, http.StatusNotFound, "Simulation does not exist")
			return
		}
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	// Fetching Battle Result from DB
	err = db.Select("result, winner_trophies, loser_trophies").Where(" match_id = ? ", matchID).
		First(&BattleResult).Error

	if err != nil {
		// When Log does not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendError(c, http.StatusNotFound, "Battle Result does not exist")
			return
		}
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	// Fetching Opponent Details from DB
	var opponent model.User
	var opponentID uint
	isAttacker := true
	result := BattleResult.Result
	if MatchmakingDetails.AttackerID == userID {
		opponentID = MatchmakingDetails.DefenderID
	} else {
		opponentID = MatchmakingDetails.AttackerID
		result *= -1
		isAttacker = false
	}

	err = db.Preload("Character").First((&opponent), opponentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.SendError(c, http.StatusNotFound, "Opponent does not exist")
			return
		}
		helper.SendError(c, http.StatusInternalServerError, "Unknown error, Try again")
		return
	}

	// Debug Information
	DebugInfo := Debug{SimulationLogDetails.AttackerSurvivors, SimulationLogDetails.DefenderSurvivors}
	SimulationLog := MakeSimulationLog(SimulationLogDetails.SimulationLog)
	//Response being sent to frontend
	response := SimulationLogResponse{
		AttackerLineUp:  (MatchmakingDetails.AttackerLineUp),
		DefenderLineUp:  (MatchmakingDetails.DefenderLineUp),
		SimulationLog:   SimulationLog,
		Debug:           DebugInfo,
		WinnerTrophies:  BattleResult.WinnerTrophies,
		LoserTrophies:   BattleResult.LoserTrophies,
		Result:          result,
		OpponentName:    opponent.Username,
		OpponentCharURL: opponent.Character.ImageURL,
		IsAttacker:      isAttacker,
	}

	helper.SendResponse(c, http.StatusOK, response)
}
