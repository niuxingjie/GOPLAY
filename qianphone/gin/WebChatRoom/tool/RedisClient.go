package tool

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

type RedisConfig struct {
	Addr     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

func RedisClient(redis_config RedisConfig) *redis.Client {
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

func GetRedisClient() *redis.Client {
	// 读取配置文件
	config := GetAppConfig()

	// redis client
	redis := RedisClient(config.RedisConfig)
	return redis
}
