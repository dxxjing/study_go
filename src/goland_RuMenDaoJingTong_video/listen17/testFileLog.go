package main

import "logger"

func test(){
	log := logger.NewFileLog()

	log.LogDebug("user id[%d] is come from china",123)
	log.LogError("test error")
	log.LogWarn("test warn log")
	log.LogFatal("test fatal")
}

func main(){

	test()

	//time.Sleep(3*time.Second)
}
