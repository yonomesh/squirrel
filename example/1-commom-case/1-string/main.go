package main

import (
	"fmt"
	"os"
	"testsq/log"

	sq "github.com/yonomesh/squirrel"
)

func main() {
	logFile, err := os.OpenFile("my.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, sq.Info, "test", []string{}, "")

	logger.Print("Hello World")
	logger.Print("Hello Squirrel")
	logger.Print("Hello Everyone")
	fmt.Println("OK")
}
