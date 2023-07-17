package helper

import (
	"encoding/json"

	"github.com/delta/arcadia-backend/server/model"
	"gorm.io/gorm"
)

type PlayerLineUp struct {
	LineupMinicons []LineupMinicon
}

type LineupMinicon struct {
	Name     string
	Health   uint
	Attack   uint
	Type     string
	ImageURL string
}

func GenerateMiniconObject(Name string, Health uint, Attack uint, Type string, URL string) LineupMinicon {
	LineupOwnedMinicon := LineupMinicon{Name: Name, Health: Health, Attack: Attack, Type: Type, ImageURL: URL}
	return LineupOwnedMinicon
}

func ReturnPlayerMiniconLineup(tx *gorm.DB, PlayerID uint) (JSONDBOwnedMinicon []byte, err error) {
	var PlayerDBMiniconLineups []model.Lineup
	var PlayerOwnedMiniconLinup PlayerLineUp
	// Fetching Minicon Lineup from DB
	err = tx.Select("position_no, owned_minicon_id").Preload("OwnedMinicon.Minicon.Type").
		Preload("OwnedMinicon.Minicon").Preload("OwnedMinicon").
		Where("creator_id = ?", PlayerID).Find(&PlayerDBMiniconLineups).Error
	if err != nil {
		return JSONDBOwnedMinicon, err
	}

	// Generate Minicon Objects
	for _, Lineup := range PlayerDBMiniconLineups {
		DBOwnedMinicon := Lineup.OwnedMinicon
		PlayerOwnedMiniconLinup.LineupMinicons = append(PlayerOwnedMiniconLinup.LineupMinicons,
			GenerateMiniconObject(DBOwnedMinicon.Minicon.Name,
				DBOwnedMinicon.Health, DBOwnedMinicon.Attack, DBOwnedMinicon.Minicon.Type.Name, DBOwnedMinicon.Minicon.ImageLink))
	}
	JSONDBOwnedMinicon, err = json.Marshal(PlayerOwnedMiniconLinup.LineupMinicons)
	return JSONDBOwnedMinicon, err
}
