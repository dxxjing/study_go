package logger

//日志级别常量
const (
	LogLevelDebug = iota
	LogLevelTrace
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

const (
	//日志时间格式模板
	LogFormat = "2006-01-02 15:04:05"
)
