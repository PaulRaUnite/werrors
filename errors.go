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

// tracker wraps errors to track error path.
type tracker struct {
	annotation string
	err        error
}

func (w tracker) Error() string {
	return fmt.Sprintf("%s *> %s", w.annotation, w.err.Error())
}

// Wrap annotates the error with the string.
// Use it to track paths of errors, as:
// main *> calc *> error text.
func Wrap(annotation string, err error) error {
	return tracker{
		annotation: annotation,
		err:        err,
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

// DefWrap is expected to use in
// defer statements for tracking
// errors, as follow:
// func f() (err error) {
//   defer DefWrap("text", &err)
//   //some actions
//   return
// }
func DefWrap(annotation string, errp *error) {
	if errp != nil {
		err := *errp
		if err != nil {
			*errp = Wrap(annotation, err)
		}
	}
}
