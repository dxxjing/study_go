package logger
//文件日志系统
import (
	"fmt"
	"os"
	"path"
)
//控制台日志  输出到控制台
type ConsoleLog struct{
	level int			//日志级别
}
//构造函数 返回接口 LoggerInterface
func NewConsoleLog() LoggerInterface {
	logger := &ConsoleLog{
	}
	return logger
}

//整合信息,写入日志文件
func (f *ConsoleLog) writeLog(format string,args ...interface{}){
	nowTime := getNowDate()
	levelName := getLogLevelName(f.level)
	fileName,funcName,lineNo := getCallerInfo(2)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format,args...)
	logMsg := fmt.Sprintf("%s %s %s:%s[line:%d] :%s\n",nowTime,levelName,fileName,funcName,lineNo,msg)
	fmt.Fprintf(os.Stdout,logMsg)
}

func (f *ConsoleLog) SetLevel(logLevel int){
	if logLevel < LogLevelDebug || logLevel > LogLevelFatal {
		logLevel = LogLevelDebug
	}
	f.level = logLevel
}

func (f *ConsoleLog) LogDebug(format string,args ...interface{}){
	f.SetLevel(LogLevelDebug)
	f.writeLog(format,args...)
}

func (f *ConsoleLog) LogTrace(format string,args ...interface{}){
	f.SetLevel(LogLevelTrace)
	f.writeLog(format,args...)
}
func (f *ConsoleLog) LogInfo(format string,args ...interface{}){
	f.SetLevel(LogLevelInfo)
	f.writeLog(format,args...)
}
func (f *ConsoleLog) LogWarn(format string,args ...interface{}){
	f.SetLevel(LogLevelWarn)
	f.writeLog(format,args...)
}
func (f *ConsoleLog) LogError(format string,args ...interface{}){
	f.SetLevel(LogLevelError)
	f.writeLog(format,args...)
}
func (f *ConsoleLog) LogFatal(format string,args ...interface{}){
	f.SetLevel(LogLevelFatal)
	f.writeLog(format,args...)
}



