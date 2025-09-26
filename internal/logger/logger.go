package logger

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"
)

var (
	Log     *LogWrapper
	LogFile *os.File
)

type LogWrapper struct {
	*log.Logger
}

func (lw *LogWrapper) Println(v ...any) {
	lw.Info(fmt.Sprint(v...))
}

func (lw *LogWrapper) Printf(format string, v ...any) {
	lw.Infof(format, v...)
}

// NewLogger creates a new logger instance
func NewLogger() *LogWrapper {
	var err error
	LogFile, err = os.OpenFile("myapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to open log file: %w", err))
	}

	logger := log.NewWithOptions(
		LogFile,
		log.Options{
			Level: log.DebugLevel,
		},
	)

	return &LogWrapper{Logger: logger}
}

// Close closes the log file
func Close() error {
	return LogFile.Close()
}

