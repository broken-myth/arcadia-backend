package helper

import (
	"errors"
	"strconv"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/go-redis/redis/v7"
)

// DO not use this function directly, use UpdateUserTrophies instead,
// to avoid matching up against users without a lineup
func InsertNewUserRedis(userID uint) (err error) {

	var log = utils.GetFunctionLogger("InsertNewUserRedis")

	defaultTrophies, err := GetConstant("default_trophy_count")
	if err != nil {
		log.Error(err)
		return err
	}

	err = UpdateUserTrophies(userID, defaultTrophies)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("New user added to leaderboard")

	return nil
}

func UpdateUserTrophies(userID uint, newTrophies int) (err error) {

	var log = utils.GetFunctionLogger("UpdateUserTrophies")

	redisDB := database.GetRedisDB()

	_, err = redisDB.ZAdd("leaderboard", &redis.Z{
		Score:  float64(newTrophies),
		Member: userID}).Result()

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func GetUserRank(userID uint) (rank uint, err error) {
	var log = utils.GetFunctionLogger("GetUserRank")

	redisDB := database.GetRedisDB()

	result, err := redisDB.ZRevRank("leaderboard", strconv.FormatUint(uint64(userID), 10)).Result()

	if err != nil {
		if err == redis.Nil {
			db := database.GetDB()
			lineup := []model.Lineup{}
			if err := db.Model(&lineup).Where("creator_id = ?", userID).Error; err != nil {
				log.Errorln("giving 0")
				return 0, nil
			}
			if len(lineup) == 0 { // don't add user to leaderboard if they don't have a lineup. Return rank = 0 instead
				return 0, nil
			}

			log.Warn("User not found in leaderboard. Adding them again")
			var user model.User
			db.First(&user, userID)
			err1 := UpdateUserTrophies(userID, user.Trophies)
			if err1 != nil {
				log.Error(err1)
				return 0, err1
			}
		} else {
			log.Error(err)
			return 0, err
		}
	}

	return uint(result) + 1, nil
}

func GetUsersWithRankInRange(startRank uint, count uint) (results []redis.Z, err error) {

	var log = utils.GetFunctionLogger("GetUsersWithRankInRange")

	redisDB := database.GetRedisDB()
	results, err = redisDB.ZRevRangeWithScores("leaderboard", int64(startRank-1), int64(startRank+count-2)).Result()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return results, nil
}

var ErrNoLineup = errors.New("No lineup set")

// Called by Matchmaking
func FindSuitors(userID uint) (suitorIDs []uint, err error) {
	var log = utils.GetFunctionLogger("FindSuitors")

	rank, err := GetUserRank(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if rank == 0 {
		return nil, ErrNoLineup
	}

	rankrange, err := GetConstant("matchmaking_rank_range")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	startRank := int(rank) - rankrange
	countRanks := uint(2*rankrange) + 1
	if startRank < 1 {
		countRanks = uint(rankrange) + rank
		startRank = 1
	}

	results, err := GetUsersWithRankInRange(uint(startRank), countRanks)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// ensure player doesnt fight himself
	for _, result := range results {
		id, err := strconv.ParseUint(result.Member.(string), 10, 64)
		if err != nil {
			return []uint{}, err
		}
		if uint(id) != userID {
			suitorIDs = append(suitorIDs, uint(id))
		}
	}

	return suitorIDs, nil
}

// function to update/add ALL users ranks (those who have created a lineup)
// useful if redis were to fail and call when server starts
func UpdateRedis() (err error) {
	var log = utils.GetFunctionLogger("UpdateRedis")

	redisDB := database.GetRedisDB()
	redisDB.FlushAll()

	db := database.GetDB()
	var userIDs []int
	if err := db.Model(&model.Lineup{}).Distinct().Pluck("creator_id", &userIDs).Error; err != nil {
		log.Error("Error while getting distinct UserIDs. Error = ", err)
		return err
	}

	var allUsers []model.User
	if err := db.Where("ID IN ?", userIDs).Find(&allUsers).Error; err != nil {
		log.Error("Error while fetching Users. Error = ", err)
		return err
	}

	for _, user := range allUsers {
		err := UpdateUserTrophies(user.ID, user.Trophies)
		if err != nil {
			log.Error("Error while updating trophies. Error = ", err)
			return err
		}
	}

	log.Info("Entire leaderboard updated")
	return nil
}
