package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()

type cache struct {
}

func newCache() *cache {
	return &cache{}
}

var Cache = newCache()

func (c *cache) HSet(cmd redis.Cmdable, key string, values interface{}, ttl int) (err error) {
	err = cmd.HSet(ctx, key, values).Err()
	if err != nil {
		return err
	}
	if ttl > 0 {
		err = cmd.Expire(ctx, key, time.Second*time.Duration(ttl)).Err()
		if err != nil {
			return err
		}
	}
	return
}

func (c *cache) GetValue(cmd redis.Cmdable, key string) (string, error) {
	rowData, err := cmd.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return rowData, nil
}

func (c *cache) HGetAll(cmd redis.Cmdable, key string) (map[string]string, error) {
	res, err := cmd.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *cache) HGetValue(cmd redis.Cmdable, key, field string) (string, error) {
	res, err := cmd.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return res, nil
}
