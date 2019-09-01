package logger
//文件日志系统
import (
	"fmt"
	"os"
)

type FileLog struct{
	level int			//日志级别
	logPath string		//文件目录
	logName string		//文件名
	file *os.File		//存储 warn及warn以下的日志
	errFile *os.File	//保存 error及error以上的的日志


}
//构造函数 返回接口 LoggerInterface
func NewFileLog(logLevel int,logPath,logName string) LoggerInterface {
	logger := &FileLog{
		level : logLevel,
		logPath : logPath,
		logName : logName,
	}

	logger.init()

	return logger
}

func (f *FileLog) init(){
	//目录需提前创建好 os.O_CREAT只会创建不存在的文件不会创建目录
	//或者使用os.MkdirAll 进行创建
	fileName := fmt.Sprintf("%s%s.log",f.logPath,f.logName)
	file,err := os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
	if err != nil {
		panic(fmt.Sprintf("open %s failed,err:%v",fileName,err))
	}
	f.file = file
	//错误日志文件
	fileName = fmt.Sprintf("%s%s.err.log",f.logPath,f.logName)
	file,err = os.OpenFile(fileName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777)
	if err != nil {
		panic(fmt.Sprintf("open %s failed,err:%v",fileName,err))
	}
	f.errFile = file
}

func (f *FileLog) SetLevel(logLevel int){
	if logLevel < LogLevelDebug || logLevel > LogLevelFatal {
		logLevel = LogLevelDebug
	}
	f.level = logLevel
}

func (f *FileLog) LogDebug(format string,args ...interface{}){
	fmt.Fprintln(f.file,fmt.Sprintf(format,args...))
}

func (f *FileLog) LogTrace(format string,args ...interface{}){
	fmt.Fprintln(f.file,fmt.Sprintf(format,args...))
}
func (f *FileLog) LogInfo(format string,args ...interface{}){
	fmt.Fprintln(f.file,fmt.Sprintf(format,args...))
}
func (f *FileLog) LogWarn(format string,args ...interface{}){
	fmt.Fprintln(f.file,fmt.Sprintf(format,args...))
}
func (f *FileLog) LogError(format string,args ...interface{}){
	fmt.Fprintln(f.errFile,fmt.Sprintf(format,args...))
}
func (f *FileLog) LogFatal(format string,args ...interface{}){
	fmt.Fprintln(f.errFile,fmt.Sprintf(format,args...))
}