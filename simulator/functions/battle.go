package simulator

// Each Round Comprises of 3 stages: Start, Attack and Hurt stage.
// Each stage Comprises of several Actions which describe the action of each minicon
type Round struct {
	StartStage  [][][]string
	AttackStage [][][]string
	HurtStage   [][][]string
}

// DKHMinicons Keeps track of all the Minicon which got Hurt, a Kill or is Dead
// This is used for keeping track of minicons for perks which have conditional triggers
type DKHMinicons struct {
	KillMinicons []int // indices of minicons which got a kill
	HurtMinicons []int // indices of minicons which got hurt
	DeadMinicons []int // indices of minicons which died
}

// Checks if DKHMinicons Object has any killed or dead minicons
func (DKHMiniconsObject *DKHMinicons) HasDeadOrKillMincons() (exist bool) {
	return (len(DKHMiniconsObject.DeadMinicons) > 0) || (len(DKHMiniconsObject.KillMinicons) > 0)
}

func (DKHMiniconsObject *DKHMinicons) InitialiseDKHMinicons() {
	DKHMiniconsObject.KillMinicons = make([]int, 0)
	DKHMiniconsObject.HurtMinicons = make([]int, 0)
	DKHMiniconsObject.DeadMinicons = make([]int, 0)
}

// Checks if DKHMinicons Object has any hurt minicons
func (DKHMiniconsObject *DKHMinicons) HasHurtMinicons() (exist bool) {
	return len(DKHMiniconsObject.HurtMinicons) > 0
}

// make a copy of the DKMinicons to avoid pointer issues
func makeDKMiniconsCopy(TheArrayToBeCopied []int) (DKMiniconsCopy []int) {

	DKMiniconsCopy = make([]int, 0)
	copy(DKMiniconsCopy, TheArrayToBeCopied)
	return DKMiniconsCopy
}

// make a copy of the DKHMinicons to avoid pointer issues
func MakeDKHMiniconsCopy(OriginalArray DKHMinicons) (DKHMiniconsCopy DKHMinicons) {

	DKHMiniconsObject := DKHMinicons{}
	DKHMiniconsObject.DeadMinicons = makeDKMiniconsCopy(OriginalArray.DeadMinicons)
	DKHMiniconsObject.KillMinicons = makeDKMiniconsCopy(OriginalArray.KillMinicons)
	DKHMiniconsObject.HurtMinicons = append(DKHMiniconsObject.HurtMinicons, OriginalArray.HurtMinicons...)
	return DKHMiniconsObject
}

func MakeRoundObj(StartStage, AttackStage, HurtStage [][][]string) (RoundCopy Round) {
	RoundObj := Round{StartStage, AttackStage, HurtStage}
	return RoundObj
}
