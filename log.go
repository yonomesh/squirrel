package squirrel

import (
	"runtime"
	"strconv"
)

// Log represents the log data format.
type LogEntry struct {
	Time     string   `json:"ts"`       // Timestamp of the log entry
	Level    string   `json:"level"`    // Log level (e.g., Trace, Debug, Info, Warning, Error, Fataland Panic)
	Category string   `json:"category"` // Category or type of the log (e.g., user-action)
	Tags     []string `json:"tags"`     // Tags related to the log
	Msg      Msger    `json:"msg"`      // Msg content, implemented via the interface for customization
	Extra    Extra    `json:"extra"`    // Ectra content, implemented via the interface for customization
}

type Msger interface {
	MsgToString() (string, error)
}

type Extra interface {
	ExtraToString() (string, error)
}

func GetCallFullPosition(pc uintptr, call_depth int) string {
	var file string
	var line int
	if pc == 0 {
		var ok bool
		_, file, line, ok = runtime.Caller(call_depth)
		if !ok {
			file = "???"
			line = 0
		}
	} else {
		fs := runtime.CallersFrames([]uintptr{pc})
		f, _ := fs.Next()
		file = f.File
		if file == "" {
			file = "???"
		}
		line = f.Line
	}
	return file + ":" + strconv.Itoa(line)
}

func GetCallShortPosition(pc uintptr, call_depth int) string {
	var file string
	var line int
	if pc == 0 {
		var ok bool
		_, file, line, ok = runtime.Caller(call_depth)
		if !ok {
			file = "???"
			line = 0
		}
	} else {
		fs := runtime.CallersFrames([]uintptr{pc})
		f, _ := fs.Next()
		file = f.File
		if file == "" {
			file = "???"
		}
		line = f.Line
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file + ":" + strconv.Itoa(line)
}
