package redis

import "github.com/go-redis/redis/v8"

var (
	redisInstance *redis.Client
)

func GetRedis() (instance *redis.Client, err error) {
	if redisInstance == nil {
		instance, err = NewRedis()
		if err != nil {
			return nil, err
		}
		redisInstance = instance
	}
	return redisInstance, nil
}

func NewRedis() (instance *redis.Client, err error) {
	opt := &redis.Options{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		Password: "DbWV0cfe",
		DB:       0, // use default DB
	}
	return redis.NewClient(opt), nil
}
