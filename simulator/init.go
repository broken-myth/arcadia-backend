package simulator

import (
	"github.com/delta/arcadia-backend/database"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/delta/arcadia-backend/server/model"
	functions "github.com/delta/arcadia-backend/simulator/functions"
	"github.com/delta/arcadia-backend/utils"
)

type MiniconDetails struct {
	Type string
	Name string
}

type PerkDetails struct {
	PerkTrigger string
	Effect      string
	Target      []int
}

type PerkDBResponse struct {
	ID          uint
	TargetVal   int
	PerkTrigger string
	Effect      string
}

// MiniconIDMap maps the minicon ID to the minicon details required by the simulator
var MiniconIDMap map[uint]MiniconDetails

// PerkIDMap maps the perk ID to the perk details required by the simulator
var PerkIDMap map[uint]PerkDetails

// Fetching Data from the DB for generating the minicon details required by the simulator
func generateMiniconMap() error {

	var log = utils.GetFunctionLogger("generateMiniconMap")

	db := database.GetDB()
	var DBMinicons []model.Minicon
	MiniconIDMap = make(map[uint]MiniconDetails)

	err := db.Preload("Type").Find(&DBMinicons).Error
	if err != nil {
		log.Error(err)
		return err
	}

	for _, minicon := range DBMinicons {
		MiniconIDMap[minicon.ID] = MiniconDetails{minicon.Type.Name, minicon.Name}
	}

	return nil
}

// Fetching Data from the DB for generating the perk details required by the simulator
func generatePerkMap() error {

	var log = utils.GetFunctionLogger("generateMiniconMap")

	db := database.GetDB()
	var DBPerks []PerkDBResponse
	PerkIDMap = make(map[uint]PerkDetails)

	err := db.Model([]model.Perk{}).
		Select("targets.target_val as target_val",
			"perks.perk_trigger as perk_trigger, perks.effect as effect, perks.id as id",
		).
		Joins("JOIN targets on targets.perk_id = perks.id").
		Find(&DBPerks).Error

	if err != nil {
		log.Error(err)
		return err
	}

	for _, perk := range DBPerks {
		copy, ok := PerkIDMap[perk.ID]
		if !ok {
			PerkIDMap[perk.ID] = PerkDetails{perk.PerkTrigger, perk.Effect, []int{perk.TargetVal}}
		} else {
			copy.Target = append(PerkIDMap[perk.ID].Target, perk.TargetVal)
			PerkIDMap[perk.ID] = copy
		}
	}

	return nil
}

func initSimulatorConstants() error {
	var log = utils.GetFunctionLogger("initSimulatorConstants")

	var numerator, denominator float32

	intNumerator, err := (helper.GetConstant("type_multiplier_numerator"))
	if err != nil {
		log.Error(err)
		return err
	}

	intDenominator, err := helper.GetConstant("type_multiplier_denominator")
	if err != nil {
		log.Error(err)
		return err
	}

	denominator = float32(intDenominator)
	numerator = float32(intNumerator)

	// Type Multiplier Constants
	functions.Strong = numerator / denominator
	functions.Equal = 1.0
	functions.Weak = denominator / numerator

	// Type Multiplier for each minicon
	functions.TypeMultipliers = map[string]map[string]float32{
		"FIRE": {
			"FIRE":    functions.Equal,
			"NORMAL":  functions.Equal,
			"WATER":   functions.Strong,
			"THUNDER": functions.Weak,
		},
		"WATER": {
			"FIRE":    functions.Weak,
			"NORMAL":  functions.Equal,
			"WATER":   functions.Equal,
			"THUNDER": functions.Strong,
		},
		"THUNDER": {
			"FIRE":    functions.Strong,
			"NORMAL":  functions.Equal,
			"WATER":   functions.Weak,
			"THUNDER": functions.Equal,
		},
		"NORMAL": {
			"FIRE":    functions.Equal,
			"NORMAL":  functions.Equal,
			"WATER":   functions.Equal,
			"THUNDER": functions.Equal,
		},
	}

	return nil
}

// Initing the simulator and preloading values required for the simulator
func Init() error {
	var log = utils.GetFunctionLogger("(Simulator) Init")

	err := generateMiniconMap()
	if err != nil {
		log.Error("Error while generating minicon map. Error = ", err)
		return err
	}

	err = generatePerkMap()
	if err != nil {
		log.Error("Error while generating perk map. Error = ", err)
		return err
	}

	err = initSimulatorConstants()
	if err != nil {
		log.Error("Error while initialising constants. Error = ", err)
		return err
	}

	log.Info("Simulator was initialised successfully")
	return nil
}
