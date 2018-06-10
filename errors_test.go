package werrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestTracker_Error(t *testing.T) {
	err := errors.New("error")
	err = Wrap(err, "ann1")
	err = Wrap(err, "ann2")
	if err.Error() != "ann2 *> ann1 *> error" {
		t.Fatal("error message must be ann2 *> ann1 *> error")
	}
}

func TestWrap(t *testing.T) {
	err := errors.New("message")
	err = Wrap(err, "test")
	if err.Error() != "test *> message" {
		t.Fatal("the text is incorrect", err)
	}
}

func TestCause(t *testing.T) {
	err := errors.New("message")
	err = Wrap(err, "test")
	if Cause(err).Error() != "message" {
		t.Fatal("cause of error is not match", err)
	}
}

func f() (err error) {
	defer DefWrap(&err, "test")
	err = errors.New("message")
	return
}

func TestDefWrap(t *testing.T) {
	err := f()
	if err.Error() != "test *> message" {
		t.Fatal("DefWrap works not right", err)
	}
}

func TestWrapf(t *testing.T) {
	err := errors.New("message")
	err = Wrapf(err, "test(a:%s)", "hello")
	if err.Error() != "test(a:hello) *> message" {
		t.Fatal("DefWrapf doesn't format", err)
	}
}

func g(arg string) (err error) {
	defer DefWrapf(&err, "test(arg:%s)", arg)
	err = errors.New("message")
	return
}

func TestDefWrapf(t *testing.T) {
	err := g("cool")
	if err.Error() != "test(arg:cool) *> message" {
		t.Fatal("DefWrapf doesn't format with deferring", err)
	}
}

func ExampleWrap() {
	err := errors.New("error")
	err = Wrap(err, "annotation")
	fmt.Println(err)
	// Output: annotation *> error
}

func ExampleDefWrap() {
	g := func() error {
		return errors.New("g(): wrong")
	}
	f := func() (err error) {
		defer DefWrap(&err, "f()")
		return g()
	}

	err := f()
	fmt.Println(err)
	// Output: f() *> g(): wrong
}

func benchmarkTrackerError(n int, b *testing.B) {
	err := errors.New("everything is bad")
	a := "a"
	for i := 0; i < n; i++ {
		a = a + "a"
		err = Wrap(err, a)
	}
	b.ResetTimer()
	var out string
	for i := 0; i < b.N; i++ {
		out = err.Error()
	}
	_ = out
}

func BenchmarkTracker_Error1(b *testing.B)   { benchmarkTrackerError(1, b) }
func BenchmarkTracker_Error4(b *testing.B)   { benchmarkTrackerError(4, b) }
func BenchmarkTracker_Error8(b *testing.B)   { benchmarkTrackerError(8, b) }
func BenchmarkTracker_Error32(b *testing.B)  { benchmarkTrackerError(32, b) }
func BenchmarkTracker_Error128(b *testing.B) { benchmarkTrackerError(128, b) }
