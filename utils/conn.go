package utils

import (
	"log"
	"net"

	"github.com/gomodule/redigo/redis"
)

//Connection() creates a connection to the local redis cluster
func Connection() net.Conn {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
