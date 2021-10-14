package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
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
		fileObj, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("open file failed %s", err)
			return
		}
		fileInfo, err := fileObj.Stat()
		if err != nil {
			fmt.Printf("get open file info failed %s", err)
			return
		}
		fileSize := fileInfo.Size()
		if fileSize >= f.maxFileSize {
			newName := path.Join(f.FilePath, f.fileName)
			newName += "." + time.Now().Format("20060102150304")
			oldName := path.Join(f.FilePath, f.fileName)
			fmt.Println(oldName, newName)
			f.fileObj.Close()
			err := os.Rename(oldName, newName)
			if err != nil {
				fmt.Println("rename failed", err)
				return

			}
			fileObj, err := os.OpenFile(newName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
			if err != nil {
				fmt.Println("open file failed ", err)
				return
			}
			f.fileObj = fileObj

		}
		msg := fmt.Sprintf(format, arg...)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%d:%s] %s\n", GetNow(), levelStr, fucName, line, path.Base(fileName), msg)
		if lv >= ERROR {
			//fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%d:%s] %s\n", GetNow(), levelStr, fucName, line, path.Base(fileName), msg)
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
