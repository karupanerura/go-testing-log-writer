package tlogw_test

import (
	"fmt"
	"io"
	"testing"

	tlogw "github.com/karupanerura/go-testing-log-writer"
)

var t testing.TB = &testingTBImpl{}

type testingTBImpl struct{ *testing.T }

func (t *testingTBImpl) Log(args ...interface{}) {
	fmt.Print(args...)
}

func ExampleNewTestingLogWriter() {
	w := tlogw.NewTestingLogWriter(t)
	io.WriteString(w, "testing log!\n")
	// Output: testing log!
}

func TestNew(t *testing.T) {
	w := tlogw.NewTestingLogWriter(t)
	w.Write([]byte("successfully create a new log writer!\n"))
}
