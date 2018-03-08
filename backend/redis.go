package backend

import (
	"github.com/go-redis/redis"
)

type RedisBackend struct {
	client *redis.Client
}

func NewRedisBackend(redisURL string) (*RedisBackend, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)

	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisBackend{
		client: client,
	}, nil
}

func (r *RedisBackend) Write(key, value string) error {
	return nil
}
