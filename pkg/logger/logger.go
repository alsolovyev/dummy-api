// Package implements logging functionality
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger

	defaultFormat = log.LstdFlags
	defaultOutput = os.Stdout
)

type Logger struct{}

func New() *Logger {
	infoLog = log.New(defaultOutput, "INFO: ", defaultFormat)
	warningLog = log.New(defaultOutput, "WARNING: ", defaultFormat)
	errorLog = log.New(defaultOutput, "ERROR: ", defaultFormat)

	return &Logger{}
}

func SetOutput(o io.Writer) {
	infoLog.SetOutput(o)
	warningLog.SetOutput(o)
	errorLog.SetOutput(o)
}

func SetFlags(f int) {
	infoLog.SetFlags(f)
	warningLog.SetFlags(f)
	errorLog.SetFlags(f)
}

// Info logs a message with info level
func (*Logger) Info(message string) {
	infoLog.Println(message)
}

// Warning logs a message with warning level
func (*Logger) Warning(message string) {
	warningLog.Println(message)
}

// Error logs a message with error level
func (*Logger) Error(message string) {
	errorLog.Println(message)
}

// Infof logs a formatted message with info level using the default format
func (*Logger) Infof(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	infoLog.Printf("%s\n", message)
}

// Warningf logs a formatted message with warning level using the default format
func (*Logger) Warningf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	warningLog.Printf("%s\n", message)
}

// Errorf logs a formatted message with error level using the default format
func (*Logger) Errorf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	errorLog.Printf("%s\n", message)
}
