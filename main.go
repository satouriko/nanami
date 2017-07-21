package main

import (
	"github.com/hudson6666/nanami/handler"
	"time"
)

func main() {
	handler.Init()
	for {
		time.Sleep(time.Second)
	}
}