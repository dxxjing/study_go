package logger

var log LoggerInterface

//日志库 包含文件日志 和 控制台日志

func InitLog(logType string){
	switch logType {
	case "file":
		log = NewFileLog()
	case "console":
		log = NewConsoleLog()
	}
}

//由于全局变量log 为私有的 定义方法

func LogDebug(format string,args ...interface{}){
	log.LogDebug(format,args...)
}

func LogTrace(format string,args ...interface{}){
	log.LogTrace(format,args...)
}

func LogInfo(format string,args ...interface{}){
	log.LogInfo(format,args...)
}

func LogWarn(format string,args ...interface{}){
	log.LogWarn(format,args...)
}

func LogError(format string,args ...interface{}){
	log.LogError(format,args...)
}

func LogFatal(format string,args ...interface{}){
	log.LogFatal(format,args...)
}
