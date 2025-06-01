package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Msger interface {
	ToJSON() (any, error)
}

type Log struct {
	Time     time.Time `json:"ts"`       // Timestamp of the log entry
	Level    string    `json:"level"`    // Log level (e.g., INFO, ERROR)
	Category string    `json:"category"` // Category or type of the log (e.g., user-action)
	Tags     []string  `json:"tags"`     // Tags related to the log
	Msg      Msger     `json:"msg"`      // Msg content, implemented via the interface for customization
	Extra    string    `json:"extra"`    // Ectra content, implemented via the interface for customization
}

type MyMsg struct {
	Action string `json:"action"` // The actual message content
}

func (m MyMsg) ToJSON() (any, error) {
	return map[string]string{"action": m.Action}, nil
}

// StringMsg is a new type for handling string messages directly
type StringMsg string

// ToJSON method for StringMsg to implement Msger interface
func (m StringMsg) ToJSON() (any, error) {
	return string(m), nil
}

func main() {
	logData2 := Log{
		Time:     time.Now(),
		Level:    "INFO",
		Category: "user-action",
		Tags:     []string{"login", "success"},
		Msg:      StringMsg("What are you doing"),
		Extra:    "",
	}
	logJSON2, err := json.MarshalIndent(logData2, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling log:", err)
		return
	}
	fmt.Println(string(logJSON2))

}
