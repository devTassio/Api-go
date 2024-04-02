package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger type definition
type Logger struct {
	debug   *log.Logger // Logger for debug messages
	info    *log.Logger // Logger for information messages
	warning *log.Logger // Logger for warning messages
	err     *log.Logger // Logger for error messages
	writer  io.Writer   // Output writer for loggers
}

// NewLogger function creates a new instance of Logger
// Parameter p is the prefix for the loggers
func NewLogger(p string) *Logger {
	// Set the output writer to standard output (stdout)
	writer := io.Writer(os.Stdout)
	// Create a new logger with output writer, prefix, and formatting options
	logger := &Logger{
		debug:   log.New(writer, p+" DEBUG: ", log.Ldate|log.Ltime),   // Logger for debug messages
		info:    log.New(writer, p+" INFO: ", log.Ldate|log.Ltime),    // Logger for information messages
		warning: log.New(writer, p+" WARNING: ", log.Ldate|log.Ltime), // Logger for warning messages
		err:     log.New(writer, p+" ERROR: ", log.Ldate|log.Ltime),   // Logger for error messages
		writer:  writer,                                               // Set the output writer for the logger
	}
	// Return the new instance of logger
	return logger
}

// Create Non-Format Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enable Logs

func (l *Logger) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.debug.Println(msg)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.info.Println(msg)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.warning.Println(msg)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.err.Println(msg)
}
