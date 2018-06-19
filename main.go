package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/brave/scproxy/scproxy"
	"github.com/go-redis/redis"
	"github.com/tidwall/redcon"
)

var listen = ":6379"

func main() {
	var mu sync.RWMutex

	var client scproxy.Client
	var remote_url *url.URL
	var local_url *url.URL
	var err error

	if remote_url, err = remote_url.Parse(os.Getenv("SCPROXY_BACKEND")); err != nil {
		log.Fatal("SCPROXY_BACKEND not set")
	}

	if remote_url.Scheme == "scproxy" {
		if local_url, err = local_url.Parse(os.Getenv("SCPROXY_LOCAL_BACKEND")); err != nil {
			log.Fatal("SCPROXY_BACKEND set to use scproxy and SCPROXY_LOCAL_BACKEND not set")
		}
	}

	var remote_host string
	if remote_url.Port() == "" {
		remote_host = fmt.Sprintf("%s:%s", remote_url.Hostname(), remote_url.Port())
	} else {
		remote_host = remote_url.Host
	}

	var local_host string
	if local_url.Port() == "" {
		local_host = fmt.Sprintf("%s:%s", local_url.Hostname(), local_url.Port())
	} else {
		local_host = local_url.Host
	}

	if s := remote_url.Scheme; s == "redis" {
		var db int
		if i, err := strconv.Atoi(remote_url.Path); err == nil {
			db = i
		} else {
			db = 0
		}

		pass, _ := remote_url.User.Password()
		client = redis.NewClient(&redis.Options{
			Addr:     remote_url.Host,
			Password: pass,
			DB:       db,
		})
	} else if s == "scproxy" {
		client = scproxy.NewCachingClient(local_host, remote_host)
	} else {
		log.Fatal("SCPROXY_BACKEND scheme must be redis or scproxy")
	}

	err = client.Ping().Err()
	if err != nil {
		log.Println("Could not connect to backend")
		log.Fatal(err)
	}

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
			case "flushall":
				if err = client.FlushAll().Err(); err != nil {
					conn.WriteError(fmt.Sprint(err))
				}
				conn.WriteString("OK")
			case "keys":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				res := client.Keys(string(cmd.Args[1]))
				mu.RUnlock()
				if err := res.Err(); err == redis.Nil {
					conn.WriteNull()
				} else if err != nil {
					panic(err)
				} else {
					r := res.Val()
					conn.WriteArray(len(r))
					for _, v := range r {
						conn.WriteBulkString(v)
					}
				}
			case "del":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				res := client.Del(string(cmd.Args[1]))
				mu.RUnlock()
				if err := res.Err(); err == redis.Nil {
					conn.WriteNull()
				} else if err != nil {
					panic(err)
				} else {
					conn.WriteInt64(res.Val())
				}
			case "get":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				res := client.Get(string(cmd.Args[1]))
				mu.RUnlock()
				if err := res.Err(); err == redis.Nil {
					conn.WriteNull()
				} else if err != nil {
					panic(err)
				} else {
					conn.WriteBulk([]byte(res.Val()))
				}
			case "set":
				if 3 < len(cmd.Args) || len(cmd.Args) > 4 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}

				var expire time.Duration
				if len(cmd.Args) > 3 {
					sec, err := strconv.Atoi(string(cmd.Args[3]))
					if err != nil {
						conn.WriteError("ERR argument 3 should be a number")
						return
					}
					expire = time.Second * time.Duration(sec)
				}

				mu.RLock()
				res := client.Set(string(cmd.Args[1]), cmd.Args[2], expire)
				mu.RUnlock()
				if err := res.Err(); err == redis.Nil {
					conn.WriteNull()
				} else if err != nil {
					panic(err)
				} else {
					conn.WriteBulk([]byte("OK"))
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
