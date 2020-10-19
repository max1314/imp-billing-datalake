// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Panic-related error handling methods.

package errors

import (
	"runtime"
)

// MaybePanicRecover tries to perform a recover() and if there was a panic, returns an Internal error.
func MaybePanicRecover(errorOutput *error) {
	if v := recover(); v != nil {
		*errorOutput = fmtPanic(v)
	}
}

// MaybePanicRecoverCh tries to perform a recover() and if there was a panic, writes an error to the channel.
// example:
// testErr := make(chan error, 1)
// go func() {
// 		defer close(testErr)
//		defer errors.MaybePanicRecoverCh(testErr)
//		some code
// }
// err = <-testErr
func MaybePanicRecoverCh(errorOutput chan<- error) {
	if v := recover(); v != nil {
		errorOutput <- fmtPanic(v)
	}
}

func fmtPanic(v interface{}) error {
	trace := make([]byte, 1<<16)
	n := runtime.Stack(trace, true)
	return Newf(nil, Internal, "panic recover\n %v\n stack trace %d bytes\n %s", v, n, trace[:n])
}
