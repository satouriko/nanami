package main

import (
	"github.com/hudson6666/nanami/handler"
	"time"
	"github.com/hudson6666/nanami/database"
)

func main() {
	database.Init()
	handler.Init()
	for {
		time.Sleep(time.Second)
	}
}