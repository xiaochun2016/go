package log

import (
	"fmt"
	"path"
	"runtime"
)

type Loglevel uint16

type Logger struct {
	level Loglevel
}

func NewFileLog(levelStr string) Logger {
	level, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: level,
	}
}
func (l Logger) log(lv Loglevel, format string, arg ...interface{}) {
	if l.level <= lv {
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
		fmt.Printf("[%s] [%s] [%s:%d:%s] %s\n", GetNow(), levelStr, fucName, line, path.Base(fileName), msg)
	}
}

func (l Logger) Debug(format string, arg ...interface{}) {
	l.log(DEBUG, format, arg...)
}

func (l Logger) Info(format string, arg ...interface{}) {
	l.log(INFO, format, arg...)
}

func (l Logger) Warning(format string, arg ...interface{}) {
	l.log(WARNING, format, arg...)
}

func (l Logger) Error(format string, arg ...interface{}) {
	l.log(ERROR, format, arg...)
}

func (l Logger) Fatal(format string, arg ...interface{}) {
	l.log(FATAL, format, arg...)
}
