package squirrel

import (
	"time"
)

// Log Level
const (
	Trace   = "trace"
	Debug   = "debug"
	Info    = "info"
	Warning = "warning"
	Error   = "error"
	Fatal   = "fatal"
	Panic   = "painc"
)

// Log represents the log data format.
type Log struct {
	Time     time.Time `json:"ts"`       // Timestamp of the log entry
	Level    string    `json:"level"`    // Log level (e.g., Trace, Debug, Info, Warning, Error, Fataland Panic)
	Category string    `json:"category"` // Category or type of the log (e.g., user-action)
	Tags     []string  `json:"tags"`     // Tags related to the log
	Msg      Msger     `json:"msg"`      // Msg content, implemented via the interface for customization
	Extra    Extra     `json:"extra"`    // Ectra content, implemented via the interface for customization
}

type Msger interface {
	MsgToString() (string, error)
}

type Extra interface {
	EctraToString() (string, error)
}
