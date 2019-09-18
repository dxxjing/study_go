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

func Debug(format string,args ...interface{}){
	log.LogDebug(format,args...)
}

func Trace(format string,args ...interface{}){
	log.LogTrace(format,args...)
}

func Info(format string,args ...interface{}){
	log.LogInfo(format,args...)
}

func Warn(format string,args ...interface{}){
	log.LogWarn(format,args...)
}

func Error(format string,args ...interface{}){
	log.LogError(format,args...)
}

func Fatal(format string,args ...interface{}){
	log.LogFatal(format,args...)
}
