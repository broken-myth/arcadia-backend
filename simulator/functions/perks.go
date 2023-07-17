package simulator

import (
	"strconv"
)

// Perks are divided into Self and Non Self Perks
// Self Perks which Boost the Player's Minicon team minicon stats i.e Health and Attack
// Non Self Perks which Lowers the Opponents's Minicon team minicons stats i.e Health and Attack

var SelfPerks = []string{"HEAL", "ATKBUFF"}
var NonSelfPerks = []string{"DMG", "ATKDEBUFF"}

type Perk struct {
	PerkValue     uint
	PerkTrigger   string
	Effect        string
	TargetIndices []int
}

// StashOfUpdates is used to implement simulataneous updation of minicon Teams
// It comprises of stashing the 2 primary methods of changing stats of a minicon
type StashOfUpdates struct {
	StashedPerkUpdates   []PerkUpdates
	StashedAttackUpdates []AttackUpdates
}

// PerkUpdates are to stash and apply perks when stashing is present
// AttackUpdates are used to apply Attacks when Stashing is present
type PerkUpdates struct {
	Perk         Perk
	SenderIndex  int
	ReceiverTeam Team
	Type         string
	IsAttacker   bool // Sender's PoV
}
type AttackUpdates struct {
	TargetIndex    int
	LeadingMinicon *Minicon
	ReceiverTeam   *Team
	SenderIndex    int
	IsAttacker     bool // Sender's PoV
}

func GeneratePerkFromDB(PerkValue uint, Trigger string, Effect string, Targets []int) (DBPerkObject Perk) {
	DBPerkObject.PerkValue = PerkValue
	DBPerkObject.PerkTrigger = Trigger
	DBPerkObject.Effect = Effect
	DBPerkObject.TargetIndices = Targets
	return DBPerkObject
}

func GetAllPerks(Perks map[string]Perk) (PerksArray []Perk) {
	PerksArray = make([]Perk, 0)
	for _, Val := range Perks {
		PerksArray = append(PerksArray, Val)
	}
	return PerksArray
}

func MakePerkObject(PerkObj Perk) (NewPerkObject Perk) {
	NewPerkObject = Perk{
		PerkValue:     PerkObj.PerkValue,
		PerkTrigger:   PerkObj.PerkTrigger,
		Effect:        PerkObj.Effect,
		TargetIndices: PerkObj.TargetIndices}
	return NewPerkObject
}

func (PerkObject *Perk) IsSelfPerk() bool {
	return contains(SelfPerks, PerkObject.Effect)
}
func (PerkObject *Perk) IsOfEffect(PerkEffect string) bool {
	return PerkObject.Effect == PerkEffect
}

func retrievePerkIfTriggerExists(MiniconObject Minicon, PerkTrigger string) (perk *Perk) {
	Perk, exist := MiniconObject.Perk[PerkTrigger]
	if !exist {
		return nil
	}
	return &Perk
}

// Stashing: In order to achieve simultaneous updation , we need to store all changes,
// then apply the changes
// this is done to avoid cases where the previous perk can kill a minicon and
// prevent that perk from applying despite being simultaneous
//
// Used to stash the minicons self and non self perks in the Stash Of Perk objects
func StashPerks(MiniconIndex int, ExecutingTeam *Team, OpponentTeam *Team, StashOfPerks *StashOfUpdates,
	PerkTrigger string, IsAttacker bool) {

	if MiniconIndex < len(ExecutingTeam.MiniconTeam) {
		MiniconObj := ExecutingTeam.MiniconTeam[MiniconIndex]
		Perk := retrievePerkIfTriggerExists(MiniconObj, PerkTrigger) // retreives the perk corresponding to that trigger
		if Perk != nil {
			var receivingTeam = *OpponentTeam
			if Perk.IsSelfPerk() {
				receivingTeam = *ExecutingTeam
			}
			StashOfPerks.StashedPerkUpdates = append(StashOfPerks.StashedPerkUpdates,
				PerkUpdates{*Perk, MiniconIndex, receivingTeam, MiniconObj.Type, IsAttacker})
		}
	}
}

// Stashes the Attack perks
func StashAttackPerks(SenderIndex int, ReceiverIndex int, ReceiverTeam *Team, LeadingMincion *Minicon,
	StashOfAttackUpdates *StashOfUpdates, IsAttacker bool) {

	StashOfAttackUpdates.StashedAttackUpdates = append(StashOfAttackUpdates.StashedAttackUpdates,
		AttackUpdates{ReceiverIndex, LeadingMincion, ReceiverTeam, SenderIndex, IsAttacker})
}

// applies perks
func ApplyPerk(Perk Perk, SenderIndex int, ReceiverTeam *Team,
	Type string, IsAttacker bool, HurtMiniconsSet *map[int]bool,
) (DKHMiniconObject DKHMinicons, ArrayOfNonSelfBattleAction [][]string) {
	DKHMiniconObject.InitialiseDKHMinicons()
	IsSelfPerk := Perk.IsSelfPerk()       // Checks if it's a self perk
	IsDMGEffect := Perk.IsOfEffect("DMG") //Check if the perk effect is dmg
	// Initial Minicon state keeps track of alive and dead minicons in each state
	InitialMiniconState := ReceiverTeam.GenerateAliveAndDeadStateSlice()

	for _, TargetIndex := range Perk.TargetIndices {
		if IsSelfPerk {
			//Calculating the relative Index of the minicon which receives effect
			TargetIndex = InitialMiniconState.GetMiniconPosition(TargetIndex, SenderIndex)
		} else {
			StartingIndex, isAlive := InitialMiniconState.GetLeadingMiniconIndex()
			if !isAlive {
				return DKHMiniconObject, ArrayOfNonSelfBattleAction
			}
			//Calculates Index of the target Minicon WRT the first Alive Minicon of opponent Team
			TargetIndex = InitialMiniconState.GetMiniconPosition(TargetIndex, StartingIndex)
		}
		if TargetIndex >= len(ReceiverTeam.MiniconTeam) || TargetIndex < 0 {
			continue //Checks for edge cases and skips the index
		}

		DeltaValue, IsDead := ReceiverTeam.UpdateTargetMinicon(TargetIndex, Perk, Type)
		if DeltaValue == 0 { // if there is a change in stats, there was some action -> minicon is updated
			continue // If minicon Wasn't updated skip the rest
		}

		// Updates Hurt Minicons present in DKHMiniconObject for Attacker And Defender if updated
		// (not valid for selfperks)

		if !IsSelfPerk && IsDMGEffect {
			targetPosition := TargetIndex + 1
			if IsAttacker {
				targetPosition *= -1
			}
			_, ok := (*HurtMiniconsSet)[targetPosition] // checks if the minicon is already present in the Hurt set
			if !ok {
				DKHMiniconObject.HurtMinicons = append(DKHMiniconObject.HurtMinicons, (targetPosition))
				(*HurtMiniconsSet)[targetPosition] = true
			}
		}

		// Updates the DeadMinicons and KilledMinicons Present in the DKHMiniconObject
		// Let PreIndex be the index of the minicon that got died or got a kill
		// For Defender , the indices are stored as -1*PreIndex -1
		// For Attacker , the indices are stored as PreIndex +1
		if IsDead {
			if IsAttacker {
				DKHMiniconObject.DeadMinicons = append(DKHMiniconObject.DeadMinicons,
					TargetIndex*-1-1)
				DKHMiniconObject.KillMinicons = append(DKHMiniconObject.KillMinicons,
					SenderIndex+1)
			} else {
				DKHMiniconObject.DeadMinicons = append(DKHMiniconObject.DeadMinicons,
					TargetIndex+1)
				DKHMiniconObject.KillMinicons = append(DKHMiniconObject.KillMinicons,
					SenderIndex*-1-1)
			}
		}

		// Generates the Action for the Non Self Perks
		if IsSelfPerk {
			if IsAttacker {
				ArrayOfNonSelfBattleAction = append(ArrayOfNonSelfBattleAction,
					[]string{Perk.Effect, Type, Perk.PerkTrigger,
						strconv.Itoa(SenderIndex + 1), strconv.Itoa(TargetIndex + 1), strconv.Itoa(DeltaValue)})
			} else {
				ArrayOfNonSelfBattleAction = append(ArrayOfNonSelfBattleAction,
					[]string{Perk.Effect, Type, Perk.PerkTrigger,
						"-" + strconv.Itoa(SenderIndex+1), "-" + strconv.Itoa(TargetIndex+1), strconv.Itoa(DeltaValue)})
			}
		} else {
			if IsAttacker {
				ArrayOfNonSelfBattleAction = append(ArrayOfNonSelfBattleAction,
					[]string{Perk.Effect, Type, Perk.PerkTrigger,
						strconv.Itoa(SenderIndex + 1), "-" + strconv.Itoa(TargetIndex+1), strconv.Itoa(DeltaValue)})
			} else {
				ArrayOfNonSelfBattleAction = append(ArrayOfNonSelfBattleAction,
					[]string{Perk.Effect, Type, Perk.PerkTrigger,
						"-" + strconv.Itoa(SenderIndex+1), "" + strconv.Itoa(TargetIndex+1), strconv.Itoa(DeltaValue)})
			}
		}

	}

	return DKHMiniconObject, ArrayOfNonSelfBattleAction
}

// Used to Attack to the Opponent's Team Leading Minicon
// Returns All the Action Caused by the Attack Move
func ApplyAttack(Target int, LeadingMinicon *Minicon, ReceiverTeam Team,
	Index int, IsAttacker bool, HurtMiniconsSet *map[int]bool,
) (DKHMiniconObject DKHMinicons, ArrayOfAttackBattleAction [][]string) {
	DKHMiniconObject.InitialiseDKHMinicons()
	Updated := false
	DeltaValue, IsNewlyDead := ReceiverTeam.AttackAndUpdateMinicon( //Attack Updates The minicon stats
		(LeadingMinicon.Attack), Target, LeadingMinicon.Type)

	if DeltaValue != 0 {
		Updated = true
	}
	if !Updated {
		return DKHMiniconObject, ArrayOfAttackBattleAction //If no changes Exit
	}

	// Updates Hurt Minicons present in DKHMiniconObject for Attacker And Defender if Updated
	if Updated {
		targetPosition := Target + 1
		if IsAttacker { // If is Attacker, the defender receives the attack, so make index negative
			targetPosition *= -1
		}
		_, ok := (*HurtMiniconsSet)[targetPosition]
		if !ok {
			DKHMiniconObject.HurtMinicons = append(DKHMiniconObject.HurtMinicons, (targetPosition))
			(*HurtMiniconsSet)[targetPosition] = true
		}
	}

	// Updates the Dead and Kill Minicons Present in the DKHMiniconObject if Any is found Dead
	// Let PreIndex be the index of the minicon that got died or got a kill
	// For Defender , the indices are stored as -1*PreIndex -1
	// For Attacker , the indices are stored as PreIndex +1
	if IsNewlyDead {
		if IsAttacker {
			DKHMiniconObject.DeadMinicons = append(DKHMiniconObject.DeadMinicons,
				Target*-1-1)
			DKHMiniconObject.KillMinicons = append(DKHMiniconObject.KillMinicons,
				Index+1)
		} else {
			DKHMiniconObject.DeadMinicons = append(DKHMiniconObject.DeadMinicons,
				Target+1)
			DKHMiniconObject.KillMinicons = append(DKHMiniconObject.KillMinicons,
				Index*-1-1)
		}
	}

	// Generates the Action for the Attack Move
	if IsAttacker {
		ArrayOfAttackBattleAction = append(ArrayOfAttackBattleAction, []string{
			"ATTACK",
			LeadingMinicon.Type,
			"ATTACK",
			"" + strconv.Itoa(Index+1), "-" + strconv.Itoa(Target+1),
			strconv.Itoa(DeltaValue),
		})
	} else {
		ArrayOfAttackBattleAction = append(ArrayOfAttackBattleAction, []string{
			"ATTACK",
			LeadingMinicon.Type,
			"ATTACK",
			"-" + strconv.Itoa(Index+1), "" + strconv.Itoa(Target+1),
			strconv.Itoa(DeltaValue),
		})
	}
	return DKHMiniconObject, ArrayOfAttackBattleAction
}

// Calls ApplyPerk
func RetrieveAndApplyPerk(ExecutingTeam *Team, OpponentTeam *Team, SenderIndex int, PerkTrigger string,
	CurrentDKHMiniconObject *DKHMinicons, IsAttacker bool, HurtMiniconsSet *map[int]bool,
) (BattleAction [][]string) {
	MiniconObj := ExecutingTeam.MiniconTeam[SenderIndex]
	Perk := retrievePerkIfTriggerExists(MiniconObj, PerkTrigger)

	if Perk != nil {
		var PerkDKHMinicons DKHMinicons
		var ArrayOfPerkAction [][]string
		if Perk.IsSelfPerk() {
			PerkDKHMinicons, ArrayOfPerkAction = ApplyPerk(*Perk, SenderIndex,
				ExecutingTeam, MiniconObj.Type, IsAttacker, HurtMiniconsSet)
		} else {
			PerkDKHMinicons, ArrayOfPerkAction = ApplyPerk(*Perk, SenderIndex,
				OpponentTeam, MiniconObj.Type, IsAttacker, HurtMiniconsSet)
		}

		BattleAction = append(BattleAction, ArrayOfPerkAction...)
		*CurrentDKHMiniconObject = MergeAllDKHMiniconObjects(*CurrentDKHMiniconObject,
			PerkDKHMinicons)
	}
	return BattleAction
}
