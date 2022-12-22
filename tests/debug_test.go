//go:build !release
// +build !release

package tests

import (
	"testing"

	"github.com/gabehardgrave/expect"
	"github.com/stretchr/testify/require"
)

func TestDoesntPanic(t *testing.T) {
	expect.True(true)
	expect.False(false)

	x := 1
	expect.Nil(nil)
	expect.NotNil(&x)

	expect.Zero(0)
}

func TestPanics(t *testing.T) {
	require.PanicsWithError(t, "should be true", func() {
		expect.True(false, "should be true")
	})

	require.PanicsWithError(t, "should be false", func() {
		expect.False(true, "should be false")
	})

	require.PanicsWithError(t, "1 isn't nil", func() {
		x := 1
		expect.Nil(&x, "%v isn't nil", x)
	})

	require.PanicsWithError(t, "is nil", func() {
		expect.NotNil(nil, "is nil")
	})

	require.PanicsWithError(t, "isn't zero", func() {
		expect.Zero(1, "isn't zero")
	})
}
