package redis

import (
	"github.com/go-redis/redis"
)

func NewRedis(opts *redis.Options) *redis.Client {
	cli := redis.NewClient(opts)
	if err := cli.Ping().Err(); err != nil {
		panic(err)
	}
	return cli
}
