package logger

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNounVerbToDoingDone(t *testing.T) {
	for _, tc := range []struct {
		noun        string
		verb        []string
		doing, done string
	}{
		{"file thing", []string{}, "Closing file thing", "File thing closed"},
		{"service", []string{"stop"}, "Stopping service", "Service stopped"},
		{"throp", []string{"funt"}, "Funting throp", "Throp funted"},
	} {
		t.Run(tc.doing, func(t *testing.T) {
			doing, done := nounVerbToDoingDone(tc.noun, tc.verb)
			assert.Check(t, doing == tc.doing, doing)
			assert.Check(t, done == tc.done, done)
		})
	}
}
