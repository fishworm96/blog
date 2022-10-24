package redis

import "errors"

var (
	ErrorCodeInvalid = errors.New("验证码已失效，请重新获取验证码")
	ErrorCodeIncorrect = errors.New("验证码不正确")
)