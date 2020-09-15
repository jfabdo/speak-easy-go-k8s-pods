package utils

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
)

//PostMessage is the format for committing to the redis post queue

type PostMessage struct {
	BID     string              `redis:bid`     //Browser ID
	LID     string              `redis:lid`     //Login ID
	SID     string              `redis:sid`     //Session ID
	command string              `redis:command` //What this message is supposed to do, e.g. report, post comment, etc.
	comment string              `redis:comment` //What this message is acting on, i.e. the URI
	body    string              `redis:body`    //The comment, report, etc.
	writer  http.ResponseWriter //Where information is being returned to
}

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
