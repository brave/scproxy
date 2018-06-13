package scproxy

import	(
	"time"
	"github.com/go-redis/redis"
)

type Client interface {
	Set(string, interface {}, time.Duration) *redis.StatusCmd
	Get(string) *redis.StringCmd
	Ping() *redis.StatusCmd
}