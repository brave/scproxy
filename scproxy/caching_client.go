package scproxy

import	(
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

func NewCachingClient(local, remote string) *CachingClient {
	client := &CachingClient{local_host: local, remote_host: remote}
	client.Connect()
  return client
}

type CachingClient struct {
	local_host string
	remote_host string
	local_client *redis.Client
	remote_client *redis.Client
}

func (s *CachingClient) Connect() {
	s.local_client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", s.local_host),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	s.remote_client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", s.remote_host),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := s.local_client.Ping().Result()
	fmt.Println(pong, err)

	pong, err = s.remote_client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func (s *CachingClient) Ping() *redis.StatusCmd {
	return s.remote_client.Ping()
}

func (s *CachingClient) Set(key string, args interface {}, time time.Duration) *redis.StatusCmd {
	return s.local_client.Set(key, args, time)
}

func (s *CachingClient) Get(key string) (*redis.StringCmd) {
	res := s.local_client.Get(key)
	if _, err := res.Result(); (err != nil) {
		return res
	}
	return s.remote_client.Get(key)
}

//func (s *CachingClient) Delete(key, value string) error {
//	err := s.local_client.Del(key).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}