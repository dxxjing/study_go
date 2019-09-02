package logger
//文件日志系统
import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLog struct{
	level int			//日志级别
	file *os.File		//存储 warn及warn以下的日志
	errFile *os.File	//保存 error及error以上的的日志
	chanData chan *LogChanData //channel 存储指针提升性能


}
//构造函数 返回接口 LoggerInterface
func NewFileLog() LoggerInterface {
	logger := &FileLog{
		chanData : make(chan *LogChanData,MaxChanSize),
	}

	return logger
}

func (f *FileLog) init(fileName string){
	//暂时将日子存放此目录下
	// todo 根据调用者目录 自动生成log 层级目录 e:/study_go/src/listen17/testFileLog.go => e:/study_go/log/listen17/testFileLog.go_20180903.log
	logPath := LogPath //可以写在配置中
	_ = os.MkdirAll(logPath,0777)
	fileName = path.Base(fileName)
	//如果是非错误日志则根据调用文件创建log存储文件 fileName_20180903.log
	if f.level >= LogLevelDebug && f.level <= LogLevelWarn {
		tmpLogFile := fileName + "_" + time.Now().Format(LogDateFormat) + ".log"
		tmpLogFile = fmt.Sprintf("%s%s",logPath,tmpLogFile)
		logFile,err := os.OpenFile(tmpLogFile,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
		if err != nil {
			panic(fmt.Sprintf("open %s failed,err:%v",tmpLogFile,err))
		}
		f.file = logFile
	}

	logFile,err := os.OpenFile(LogErrorFilePath,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
	if err != nil {
		panic(fmt.Sprintf("open %s failed,err:%v",LogErrorFilePath,err))
	}
	f.errFile = logFile
	go f.writeLog()
}

func (f *FileLog) insertLogDataToChan(format string,args ...interface{}){
	nowTime := getNowDate()
	levelName := getLogLevelName(f.level)
	fileName,funcName,lineNo := getCallerInfo(3)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format,args...)

	f.init(fileName)

	//组装数据 塞进channel
	logData := &LogChanData{
		Msg : msg,
		TimeStr : nowTime,
		LevelName:levelName,
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
		fmt.Println(data.Msg)
		logFile := f.file
		if f.level == LogLevelError || f.level == LogLevelFatal {
			logFile = f.errFile
		}

		fmt.Fprintf(logFile, "%s %s %s:%s[line:%d] :%s\n",data.TimeStr,data.LevelName,data.FileName,data.FuncName,data.LineNo,data.Msg)
	}
}

func (f *FileLog) SetLevel(logLevel int){
	if logLevel < LogLevelDebug || logLevel > LogLevelFatal {
		logLevel = LogLevelDebug
	}
	f.level = logLevel
}

func (f *FileLog) LogDebug(format string,args ...interface{}){
	f.SetLevel(LogLevelDebug)
	f.insertLogDataToChan(format,args...)
}

func (f *FileLog) LogTrace(format string,args ...interface{}){
	f.SetLevel(LogLevelTrace)
	f.insertLogDataToChan(format,args...)
}
func (f *FileLog) LogInfo(format string,args ...interface{}){
	f.SetLevel(LogLevelInfo)
	f.insertLogDataToChan(format,args...)
}
func (f *FileLog) LogWarn(format string,args ...interface{}){
	f.SetLevel(LogLevelWarn)
	f.insertLogDataToChan(format,args...)
}
func (f *FileLog) LogError(format string,args ...interface{}){
	f.SetLevel(LogLevelError)
	f.insertLogDataToChan(format,args...)
}
func (f *FileLog) LogFatal(format string,args ...interface{}){
	f.SetLevel(LogLevelFatal)
	f.insertLogDataToChan(format,args...)
}

func (f *FileLog) Close(){
	f.file.Close()
	f.errFile.Close()
}

