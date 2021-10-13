package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

type FileLog struct {
	level       Loglevel
	fileName    string
	FilePath    string
	errFileName string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLog {
	loglevel, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLog{
		level:       loglevel,
		fileName:    fn,
		FilePath:    fp,
		maxFileSize: maxSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLog) initFile() (err error) {
	allFile := path.Join(f.FilePath, f.fileName)
	fileObj, err := os.OpenFile(allFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open log file failed %v\n", err)
		return err
	}
	allErrFile := path.Join(f.FilePath, f.fileName)
	errFileObj, err := os.OpenFile(allErrFile+".err", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open err log file failed %v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f FileLog) log(lv Loglevel, format string, arg ...interface{}) {
	if f.level <= lv {
		pc, fileName, line, ok := runtime.Caller(2)
		if !ok {
			fmt.Printf("获取失败")
			return
		}
		fucName := runtime.FuncForPC(pc).Name()
		levelStr, err := UnParseLogLevel(lv)
		if err != nil {
			fmt.Printf("获取失败")
			return
		}
		msg := fmt.Sprintf(format, arg...)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%d:%s] %s\n", GetNow(), levelStr, fucName, line, path.Base(fileName), msg)
		if lv >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%d:%s] %s\n", GetNow(), levelStr, fucName, line, path.Base(fileName), msg)
		}
	}
}

func (f FileLog) Debug(format string, arg ...interface{}) {
	f.log(DEBUG, format, arg...)
}

func (f FileLog) Info(format string, arg ...interface{}) {
	f.log(INFO, format, arg...)
}

func (f FileLog) Warning(format string, arg ...interface{}) {
	f.log(WARNING, format, arg...)
}

func (f FileLog) Error(format string, arg ...interface{}) {
	f.log(ERROR, format, arg...)
}

func (f FileLog) Fatal(format string, arg ...interface{}) {
	f.log(FATAL, format, arg...)
}

func (f *FileLog) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
