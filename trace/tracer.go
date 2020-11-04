package trace

import (
	"fmt"
	"io"
)

// Tracer saves what happened in object
type Tracer interface {
	Trace(...interface{})
}

// New returns new tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
