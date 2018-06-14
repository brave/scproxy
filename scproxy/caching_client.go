package scproxy

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func NewCachingClient(local, remote string) *CachingClient {
	client := &CachingClient{local_host: local, remote_host: remote}
	client.Connect()
	return client
}

type CachingClient struct {
	local_host    string
	remote_host   string
	local_client  *redis.Client
	remote_client *redis.Client
}

func (s *CachingClient) Connect() {
	s.local_client = redis.NewClient(&redis.Options{
		Addr:     s.local_host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	s.remote_client = redis.NewClient(&redis.Options{
		Addr:     s.remote_host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := s.local_client.Ping().Result(); err != nil {
		log.Fatal(fmt.Sprintf("Could not ping local redis instance at %s: %s", s.local_host, err))
	} else {
		log.Printf("Connected to local cache at %s", s.local_host)
	}

	if _, err := s.remote_client.Ping().Result(); err != nil {
		log.Fatal(fmt.Sprintf("Could not ping remote redis instance at %s: %s", s.remote_host, err))
	} else {
		log.Printf("Connected to remote cache at %s", s.remote_host)
	}
}

func (s *CachingClient) Ping() *redis.StatusCmd {
	return s.remote_client.Ping()
}

func (s *CachingClient) Set(key string, args interface{}, time time.Duration) *redis.StatusCmd {
	return s.local_client.Set(key, args, time)
}

func (s *CachingClient) Get(key string) *redis.StringCmd {
	res := s.local_client.Get(key)
	if _, err := res.Result(); err != redis.Nil {
		if err != nil {
			log.Printf("Error running get on local cache")
		} else {
			return res
		}
	}
	return s.remote_client.Get(key)
}

func (s *CachingClient) Del(key ...string) *redis.IntCmd {
	return s.local_client.Del(key...)
}

func (s *CachingClient) FlushAll() *redis.StatusCmd {
	return s.local_client.FlushAll()
}

func (s *CachingClient) Keys(pattern string) *redis.StringSliceCmd {
	return s.local_client.Keys(pattern)
}
