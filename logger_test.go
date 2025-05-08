package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"gotest.tools/v3/assert"
)

// Ctx() returns an appropriate default if no logger is attached.
func TestCtxDefault(t *testing.T) {
	logger := Ctx(t.Context())
	assert.Check(t, logger.GetLevel() == zerolog.Disabled)
	logger.Error().Msg("this shouldn't display (or crash)")
}

// TestContext() returns a logger that may or may not be disabled
// depending on what environment variables are in use.
func TestTestContext(t *testing.T) {
	logger := Ctx(TestContext(t))

	logger.Debug().
		Str("if", "LL in [debug,trace]").
		Msg("This should only display")
	logger.Trace().
		Str("if", "LL=trace").
		Msg("This should only display")
	logger.Info().
		Str("unless", "LL in [disabled,warn,error,panic]").
		Msg("This should display")
}

// Deprecated
func TestTestCtx(t *testing.T) {
	logger := Ctx(TestCtx(t))
	assert.Check(t, logger.GetLevel() == zerolog.DebugLevel)
	logger.Info().Msg("Two messages should follow regardless of LOG_LEVEL")
	logger.Debug().Str("at_level", "debug").Msg("this should display")
	logger.Trace().Msg("THIS SHOULDN'T BE DISPLAYED!!!")
	logger.Info().Str("at_level", "info").Msg("so should this")
}
