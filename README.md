# go-testing-log-writer

[![Actions Status](https://github.com/karupanerura/go-testing-log-writer/workflows/test/badge.svg)](https://github.com/karupanerura/go-testing-log-writer/actions)
[![codecov](https://codecov.io/gh/karupanerura/go-testing-log-writer/branch/master/graph/badge.svg)](https://codecov.io/gh/karupanerura/go-testing-log-writer)
[![GoDoc](https://godoc.org/github.com/karupanerura/go-testing-log-writer?status.svg)](http://godoc.org/github.com/karupanerura/go-testing-log-writer)

 --
    import tlogw "github.com/karupanerura/go-testing-log-writer"

## Usage

#### type TestingLogWriter

```go
type TestingLogWriter struct {
}
```

TestingLogWriter provides io.Writer interface for (*testing.T).Log

#### func  NewTestingLogWriter

```go
func NewTestingLogWriter(t testing.TB) *TestingLogWriter
```
NewTestingLogWriter creates a new TestingLogWriter instance.

#### func (*TestingLogWriter) Write

```go
func (w *TestingLogWriter) Write(b []byte) (int, error)
```
Write a log to (*testing.T).Log with line buffering.

#### func (*TestingLogWriter) WriteString

```go
func (w *TestingLogWriter) WriteString(s string) (int, error)
```
WriteString is for io.WriteString compatibility.
