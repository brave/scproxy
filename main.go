package main

import (
	"log"
	"strings"
	"sync"
	"os"
	"fmt"
	"net/url"
	"strconv"

	"github.com/tidwall/redcon"
	"github.com/RyanJarv/scproxy/scproxy"
	"github.com/go-redis/redis"
)

var listen = ":6379"

func main() {
	var mu sync.RWMutex

	var client scproxy.Client

	if u, err := url.Parse(os.Getenv("SCPROXY_BACKEND")); err == nil {
		var db int
		if i, err := strconv.Atoi(u.Path); err == nil {
			db = i
		} else {
			db = 0
		}

		pass, _ := u.User.Password()
		client = redis.NewClient(&redis.Options{
			Addr:     u.Host,
			Password: pass,
			DB:       db,
		})
	} else {
		client = scproxy.NewCachingClient("local_redis", "remote_redis")
	}

	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("Could not connect to backend")
		log.Fatal(err)
	}
	fmt.Println(pong)

	go log.Printf("started server at %s", listen)
	err = redcon.ListenAndServe(listen,
		func(conn redcon.Conn, cmd redcon.Command) {
			fmt.Sprint(cmd.Args)
			switch strings.ToLower(string(cmd.Args[0])) {
			default:
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
			case "ping":
				if err = client.Ping().Err(); err != nil {
					conn.WriteError(fmt.Sprint(err))
				}
				conn.WriteString("PONG")
			case "get":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				res := client.Get(string(cmd.Args[1]))
				mu.RUnlock()
				if err := res.Err(); err != nil {
					if err != redis.Nil {
						conn.WriteNull()
					} else {
						panic(err)
					}
				} else {
					conn.WriteBulk([]byte(res.Val()))
				}
			}
		},
		func(conn redcon.Conn) bool {
			// use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// this is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}