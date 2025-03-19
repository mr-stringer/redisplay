package main

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func main() {

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	/*Put some stuff in Redis*/
	err := rdb.Set(ctx, "m1", "This is the primary message", 0).Err()
	if err != nil {
		slog.Error("redis error", "error string", err.Error())
		return
	}

	val, err := rdb.Get(ctx, "m1").Result()
	if err != nil {
		slog.Error("redis error", "error string", err.Error())
		return
	}

	slog.Info("redis get", "key", val)
	return
}
