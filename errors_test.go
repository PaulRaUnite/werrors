package werrors

import (
	"testing"
)

func TestNew(t *testing.T) {
	err := New("message")
	if err.Error() != "message" {
		t.Fatal("string corruption")
	}
}

func TestWrap(t *testing.T) {
	err := New("message")
	err = Wrap("test", err)
	if err.Error() != "test *> message" {
		t.Fatal("the text is incorrect")
	}
}

func TestCause(t *testing.T) {
	err := New("message")
	err = Wrap("test", err)
	if Cause(err).Error() != "message" {
		t.Fatal("cause of error is not match")
	}
}

func f() (err error) {
	defer DefWrap("test", &err)
	err = New("message")
	return
}
func TestDefWrap(t *testing.T) {
	err := f()
	if err.Error() != "test *> message" {
		t.Fatal("DefWrap works not right")
	}
}
