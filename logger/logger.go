package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// LogLevel defines the level of logging
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var levelStrings = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

var levelColors = []string{
	"\033[36m", // DEBUG - Cyan
	"\033[32m", // INFO - Green
	"\033[33m", // WARNING - Yellow
	"\033[31m", // ERROR - Red
	"\033[35m", // FATAL - Magenta
}

const resetColor = "\033[0m"

type Logger struct {
	mu       sync.Mutex  // Mutex for serializing access to the logger
	logLevel LogLevel    // The level of logging to use
	logger   *log.Logger // The logger to use
	jsonMode bool        // Whether to output in JSON format
}

// LogMessage defines a single log message
type LogMessage struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

// NewLogger creates a new Logger instance
func NewLogger(level LogLevel, jsonMode bool) *Logger {
	return &Logger{
		logLevel: level,
		logger:   log.New(os.Stdout, "", 0),
		jsonMode: jsonMode,
	}
}

func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level >= l.logLevel {
		l.mu.Lock()
		defer l.mu.Unlock()

		prefix := levelStrings[level]
		color := levelColors[level]

		message := fmt.Sprintf(format, v...)
		timestamp := time.Now().Format(time.RFC3339)
		if l.jsonMode {
			logMsg := LogMessage{
				Timestamp: timestamp,
				Level:     prefix,
				Message:   message,
			}
			logData, _ := json.Marshal(logMsg)
			l.outputLog(string(logData))
		} else {
			logLine := fmt.Sprintf("%s[%s] [%s] %s%s", color, timestamp, prefix, message, resetColor)
			l.outputLog(logLine)
		}

		if level == FATAL {
			os.Exit(1)
		}
	}
}

func (l *Logger) outputLog(logLine string) {
	l.logger.Println(logLine)
}

// Debug logs a message at DEBUG level
func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

// Info logs a message at INFO level
func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

// Warning logs a message at WARNING level
func (l *Logger) Warning(format string, v ...interface{}) {
	l.log(WARNING, format, v...)
}

// Error logs a message at ERROR level
func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

// Fatal logs a message at FATAL level and exits
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
}
