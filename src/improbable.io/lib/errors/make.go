// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// New returns a new error, or nil if code is OK. err describes the underlying error if there was
// one. message provides the description of the error. Use Newf to provide format arguments. The
// underlying error can be accessed using the Cause() function.
func New(err error, code codes.Code, message string) error {
	if code == OK {
		return nil
	}
	return newFrame(err, code, message)
}

// Newf returns a new error, or nil if code is OK. err describes the underlying error if there was
// one. The format and arguments will be combined to the error message. The underlying error can be
// accessed using the Cause() function.
func Newf(err error, code codes.Code, format string, args ...interface{}) error {
	if code == OK {
		return nil
	}
	return newFrame(err, code, fmt.Sprintf(format, args...))
}

// Wrap returns an error that wraps and replaces an existing error. The error code of the underlying
// error is kept intact. If no code is available, `Unknown` will be used. message provides the
// description of the returned error. The wrapped error can be accessed using the Cause() function.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return newFrame(err, Code(err), message)
}

// Wrapf returns an error that wraps and replaces an existing error. The error code of the
// underlying error is kept intact. If no code is available, `Unknown` will be used. The format and
// arguments will be combined to the error message. The wrapped error can be accessed using the
// Cause() function.
func Wrapf(e error, format string, a ...interface{}) error {
	if e == nil {
		return nil
	}
	return newFrame(e, Code(e), fmt.Sprintf(format, a...))
}

// Concat returns an error that wraps an existing error, combining the descriptions. The error code
// of the underlying error is kept intact. If no code is available, `Unknown` will be used. message
// is concatenated with the wrapped error's description to provide the description of the returned
// error. The wrapped error can be accessed using the Cause() function.
func Concat(e error, message string) error {
	if e == nil {
		return nil
	}
	return newFrame(e, Code(e), message+"; "+Desc(e))
}

// Concat returns an error that wraps an existing error, combining the descriptions. The error code
// of the underlying error is kept intact. If no code is available, `Unknown` will be used. The
// format and arguments will be combined and concatenated with the wrapped error's description to
// provide the description of the returned error. The wrapped error can be accessed using the
// Cause() function.
func Concatf(e error, format string, a ...interface{}) error {
	if e == nil {
		return nil
	}
	return newFrame(e, Code(e), fmt.Sprintf(format, a...)+"; "+Desc(e))
}
