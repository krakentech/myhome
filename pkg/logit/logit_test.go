package logit

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

type LogItTest struct {
	Name    string
	Debug   bool
	Pattern string
	Values  []any
	Err     error
	Want    string
}

func TestDebug(t *testing.T) {

	tests := []LogItTest{
		{Name: "debug test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: ""},
		{Name: "debug test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🐛 - test value\n", Debug: true},
		{Name: "debug test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00   🐛 - the number is 123\n", Debug: true},
	}

	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			IsDebug = tt.Debug
			Debug(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}

}

func TestInfo(t *testing.T) {
	tests := []LogItTest{
		{Name: "info test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🧠 - test value\n"},
		{Name: "info test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🧠 - test value\n", Debug: true},
		{Name: "info test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00   🧠 - the number is 123\n", Debug: true},
	}

	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			IsDebug = tt.Debug
			Info(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}

}

func TestWarn(t *testing.T) {
	tests := []LogItTest{
		{Name: "warn test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🚧 - test value\n"},
		{Name: "warn test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🚧 - test value\n", Debug: true},
		{Name: "warn test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00   🚧 - the number is 123\n", Debug: true},
	}

	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			IsDebug = tt.Debug
			Warn(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}

}

func TestErr(t *testing.T) {
	tests := []LogItTest{
		{Name: "err test debug off", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🛑 - test value\n"},
		{Name: "err test only pattern", Pattern: "test value", Values: []any{}, Err: nil, Want: "00.00.00-00:00:00   🛑 - test value\n", Debug: true},
		{Name: "err test  pattern and values", Pattern: "the number is %d", Values: []any{123}, Err: nil, Want: "00.00.00-00:00:00   🛑 - the number is 123\n", Debug: true},
	}

	var buf bytes.Buffer
	SetWriter(&buf)
	isTesting = true
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			IsDebug = tt.Debug
			Err(tt.Pattern, tt.Values...)
			assert.Equal(t, tt.Want, buf.String())
			buf.Reset()
		})
	}

}

func TestWriter(t *testing.T) {
	storeStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	SetWriter(MockIoWriter{})
	Info("this is just  a test")

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = storeStdout

	assert.Equal(t, "failed to print line: on purpouse fail for test", string(out))
}

type MockIoWriter struct{}

func (mw MockIoWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("on purpouse fail for test")
}
