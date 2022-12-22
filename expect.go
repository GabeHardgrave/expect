package expect

func True(cond bool, msgArgs ...any) {
	eval(cond, msgArgs...)
}

func False(cond bool, msgArgs ...any) {
	eval(!cond, msgArgs...)
}

func Nil(ptr any, msgArgs ...any) {
	eval(ptr == nil, msgArgs...)
}

func NotNil(ptr any, msgArgs ...any) {
	eval(ptr != nil, msgArgs...)
}

func Zero(x int, msgArgs ...any) {
	eval(x == 0, msgArgs...)
}
