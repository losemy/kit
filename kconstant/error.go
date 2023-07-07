package kconstant

import (
	"github.com/losemy/kit/errors"
)

// 使用负数定义非业务异常
const (
	CodeSystem    = -1
	CodeTimestamp = -2
	CodeSign      = -3
)

var (
	ErrSystem    = errors.NewErrorf(CodeSystem, "system err")
	ErrTimestamp = errors.NewError(CodeTimestamp)
	ErrSign      = errors.NewError(CodeSign)
)
