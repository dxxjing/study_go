package logger
//文件日志系统
import (
	"fmt"
	"os"
	"path"
	"time"
)

//todo 该文件日志 有个问题：当主程序退出，写日志到文件的协程也被强制退出，导致日志不能写完

type FileLog struct{
	//level int			//日志级别  废弃原因：多协程并发执行 最后的SetLevel 会覆盖前面的，导致前面的不能正确取值
	//由上得知：成员变量file 也会有覆盖的问题 所以将init方法放在 写日志到文件的协程中调用来规避,
	//file 成员变量可有可无
	file *os.File		//存储 warn及warn以下的日志
	errFile *os.File	//保存 error及error以上的的日志
	chanData chan *LogChanData //channel 存储指针提升性能


}
//构造函数 返回接口 LoggerInterface
func NewFileLog() LoggerInterface {
	logger := &FileLog{
		chanData : make(chan *LogChanData,MaxChanSize),
	}
	//启动一个协程 从chanData中拿数据并将日志写入文件
	go logger.writeLog()

	return logger
}

func (f *FileLog) init(fileName string,level int){
	//暂时将日子存放此目录下
	// todo 根据调用者目录 自动生成log 层级目录 e:/study_go/src/listen17/testFileLog.go => e:/study_go/log/listen17/testFileLog.go_20180903.log
	logPath := LogPath //可以写在配置中
	_ = os.MkdirAll(logPath,0777)
	fileName = path.Base(fileName)
	//如果是非错误日志则根据调用文件创建log存储文件 fileName_20180903.log

	if level >= LogLevelDebug && level <= LogLevelWarn {
		tmpLogFile := fileName + "_" + time.Now().Format(LogDateFormat) + ".log"
		tmpLogFile = fmt.Sprintf("%s%s",logPath,tmpLogFile)
		logFile,err := os.OpenFile(tmpLogFile,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
		if err != nil {
			panic(fmt.Sprintf("open %s failed,err:%v",tmpLogFile,err))
		}
		//打开已打开的文件不会报错
		//logFile,err = os.OpenFile(tmpLogFile,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
		f.file = logFile
	}

	errLogFile,err := os.OpenFile(LogErrorFilePath,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
	if err != nil {
		panic(fmt.Sprintf("open %s failed,err:%v",LogErrorFilePath,err))
	}
	f.errFile = errLogFile
}

func (f *FileLog) insertLogDataToChan(level int,format string,args ...interface{}){
	nowTime := getNowDate()
	fileName,funcName,lineNo := getCallerInfo(3)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format,args...)
	//组装数据 塞进channel
	logData := &LogChanData{
		Msg : msg,
		TimeStr : nowTime,
		Level:level,
		FileName:fileName,
		FuncName:funcName,
		LineNo:lineNo,
	}

	//log数据塞进channel 用select 防止到达MaxChanSize后阻塞
	select{
	case f.chanData <- logData:
		//fmt.Println("success:",*logData)
	default:
		//fmt.Println("failed:",*logData)
	}
}

//整合信息,写入日志文件
func (f *FileLog) writeLog(){
	for data := range f.chanData {
		f.init(data.FileName,data.Level)
		logFile := f.file
		if data.Level == LogLevelError || data.Level == LogLevelFatal {
			logFile = f.errFile
		}

		LevelName := getLogLevelName(data.Level)
		fmt.Fprintf(logFile, "%s %s %s:%s[line:%d] :%s\n",data.TimeStr,LevelName,data.FileName,data.FuncName,data.LineNo,data.Msg)

		f.Close()
	}
}

//注意：当多协程并发调用 该level 会被最后一个给覆盖
/*func (f *FileLog) SetLevel(logLevel int){
	if logLevel < LogLevelDebug || logLevel > LogLevelFatal {
		logLevel = LogLevelDebug
	}
	f.level = logLevel
}*/

func (f *FileLog) LogDebug(format string,args ...interface{}){
	f.insertLogDataToChan(LogLevelDebug,format,args...)
}

func (f *FileLog) LogTrace(format string,args ...interface{}){
	f.insertLogDataToChan(LogLevelTrace,format,args...)
}
func (f *FileLog) LogInfo(format string,args ...interface{}){
	f.insertLogDataToChan(LogLevelInfo,format,args...)
}
func (f *FileLog) LogWarn(format string,args ...interface{}){
	f.insertLogDataToChan(LogLevelWarn,format,args...)
}
func (f *FileLog) LogError(format string,args ...interface{}){
	f.insertLogDataToChan(LogLevelError,format,args...)
}
func (f *FileLog) LogFatal(format string,args ...interface{}){
	f.insertLogDataToChan(LogLevelFatal,format,args...)
}

func (f *FileLog) Close(){
	f.file.Close()
	f.errFile.Close()
}

