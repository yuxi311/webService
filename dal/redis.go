package dal

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/yuxi311/webService/internal/config"
)

var internal_rdb *redis.Client

func InitRedis() error {
	cfg := config.Redis()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Server,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()

	_, err := rdb.Do(ctx, "SELECT", cfg.DB).Result()
	if err != nil {
		return err
	}

	internal_rdb = rdb
	return nil
}

func RDB() *redis.Client {
	return internal_rdb
}
