package logger

//单元测试 直接go test 就会运行
import (
	"testing"
)


func TestLogger(t *testing.T){
	//InitLog("file")
	InitLog("console")
	LogDebug("logger->user id[%d] is come from china",123)
}
