package helpers

import (
	"strconv"

	"github.com/delta/arcadia-backend/database"
	"github.com/delta/arcadia-backend/server/model"
	"github.com/delta/arcadia-backend/utils"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

func InsertNewUserRedis(userID uint) (err error) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "InsertNewUserRedis",
	})

	defaultTrophies, err := GetConstant("default_trophy_count")
	if err != nil {
		l.Error(err)
		return err
	}

	err = UpdateUserTrophies(userID, uint(defaultTrophies))
	if err != nil {
		l.Error(err)
		return err
	}

	l.Info("New user added to leaderboard")

	return nil
}

func UpdateUserTrophies(userID uint, newTrophies uint) (err error) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":            "UpdateUserTrophies",
		"param_userID":      userID,
		"param_newTrophies": newTrophies,
	})

	redisDB := database.GetRedisDB()

	_, err = redisDB.ZAdd("leaderboard", &redis.Z{
		Score:  float64(newTrophies),
		Member: userID}).Result()

	if err != nil {
		l.Error(err)
		return err
	}

	l.Info("User's trophies updated")

	return nil
}

func GetUserRank(userID uint) (rank uint, err error) {
	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":       "GetUserRank",
		"param_userID": userID,
	})

	redisDB := database.GetRedisDB()

	result, err := redisDB.ZRevRank("leaderboard", strconv.FormatUint(uint64(userID), 10)).Result()

	if err != nil {
		if err == redis.Nil {
			l.Warn("User not found in leaderboard. Adding them again")
			db := database.GetDB()
			var user model.User
			db.First(&user, userID)
			err1 := UpdateUserTrophies(userID, user.Trophies)
			if err1 != nil {
				l.Error(err1)
				return 0, err1
			}
		} else {
			l.Error(err)
			return 0, err
		}
	}

	return uint(result) + 1, nil
}

func GetUsersWithRankInRange(startRank uint, count uint) (results []redis.Z, err error) {

	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":      "GetUsersWithRankInRange",
		"param_start": startRank,
		"param_count": count,
	})

	redisDB := database.GetRedisDB()
	results, err = redisDB.ZRevRangeWithScores("leaderboard", int64(startRank-1), int64(startRank+count-2)).Result()
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return results, nil
}

// Called by Matchmaking
func FindSuitors(userID uint) (suitorIDs []uint, err error) {
	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method":       "FindSuitors",
		"param_userID": userID,
	})

	rank, err := GetUserRank(userID)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	rankrange, err := GetConstant("matchmaking_rank_range")
	if err != nil {
		l.Error(err)
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
		l.Error(err)
		return nil, err
	}

	for _, result := range results {
		// ensure player doesnt fight himself
		if result.Member.(uint) != userID {
			suitorIDs = append(suitorIDs, result.Member.(uint))
		}
	}

	return suitorIDs, nil
}

// function to update/add ALL users ranks.. useful if redis were to fail and call when server starts
func UpdateRedis() (err error) {
	var l = utils.GetLogger().WithFields(logrus.Fields{
		"method": "UpdateRedis",
	})

	redisDB := database.GetRedisDB()
	redisDB.FlushAll()

	var allUsers []model.User
	db := database.GetDB()
	db.Find(&allUsers)

	for _, user := range allUsers {
		err := UpdateUserTrophies(user.ID, user.Trophies)
		if err != nil {
			l.Error(err)
			return err
		}
	}

	l.Info("Entire leaderboard updated")
	return nil
}
