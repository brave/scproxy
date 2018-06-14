package scproxy

import (
	"time"

	"github.com/go-redis/redis"
)

type Client interface {
	Ping() *redis.StatusCmd
	FlushAll() *redis.StatusCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	Get(string) *redis.StringCmd
	Del(...string) *redis.IntCmd
	Keys(string) *redis.StringSliceCmd
}
