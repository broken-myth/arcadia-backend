package simulator

func contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

type ApplyStartPerkParams struct {
	AttackerTeam        *Team
	DefenderTeam        *Team
	Priority            *bool
	HurtMiniconsSet     *map[int]bool
	ArrayOfBattleAction *[][][]string
	CurrentDKHMinicons  *DKHMinicons
	StashOfUpdates      *StashOfUpdates
}

func generateStartStageParams(Attacker *Team, Defender *Team, Priority *bool,
	HurtMinicons *map[int]bool,
) (parameters *ApplyStartPerkParams) {
	var ArrayOfBattleAction [][][]string
	var StashOfStartUpdates StashOfUpdates // Stashes Updates
	var CurrentDKHMinicons DKHMinicons
	CurrentDKHMinicons.InitialiseDKHMinicons()
	params := ApplyStartPerkParams{Attacker, Defender, Priority, HurtMinicons, &ArrayOfBattleAction,
		&CurrentDKHMinicons, &StashOfStartUpdates}
	return &params
}

func (params *ApplyStartPerkParams) ResetStash() {
	params.StashOfUpdates = &StashOfUpdates{}
}

// Runs the (DK)* perks
func executeDeadAndKillPerkStage(AttackerTeam *Team, DefenderTeam *Team, DKHMiniconsObject *DKHMinicons,
	HurtMiniconsSet *map[int]bool,
) (UpdatedDKHMinicons *DKHMinicons, ArrayOfBattleAction [][][]string) {

	for DKHMiniconsObject.HasDeadOrKillMincons() {
		var AP int // AttackerTeam Pointer
		var DP int // DefenderTeam Pointer
		var CurrentDKHMinicons DKHMinicons
		CurrentDKHMiniconsCopy := *DKHMiniconsObject
		CurrentDKHMinicons.HurtMinicons = DKHMiniconsObject.HurtMinicons
		//Dead Perks
		for _, index := range CurrentDKHMiniconsCopy.DeadMinicons {
			var CurrentMoveBattleAction [][]string // AttackerTeam
			if index > 0 {
				AP = index - 1 // Contains merged array indexing to attacker index
				CurrentMoveBattleAction = RetrieveAndApplyPerk(AttackerTeam,
					DefenderTeam, AP, "DEAD", &CurrentDKHMinicons, true, HurtMiniconsSet)
			} else { //DefenderTeam Case
				DP = (-1 * index) - 1 // Contains merged array indexing to defender index
				CurrentMoveBattleAction = RetrieveAndApplyPerk(DefenderTeam,
					AttackerTeam, DP, "DEAD", &CurrentDKHMinicons, false, HurtMiniconsSet)
			}
			if len(CurrentMoveBattleAction) > 0 {
				ArrayOfBattleAction = append(ArrayOfBattleAction, CurrentMoveBattleAction) //Updates the actions in game if any
			}
		}

		//Kill Perks
		for _, index := range CurrentDKHMiniconsCopy.KillMinicons {
			var CurrentMoveBattleAction [][]string
			//AttackerTeam
			if index > 0 {
				AP := index - 1 // Contains merged array indexing to attacker index
				CurrentMoveBattleAction = RetrieveAndApplyPerk(AttackerTeam,
					DefenderTeam, AP, "KILL", &CurrentDKHMinicons, true, HurtMiniconsSet)
			} else { //DefenderTeam
				DP = (-1 * index) - 1 // Contains merged array indexing to defender index
				CurrentMoveBattleAction = RetrieveAndApplyPerk(DefenderTeam,
					AttackerTeam, DP, "KILL", &CurrentDKHMinicons, false, HurtMiniconsSet)
			}
			if len(CurrentMoveBattleAction) > 0 {
				ArrayOfBattleAction = append(ArrayOfBattleAction, CurrentMoveBattleAction) //Updates the actions in game if any

			}
		}
		//Cleans the DKHMinicon object (Removes any overlapping hurt and dead perks)
		*DKHMiniconsObject = CurrentDKHMinicons
	}
	return DKHMiniconsObject, ArrayOfBattleAction
}

// Run the (S(DK))
func executeStartStage(AttackerTeam *Team, DefenderTeam *Team,
	Priority *bool, HurtMiniconsSet *map[int]bool,
) (ArrayOfBattleAction [][][]string, UpdatedDKHMinicons DKHMinicons) {

	//var StashOfStartUpdates StashOfUpdates // Stashes Updates
	StartStageStashParams := generateStartStageParams(AttackerTeam, DefenderTeam, Priority, HurtMiniconsSet)
	AP, IsAttackerTeamAlive := AttackerTeam.GetLeadingMiniconIndex()
	DP, IsDefenderTeamAlive := DefenderTeam.GetLeadingMiniconIndex()
	if !IsAttackerTeamAlive || !IsDefenderTeamAlive {
		return *StartStageStashParams.ArrayOfBattleAction, *StartStageStashParams.CurrentDKHMinicons
	}

	//Iterates through the whole array
	for AP < len(AttackerTeam.MiniconTeam) && DP < len(DefenderTeam.MiniconTeam) {
		StartStageStashParams.ResetStash()
		AP = AttackerTeam.GetNextAliveMiniconIndex(AP)
		DP = DefenderTeam.GetNextAliveMiniconIndex(DP)
		ShouldFlipPriority := false
		if AP == -1 || DP == -1 {
			return *StartStageStashParams.ArrayOfBattleAction, *StartStageStashParams.CurrentDKHMinicons
		}

		if AttackerTeam.CheckIfMiniconAtIndexHasPerk(AP, "START") && DefenderTeam.CheckIfMiniconAtIndexHasPerk(DP, "START") {
			ShouldFlipPriority = true
		}

		// To incorporate randomness when two same perks run
		if *Priority {
			//AttackerTeam
			StashPerks(AP, AttackerTeam, DefenderTeam, StartStageStashParams.StashOfUpdates, "START", true) //Stashes Updates
			//DefenderTeam
			StashPerks(DP, DefenderTeam, AttackerTeam, StartStageStashParams.StashOfUpdates, "START", false) //Stashes Updates
		} else {
			//DefenderTeam
			StashPerks(DP, DefenderTeam, AttackerTeam, StartStageStashParams.StashOfUpdates, "START", false) //Stashes Updates
			//AttackerTeam
			StashPerks(AP, AttackerTeam, DefenderTeam, StartStageStashParams.StashOfUpdates, "START", true) //Stashes Updates
		}
		AP, DP = AP+1, DP+1
		applyStartPerkStash(StartStageStashParams)

		//Flip Priority
		if ShouldFlipPriority {
			*Priority = !(*Priority)
		}

	}

	// if attacker team is alive
	for AP < len(AttackerTeam.MiniconTeam) {
		StartStageStashParams.ResetStash()
		AP = AttackerTeam.GetNextAliveMiniconIndex(AP)
		if AP == -1 {
			return *StartStageStashParams.ArrayOfBattleAction, *StartStageStashParams.CurrentDKHMinicons
		}
		StashPerks(AP, AttackerTeam, DefenderTeam, StartStageStashParams.StashOfUpdates, "START", true) //Stashes Updates
		AP++
		//Apply start perks from stash
		applyStartPerkStash(StartStageStashParams)
	}

	// if defender team is alive
	for DP < len(DefenderTeam.MiniconTeam) {
		StartStageStashParams.ResetStash()
		DP = DefenderTeam.GetNextAliveMiniconIndex(DP)
		if DP == -1 {
			return *StartStageStashParams.ArrayOfBattleAction, *StartStageStashParams.CurrentDKHMinicons
		}
		StashPerks(DP, DefenderTeam, AttackerTeam, StartStageStashParams.StashOfUpdates, "START", false) //Stashes Updates
		DP++
		//Apply start perks from stash
		applyStartPerkStash(StartStageStashParams)
	}
	return *StartStageStashParams.ArrayOfBattleAction, *StartStageStashParams.CurrentDKHMinicons
}

func applyStartPerkStash(Params *ApplyStartPerkParams) {
	var BattleAction [][]string
	for _, object := range Params.StashOfUpdates.StashedPerkUpdates {
		PerkDKHMinicons, ArrayOfSelfBattleAction := ApplyPerk(
			object.Perk,
			object.SenderIndex,
			&object.ReceiverTeam,
			object.Type,
			object.IsAttacker, Params.HurtMiniconsSet)
		BattleAction = append(BattleAction, ArrayOfSelfBattleAction...)
		if len(BattleAction) > 0 {
			*Params.CurrentDKHMinicons = MergeAllDKHMiniconObjects(*Params.CurrentDKHMinicons,
				PerkDKHMinicons)
		}
	}
	if len(BattleAction) > 0 {
		*Params.ArrayOfBattleAction = append(*Params.ArrayOfBattleAction, BattleAction)
	}
	//Storing all the start perks Action in the Battle Action
	//Apply Dead And Kill Perks
	_, MiniconHurtStageAction := executeDeadAndKillPerkStage(Params.AttackerTeam, Params.DefenderTeam,
		Params.CurrentDKHMinicons, Params.HurtMiniconsSet)
	if len(MiniconHurtStageAction) != 0 {
		*Params.ArrayOfBattleAction = append(*Params.ArrayOfBattleAction, MiniconHurtStageAction...)
	}
}

// Runs the A(DK)* Perks
func executeAttackStage(AttackerTeam *Team, DefenderTeam *Team, DKHMiniconObject *DKHMinicons,
	Priority *bool, hurtMiniconsSet *map[int]bool) (ArrayOfBattleAction [][][]string, CurrentDKHMinicons DKHMinicons) {

	CurrentDKHMinicons = *DKHMiniconObject
	AP, IsAttackerTeamAlive := AttackerTeam.GetLeadingMiniconIndex()
	DP, IsDefenderTeamAlive := DefenderTeam.GetLeadingMiniconIndex() //Get the Leading minicon Indexes
	if !IsAttackerTeamAlive || !IsDefenderTeamAlive {
		return ArrayOfBattleAction, CurrentDKHMinicons
	}

	//Get the Leading minicon from indexes
	AttackerLeadingMinicon := AttackerTeam.MiniconTeam[AP]
	DefenderLeadingMinicon := DefenderTeam.MiniconTeam[DP]
	var StashOfAttackUpdates StashOfUpdates
	// To incorporate randomness and prevent advantage to one team over the other
	if *Priority {
		//AttackerTeam
		StashAttackPerks(AP, DP, DefenderTeam, &AttackerLeadingMinicon, &StashOfAttackUpdates, true) //Stashes Attack Perks
		//DefenderTeam
		StashAttackPerks(DP, AP, AttackerTeam, &DefenderLeadingMinicon, &StashOfAttackUpdates, false) //Stashes Attack Perks
	} else {
		//DefenderTeam
		StashAttackPerks(DP, AP, AttackerTeam, &DefenderLeadingMinicon, &StashOfAttackUpdates, false) //Stashes Attack Perks
		//AttackerTeam
		StashAttackPerks(AP, DP, DefenderTeam, &AttackerLeadingMinicon, &StashOfAttackUpdates, true) //Stashes Attack Perks
	}
	*Priority = !(*Priority)

	//Applying Attack Perks from stash and storing updates
	var AttackAction [][]string
	for _, object := range StashOfAttackUpdates.StashedAttackUpdates {
		AttackDKHMinicons, BattleAttackAction := ApplyAttack(object.TargetIndex, object.LeadingMinicon,
			*object.ReceiverTeam, object.SenderIndex, object.IsAttacker, hurtMiniconsSet)

		AttackAction = append(AttackAction, BattleAttackAction...)
		CurrentDKHMinicons = MergeAllDKHMiniconObjects(CurrentDKHMinicons, AttackDKHMinicons)
	}
	if len(AttackAction) > 0 {
		ArrayOfBattleAction = append(ArrayOfBattleAction, AttackAction)
	}

	DeadKillPerkDKHMinicons, ArrayOfDeadAndKillBattleAction :=
		executeDeadAndKillPerkStage(AttackerTeam, DefenderTeam, &CurrentDKHMinicons, hurtMiniconsSet)
	if len(ArrayOfDeadAndKillBattleAction) > 0 {
		ArrayOfBattleAction = append(ArrayOfBattleAction, ArrayOfDeadAndKillBattleAction...)
	}

	return ArrayOfBattleAction, *DeadKillPerkDKHMinicons
}

// Hurt Minicon Perk
func executeHurtStage(AttackerTeam *Team, DefenderTeam *Team, CurrentDKHMinicons *DKHMinicons,
	HurtMiniconSet *map[int]bool) (ArrayOfBattleAction [][][]string) {

	var NewDKHMiniconObject DKHMinicons
	KeepingTrackOfDKHMinicons := MakeDKHMiniconsCopy(*CurrentDKHMinicons)

	for CurrentDKHMinicons.HasHurtMinicons() {
		var CurrentDeadAndKillMinicon [][][]string
		HP := 0
		for HP < len(CurrentDKHMinicons.HurtMinicons) {
			index := CurrentDKHMinicons.HurtMinicons[HP]
			if index > 0 {
				AttackerIndex := (index - 1)
				if AttackerTeam.CheckIfMiniconAtIndexIsDead(AttackerIndex) {
					HP++
					continue
				}
				(*HurtMiniconSet)[index] = true
				ArrayOfHurtAction := RetrieveAndApplyPerk(AttackerTeam, DefenderTeam, AttackerIndex,
					"HURT", CurrentDKHMinicons, true, HurtMiniconSet)
				if len(ArrayOfHurtAction) > 0 {
					ArrayOfBattleAction = append(ArrayOfBattleAction, ArrayOfHurtAction)
				}
			} else {
				DefenderIndex := (index + 1) * -1
				if DefenderTeam.CheckIfMiniconAtIndexIsDead(DefenderIndex) {
					HP++
					continue
				}
				(*HurtMiniconSet)[index] = true
				ArrayOfHurtAction := RetrieveAndApplyPerk(DefenderTeam, AttackerTeam, DefenderIndex,
					"HURT", CurrentDKHMinicons, false, HurtMiniconSet)
				if len(ArrayOfHurtAction) > 0 {
					ArrayOfBattleAction = append(ArrayOfBattleAction, ArrayOfHurtAction)
				}
			}
			HP++
		}

		KeepingTrackOfDKHMinicons = MergeAllDKHMiniconObjects(KeepingTrackOfDKHMinicons,
			NewDKHMiniconObject)

		CurrentDKHMinicons = &NewDKHMiniconObject
		CurrentDKHMinicons, CurrentDeadAndKillMinicon = executeDeadAndKillPerkStage(AttackerTeam,
			DefenderTeam, CurrentDKHMinicons, HurtMiniconSet)

		if len(CurrentDeadAndKillMinicon) > 0 {
			ArrayOfBattleAction = append(ArrayOfBattleAction, CurrentDeadAndKillMinicon...)
		}
	}
	return ArrayOfBattleAction
}

//Running All the minicon Functionality

func ExecuteRound(AttackerTeam *Team, DefenderTeam *Team, Priority *bool) (BattleRoundIteration Round) {

	HurtMiniconSet := make(map[int]bool)
	//Making the Dead, Hurt And Killed Minicon Objects as well as a copy of the initial minicon lineup
	//Running the game with Start,Attack, Hurt Move
	StartBattleAction, CurrentStartHurtMinicon := executeStartStage(AttackerTeam,
		DefenderTeam, Priority, &HurtMiniconSet)

	AttackBattleAction, CurrentAttackHurtMinicon := executeAttackStage(AttackerTeam,
		DefenderTeam, &CurrentStartHurtMinicon, Priority, &HurtMiniconSet)

	HurtBattleAction := executeHurtStage(AttackerTeam, DefenderTeam,
		&CurrentAttackHurtMinicon, &HurtMiniconSet)

	//Making the Round object
	BattleRoundIteration = MakeRoundObj(StartBattleAction,
		AttackBattleAction, HurtBattleAction)

	return BattleRoundIteration
}
