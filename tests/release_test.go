//go:build release
// +build release

package tests

import (
	"testing"

	"github.com/gabehardgrave/expect"
)

func TestDoesntPanic(t *testing.T) {
	expect.True(false, "should be true")
	expect.False(true, "should be false")

	x := 1
	expect.Nil(&x, "isn't nil")
	expect.NotNil(nil, "is nil")

	expect.Zero(1, "isn't zero")
}
