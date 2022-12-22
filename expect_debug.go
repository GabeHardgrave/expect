//go:build !release
// +build !release

package expect

import "fmt"

func fail(msgArgs ...any) {
	format, args := msgArgs[0], msgArgs[1:]
	if fmtStr, ok := format.(string); ok {
		panic(fmt.Errorf(fmtStr, args...))
	} else {
		panic(fmt.Errorf("%v", msgArgs))
	}
}

func eval(cond bool, msgArgs ...any) {
	if !cond {
		fail(msgArgs...)
	}
}
