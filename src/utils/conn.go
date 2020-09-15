package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Connection creates a connection to the local redis cluster
func Connection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

// Subscribe subscribes to the topic of your choice.
// It pulls in a client automatically or one can be provided
func Subscribe(topic string, client *redis.Client) redis.Client {
	sub := client.Subscribe(ctx)
	iface, err := sub.Receive(ctx)
	if err != nil {
		panic(err)
	}
	return iface.(redis.Client)
}
