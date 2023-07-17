package simulator

// Minicon Object
type Minicon struct {
	Name       string
	Health     uint
	Attack     uint
	Type       string
	Dead       bool
	Perk       map[string]Perk
	BaseHealth uint
}

// Team Comprises of the Minicons sent to battle
type Team struct {
	MiniconTeam []Minicon
}

// AliveTeam Comprises of the true/false values correspondign whether Team Minicon is alive or dead
type AliveTeam struct {
	MiniconBoolTeam []bool
}

// Type Constant
var Strong float32
var Equal float32
var Weak float32
var TypeMultipliers map[string]map[string]float32

func getTypeMultiplier(receiverType string, senderType string) (multiplier float32) {
	return TypeMultipliers[receiverType][senderType]
}

// Used to update minicon Objects with the given changes
// DeltaValue: how much did it change by
// IsNewlyDead: whether it died because of the update
func (MiniconObject *Minicon) UpdateMiniconStats(Effect string, PreDeltaValue uint,
	Type string) (DeltaValue int, IsNewlyDead bool) {
	if MiniconObject.Health <= 0 {
		return 0, false
	}
	DeltaValue = int(PreDeltaValue) // PreDeltaValue is value before type multiplier, if applicable
	switch Effect {
	case "DMG":
		DeltaValue = -1 * int(float32(PreDeltaValue))
		if int(MiniconObject.Health)+DeltaValue <= 0 {
			MiniconObject.Health = 0 //Handles edge cases when uint is lesser than 0
		} else {
			MiniconObject.Health = uint(DeltaValue + int(MiniconObject.Health))
		}
	case "ATKBUFF":
		MiniconObject.Attack += uint(DeltaValue)
	case "HEAL":
		if int(MiniconObject.Health)+DeltaValue >= int(MiniconObject.BaseHealth) {
			DeltaValue = int(MiniconObject.BaseHealth) - int(MiniconObject.Health)
			MiniconObject.Health = MiniconObject.BaseHealth //Handles edge cases when health crosses base health
		} else {
			MiniconObject.Health += uint(DeltaValue)
		}
	case "ATKDEBUFF":
		DeltaValue *= -1
		if int(MiniconObject.Attack)+DeltaValue <= 0 { //Handles edge cases when uint is lesser than 0
			DeltaValue = int(MiniconObject.Attack) - 1
			MiniconObject.Attack = 1 //We set up the Attack to be a minimum of 1
		} else {
			MiniconObject.Attack = uint(DeltaValue + int(MiniconObject.Attack))
		}
	case "ATTACK":
		multiplier := getTypeMultiplier(MiniconObject.Type, Type)
		DeltaValue = -1 * int(multiplier*float32(PreDeltaValue))
		if int(MiniconObject.Health)+DeltaValue <= 0 {
			MiniconObject.Health = 0 //Handles edge cases when uint is lesser than 0
		} else {
			MiniconObject.Health = uint(DeltaValue + int(MiniconObject.Health))
		}
	default:
		return
	}
	if MiniconObject.Health <= 0 { //Checks and updates if the minicon is dead
		MiniconObject.Dead = true
		IsNewlyDead = true
	}
	return DeltaValue, IsNewlyDead
}

func (MiniconObject *Minicon) checkIfMiniconHasPerkOfTrigger(Trigger string) (exist bool) {
	_, PerkPresent := MiniconObject.Perk[Trigger]
	return PerkPresent
}

func GenerateMiniconObjectFromDB(Name string, Health uint, Attack uint, Type string) (minicon Minicon) {
	DBOwnedMiniconObj := Minicon{}
	DBOwnedMiniconObj.Name = Name
	DBOwnedMiniconObj.Health = Health
	DBOwnedMiniconObj.BaseHealth = Health
	DBOwnedMiniconObj.Attack = Attack
	DBOwnedMiniconObj.Type = Type
	DBOwnedMiniconObj.Perk = make(map[string]Perk)
	return DBOwnedMiniconObj
}

func makeMiniconObject(Name string, Health uint,
	Attack uint, Type string, Dead bool, Perks []Perk) (minicon Minicon) {
	MiniconObj := Minicon{
		Name:       Name,
		Health:     Health,
		BaseHealth: Health,
		Attack:     Attack,
		Type:       Type,
		Dead:       Dead,
		Perk:       make(map[string]Perk)}
	for _, PerkObj := range Perks {
		MiniconObj.Perk[PerkObj.Effect] = MakePerkObject(PerkObj)
	}
	return MiniconObj
}

// Iterates through the array and returns true whether there is atleast one alive minicon or not
func (team *Team) IsAlive() (alive bool) {
	for _, Minicon := range team.MiniconTeam {
		if !Minicon.Dead {
			return true
		}
	}
	return false
}

// Applies the Perk to the target minicon in the Array
func (team *Team) UpdateTargetMinicon(target int, PerkUsed Perk, Type string,
) (DeltaValue int, IsNewlyDead bool) {

	DeltaValue, IsNewlyDead = team.MiniconTeam[target].
		UpdateMiniconStats(PerkUsed.Effect, PerkUsed.PerkValue, Type)

	return DeltaValue, IsNewlyDead
}

// Applies the attack to the target minicon in the Array
func (team *Team) AttackAndUpdateMinicon(AttackValue uint, Receiver int, Type string,
) (DeltaValue int, IsNewlyDead bool) {

	DeltaValue, IsNewlyDead = team.MiniconTeam[Receiver].
		UpdateMiniconStats("ATTACK", AttackValue, Type)

	return DeltaValue, IsNewlyDead
}

func MakeMiniconTeam(MiniconTeam []Minicon) Team {
	MiniconRowObj := Team{}
	for _, ActiveMiniconObj := range MiniconTeam {
		MiniconRowObj.MiniconTeam = append(MiniconRowObj.MiniconTeam,
			makeMiniconObject(ActiveMiniconObj.Name, ActiveMiniconObj.Health,
				ActiveMiniconObj.Attack, ActiveMiniconObj.Type,
				ActiveMiniconObj.Dead, GetAllPerks(ActiveMiniconObj.Perk)))
	}

	return MiniconRowObj
}

func (team *Team) GetLeadingMiniconIndex() (minconIndex int, exist bool) {
	StartingIndex := 0
	for StartingIndex < len(team.MiniconTeam) {
		if !team.MiniconTeam[StartingIndex].Dead {
			return StartingIndex, true
		}
		StartingIndex++
	}
	return -1, false
}

// Returns the index of the first occurrence of an alive minicon
// It iterates from index till end of array
func (team *Team) GetNextAliveMiniconIndex(Index int) (miniconIndex int) {
	StartingIndex := Index
	for StartingIndex < len(team.MiniconTeam) {
		if !team.MiniconTeam[StartingIndex].Dead {
			return StartingIndex
		}
		StartingIndex++
	}
	return -1
}

func (team *Team) CheckIfMiniconAtIndexIsDead(Index int) (dead bool) {
	MiniconObj := team.MiniconTeam[Index]
	return MiniconObj.Dead
}

func (team *Team) ReturnNumberOfSurvivors() (survivors uint) {
	for _, minicon := range team.MiniconTeam {
		if !minicon.Dead {
			survivors++
		}
	}
	return survivors
}

func (team *Team) CheckIfMiniconAtIndexHasPerk(Index int, PerkTrigger string) (exist bool) {
	MiniconObj := team.MiniconTeam[Index]
	return MiniconObj.checkIfMiniconHasPerkOfTrigger(PerkTrigger)
}

func (team *Team) GenerateAliveAndDeadStateSlice() (AliveTeam AliveTeam) {
	for _, minicon := range team.MiniconTeam {
		if minicon.Dead {
			AliveTeam.MiniconBoolTeam = append(AliveTeam.MiniconBoolTeam, false)
		} else {
			AliveTeam.MiniconBoolTeam = append(AliveTeam.MiniconBoolTeam, true)
		}
	}
	return AliveTeam
}

func (team *AliveTeam) GetLeadingMiniconIndex() (index int, exist bool) {
	StartingIndex := 0
	for StartingIndex < len(team.MiniconBoolTeam) {
		if team.MiniconBoolTeam[StartingIndex] {
			return StartingIndex, true
		}
		StartingIndex++
	}
	return -1, false
}

// Used to obtain the Minicon position wrt to the targets of the perks
func (team *AliveTeam) GetMiniconPosition(Position int, StartingIndex int) (index int) {
	if Position < 0 {
		Jumps := Position * -1 //Keeping track of jumps required to reach required position
		for JumpIndex := StartingIndex - 1; JumpIndex >= 0; JumpIndex-- {
			if team.MiniconBoolTeam[JumpIndex] {
				Jumps-- // if the minicon is alive, Decrement the number of jumps left
			}
			if Jumps == 0 { // Return index when required number of jumps is 0
				return JumpIndex
			}
		}
		if Jumps != 0 {
			return -1 // index is not present , return -1
		}
		// When the position is negative, we have to traverse forwards
	} else if Position > 0 {
		Jumps := Position //Keeping track of jumps required to reach required position
		for JumpIndex := StartingIndex + 1; JumpIndex < len(team.MiniconBoolTeam); JumpIndex++ {
			if team.MiniconBoolTeam[JumpIndex] {
				Jumps-- // if the minicon is alive, Decrement the number of jumps left
			}
			if Jumps == 0 { // Return index when required number of jumps is 0
				return JumpIndex
			}
		}
		if Jumps != 0 {
			return -1 // index is not present , return -1
		}
	}
	// Return starting index when the Position is 0
	return StartingIndex
}
