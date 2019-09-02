package logger

//单元测试 直接go test 就会运行
import (
	"testing"
)


func TestFileLogger(t *testing.T){
	log := NewFileLog()

	log.LogDebug("user id[%d] is come from china",123)
	log.LogError("test error")
	log.LogWarn("test warn log")
	log.LogFatal("test fatal")
}

