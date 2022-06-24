package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const (
	redisServiceName     string = "redis-service"
	redisServicePort     int    = 6379
	numberOfVisistsRedis string = "numberOfVisits"
)

type RedisWrapper struct {
	rdb *redis.Client
}

func GetRedisWrapper(ctx context.Context) (*RedisWrapper, error) {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%d", redisServiceName, redisServicePort),
		},
	)

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisWrapper{
		rdb: redisClient,
	}, nil
}

func (w *RedisWrapper) NumberOfVisists(ctx context.Context) (int, error) {
	val, err := w.rdb.Get(ctx, numberOfVisistsRedis).Result()
	switch err {
	case nil:
		// convert to an int and proceed with incrementing
		res, err := strconv.Atoi(val)
		if err != nil {
			return -1, err
		}

		res++
		err = w.rdb.Set(ctx, numberOfVisistsRedis, res, 0).Err()
		if err != nil {
			return -1, err
		}

		return res, nil
	case redis.Nil:
		// create a variable for the first visiting
		err = w.rdb.Set(ctx, numberOfVisistsRedis, 1, 0).Err()
		if err != nil {
			return -1, err
		}
		return 1, nil
	}
	// default
	return -1, err
}
