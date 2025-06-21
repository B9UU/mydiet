package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Log *log.Logger
var LogFile *os.File

func NewLogger() *log.Logger {
	var err error
	LogFile, err = os.OpenFile("myapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return log.NewWithOptions(
		LogFile,
		log.Options{
			ReportCaller:    true,
			ReportTimestamp: true,
			TimeFormat:      time.RFC3339,
			Prefix:          "->",
			Level:           log.DebugLevel,
		},
	)
}
