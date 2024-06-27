package microlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New creates a new Logger instance with the specified logging threshold and options.
// By default, logs are written to os.Stdout.
func New(threshold Level, opts ...Option) *Logger {
	l := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(l)
	}

	return l
}

// Tracef formats and prints a message if the log level is trace or higher.
func (l *Logger) Tracef(format string, args ...any) {
	if l.threshold > LevelTrace {
		return
	}

	l.logf(LevelTrace, format, args...)
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf(LevelInfo, format, args...)
}

// Warnf formats and prints a message if the log level is warn or higher.
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	l.logf(LevelWarn, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf(LevelError, format, args...)
}

// Fatalf formats and prints a message if the log level is error or higher.
func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}

	l.logf(LevelFatal, format, args...)
}

// Logf logs a formatted message at the specified level if the log level is low enough.
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args...)
}

// logf formats and writes a log message to the output.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(l.output, "%s %s\n", lvl, message)
}

// TODO: Implement message size limits to ensure logged messages do not exceed a specified size in bytes/characters.
// TODO: Enhance log functions to output logs in JSON format. Example:
// {
//    "time": "2024-06-27 18:06:30.148845Z",
//    "level": "fatal",
//    "message": "the program quit unexpectedly"
// }
