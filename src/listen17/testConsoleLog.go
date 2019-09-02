package main

import (
	"logger"
	"time"
)

func main(){

	log := logger.NewConsoleLog()

	for{
		log.LogDebug("user id[%d] is come from china",123)
		time.Sleep(time.Second)
	}
}
