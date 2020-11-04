package tlogw_test

import (
	"fmt"
	"io"
	"testing"

	tlogw "github.com/karupanerura/go-testing-log-writer"
)

var t testing.TB = &testingTBImpl{} // emulate *testing.T

type testingTBImpl struct{ *testing.T }

func (t *testingTBImpl) Log(args ...interface{}) {
	fmt.Print(args...) // emulate (*testing.T).Log
}

func ExampleNewTestingLogWriter() {
	w := tlogw.NewTestingLogWriter(t)
	io.WriteString(w, "testing log!\n")
	// Output: testing log!
}
