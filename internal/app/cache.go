package app

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Redis Redis
}

type Redis struct {
	Host     string
	Port     uint32
	DB       uint8
	Username string
	Password string
	Client   *redis.Client
}

func (c *Cache) prepare() error {
	return c.Redis.connect()
}

func (r *Redis) connect() error {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		DB:       int(r.DB),
		Username: r.Username,
		Password: r.Password,
	})

	stat := r.Client.Ping(context.Background())

	return stat.Err()
}
