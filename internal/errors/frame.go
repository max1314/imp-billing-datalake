// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/max1314/imp-billing-datalake/internal/sharedflags"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var fEnableFilenameAndLineNumber = sharedflags.Set.Bool("errors_include_debug_info", false, "show filename and line number in error trace")
var fEnableStackTrace = sharedflags.Set.Bool("errors_include_stack_trace", false, "show stack trace in error trace")

// This function must not return nil as it uses a concrete error type in its signature
func newFrame(err error, c codes.Code, message string) frame {
	// you never want to remap context cancelled into anything else
	if Code(err) == Canceled {
		c = Canceled
	}
	e := frame{
		code:    c,
		Err:     err,
		message: message,
	}
	e.setLocation(2)
	if *fEnableStackTrace {
		// Add stack trace only to the first error in the cause list.
		if _, ok := err.(frame); !ok {
			e.stackTrace = string(debug.Stack())
		}
	}
	return e
}

// frame is a frame in a stack of errors with optional information about location and underlying cause
type frame struct {
	// code categorises the error and makes it amenable for use by grpc clients/servers
	code codes.Code

	// Err is the underlying error that can be nil if this is a base error
	// The field name matches the convention of Go standard library types like os.SyscallError
	Err error

	// message is the error string
	message string

	// file is the name of the file that created this error
	file string

	// line is the line number in the file that created this error
	line int

	stackTrace string
}

type list struct {
	frame

	Errs []error
}

func (e frame) Error() string {
	// This follows the format of grpc's error desc
	return fmt.Sprintf("code = %s desc = %s", e.code.String(), e.message)
}

func (e frame) FileAndLine() (string, int) {
	return e.file, e.line
}

func (e frame) Code() codes.Code {
	return e.code
}

// ToGRPC creates a rpcError from the given error based on the code and message
// This can be used together with the Remapper, returning a grpc compatible error
func (e frame) ToGRPC() error {
	return grpc.Errorf(e.code, e.message)
}

// setLocation records the source location of the error by setting
// e.Location, at `callDepth` stack frames above the call.
func (e *frame) setLocation(callDepth int) {
	if *fEnableFilenameAndLineNumber {
		_, file, line, _ := runtime.Caller(callDepth + 1)
		e.file, e.line = file, line
	}
}

func (l list) Len() int {
	return len(l.Errs)
}
