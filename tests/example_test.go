package tests

import (
	"testing"

	"github.com/gabehardgrave/expect"
)

// These should all panic, unless ran with
// `go test -v ./... -tags=release`

func TestTrue(t *testing.T) {
	expect.True(false, "should be true")
}

func TestFalse(t *testing.T) {
	expect.False(true, "should be false")
}

func TestNil(t *testing.T) {
	x := 1
	expect.Nil(&x, "isn't nil")
}

func TestNotNil(t *testing.T) {
	expect.NotNil(nil, "is nil")
}

func TestZero(t *testing.T) {
	expect.Zero(1, "isn't zero")
}
