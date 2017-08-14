package werrors

import (
	"errors"
	"fmt"
	"testing"
)

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

func BenchmarkTracker_Error(b *testing.B) {
	err := errors.New("everything is bad")
	a := "a"
	for i := 0; i < 100; i++ {
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