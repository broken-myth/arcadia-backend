package simulator

import helper "github.com/delta/arcadia-backend/simulator/functions"

// Healer
var TrialStartHeal = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "START",
	Effect:        "HEAL",
	TargetIndices: []int{0, -2},
}

// Weaker Trial Perk

var TrialStartW = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "START",
	Effect:        "DMG",
	TargetIndices: []int{0, 1},
}

var TrialStart2W = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "START",
	Effect:        "DMG",
	TargetIndices: []int{1},
}

var TrialHurtW = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 2, 3, 4},
}
var TrialHurtAllW = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 2, 3, 4},
}
var TrialDeadW = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "DEAD",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 2, 3},
}
var TrialKillW = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "KILL",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 2, 3},
}

// Overpowered Perks
var TrialStartOP = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "START",
	Effect:        "DMG",
	TargetIndices: []int{0, 2},
}

var TrialDeadOP = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "DEAD",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 3, 5, 7, 9, 10},
}
var TrialKillOP = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "KILL",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 3, 5, 7, 9, 10},
}

// Trial Perks

var TrialStart = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "START",
	Effect:        "DMG",
	TargetIndices: []int{0, 1},
}

var TrialDeadstrong = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "DEAD",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 3, 5, 7, 9},
}
var TrialKillstrong = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "KILL",
	Effect:        "DMG",
	TargetIndices: []int{2, 4, 6, 8, 10},
}

var TrialStart2 = helper.Perk{
	PerkValue:     1000,
	PerkTrigger:   "START",
	Effect:        "DMG",
	TargetIndices: []int{1, 2},
}

var TrialHurt = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{0, 1},
}
var TrialHurtAll = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 2, 3, 4},
}
var TrialHurt2 = helper.Perk{
	PerkValue:     20,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{0, 1, 2, 3, 4},
}

var TrialDead = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "DEAD",
	Effect:        "DMG",
	TargetIndices: []int{0, 1},
}
var TrialKill = helper.Perk{
	PerkValue:     10,
	PerkTrigger:   "KILL",
	Effect:        "DMG",
	TargetIndices: []int{0, 1},
}

// Perks
var VisiciousVenom = helper.Perk{PerkValue: 10,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{1, 2, 4},
}
var SoothingAid = helper.Perk{PerkValue: 02,
	PerkTrigger:   "START",
	Effect:        "HEAL",
	TargetIndices: []int{-2, -1, 0, 1, 2},
}

var MightyStrike = helper.Perk{PerkValue: 2,
	PerkTrigger:   "DEAD",
	Effect:        "ATKBUFF",
	TargetIndices: []int{1},
}
var WeakenedBlow = helper.Perk{PerkValue: 11,
	PerkTrigger:   "KILL",
	Effect:        "ATKDEBUFF",
	TargetIndices: []int{0, 2},
}
var NonStop = helper.Perk{PerkValue: 12,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{0, 1}}
var FlyAhead = helper.Perk{PerkValue: 15,
	PerkTrigger:   "START",
	Effect:        "SWAP",
	TargetIndices: []int{-2},
}
var MachPunch = helper.Perk{PerkValue: 30,
	PerkTrigger:   "HURT",
	Effect:        "DMG",
	TargetIndices: []int{1, 2},
}

// Trial Minicons

var TrialMinicon = helper.Minicon{Name: "TrialMinicon",
	Health:     1000,
	BaseHealth: 1000,
	Attack:     100,
	Type:       "NORMAL",
	Dead:       false,
	Perk: map[string]helper.Perk{
		"START": TrialStart,
		"HURT":  TrialHurt,
		"KILL":  TrialKillOP,
		"DEAD":  TrialDeadOP,
	}}
var TrialMinicon2 = helper.Minicon{Name: "TrialMinicon",
	Health:     1000,
	BaseHealth: 1000,
	Attack:     100,
	Type:       "NORMAL",
	Dead:       false,
	Perk:       map[string]helper.Perk{"DEAD": TrialDead}}
var TrialMinicon3 = helper.Minicon{Name: "TrialMinicon",
	Health:     1000,
	BaseHealth: 1000,
	Attack:     100,
	Type:       "NORMAL",
	Dead:       false,
	Perk:       map[string]helper.Perk{"START": TrialStartHeal}}

// Very Strong Attack Trial Minicon

var TrialMiniconStrong = helper.Minicon{Name: "TrialMinicon",
	Health:     1000,
	BaseHealth: 1000,
	Attack:     1000,
	Type:       "NORMAL",
	Dead:       false,
	Perk: map[string]helper.Perk{"START": TrialStart,
		"HURT": TrialHurt}}
var TrialMiniconStrongAll = helper.Minicon{Name: "TrialMinicon",
	Health:     1000,
	BaseHealth: 1000,
	Attack:     1000,
	Type:       "NORMAL",
	Dead:       false,
	Perk: map[string]helper.Perk{
		"START": TrialStart,
		"HURT":  TrialHurt,
		"KILL":  TrialKill,
		"DEAD":  TrialDead}}
var TrialMiniconScout = helper.Minicon{Name: "TrialMinicon",
	Health:     1000,
	BaseHealth: 1000,
	Attack:     100,
	Type:       "NORMAL",
	Dead:       false,
	Perk: map[string]helper.Perk{"START": TrialStart,
		"HURT": TrialHurtAll,
		"KILL": TrialKill,
		"DEAD": TrialDead}}

// Minicons

var Sparkler = helper.Minicon{Name: "Sparkler",
	Health:     500,
	BaseHealth: 100,
	Attack:     60,
	Type:       "THUNDER",
	Dead:       false,
	Perk: map[string]helper.Perk{"START": SoothingAid,
		"HURT": VisiciousVenom,
		"KILL": WeakenedBlow,
		"DEAD": MightyStrike}}

var Phoenixa = helper.Minicon{Name: "Phoenixa",
	Health:     78,
	BaseHealth: 78,
	Attack:     99,
	Type:       "FIRE",
	Dead:       false,
	Perk:       map[string]helper.Perk{"START": SoothingAid, "KILL": WeakenedBlow}}
var Aquaray = helper.Minicon{Name: "Aquaray",
	Health:     45,
	Attack:     67,
	BaseHealth: 45,
	Type:       "WATER",
	Dead:       false,
	Perk:       map[string]helper.Perk{"HURT": MachPunch, "START": SoothingAid}}
var Galliard = helper.Minicon{Name: "Galliard",
	Health:     54,
	Attack:     50,
	BaseHealth: 54,
	Type:       "NORMAL",
	Dead:       false,
	Perk:       map[string]helper.Perk{"HURT": NonStop, "DEAD": MightyStrike}}
