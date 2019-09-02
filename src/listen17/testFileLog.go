package main

import (
	"logger"
)

func main(){

	log := logger.NewFileLog()

	log.LogDebug("user id[%d] is come from china",123)
	log.LogError("test error")
	log.LogWarn("test warn log")
	log.LogFatal("test fatal")


}
