package log

import (
	"bytes"
	"errors"
	"testing"

	"github.com/rs/zerolog"
	"gotest.tools/v3/assert"
)

func TestTrace(t *testing.T) {
	Trace().Msg("this shouldn't crash")
}

func TestErr(t *testing.T) {
	assert.Equal(t, caplog(func() {
		Err(errors.New("problem")).Msg("")
	}),
		`{"level":"error","error":"problem"}`+"\n",
	)
}

func TestNilErr(t *testing.T) {
	assert.Equal(t, caplog(func() {
		Err(nil).Msg("")
	}),
		`{"level":"info"}`+"\n",
	)
}

func TestWarnErr(t *testing.T) {
	assert.Equal(t, caplog(func() {
		WarnErr(errors.New("problem")).Msg("")
	}),
		`{"level":"warn","error":"problem"}`+"\n",
	)
}

func TestWarnNilErr(t *testing.T) {
	assert.Equal(t, caplog(func() {
		WarnErr(nil).Msg("er?")
	}),
		`{"level":"info","message":"er?"}`+"\n",
	)
}

func TestWarnNotImplemented(t *testing.T) {
	assert.Equal(t, caplog(func() {
		WarnNotImplemented("things")
	}),
		`{"level":"warn","what":"things","message":"Not implemented"}`+"\n",
	)
}

func caplog(f func()) string {
	var b bytes.Buffer
	testLogger := zerolog.New(&b)
	saved := defaultLogger
	defer func() { defaultLogger = saved }()
	defaultLogger = func() *Logger {
		return &testLogger
	}
	f()
	return string(b.Bytes())
}
