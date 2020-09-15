package main

import (
	"net/http"

	"github.com/jfabdo/speak-easy-go-k8s-pods/utils"
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

func main() {
	utils.Subscribe("ingress")
}
