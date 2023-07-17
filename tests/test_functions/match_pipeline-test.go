package tests

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/delta/arcadia-backend/database"
	matchHelpers "github.com/delta/arcadia-backend/server/helper/match"
	"github.com/delta/arcadia-backend/server/model"
	simulator "github.com/delta/arcadia-backend/simulator"
)

func TestMatchPipeline() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	attackerID := uint(r.Intn(49))
	defenderID, err := matchHelpers.MatchMaker(attackerID)
	fmt.Println("Attacker ID: ", attackerID, " Defender ID: ", defenderID)
	if err != nil {
		panic(err)
	}
	// var attackerID uint = 23
	// var defenderID uint = 29
	tx := database.GetDB().Begin()
	err = simulator.Init()
	if err != nil {
		panic(err)
	}
	matchID, err := matchHelpers.CreateMatchmakingDetails(tx, (attackerID), (defenderID))
	if err != nil {
		panic(err)
	}

	attSurvivors, defSurvivors, err := simulator.Start(tx, matchID)
	if err != nil {
		fmt.Println("Error in Simulator")
		tx.Rollback()
		return
	}
	var attacker model.User
	if err := tx.First(&attacker, attackerID).Error; err != nil {
		fmt.Println("Error in fetching Attacker User Details")
		tx.Rollback()
		return
	}
	var defender model.User
	if err := tx.Preload("UserRegistration").Where("id = ?", defenderID).First(&defender).Error; err != nil {
		fmt.Println("Error in fetching Defender User Details")
		tx.Rollback()
		return
	}

	err = matchHelpers.CreateBattleResult(tx, matchID, attacker, defender,
		attSurvivors, defSurvivors)
	if err != nil {
		fmt.Println("Error in creating Battle Result")
		tx.Rollback()
		return
	}
	fmt.Print("\n\n Attacker Survivors: \t", attSurvivors, "\tDefender Survivors:", defSurvivors)
	tx.Rollback()
	fmt.Println("\tGame done")
}
