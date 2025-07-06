package cache

import "github.com/redis/go-redis/v9"

type RedisOptions struct {
	Addr     string
	Password string
	DB       int
}

// NewRedisClient создание нового клиента redis
func NewRedisClient(client RedisOptions) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     client.Addr,
		Password: client.Password,
		DB:       client.DB,
	})
	return rdb
}
