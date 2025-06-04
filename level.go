package squirrel

import (
	"errors"
	"fmt"
)

type LogLevel = uint8

// Log Level
const (
	Trace   LogLevel = 1 << iota // 1   0000001
	Debug                        // 2   0000010
	Info                         // 4   0000100
	Warning                      // 8   0001000
	Error                        // 16  0010000
	Fatal                        // 32  0100000
	Panic                        // 64  1000000

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
