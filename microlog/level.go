package microlog

// Level represents an available logging level.
type Level byte

const (
	// LevelTrace represents a logging level for tracing code execution at a very fine-grained level.
	LevelTrace Level = iota
	// LevelDebug represents a logging level that is diagnostically helpful for IT and sysadmins.
	LevelDebug
	// LevelInfo represents a logging level that contains information deemed valuable, such as service start/stop.
	LevelInfo
	// LevelWarn represents a logging level for potential issues that a system can recover from automatically.
	LevelWarn
	// LevelError represents a logging level for errors that require user intervention but do not stop the application.
	LevelError
	// LevelFatal represents a logging level for errors that force an application to shut down.
	LevelFatal
)

// String returns the string representation of the logging level.
func (lvl Level) String() string {
	switch lvl {
	case LevelTrace:
		return "[TRACE]"
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelWarn:
		return "[WARN]"
	case LevelError:
		return "[ERROR]"
	case LevelFatal:
		return "[FATAL]"
	default:
		return "[UNKNOWN]"
	}
}
