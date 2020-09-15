package utils

import (
	"log"
	"net"

	"github.com/gomodule/redigo/redis"
)

// Connection creates a connection to the local redis cluster
func Connection() net.Conn {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}


// Subscribe subscribes to the topic of your choice. 
// It pulls in a client automatically or one can be provided
func Subscribe(topic string, client net.Conn = Connection()) client.Channel {
	sub := client.Subscribe(queryResp)
	iface, err := sub.Receive()
	ch := sub.Channel()
	return ch
}