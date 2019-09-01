package logger

//单元测试 直接go test 就会运行
import (
	"os"
	"testing"
)


func TestFileLogger(t *testing.T){
	logPath := "e:/study_go/log/"
	_ = os.MkdirAll(logPath,0777)
	logName := "file_test"
	fileLogger := NewFileLog(LogLevelDebug,logPath,logName)

	fileLogger.LogDebug("user id[%d] is come from china",123)
	fileLogger.LogError("test error")
	fileLogger.LogWarn("test warn log")
	fileLogger.LogFatal("test fatal")
}

