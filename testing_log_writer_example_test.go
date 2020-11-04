package tlogw_test

import (
	"fmt"
	"testing"

	tlogw "github.com/karupanerura/go-testing-log-writer"
)

var t testing.TB = &testingTBImpl{}

type testingTBImpl struct{ *testing.T }

func (t *testingTBImpl) Log(args ...interface{}) {
	fmt.Print(args...)
}

func ExampleLog() {
	w := tlogw.NewTestingLogWriter(t)
	w.Write([]byte("testing "))
	w.WriteString("log!\n")
	// Output: testing log!
}
