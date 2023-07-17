package model

import (
	"gorm.io/gorm"
)

type ConstantType string

const (
	MiniconsInLineup                 ConstantType = "minicons_in_lineup"
	MatchmakingRankRange             ConstantType = "matchmaking_rank_range"
	MinTrophyGain                    ConstantType = "min_trophy_gain"
	TrophyGainRange                  ConstantType = "trophy_gain_range"
	TrophyDiffLoser                  ConstantType = "trophy_diff_loser"
	SurvivorTrophyRange              ConstantType = "survivor_trophy_range"
	DefaultTrophyCount               ConstantType = "default_trophy_count"
	SuccessiveDuplicateMatchLimit    ConstantType = "successive_duplicate_match_limit"
	DailyAttackLimit                 ConstantType = "daily_attack_limit"
	IncrXpMinicon                    ConstantType = "incr_xp_minicon"
	IncrXpUser                       ConstantType = "incr_xp_user"
	XpLevelMultiplier                ConstantType = "xp_level_multiplier"
	XpBaseCount                      ConstantType = "xp_base_count"
	LevelUpStatMultiplierNumerator   ConstantType = "level_up_stat_multiplier_numerator"
	LevelUpStatMultiplierDenominator ConstantType = "level_up_stat_multiplier_denominator"
	IsArenaOpen                      ConstantType = "is_arena_open"
	MaxMiniconLevel                  ConstantType = "max_minicon_level"
	MaxUnlockedMinicons              ConstantType = "max_unlocked_minicons"
	TypeMultiplierNumerator          ConstantType = "type_multiplier_numerator"
	TypeMultiplierDenominator        ConstantType = "type_multiplier_denominator"
)

type Constant struct {
	gorm.Model
	Name  ConstantType `gorm:"not null; unique; type:varchar(255);"`
	Value int          `gorm:"not null;"`
}
