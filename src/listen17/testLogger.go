package main

import "logger"

func main(){

	//logger.InitLog("file")
	logger.InitLog("console")

	logger.LogDebug("user id[%d] is come from china",123)
	logger.LogError("test error")
	logger.LogWarn("test warn log")
	logger.LogFatal("test fatal")

}
