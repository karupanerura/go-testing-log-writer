package tlogw_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	tlogw "github.com/karupanerura/go-testing-log-writer"
)

func TestNew(t *testing.T) {
	w := tlogw.NewTestingLogWriter(t)
	w.Write([]byte("successfully create a new log writer!\n"))
}

type testingLoggerImpl struct {
	*testing.T
	buf bytes.Buffer
}

func (t *testingLoggerImpl) Helper() {} // emulate (*testing.T).Helper

func (l *testingLoggerImpl) Log(args ...interface{}) {
	l.buf.WriteString(fmt.Sprint(args...) + "\n") // add \n for suffix to emulate a log line
}

func TestBuffering(t *testing.T) {
	l := &testingLoggerImpl{}
	w := tlogw.NewTestingLogWriter(l)

	io.WriteString(w, "hello ")
	if l.buf.Len() != 0 {
		t.Error("hello should be buffered")
		t.Logf("outbuf: %v", l.buf.String())
	}

	io.WriteString(w, "world!\n")
	if l.buf.String() != "hello world!\n" {
		t.Error("hello world! should be outputed")
		t.Logf("outbuf: %v", l.buf.String())
	}

	io.WriteString(w, "go testing users!\nhey!")
	if l.buf.String() != "hello world!\ngo testing users!\n" {
		t.Error("go testing users! should be outputed")
		t.Logf("outbuf: %v", l.buf.String())
	}

	io.WriteString(w, "\n")
	if l.buf.String() != "hello world!\ngo testing users!\nhey!\n" {
		t.Error("hey! should be outputed")
		t.Logf("outbuf: %v", l.buf.String())
	}
}
