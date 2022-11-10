package tool

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedisClient(config *Config) *redis.Client {
	redis_config := config.RedisConfig
	redis_client := redis.NewClient(&redis.Options{
		Addr: redis_config.Addr + ":" + redis_config.Port,
		// Password: redis_config.Password,
		DB: redis_config.Db,
	})

	// err := redis_client.Set("key", "value", time.Minute*1).Err()
	// if err != nil {
	// 	panic(err)
	// }
	Redis = redis_client
	return Redis
}
