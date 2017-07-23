// Package provides simplified errors
// and error tracking.
package werrors

import (
	"fmt"
)

// errorString is a simplified version of standard
// Golang errors.errorString type.
type errorString string

func (e errorString) Error() string {
	return string(e)
}

// New is a constructor for errorString type.
func New(message string) error {
	return errorString(message)
}

// tracker is a struct that is used
// to wrap errors for tracking
// error path.
type tracker struct {
	annotation string
	err        error
}

func (w tracker) Error() string {
	return fmt.Sprintf("%s *> %s", w.annotation, w.err.Error())
}

// Wrap annotates an error with a string.
// Use it to track paths of errors, as:
// main *> calc *> error text.
func Wrap(err error, annotation string) error {
	return tracker{
		annotation: annotation,
		err:        err,
	}
}

// Wrapf annotates an error using formatting.
func Wrapf(err error, format string, args ...interface{}) error {
	return tracker{
		annotation: fmt.Sprintf(format, args...),
		err: err,
	}
}

// Cause returns an initial error if
// err is a tracker.
func Cause(err error) error {
	c, ok := err.(tracker)
	if !ok {
		return err
	}
	return Cause(c.err)
}

// DefWrap annotates an POINTER of
// error to allow an user use it
// in defer statements.
// Example:
// func f() (err error) {
//   defer DefWrap("text", &err)
//   //some actions
//   return
// }
func DefWrap(errp *error, annotation string) {
	if errp != nil {
		err := *errp
		if err != nil {
			*errp = Wrap(err, annotation)
		}
	}
}

// DefWrapf does everything that DefWrap
// and Wrapf do.
func DefWrapf(errp *error, format string, args ...interface{}) {
	if errp != nil {
		err := *errp
		if err != nil {
			*errp = Wrapf(err, format, args...)
		}
	}
}
