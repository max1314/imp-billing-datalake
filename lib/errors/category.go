// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// Category categorises the error. It is very closely related to the remapper interceptor. Errors of type System and Transient will be wrapped in another error to hide the underlying cause.
type Category int8

//go:generate stringer -type Category

const (
	// These codes are grouped together to express what was the underlying cause of the error.
	_ Category = iota
	User
	Transient
	System
	None
)

// ErrorCause identifies (based on error Codes) who is responsible for the problem.
func Categorise(e error) Category {
	if e == nil {
		return None
	}

	switch Code(e) {
	case codes.InvalidArgument,
		codes.NotFound,
		codes.AlreadyExists,
		codes.PermissionDenied,
		codes.Unauthenticated,
		codes.OutOfRange,
		codes.FailedPrecondition,
		codes.ResourceExhausted,
		codes.Canceled,
		codes.DeadlineExceeded:
		return User

	case codes.Aborted,
		codes.Unavailable:
		return Transient
	case codes.OK:
		return None
	}
	return System
}

func CategoriseHTTPStatus(httpStatus int) Category {
	switch httpStatus {
	case http.StatusRequestTimeout, http.StatusGatewayTimeout, http.StatusConflict, http.StatusServiceUnavailable:
		return Transient
	case http.StatusOK:
		return None
	}
	if httpStatus >= 200 && httpStatus < 400 {
		return None
	}
	if httpStatus >= 400 && httpStatus < 500 {
		return User
	}
	return System
}
