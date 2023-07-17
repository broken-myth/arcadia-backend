package helper

import (
	"math"

	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"gorm.io/gorm"
)

var levelUpStatMultiplier float32

func UpdateXpUserLineup(tx *gorm.DB, userID uint) (err error) {
	params := map[string]interface{}{
		"userID": userID,
	}

	var log = utils.GetFunctionLoggerWithFields("UpdateXpUserLineup", params)

	incrXpMiniconInt, err := GetConstant("incr_xp_minicon")
	if err != nil {
		log.Error("Error while fetching constant incr_xp_minicon. Error = ", err)
		tx.Rollback()
		return err
	}
	incrXpMinicon := uint(incrXpMiniconInt)

	incrXpUserInt, err := GetConstant("incr_xp_user")
	if err != nil {
		log.Error("Error while fetching constant incr_xp_user. Error = ", err)
		tx.Rollback()
		return err
	}
	incrXpUser := uint(incrXpUserInt)
	// Get from constants

	var lineup []model.Lineup

	err = tx.Preload("OwnedMinicon").Preload("OwnedMinicon.Owner").Where("creator_id = ?", userID).Find(&lineup).Error
	if err != nil {
		log.Error("Error while loading OwnedMinicon. Error = ", err)
		tx.Rollback()
		return err
	}

	for _, lineupMinicon := range lineup {
		userXpDeltaValue, err := UpdateXpMinicon(tx, &lineupMinicon.OwnedMinicon, "increment", int(incrXpMinicon))
		if err != nil {
			log.Error("Error while updating xp of minicon. Error = ", err)
			tx.Rollback()
			return err
		}
		lineupMinicon.OwnedMinicon.Owner.XP = uint(int(lineupMinicon.OwnedMinicon.Owner.XP) + userXpDeltaValue)

	}

	var user model.User
	err = tx.Find(&user, userID).Error
	if err != nil {
		log.Error("Error while loading user. Error = ", err)
		tx.Rollback()
		return err
	}

	// TODO: implement locking properly
	user.XP += (uint(len(lineup)) * incrXpMinicon) + incrXpUser
	err = tx.Save(user).Error
	if err != nil {
		log.Error("Error while saving user. Error = ", err)
		tx.Rollback()
		return err
	}
	return nil
}

// updateType should be  "increment" or "update". Remember to Preload "Owner" in ownedMinicon
func UpdateXpMinicon(tx *gorm.DB, ownedMinicon *model.OwnedMinicon, updateType string, deltaValue int) (
	userXpDeltaValue int, err error) {

	var log = utils.GetFunctionLogger("UpdateXpMinicon")

	if updateType == "update" {
		deltaValue = deltaValue - int(ownedMinicon.XP)
	}

	maxMiniconLevelInt, err := GetConstant("max_minicon_level")
	if err != nil {
		log.Error("Error while fetching constant max_minicon_level. Error = ", err)
		tx.Rollback()
		return 0, err
	}
	maxMiniconLevel := uint(maxMiniconLevelInt)

	if ownedMinicon.Level >= maxMiniconLevel {
		return deltaValue, nil
	}

	xpLevelMultiplierInt, err := GetConstant("xp_level_multiplier")
	if err != nil {
		log.Error("Error while fetching constant xp_level_multiplier. Error = ", err)
		tx.Rollback()
		return 0, err
	}
	xpLevelMultiplier := uint(xpLevelMultiplierInt)

	xpBaseCountInt, err := GetConstant("xp_base_count")
	if err != nil {
		log.Error("Error while fetching constant xp_base_count. Error = ", err)
		tx.Rollback()
		return 0, err
	}
	xpBaseCount := uint(xpBaseCountInt)
	// Get from constants

	// Look into implementing locking properly

	ownedMinicon.XP = uint(int(ownedMinicon.XP) + deltaValue)

	if ownedMinicon.XP >= uint(math.Pow(float64(xpLevelMultiplier), float64(ownedMinicon.Level)))*xpBaseCount {
		err = UpdateLevelMinicon(tx, ownedMinicon)
		if err != nil {
			log.Error("Error while updating level of minicon. Error = ", err)
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Save(&ownedMinicon).Error
	if err != nil {
		log.Error("Error while saving ownedMinicon. Error = ", err)
		tx.Rollback()
		return 0, err
	}

	return deltaValue, nil
}

func UpdateLevelMinicon(tx *gorm.DB, ownedMinicon *model.OwnedMinicon) (err error) {
	var log = utils.GetFunctionLogger("UpdateLevelMinicon")

	var OwnedMiniconPerks []model.OwnedPerk

	// Get from constant if doesn't exist
	if levelUpStatMultiplier <= 0 {

		levelUpStatMultiplierNumerator, err := GetConstant("level_up_stat_multiplier_numerator")
		if err != nil {
			log.Error("Error while fetching constant level_up_stat_multiplier_numerator. Error = ", err)
			tx.Rollback()
			return err
		}

		levelUpStatMultiplierDenominator, err := GetConstant("level_up_stat_multiplier_denominator")
		if err != nil {
			log.Error("Error while fetching constant level_up_stat_multiplier_denominator. Error = ", err)
			tx.Rollback()
			return err
		}

		levelUpStatMultiplier = float32(levelUpStatMultiplierNumerator) / float32(levelUpStatMultiplierDenominator)
	}

	if err = tx.Where("owned_minicon_id = ?", ownedMinicon.ID).Find(&OwnedMiniconPerks).Error; err != nil {
		log.Error("Error while fetching Owned Minicon Perks. Error = ", err)
		tx.Rollback()
		return err
	}

	ownedMinicon.Level++
	ownedMinicon.Health = uint(levelUpStatMultiplier * float32(ownedMinicon.Health))
	ownedMinicon.Attack = uint(levelUpStatMultiplier * float32(ownedMinicon.Attack))

	err = tx.Save(ownedMinicon).Error
	if err != nil {
		log.Error("Error while saving ownedMinicon. Error = ", err)
		tx.Rollback()
		return err
	}

	for _, OwnedPerk := range OwnedMiniconPerks {
		OwnedPerk.PerkValue = uint(float32(OwnedPerk.PerkValue) * levelUpStatMultiplier)
		err = tx.Save(&OwnedPerk).Error
		if err != nil {
			log.Error("Error while saving ownedPerk. Error = ", err)
			tx.Rollback()
			return err
		}
	}

	return nil
}
