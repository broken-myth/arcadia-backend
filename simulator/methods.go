package simulator

import (
	"encoding/json"
	"fmt"

	"github.com/delta/arcadia-backend/server/model"
	functions "github.com/delta/arcadia-backend/simulator/functions"
	helper "github.com/delta/arcadia-backend/simulator/functions"
	"github.com/delta/arcadia-backend/utils"
	"gorm.io/datatypes"
)

// // Logging stuff (Uncomment when testing)

// type LogRound struct {
// 	round    int
// 	Attacker []LogMinicon
// 	Defender []LogMinicon
// }
// type LogMinicon struct {
// 	Number int
// 	Health int
// 	Type   string
// 	Perks  map[string]helper.Perk
// 	Attack int
// }

// //  printGameLog() : Used to print the details of a lineup at the end of the round (uncomment when testing)
// func printGameLog(g *GameStruct) {
// 	LogOfCurrentRound := LogRound{}
// 	LogOfCurrentRound.round = g.RoundNumber
// 	for index, minicon := range g.AttackerTeam.MiniconTeam {
// 		LogOfCurrentRound.Attacker = append(LogOfCurrentRound.Attacker,
// 			LogMinicon{index + 1, int(minicon.Health), minicon.Type, minicon.Perk, int(minicon.Attack)})
// 	}
// 	for index, minicon := range g.DefenderTeam.MiniconTeam {
// 		LogOfCurrentRound.Defender = append(LogOfCurrentRound.Defender,
// 			LogMinicon{index + 1, int(minicon.Health), minicon.Type, minicon.Perk, int(minicon.Attack)})
// 	}
// 	LogOfCurrentRoundJSON, errJSON := json.Marshal(LogOfCurrentRound)

// 	if errJSON != nil {
// 		panic("error in AttackerLineup")
// 	}
// 	fmt.Println(" \t \t \t \t \t \t \t \t The Current Round is :  ", g.RoundNumber)
// 	fmt.Println(string(LogOfCurrentRoundJSON))

// 	//Attacker Stats:
// 	fmt.Println(color.RedString("Attacker Stats : "))
// 	for _, minicon := range LogOfCurrentRound.Attacker {
// 		fmt.Println("Index : ", minicon.Number)
// 		fmt.Println("Health : ", minicon.Health)
// 		fmt.Println("Attack : ", minicon.Attack)
// 		fmt.Println("Type : ", minicon.Type)
// 		fmt.Println("Perks ")
// 		for k, v := range minicon.Perks {
// 			fmt.Println(k, " : ", v)
// 		}
// 		fmt.Println()

// 	}

// 	// Defender Lineup
// 	fmt.Println(color.GreenString("Defender Stats : "))
// 	for _, minicon := range LogOfCurrentRound.Defender {
// 		fmt.Println("Index : ", minicon.Number)
// 		fmt.Println("Health : ", minicon.Health)
// 		fmt.Println("Attack : ", minicon.Attack)
// 		fmt.Println("Type : ", minicon.Type)
// 		fmt.Println("Perks ")
// 		for k, v := range minicon.Perks {
// 			fmt.Println(k, " : ", v)
// 		}
// 		fmt.Println()
// 	}
// }

// Functions for Preload():

type MiniconStats struct {
	PerkValue  uint
	PerkID     uint
	Health     uint
	Attack     uint
	PositionNo uint
	MiniconID  uint
}

type LoadDBLineup struct {
	Lineup []MiniconStats
}

func (Result *LoadDBLineup) ReturnSizeOfPlayerLineup() (SizeOfLineup uint) {
	trackOfIndices := make(map[uint]int)
	for _, minicon := range Result.Lineup {
		_, ok := trackOfIndices[minicon.PositionNo]
		if !ok {
			SizeOfLineup++
			trackOfIndices[minicon.PositionNo] = 1
		}
	}
	return SizeOfLineup
}

// // A premade minicon lineup : Used for testing minicons (Uncomment when testing)

// func premadeLineup(g *GameStruct) {
// 	g.AttackerTeam = helper.Team{MiniconTeam: []helper.Minicon{TrialMiniconStrongAll, TrialMinicon, TrialMinicon,
// 		TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon}}
// 	g.DefenderTeam = helper.Team{MiniconTeam: []helper.Minicon{TrialMinicon, TrialMinicon, TrialMinicon,
// 		TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon, TrialMinicon}}
// }

// Functions for End():

func storeSimulatorLogInDB(Result []byte, g *GameStruct) error {
	SimulatorLog := model.SimulationDetail{
		AttackerSurvivors: g.AttackerSurvivors, // For Debuging Purposes
		DefenderSurvivors: g.DefenderSurvivors, // For Debugging Purposes
		SimulationLog:     datatypes.JSON(Result),
		MatchID:           g.MatchID,
	}
	if err := g.tx.Create(&SimulatorLog).Error; err != nil {
		return err
	}
	return nil
}

// Functions for Preload():

// // Making Clones of the player teams as the initial line up objects (Uncomment when testing)
// func makeCopyOfMiniconTeams(g *GameStruct) {
// 	AttackerClone := helper.MakeMiniconTeam(g.AttackerTeam.MiniconTeam)
// 	DefenderClone := helper.MakeMiniconTeam(g.DefenderTeam.MiniconTeam)
// 	g.AttackerTeamClone = AttackerClone
// 	g.DefenderTeamClone = DefenderClone
// }

// Generate the game object
func initialiseBattle(g *GameStruct) {
	// makeCopyOfMiniconTeams(g) // Uncomment when testing
	g.BattleResult = make([]functions.Round, 0)
}

func determinePriority(MatchmakingDetails model.MatchmakingDetails, g *GameStruct) {
	g.AttackerPriority = false
	if MatchmakingDetails.Attacker.Trophies <= MatchmakingDetails.Defender.Trophies {
		g.AttackerPriority = true
	}
}

func fetchLineUpDataAndGenerateTeams(PlayerMiniconTeam *functions.Team, PlayerID uint, g *GameStruct) error {
	//Fetching Owned_Minicons and Owned_Perks from lineups  database
	var LineupResult LoadDBLineup
	err := g.tx.Model(&model.Lineup{}).
		Select("lineups.position_no as position_no",
			"owned_perks.perk_value as perk_value, owned_perks.perk_id as perk_id",
			"owned_minicons.health as health , owned_minicons.attack as attack, owned_minicons.minicon_id as minicon_id",
		).
		Joins("JOIN owned_minicons on owned_minicons.id =lineups.owned_minicon_id").
		Joins("JOIN owned_perks on owned_minicons.id = owned_perks.owned_minicon_id").
		Where("lineups.creator_id = ?", PlayerID).
		Find(&LineupResult.Lineup).Error

	// Obtaining Size of Player Lineup
	SizeOfLineup := LineupResult.ReturnSizeOfPlayerLineup()
	PlayerMiniconTeam.MiniconTeam = make([]helper.Minicon, SizeOfLineup)

	// To keep track of minicons that have been added to lineup
	// We implement a map to keep track of minicons via their position number
	trackOfMiniconIndices := make(map[uint]int)
	for _, minicon := range LineupResult.Lineup {
		_, ok := trackOfMiniconIndices[minicon.PositionNo] // Checks if minicon is in lineup
		PerkTrigger := PerkIDMap[minicon.PerkID].PerkTrigger
		OwnedPerkBase := PerkIDMap[minicon.PerkID]

		//Case 1: When the minicon has not been visited before in the lineup
		if !ok {
			trackOfMiniconIndices[minicon.PositionNo] = 1 // We update map to include minicon in lineup
			// Generating Minicon Object
			DBOwnedMinicon := helper.GenerateMiniconObjectFromDB(MiniconIDMap[minicon.MiniconID].Name,
				minicon.Health, minicon.Attack, MiniconIDMap[minicon.MiniconID].Type)
			//Updating Perk Map if minicon is present
			DBOwnedMinicon.Perk[PerkTrigger] = helper.GeneratePerkFromDB(minicon.PerkValue, PerkTrigger,
				OwnedPerkBase.Effect, OwnedPerkBase.Target)
			//Inserting Minicon in it's correct lineup position
			PlayerMiniconTeam.MiniconTeam[minicon.PositionNo-1] = DBOwnedMinicon
		} else {

			// Case 2: When minicon has been traversed before
			// Updation of perk map to include the new perk
			PlayerMiniconTeam.MiniconTeam[minicon.PositionNo-1].Perk[PerkTrigger] = helper.
				GeneratePerkFromDB(minicon.PerkValue, PerkTrigger, OwnedPerkBase.Effect, OwnedPerkBase.Target)
		}

	}

	return err
}

func initialiseTeam(g *GameStruct) error {
	var MatchmakingDetails model.MatchmakingDetails
	err := g.tx.Preload("Defender").Preload("Attacker").First(&MatchmakingDetails, g.MatchID).Error
	if err != nil {
		return err
	}
	//Priority
	determinePriority(MatchmakingDetails, g)
	// Attacker Lineup
	err = fetchLineUpDataAndGenerateTeams(&g.AttackerTeam, MatchmakingDetails.AttackerID, g)
	if err != nil {
		return err
	}
	// Defender Lineup
	err = fetchLineUpDataAndGenerateTeams(&g.DefenderTeam, MatchmakingDetails.DefenderID, g)
	if err != nil {
		return err
	}

	var log = utils.GetFunctionLogger("(Simulator): initialiseTeam")
	log.Debug("Attacker ID: ", MatchmakingDetails.AttackerID, " Defender ID: ", MatchmakingDetails.DefenderID)
	return nil
}

func (g *GameStruct) Preload() error {
	err := initialiseTeam(g)
	if err != nil {
		return err
	}
	initialiseBattle(g)
	return nil
}

// Functions for Update():

// // Used To Determine the outcome of the game (Uncomment when testing)
// func determineOutcome(g *GameStruct) {
// 	if g.AttackerTeam.ReturnNumberOfSurvivors() > g.DefenderTeam.ReturnNumberOfSurvivors() {
// 		g.Outcome = "Attacker Wins"
// 	} else if g.AttackerTeam.ReturnNumberOfSurvivors() > g.DefenderTeam.ReturnNumberOfSurvivors() {
// 		g.Outcome = "Defender Wins"
// 	} else {
// 		g.Outcome = "Draw"
// 	}
// }

func (g *GameStruct) Update() error {

	if g.RoundNumber > 100 {
		g.AttackerSurvivors = g.AttackerTeam.ReturnNumberOfSurvivors()
		g.DefenderSurvivors = g.DefenderTeam.ReturnNumberOfSurvivors()

		// // Uncomment for testing purposes
		// determineOutcome(g)
		// return nil
	}
	//Recurse until one or both the minicon teams are found dead
	if g.AttackerTeam.IsAlive() && g.DefenderTeam.IsAlive() {
		//Run And Get Details of each round
		BattleRound := helper.ExecuteRound(&g.AttackerTeam, &g.DefenderTeam, &g.AttackerPriority)
		//Store the result of each round
		g.BattleResult = append(g.BattleResult, BattleRound)

		// // Uncomment for testing purposes
		// printGameLog(g)

		//Increment Round Number
		g.RoundNumber++

		err := g.Update()
		if err != nil {
			return err
		}
	} else {
		// Determining the outcome of the Game
		// determineOutcome(g)
		g.AttackerSurvivors = g.AttackerTeam.ReturnNumberOfSurvivors()
		g.DefenderSurvivors = g.DefenderTeam.ReturnNumberOfSurvivors()

		var log = utils.GetLogger()
		log.Debug("Total Rounds: ", g.RoundNumber)
		log.Debug("Game Result ", g.Outcome,
			" Attacker: ", g.AttackerSurvivors, " Defender Survivors: ", g.DefenderSurvivors)
	}
	return nil
}

func (g *GameStruct) End() error {
	JSONGameResult, err := json.Marshal(g.BattleResult)
	if err != nil {
		return err
	}
	err = storeSimulatorLogInDB(JSONGameResult, g)
	if err != nil {
		return err
	}
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		var log = utils.GetFunctionLogger("(Simulator): End)")
		log.Debug("EndGame Attacker Lineup", g.AttackerTeam)
		log.Debug("EndGame Defender Lineup", g.DefenderTeam)
		log.Debug("JSON Game Result: ", string(JSONGameResult))
	}

	// // Uncomment out the lines of code for testing purposes

	// AttackerTeam, err := json.Marshal(g.AttackerTeamClone)
	// if err != nil {
	// 	return err
	// }
	// DefenderTeam, err := json.Marshal(g.DefenderTeamClone)
	// if err != nil {
	// 	return err
	// }
	// BattleResult, err := json.Marshal(g.BattleResult)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("Initial Attacker Team ", string(AttackerTeam))
	// fmt.Println("Initial Defender Team ", string(DefenderTeam))
	// fmt.Println("Battle Log : ", string(BattleResult))
	return nil
}

func (g *GameStruct) Stop(err error, phase GamePhase, MatchID uint) error {
	g.log.Println("Stopping game (matchID = ", MatchID, ") due to error in phase-", phase, " : ", err)
	g.tx.Rollback()
	return err
}
