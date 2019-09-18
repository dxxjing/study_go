package main

import "logger"

func main(){

	//logger.InitLog("file")
	logger.InitLog("console")

	logger.Debug("user id[%d] is come from china",123)
	logger.Error("test error")
	logger.Warn("test warn log")
	logger.Fatal("test fatal")

}
