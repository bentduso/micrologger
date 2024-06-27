package microlog_test

import (
	"github.com/bentduso/micrologger/microlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := microlog.New(microlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: [DEBUG] Hello, world
}

const (
	traceMessage = "Entering function xyz."
	debugMessage = "Performing complex calculation..."
	infoMessage  = "Service started successfully."
	warnMessage  = "Memory usage high."
	errorMessage = "Database connection failed."
	fatalMessage = "Server crashed unexpectedly."
	message      = "An ordinary message."
)

// testWriter implements the io.Writer interface.
type testWriter struct {
	contents string
}

// Write appends the byte slice p to the contents of the testWriter instance.
// It returns the number of bytes written from p and a nil error.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	tt := map[string]struct {
		level    microlog.Level
		expected string
	}{
		"trace": {
			level: microlog.LevelTrace,
			expected: "[TRACE] " + traceMessage + "\n" + "[DEBUG] " + debugMessage + "\n" +
				"[INFO] " + infoMessage + "\n" + "[WARN] " + warnMessage + "\n" +
				"[ERROR] " + errorMessage + "\n" + "[FATAL] " + fatalMessage + "\n",
		},
		"debug": {
			level: microlog.LevelDebug,
			expected: "[DEBUG] " + debugMessage + "\n" + "[INFO] " + infoMessage + "\n" +
				"[WARN] " + warnMessage + "\n" + "[ERROR] " + errorMessage + "\n" +
				"[FATAL] " + fatalMessage + "\n",
		},
		"info": {
			level: microlog.LevelInfo,
			expected: "[INFO] " + infoMessage + "\n" + "[WARN] " + warnMessage + "\n" +
				"[ERROR] " + errorMessage + "\n" + "[FATAL] " + fatalMessage + "\n",
		},
		"warn": {
			level: microlog.LevelWarn,
			expected: "[WARN] " + warnMessage + "\n" + "[ERROR] " + errorMessage + "\n" +
				"[FATAL] " + fatalMessage + "\n",
		},
		"error": {
			level:    microlog.LevelError,
			expected: "[ERROR] " + errorMessage + "\n" + "[FATAL] " + fatalMessage + "\n",
		},
		"fatal": {
			level:    microlog.LevelFatal,
			expected: "[FATAL] " + fatalMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := microlog.New(tc.level, microlog.WithOutput(tw))

			testedLogger.Tracef(traceMessage)
			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Warnf(warnMessage)
			testedLogger.Errorf(errorMessage)
			testedLogger.Fatalf(fatalMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents: expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_Logf(t *testing.T) {
	tt := map[string]struct {
		level    microlog.Level
		expected string
	}{
		"trace": {
			level:    microlog.LevelTrace,
			expected: "[TRACE] " + message + "\n",
		},
		"debug": {
			level:    microlog.LevelDebug,
			expected: "[DEBUG] " + message + "\n",
		},
		"info": {
			level:    microlog.LevelInfo,
			expected: "[INFO] " + message + "\n",
		},
		"warn": {
			level:    microlog.LevelWarn,
			expected: "[WARN] " + message + "\n",
		},
		"error": {
			level:    microlog.LevelError,
			expected: "[ERROR] " + message + "\n",
		},
		"fatal": {
			level:    microlog.LevelFatal,
			expected: "[FATAL] " + message + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := microlog.New(tc.level, microlog.WithOutput(tw))

			testedLogger.Logf(tc.level, message)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents: expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}
