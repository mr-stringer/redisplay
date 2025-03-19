package main

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func main() {

	m1 := make(map[string]string)
	m2 := make(map[string]string)
	m1["msg1"] = "The first message"
	m1["msg2"] = "The second message"
	m1["msg3"] = "You get the idea..."

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	/*Put some stuff in Redis*/
	for k, v := range m1 {
		err := rdb.Set(ctx, k, v, 0).Err()
		if err != nil {
			slog.Error("redis error", "error string", err.Error())
			return
		}
	}

	for k, _ := range m1 {

		val, err := rdb.Get(ctx, k).Result()
		if err != nil {
			slog.Error("redis error", "error string", err.Error())
			return
		}
		m2[k] = val
	}

	/* output will be unsorted */
	for k, v := range m2 {
		slog.Info("redis get", "key", k, "value", v)
	}
}
