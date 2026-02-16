// Package logger provides a thin wrapper around [zerolog].
package logger

import (
	"context"
	"io"
	"os"

	"golang.org/x/term"

	"github.com/rs/zerolog"
)

var (
	DefaultLevel     = zerolog.LevelInfoValue
	NewConsoleWriter = zerolog.NewConsoleWriter
)

type (
	Context = zerolog.Context
	Level   = zerolog.Level
	Logger  = zerolog.Logger
	Event   = zerolog.Event
)

type Options struct {
	Writer io.Writer
	Level  string
}

// New creates a new logger.
func New(options *Options) Logger {
	if options == nil {
		options = &Options{}
	}

	log := zerolog.New(options.writer())

	level, err := zerolog.ParseLevel(options.level())
	if err != nil {
		log.Warn().Err(err).Msg("")
	}
	if level != zerolog.NoLevel {
		log = log.Level(level)
	}

	return log
}

func (o *Options) level() string {
	if s := o.Level; s != "" {
		return s
	}
	if s := os.Getenv("LOG_LEVEL"); s != "" {
		return s
	}
	if s := os.Getenv("LL"); s != "" {
		return s
	}
	return DefaultLevel
}

func (o *Options) writer() io.Writer {
	if w := o.Writer; w != nil {
		return w // caller supplied
	}
	if term.IsTerminal(int(os.Stdout.Fd())) {
		return NewConsoleWriter() // pretty
	}
	return os.Stdout // raw JSON
}

// Ctx returns the Logger associated with the given context, or
// an appropriate (non-nil) default if the given context has no
// logger associated.
func Ctx(ctx context.Context) *Logger {
	if ctx == nil {
		panic("nil context")
	}

	return zerolog.Ctx(ctx)
}

// Contexter is the subset of [testing.T] et al used by TestCtx.
type Contexter interface {
	Context() context.Context
}

// TestContext returns its receiver's context after associating a
// logger with it.  It is intended for use with [testing.T] et al,
// for example:
//
//	package pkg
//
//	import (
//		"testing"
//		"gbenson.net/go/logger"
//	)
//
//	func TestSomething(t *testing.T) {
//		ctx := logger.TestCtx(t)
//		// ...
//		logger.Ctx(ctx).Info().Msg("something happened")
//	}
func TestContext(t Contexter) context.Context {
	log := New(nil)
	return log.WithContext(t.Context())
}

// TestCtx returns its receiver's context after associating a logger
// at debug level with it.
//
// Deprecated: Use [TestContext] instead.
func TestCtx(t Contexter) context.Context {
	log := New(&Options{Level: "Debug"})
	return log.WithContext(t.Context())
}

// LevelFor returns an appropriate level to log the given error at.
func LevelFor(err error) Level {
	if IsRecoveredPanicError(err) {
		return zerolog.PanicLevel
	}
	return zerolog.ErrorLevel
}
