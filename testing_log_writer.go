package tlogw

import (
	"bytes"
	"strings"
	"testing"
)

// TestingLogWriter provides io.Writer interface for (*testing.T).Log
type TestingLogWriter struct {
	t   testingLogger
	buf bytes.Buffer
}

type testingLogger interface {
	Log(args ...interface{})
}

// NewTestingLogWriter creates a new TestingLogWriter instance.
func NewTestingLogWriter(t testing.TB) *TestingLogWriter {
	return &TestingLogWriter{t: t}
}

// Write a log to (*testing.T).Log with line buffering.
func (w *TestingLogWriter) Write(b []byte) (int, error) {
	return w.WriteString(string(b))
}

// WriteString is for io.WriteString compatibility.
func (w *TestingLogWriter) WriteString(s string) (int, error) {
	if i := strings.IndexRune(s, '\n'); i == -1 {
		return w.buf.WriteString(s)
	} else if i == len(s)-1 {
		w.t.Log(w.buf.String(), s[:i])
		w.buf.Reset()
		return len(s), nil
	} else {
		w.t.Log(w.buf.String(), s[:i])
		w.buf.Reset()
		return w.WriteString(s[i+1:])
	}
}
