package simulator

import (
	"time"

	functions "github.com/delta/arcadia-backend/simulator/functions"
	"github.com/delta/arcadia-backend/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GameStruct struct {
	tx                *gorm.DB
	log               *logrus.Entry
	MatchID           uint
	RoundNumber       int
	AttackerTeam      functions.Team
	DefenderTeam      functions.Team
	AttackerPriority  bool
	AttackerSurvivors uint
	DefenderSurvivors uint
	BattleResult      []functions.Round
	DefenderTeamClone functions.Team // Used for testing
	AttackerTeamClone functions.Team // Used for testing
	Outcome           string         // Used for testing
}

type GamePhase string

const (
	PRELOAD GamePhase = "PRELOAD"
	UPDATE  GamePhase = "UPDATE"
	STOP    GamePhase = "STOP"
	END     GamePhase = "END"
)

type Game interface {

	/* Methods for the game */

	// Preload() is called when the simulation starts, and is used to initialize any data that is needed for the simulation
	Preload() error

	// Update() is a recursive function that is called every turn of the simulation
	Update() error

	// End() is called when the simulation ends, and is used to clean up any data and send the results to the app
	End() error

	/* Utility methods */

	// Stop() is called when the simulation encounters an error
	Stop(err error, phase GamePhase) error
}

func NewGame(tx *gorm.DB, MatchID uint) *GameStruct {

	params := map[string]interface{}{
		"matchID": MatchID,
	}
	var log = utils.GetFunctionLoggerWithFields("Simulator", params)

	return &GameStruct{
		tx:      tx,
		MatchID: MatchID,
		log:     log,
	}
}

func Start(tx *gorm.DB, MatchID uint) (AttackerSurvivors uint, DefenderSurvivors uint, err error) {
	start := time.Now()
	game := NewGame(tx, MatchID)

	if err := game.Preload(); err != nil {
		return 0, 0, game.Stop(err, PRELOAD, MatchID)
	}

	if err := game.Update(); err != nil {
		return 0, 0, game.Stop(err, UPDATE, MatchID)
	}

	if err := game.End(); err != nil {
		return 0, 0, game.Stop(err, END, MatchID)
	}

	AttackerSurvivors = game.AttackerSurvivors
	DefenderSurvivors = game.DefenderSurvivors
	elapsed := time.Since(start)

	var log = utils.GetLogger()
	log.Debug("MatchID: ", MatchID, " Time: ", elapsed)
	log.Debug("Game ended")
	return AttackerSurvivors, DefenderSurvivors, nil
}
