package squirrel

import (
	"errors"
	"fmt"
)

type LogLevel = uint8

// Log Level
const (
	Trace   LogLevel = 0
	Debug   LogLevel = 1
	Info    LogLevel = 2
	Warning LogLevel = 3
	Error   LogLevel = 4
	Fatal   LogLevel = 5
	Panic   LogLevel = 6
)

func LevelToString(level LogLevel) string {
	switch level {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	case Panic:
		return "panic"
	default:
		return "unknown"
	}
}

func LevelToNumber(level string) (LogLevel, error) {
	switch level {
	case "trace":
		return Trace, nil
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warning":
		return Warning, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	case "panic":
		return Panic, nil
	default:
		errMsg := fmt.Sprintf("unknown log level: %s", level)
		return Trace, errors.New(errMsg)
	}
}
