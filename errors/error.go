package errors

import (
	"fmt"
)

type Error struct {
	Code int32
	Msg  string
}

func NewError(code int32) *Error {
	return &Error{
		Code: code,
	}
}

func NewErrorf(code int32, msg string, params ...interface{}) *Error {
	return &Error{
		Code: code,
		Msg:  fmt.Sprintf(msg, params...),
	}
}

func (c *Error) Error() string {
	return fmt.Sprintf("Error: %d: %s", c.Code, c.Msg)
}

func Code(err error) int32 {
	if err != nil {
		if errs, ok := err.(*Error); ok {
			return errs.Code
		}
	}
	return 0
}

func Msg(err error) string {
	if err != nil {
		if errs, ok := err.(*Error); ok {
			return errs.Msg
		}
	}
	return ""
}
