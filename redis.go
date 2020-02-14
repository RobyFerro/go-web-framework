package gwf

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

// Connect to Redis
func ConnectRedis(conf Conf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		ProcessError(err)
	}

	return client
}
