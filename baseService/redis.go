package baseService

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/juntao3000/fxh-spider-go-common/baseCommon"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	RedisClient *redis.Client
)

func DisposeRedis() {
	if RedisClient != nil {
		_ = RedisClient.Close()
	}
}

func InitRedis(isDex bool) error {
	if RedisClient != nil {
		return nil
	}

	if isDex {
		RedisClient = redis.NewClient(&redis.Options{
			//Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%d", baseCommon.BaseConfig.DexRedisHost, baseCommon.BaseConfig.DexRedisPort),
			Password: baseCommon.BaseConfig.DexRedisPassword,
			DB:       baseCommon.BaseConfig.DexRedisDatabase,
		})
	} else {
		RedisClient = redis.NewClient(&redis.Options{
			//Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%d", baseCommon.BaseConfig.RedisHost, baseCommon.BaseConfig.RedisPort),
			Password: baseCommon.BaseConfig.RedisPassword,
			DB:       baseCommon.BaseConfig.RedisDatabase,
		})
	}

	infoRet := RedisClient.Info(context.Background(), "Server")
	if infoRet.Err() != nil {
		return infoRet.Err()
	}

	return nil
}

func TryLock(lockKey string, ttl time.Duration) (*redislock.Lock, error) {
	locker := redislock.New(RedisClient)
	return locker.Obtain(context.TODO(), lockKey, ttl, nil)
}
