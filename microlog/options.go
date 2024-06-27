package microlog

import "io"

// Option defines a functional option to configure a Logger.
type Option func(*Logger)

// WithOutput returns an Option that sets the output destination for a Logger.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}
