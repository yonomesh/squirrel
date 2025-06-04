package log

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	sq "github.com/yonomesh/squirrel"
)

type StringMsg string

func (m StringMsg) MsgToString() (string, error) {
	return string(m), nil
}

type StringExtra string

func (e StringExtra) ExtraToString() (string, error) {
	return string(e), nil
}

type Logger struct {
	output   io.Writer
	level    sq.LogLevel
	category string
	tags     []string
	extra    StringExtra
}

func New(output io.Writer, new_level sq.LogLevel, new_category string, new_tags []string, new_extra StringExtra) *Logger {
	return &Logger{
		output:   output,
		level:    new_level,
		category: new_category,
		tags:     new_tags,
		extra:    new_extra,
	}
}

func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}

func NewLogEntry(timestamp string, level sq.LogLevel, category string, tags []string, msg StringMsg, extra StringExtra) *sq.LogEntry {
	level_str := sq.LevelToString(level)
	return &sq.LogEntry{
		Time:     timestamp,
		Level:    level_str,
		Category: category,
		Tags:     tags,
		Msg:      msg,
		Extra:    extra,
	}
}

func (l *Logger) Print(msg StringMsg) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix()) // Unix timestamp in seconds
	log_entry := NewLogEntry(timestamp, l.level, l.category, l.tags, msg, l.extra)
	log_data, err := json.Marshal(log_entry)
	if err != nil {
		fmt.Fprintf(l.output, "Error marshalling log entry: %v, failed entry: %v\n", err, log_entry)
		return
	}
	fmt.Fprintln(l.output, string(log_data))
}
