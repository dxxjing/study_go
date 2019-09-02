package logger

import (
	"runtime"
	"time"
)

//获取当前时间格式 2019-09-02 09:34:45
func getNowDate() string {
	return time.Now().Format(LogFormat)
}

//获取日志级别对用的名字
func getLogLevelName(level int) string {
	var logLevelName string
	switch level {
	case LogLevelDebug:
		logLevelName = "LogDebug"
	case LogLevelTrace:
		logLevelName = "LogTrace"
	case LogLevelInfo:
		logLevelName = "LogInfo"
	case LogLevelWarn:
		logLevelName = "LogWarn"
	case LogLevelError:
		logLevelName = "LogError"
	case LogLevelFatal:
		logLevelName = "LogFatal"
	default:
		logLevelName = "UNKNOW"
	}
	return logLevelName
}

//获取调用者信息 包括 文件名 函数名 行号
//skip :0 代表当前函数，也是调用runtime.Caller的函数。1 代表上一层调用者，以此类推。
//F:/study_go/src/logger/fileLog_test.go
//logger.TestFileLogger
func getCallerInfo(skip int)(fileName string,funcName string,lineNo int){

	pc,file,line,ok := runtime.Caller(skip)
	if ok{
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}
