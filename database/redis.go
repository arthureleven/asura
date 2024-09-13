package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

func Init() error {
	Cache = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
	})

	command := Cache.Ping(context.Background())

	if err := command.Err(); err != nil {
		return err
	}

	return nil
}
