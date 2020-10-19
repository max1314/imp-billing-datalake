// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

type HttpErr struct {
	code int
	msg  string
}

func NewHTTPError(code int, msg string) *HttpErr {
	return &HttpErr{code: code, msg: msg}
}

func (e *HttpErr) Code() int {
	return e.code
}

func (e *HttpErr) Error() string {
	return e.msg
}
