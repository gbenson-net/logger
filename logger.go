// Package logger provides a thin wrapper around [zerolog].
package logger

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var DefaultLevel = zerolog.LevelInfoValue

type (
	Level  = zerolog.Level
	Logger = zerolog.Logger
)

type Options struct {
	Writer io.Writer
	Level  string
}

// New creates a new logger.
func New(options *Options) Logger {
	if options == nil {
		panic("nil options")
	}

	writer := options.Writer
	if writer == nil {
		writer = zerolog.NewConsoleWriter()
	}
	log := zerolog.New(writer)

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

// Ctx returns the Logger associated with the given context, or
// an appropriate (non-nil) default if the given context has no
// logger associated.
func Ctx(ctx context.Context) *Logger {
	if ctx == nil {
		panic("nil context")
	}

	return zerolog.Ctx(ctx)
}

// LevelFor returns an appropriate level to log the given error at.
func LevelFor(err error) Level {
	if IsRecoveredPanicError(err) {
		return zerolog.PanicLevel
	}
	return zerolog.ErrorLevel
}
