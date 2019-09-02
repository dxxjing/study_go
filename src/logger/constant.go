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
	LogTimeFormat = "2006-01-02 15:04:05"
	LogDateFormat = "20060102"

	LogPath = "e:/study_go/log/"
	LogErrorFilePath = LogPath + "err.log"

	MaxChanSize = 50000
)
