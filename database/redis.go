package database

import (
	"fmt"
	"strconv"

	"github.com/delta/arcadia-backend/config"
	"github.com/delta/arcadia-backend/utils"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v7"
)

// redis connection object
var RedisDB *redis.Client

// GetRedisDB returns the Redis Database
func GetRedisDB() *redis.Client {
	return RedisDB
}

func ConnectRedisDB() {
	var log = utils.GetFunctionLogger("ConnectRedisDB")

	config := config.GetConfig()
	redisHost := config.RedisDb.Host
	redisPort := strconv.FormatUint(uint64(config.RedisDb.Port), 10)

	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     addr,                    // host:port of the redis server
		Password: config.RedisDb.Password, // Use password set
		DB:       0,                       // use default DB
	})

	// testing connection with redis server
	if _, err := RedisDB.Ping().Result(); err != nil {
		log.Error(err)
		panic(fmt.Errorf(color.RedString("error connecting with redis db: %+v", err)))
	} else {
		log.Info("Redis Database connected!")
		fmt.Println(color.GreenString("Redis Database connected!"))
	}
}
