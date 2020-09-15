package main

import (
	"github.com/jfabdo/speak-easy-go-k8s-pods/src/utils"
)

func main() {
	utils.Subscribe("ingress", utils.Connection())
}
