package errors

import "errors"

var (
	ErrSuiZkLoginUnsupported = errors.New("sui zklogin unsupported")
	ErrBadCaptcha            = errors.New("bad captcha")
)
