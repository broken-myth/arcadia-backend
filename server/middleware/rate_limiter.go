package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/delta/arcadia-backend/config"
	helper "github.com/delta/arcadia-backend/server/helper/general"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	helper.SendError(c, http.StatusTooManyRequests, "Too many requests, please try again later")
}

func RateLimiter() gin.HandlerFunc {

	redisHost := config.GetConfig().RedisDb.Host
	redisPort := strconv.FormatUint(uint64(config.GetConfig().RedisDb.Port), 10)
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	Limit := config.GetConfig().Ratelimit
	if Limit <= 0 {
		Limit = 1
	}

	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.GetConfig().RedisDb.Password,
			DB:       0,
		}),
		Rate:  time.Second,
		Limit: Limit,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
	return mw
}
