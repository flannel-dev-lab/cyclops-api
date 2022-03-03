package cycapi

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}

	// Set standard log format as JSON
	standardLogger.Formatter = &logrus.JSONFormatter{}

	// Default log output file name
	logFile := "log.log"
	// Get output log name from env
	if lf := os.Getenv("LOG_FILE"); lf != "" {
		logFile = lf
	}

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		standardLogger.SetOutput(file)
	} else {
		standardLogger.Error("Failed to log to file")
	}

	// Default log level = 3 (Warning and above)
	logLevel := 3
	// Get log level from env
	if ll := os.Getenv("LOG_LEVEL"); ll != "" {
		level, err := strconv.Atoi(ll)
		if err == nil {
			logLevel = level
		}
	}

	// Only log the info severity or above.
	standardLogger.SetLevel(logrus.Level(logLevel))

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
	infoMessage            = Event{4, "Info: %s"}
	debugMessage           = Event{5, "Debug: %s"}
)

// Debug is a standard error message
func (l *StandardLogger) Debug(info string) {
	l.Infof(debugMessage.message, info)
}

// Info is a standard error message
func (l *StandardLogger) Info(info string) {
	l.Infof(infoMessage.message, info)
}

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}