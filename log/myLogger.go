package log

import (
	"errors"
	"strings"
	"time"
)

const (
	UNKNOWN Loglevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func ParseLogLevel(level string) (Loglevel, error) {
	s := strings.ToUpper(level)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("未知道的level类型")

	}
}

func UnParseLogLevel(level Loglevel) (string, error) {
	switch level {
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARNING:
		return "WARNING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		return "UNKNOWN", errors.New("未知道的level类型")

	}
}

func GetNow() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
