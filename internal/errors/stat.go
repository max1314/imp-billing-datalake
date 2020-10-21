// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"context"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Desc returns the description part of the error.
func Desc(e error) string {
	if err, ok := e.(frame); ok {
		return err.message
	}
	if err, ok := e.(list); ok {
		return err.message
	}
	return grpc.ErrorDesc(e)
}

// Code returns the embedded error code.
func Code(e error) codes.Code {
	if e == nil {
		return OK
	}
	if coder, ok := e.(errorCoder); ok {
		return coder.Code()
	}

	if e == context.Canceled {
		return Canceled
	}
	if e == context.DeadlineExceeded {
		return DeadlineExceeded
	}

	if dnsErr, ok := e.(*net.DNSError); ok {
		return dnsErrCode(dnsErr)
	}

	if netOpErr, ok := e.(*net.OpError); ok {
		return netOpErrCode(netOpErr)
	}

	if os.IsNotExist(e) {
		return NotFound
	}
	if os.IsPermission(e) {
		return PermissionDenied
	}
	if os.IsTimeout(e) {
		return DeadlineExceeded
	}

	return grpc.Code(e)
}

// ToGRPC creates a rpcError from the given error based on the code and message
// If the error given implements Code and Desc, a new grpc error will be created based on that
// If the error is generic,
// This can be used together with the Remapper, returning a grpc compatible error
func ToGRPC(e error) error {
	if err, ok := e.(frame); ok {
		return err.ToGRPC()
	}
	if err, ok := e.(list); ok {
		return err.ToGRPC()
	}
	// generic error
	err := grpc.Errorf(Code(e), Desc(e))
	return err
}

// Coder allows you to return codes for errors.
type errorCoder interface {
	error
	Code() codes.Code
}

func dnsErrCode(e *net.DNSError) codes.Code {
	if e.Timeout() {
		return DeadlineExceeded
	}
	if e.Temporary() {
		return Unavailable
	}
	if strings.Contains("no such host", e.Err) {
		// see net.errNoSuchHost
		return NotFound
	}
	return Unknown
}

func netOpErrCode(e *net.OpError) codes.Code {
	if dnsErr, ok := e.Err.(*net.DNSError); ok {
		return dnsErrCode(dnsErr)
	} else if e.Timeout() {
		// i/o timeouts
		return DeadlineExceeded
	} else if e.Temporary() {
		// connection reset by peer
		return Unavailable
	} else if syscallErr, ok := e.Err.(*os.SyscallError); ok {
		// compatibility with Go 1.11 changes, as ECONNRESET/ECONNABORTED for non-accept ops are no longer temporary errors
		// https://github.com/golang/go/commit/7f6105f138b3836e9ad85b8da26d44c742bf217b
		if isConnError(syscallErr.Err) {
			return Unavailable
		}
	}
	return Unknown
}
