package assert

import (
	"fmt"
	"reflect"
	"strings"
)

// Assert panics with AssertionError if cond is false.
func Assert(cond bool, text string) {
	if cond {
		return
	}
	panic(AssertionError{Text: text})
}

// Assertf panics with AssertionError if cond is false.
func Assertf(cond bool, format string, v ...interface{}) {
	if cond {
		return
	}
	text := fmt.Sprintf(format, v...)
	panic(AssertionError{Text: text})
}

// AssertNotNil takes a pointer to a nil-able type (pointer, interface, etc)
// and panics with AssertionError if the pointed-to value is nil.
func AssertNotNil(v interface{}) {
	r0 := reflect.ValueOf(v)
	r1 := r0.Elem()
	if !r1.IsNil() {
		return
	}
	text := fmt.Sprintf("%s is nil", r1.Type().String())
	panic(AssertionError{Text: text})
}

// AssertionError is the error type for Assert failure panics.
type AssertionError struct {
	Text string
}

// Error fulfills the error interface.
func (err AssertionError) Error() string {
	var buf strings.Builder
	buf.Grow(16 + len(err.Text))
	buf.WriteString("AssertionError")
	if err.Text != "" {
		buf.WriteString(": ")
		buf.WriteString(err.Text)
	}
	return buf.String()
}

var _ error = AssertionError{}
