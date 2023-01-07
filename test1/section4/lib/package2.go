package lib

import "go-study/section4/lib/user"

func CheckNum1(c int32) bool {
	user.User()
	return c > 10
}

func CheckNum2(c int32) bool {
	return c > 100
}
