package db

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"qiuyier/blog/pkg/config"
)

var rdb redis.Cmdable

func InitRedis(rdbConfig config.Redis) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     rdbConfig.Addr,
		Password: rdbConfig.Password,
		DB:       rdbConfig.Db,
	})

	if err = rdb.Ping(context.Background()).Err(); err != nil {
		logrus.Errorf("redis connect failed: %s", err.Error())
	}
	return
}

func Rdb() redis.Cmdable {
	return rdb
}
