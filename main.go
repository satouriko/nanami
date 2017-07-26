package main

import (
	"time"
	"github.com/hudson6666/nanami/database"
	"github.com/hudson6666/nanami/handler"
)

func main() {
	database.Init()
	handler.Init()
	for {
		time.Sleep(time.Second)
	}
}