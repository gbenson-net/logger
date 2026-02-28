// Package log provides a global default logger.
package log

import (
	"io"
	"sync"

	"gbenson.net/go/logger"
)

type Logger = logger.Logger

var DefaultLoggerOptions = &logger.Options{
	Timestamp: true,
}

var defaultLogger = sync.OnceValue(func() *Logger {
	l := logger.New(DefaultLoggerOptions)
	return &l
})

// DefaultLogger returns the global default logger.
func DefaultLogger() *Logger {
	return defaultLogger()
}

// With creates a child of the default logger with the field added to its context.
func With() logger.Context {
	return defaultLogger().With()
}

// Trace starts a new message with trace level.
func Trace() *logger.Event {
	return defaultLogger().Trace()
}

// Debug starts a new message with debug level.
func Debug() *logger.Event {
	return defaultLogger().Debug()
}

// Info starts a new message with info level.
func Info() *logger.Event {
	return defaultLogger().Info()
}

// Warn starts a new message with warn level.
func Warn() *logger.Event {
	return defaultLogger().Warn()
}

// Error starts a new message with error level.
func Error() *logger.Event {
	return defaultLogger().Error()
}

// Err starts a new message with error level with err as a field if
// not nil, or with info level if err is nil.
func Err(err error) *logger.Event {
	return defaultLogger().Err(err)
}

// Panic starts a new message with panic level. The resulting event's
// Msg method will call the panic() function when invoked, stopping
// the ordinary flow of the calling goroutine.
func Panic() *logger.Event {
	return defaultLogger().Panic()
}

// WarnErr starts a new message with warn level with err as a field
// if not nil, or with info level if err is nil.
func WarnErr(err error) *logger.Event {
	if err == nil {
		return Info()
	} else {
		return Warn().Err(err)
	}
}

// NotImplemented logs a "not implemented" message with warning level.
func WarnNotImplemented(item string) {
	Warn().Str("what", item).Msg("Not implemented")
}

// LoggedClose wraps an [io.Closer] so it logs errors that would
// otherwise be silently dropped if the closer was invoked as a
// deferred function.
func LoggedClose(c io.Closer, noun string, verb ...string) {
	logger.LoggedClose(defaultLogger(), c, noun, verb...)
}
