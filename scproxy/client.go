package scproxy

import	(
	"log"
	"fmt"
	"github.com/go-redis/redis"
)

func NewClient() *Client {
	return &Client{local_host: "redis_local", remote_host: "redis_remote"}
}

type Client struct {
	local_host string
	remote_host string
	local_client *redis.Client
	remote_client *redis.Client
}

func (s *Client) Connect() {
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

func (s *Client) Set(key, value string) error {
	err := s.local_client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Client) Get(key string) (val string, err error) {
	val, err = s.local_client.Get(key).Result()
	if (err != nil && err != redis.Nil) {
		return "", err
	} else if (err != redis.Nil) {
		return val, nil
	}

	log.Printf("[DEBUG] '%s' not found in local cache", key)

	val, err = s.remote_client.Get(key).Result()
	if err != nil {
		return "", err
	} 

	s.Set(key, val)
	if err == redis.Nil {
		log.Printf("[DEBUG] '%s' not found in remote cache", key)
	} else if err != nil {
		return "", err
	} 

	return val, nil
}