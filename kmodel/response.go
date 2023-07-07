package kmodel

import (
	"github.com/losemy/kit/errors"
)

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Error(err *errors.Error) Response {
	return Response{
		Code: err.Code,
		Msg:  err.Msg,
	}
}

func Success() Response {
	return Response{
		Code: 0,
		Msg:  "success",
	}
}

func SuccessWithData(data interface{}) Response {
	return Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}
