package logger

//日志接口
type LoggerInterface interface{
	SetLevel(level int) // 设置日志级别

	LogDebug(format string,args ...interface{})
	LogTrace(format string,args ...interface{})
	LogInfo(format string,args ...interface{})
	LogWarn(format string,args ...interface{})
	LogError(format string,args ...interface{})
	LogFatal(format string,args ...interface{})
}
