package main

import (
	"logger"
	"os"
)
//使用filelog 日志
func main(){
	logPath := "e:/study_go/log/"
	_ = os.MkdirAll(logPath,0777)
	logName := "file_test"
	fileLogger := logger.NewFileLog(logPath,logName)

	fileLogger.LogDebug("user id[%d] is come from china",123)
	fileLogger.LogError("test error")
	fileLogger.LogWarn("test warn log")
	fileLogger.LogFatal("test fatal")
}
