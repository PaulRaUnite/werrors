/*
Package werrors provides error tracking and/or
adding more context to errors (something like function arguments).

Important notes

If you wrap a nil error, you get not nil error.
So check your error before use Wrap and "friends".

Current realizations of Wrapf, DefWrapf and Error
cause allocations because of fmt.Printf use.
*/
package werrors

import (
	"fmt"
)

// tracker is a struct that is used
// to wrap errors with annotation.
type tracker struct {
	annotation string
	err        error
}

// Error is the implementation of error interface.
func (t tracker) Error() string {
	return string(t.Bytes())
}

const divider = " *> "

// Bytes returns byte slice of nested annotations and
// error in the end.
func (t tracker) Bytes() []byte {
	//preallocate slice for common nesting
	annotations := make([]string, 1, 8)
	annotations[0] = t.annotation

	err := t.err
	//length of outcome slice
	needed := len(t.annotation) + len(divider)
	for {
		//if next error is tracker also
		if tr, ok := err.(tracker); ok {
			//save annotation field and it's length
			annotations = append(annotations, tr.annotation)
			needed += len(tr.annotation) + len(divider)
			err = tr.err
		} else {
			break
		}
	}
	//ending error
	errStr := err.Error()
	//calculate final length
	needed += len(errStr)

	//construct outcome
	outcome := make([]byte, 0, needed)
	for _, ann := range annotations {
		//ann1 *> ann2 *> ... *> annN *> ...
		outcome = append(outcome, ann...)
		outcome = append(outcome, divider...)
	}
	//...error
	outcome = append(outcome, errStr...)
	return outcome
}

// Wrap annotates an error with a string.
// See example.
func Wrap(err error, annotation string) error {
	return tracker{
		annotation: annotation,
		err:        err,
	}
}

// Wrapf annotates an error using formatting.
// Use it as a smart analogue of Wrap.
func Wrapf(err error, format string, args ...interface{}) error {
	return tracker{
		annotation: fmt.Sprintf(format, args...),
		err:        err,
	}
}

// Cause returns the initial
// error if err is a tracker.
func Cause(err error) error {
	c, ok := err.(tracker)
	if !ok {
		return err
	}
	return Cause(c.err)
}

// DefWrap annotates an error referenced
// by the pointer to allow using wrapping
// in defer statements.
// See example.
func DefWrap(errp *error, annotation string) {
	if errp != nil {
		err := *errp
		if err != nil {
			*errp = Wrap(err, annotation)
		}
	}
}

// DefWrapf does everything that
// DefWrap and Wrapf do.
func DefWrapf(errp *error, format string, args ...interface{}) {
	if errp != nil {
		err := *errp
		if err != nil {
			*errp = Wrapf(err, format, args...)
		}
	}
}
