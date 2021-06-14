package assert

import (
	"fmt"
	"reflect"
	"strings"
)

// Raise panics with Error.
func Raise(text string) {
	panic(Error{Text: text})
}

// Raisef panics with Error.
func Raisef(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	panic(Error{Text: text})
}

// Assert panics with Error if cond is false.
func Assert(cond bool, text string) {
	if cond {
		return
	}
	Raise(text)
}

// Assertf panics with Error if cond is false.
func Assertf(cond bool, format string, v ...interface{}) {
	if cond {
		return
	}
	Raisef(format, v...)
}

// NotNil takes a pointer to a nil-able type (pointer, interface, etc) and
// panics with Error if the pointed-to value is nil.
func NotNil(v interface{}) {
	r0 := reflect.ValueOf(v)
	r1 := r0.Elem()
	if !r1.IsNil() {
		return
	}
	Raisef("%s is nil", r1.Type().String())
}

// Error is the error type for Assert failure panics.
type Error struct {
	Text string
}

// Error fulfills the error interface.
func (err Error) Error() string {
	var buf strings.Builder
	buf.Grow(16 + len(err.Text))
	buf.WriteString("AssertionError")
	if err.Text != "" {
		buf.WriteString(": ")
		buf.WriteString(err.Text)
	}
	return buf.String()
}

var _ error = Error{}

// AssertNotNil is a compatibility alias for NotNil.
func AssertNotNil(v interface{}) {
	NotNil(v)
}

// AssertionError is a compatibility alias for Error.
type AssertionError = Error
