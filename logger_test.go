package logger

import (
	"testing"

	"github.com/rs/zerolog"
)

// Ctx() returns an appropriate default if no logger is attached.
func TestCtxDefault(t *testing.T) {
	logger := Ctx(t.Context())
	expect(t, logger.GetLevel(), zerolog.Disabled)
	logger.Error().Msg("this shouldn't display (or crash)")
}

func TestTestCtx(t *testing.T) {
	logger := Ctx(TestCtx(t))
	expect(t, logger.GetLevel(), zerolog.DebugLevel)
	logger.Debug().Str("at_level", "debug").Msg("this should display")
	logger.Trace().Msg("THIS SHOULDN'T BE DISPLAYED!!!")
	logger.Info().Str("at_level", "info").Msg("so should this")
}
