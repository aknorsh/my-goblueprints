package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New returns nil.")
	} else {
		tracer.Trace("Hello, trace package")
		if buf.String() != "Hello, trace package\n" {
			t.Errorf("'%s' is not expected.", buf.String())
		}
	}
}
