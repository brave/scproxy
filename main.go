package main

import (
	"log"
	"strings"
	"sync"

	"github.com/tidwall/redcon"
	"github.com/RyanJarv/scproxy/scproxy"
)

var addr = ":6379"

func main() {
	var mu sync.RWMutex
	client := scproxy.NewClient()
	client.Connect()
	go log.Printf("started server at %s", addr)
	err := redcon.ListenAndServe(addr,
		func(conn redcon.Conn, cmd redcon.Command) {
			switch strings.ToLower(string(cmd.Args[0])) {
			default:
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
			case "ping":
				conn.WriteString("PONG")
			case "quit":
				conn.WriteString("OK")
				conn.Close()
			case "set":
				if len(cmd.Args) != 3 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.Lock()
				err := client.Set(string(cmd.Args[1]), string(cmd.Args[2]))
				log.Printf("[INFO] %s", err)
				mu.Unlock()
				if err != nil {
					log.Printf("[ERROR] %s", err)
				}
				conn.WriteString("OK")
			case "get":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				val, err := client.Get(string(cmd.Args[1]))
				mu.RUnlock()
				if err != nil {
					log.Printf("[ERROR] %s", err)
					conn.WriteNull()
				} else {
					conn.WriteBulk([]byte(val))
				}
			case "del":
				//if len(cmd.Args) != 2 {
				//	conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				//	return
				//}
				//mu.Lock()
				//_, ok := items[string(cmd.Args[1])]
				//delete(items, string(cmd.Args[1]))
				//mu.Unlock()
				//if !ok {
				//	conn.WriteInt(0)
				//} else {
				//	conn.WriteInt(1)
				//}
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