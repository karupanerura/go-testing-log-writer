# go-testing-log-writer

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
func NewTestingLogWriter(t *testing.T) *TestingLogWriter
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
