// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"improbable.io/lib/errors"
	"improbable.io/lib/sharedflags"

	githuberrs "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNilError(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	d := error(nil)
	details := errors.Details(d)
	assert.Contains(t, details, "")
}

func TestSingleDebugError(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	d := errors.New(nil, errors.Aborted, "whatever")
	details := errors.Details(d)
	assert.NotContains(t, d.Error(), "lib/errors/errors_test.go:", "error should contain filename and line number")
	assert.Contains(t, d.Error(), "whatever", "error should contain the error message")
	assert.Contains(t, details, "lib/errors/errors_test.go:", "error should contain filename and line number")
	assert.Contains(t, details, "whatever", "error should contain the error message")
}

func TestMultipleDebugErrors(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "true")
	d := errors.New(nil, errors.Aborted, "A")
	e := errors.New(d, errors.NotFound, "B")
	details := errors.Details(e)
	assert.NotContains(t, e.Error(), "lib/errors/errors_test.go:", "error should contain filename and line number")
	assert.NotContains(t, e.Error(), "A", "error should contain the error message")
	assert.Contains(t, e.Error(), "B", "error should contain the error message")
	assert.Contains(t, details, "lib/errors/errors_test.go:", "error should contain filename and line number")
	assert.Contains(t, details, "A", "error should contain the error message")
	assert.Contains(t, details, "B", "error should contain the error message")
}

func TestNonDebugErrors(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	d := errors.New(nil, errors.Aborted, "whatever")
	assert.NotContains(t, d.Error(), "lib/errors/errors_test.go:", "error should not contain filename and line number")
	assert.Contains(t, d.Error(), "whatever", "error should contain the error message")
}

func TestWrapping_WrappingGenericError_ReturnsUnknownErrWithoutConcat(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	err := fmt.Errorf("Generic error")
	wrappedError := errors.Wrap(err, "Outer message")
	assert.Equal(t, "code = Unknown desc = Outer message", wrappedError.Error(), "Wrapping generic error does not print out the correct message")
	assert.Equal(t, errors.Code(wrappedError), errors.Unknown, "generic errors should be mapped to unknown")
}

func TestWrappingWithContext_WrappingGenericError_ReturnsUnknownErrWithConcatenatedMessage(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	err := fmt.Errorf("Generic error")
	wrappedError := errors.Concat(err, "Outer message")
	assert.Equal(t, "code = Unknown desc = Outer message; Generic error", wrappedError.Error(), "Wrapping generic error does not print out the correct message")
	assert.Equal(t, errors.Code(wrappedError), errors.Unknown, "generic errors should be mapped to unknown")
}

func TestWrapping_WrappingRPCError_RetainsErrorCode(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	err := status.Errorf(codes.Unimplemented, "rpc error")
	wrappedError := errors.Wrap(err, "Outer message")
	assert.Equal(t, "code = Unimplemented desc = Outer message", wrappedError.Error(), "Wrapping rpc error does not print out the correct message")
	assert.Equal(t, errors.Code(wrappedError), errors.Unimplemented, "wrapped rpc errors should retain error code")
}

func TestWrappingWithConcat_WrappingRPCError_RetainsErrorCodeWithConcatenatedMessage(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	err := status.Errorf(codes.Unimplemented, "rpc error")
	wrappedError := errors.Concat(err, "Outer message")
	assert.Equal(t, "code = Unimplemented desc = Outer message; rpc error", wrappedError.Error(), "Wrapping rpc error does not print out the correct message")
	assert.Equal(t, errors.Code(wrappedError), errors.Unimplemented, "wrapped rpc errors should retain error code")
}

func TestWrapping_WrappingExistingErr_RetainsErrorCode(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	err := errors.New(nil, errors.Unimplemented, "some error")
	wrappedError := errors.Wrap(err, "Outer message")
	assert.Equal(t, "code = Unimplemented desc = Outer message", wrappedError.Error(), "Wrapping existing error does not print out the correct message")
	assert.Equal(t, errors.Code(wrappedError), errors.Unimplemented, "wrapped existing errors should retain error code")
}

func TestWrappingWithConcat_WrappingExistingErr_RetainsErrorCodeWithConcatenatedMessage(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	err := errors.New(nil, errors.Unimplemented, "some error")
	wrappedError := errors.Concat(err, "Outer message")
	assert.Equal(t, "code = Unimplemented desc = Outer message; some error", wrappedError.Error(), "Wrapping existing error does not print out the correct message")
	assert.Equal(t, errors.Code(wrappedError), errors.Unimplemented, "wrapped existing errors should retain error code")
}

func TestCause_UnwrapsCommonErrorTypes(t *testing.T) {
	simple := errors.New(nil, errors.InvalidArgument, "badness")
	require.Error(t, simple)

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network string, target string) (net.Conn, error) {
				return nil, simple
			},
		},
	}
	_, httpErr := client.Get("http://invalidhost1234aa.com")
	require.Error(t, httpErr)

	for _, tc := range []struct {
		base    error
		wrapped error
	}{
		{&net.DNSError{}, &net.OpError{Err: &net.DNSError{}}},
		{simple, &json.MarshalerError{Err: simple}},
		{syscall.EINVAL, &os.SyscallError{Err: syscall.EINVAL}},
		{simple, errors.New(simple, errors.Unknown, "woops")},
		{simple, httpErr},
		// github.com/pkg/errors.Wrap double-wraps
		{simple, errors.Cause(githuberrs.Wrap(simple, "woops"))},
	} {
		assert.Equal(t, tc.base, errors.Cause(tc.wrapped))
	}
}

func TestDesc(t *testing.T) {
	comboError := &errors.ComboError{}
	comboError.Add(errors.New(nil, errors.NotFound, "Err1"))
	comboError.Add(errors.New(nil, errors.Unknown, "Err2"))

	for name, spec := range map[string]struct {
		err            error
		expectedString string
	}{
		"single error": {
			err:            errors.New(nil, errors.NotFound, "Err1"),
			expectedString: "Err1",
		},
		"chained error": {
			err:            errors.New(errors.New(nil, errors.NotFound, "Err1"), errors.Unknown, "Err2"),
			expectedString: "Err2",
		},
		"wrapped error": {
			err:            errors.Wrap(errors.New(nil, errors.NotFound, "Err1"), "Err2"),
			expectedString: "Err2",
		},
		"combo error": {
			err:            comboError.NilOrError(),
			expectedString: "Err1\n\tErr2",
		},
	} {
		actual := errors.Desc(spec.err)
		assert.Equal(t, spec.expectedString, actual, "case (%s) did not produce the right error description", name)
	}
}

func TestToGRPC(t *testing.T) {
	comboError := &errors.ComboError{}
	comboError.Add(errors.New(nil, errors.NotFound, "Err1"))
	comboError.Add(errors.New(nil, errors.Unknown, "Err2"))

	for name, spec := range map[string]struct {
		err            error
		expectedString string
	}{
		"single error": {
			err:            errors.New(nil, errors.NotFound, "Err1"),
			expectedString: "rpc error: code = NotFound desc = Err1",
		},
		"chained error": {
			err:            errors.New(errors.New(nil, errors.NotFound, "Err1"), errors.Unknown, "Err2"),
			expectedString: "rpc error: code = Unknown desc = Err2",
		},
		"wrapped error": {
			err:            errors.Wrap(errors.New(nil, errors.NotFound, "Err1"), "Err2"),
			expectedString: "rpc error: code = NotFound desc = Err2",
		},
		"wrapped with concat error": {
			err:            errors.Concat(errors.New(nil, errors.NotFound, "Err1"), "Err2"),
			expectedString: "rpc error: code = NotFound desc = Err2; Err1",
		},
		"combo error": {
			err:            comboError.NilOrError(),
			expectedString: "rpc error: code = Unknown desc = Err1\n\tErr2",
		},
	} {
		actual := errors.ToGRPC(spec.err).Error()
		assert.Equal(t, spec.expectedString, actual, "case (%s) did not produce the right rpc error", name)
	}
}

func TestDetails(t *testing.T) {
	comboError := &errors.ComboError{}
	comboError.Add(errors.New(nil, errors.NotFound, "Err1"))
	comboError.Add(errors.New(nil, errors.Unknown, "Err2"))

	for name, spec := range map[string]struct {
		err            error
		expectedString string
	}{
		"single error": {
			err:            errors.New(nil, errors.NotFound, "Err1"),
			expectedString: "{code = NotFound desc = Err1}",
		},
		"chained error": {
			err:            errors.New(errors.New(nil, errors.NotFound, "Err1"), errors.Unknown, "Err2"),
			expectedString: "{code = Unknown desc = Err2}\n{code = NotFound desc = Err1}",
		},
		"wrapped error": {
			err:            errors.Wrap(errors.New(nil, errors.NotFound, "Err1"), "Err2"),
			expectedString: "{code = NotFound desc = Err2}\n{code = NotFound desc = Err1}",
		},
		"wrapped with concat error": {
			err:            errors.Concat(errors.New(nil, errors.NotFound, "Err1"), "Err2"),
			expectedString: "{code = NotFound desc = Err2; Err1}\n{code = NotFound desc = Err1}",
		},
		"combo error": {
			err:            comboError.NilOrError(),
			expectedString: "{code = Unknown details = {\n\t{code = NotFound desc = Err1}\n\t{code = Unknown desc = Err2}}}",
		},
	} {
		actual := errors.Details(spec.err)
		assert.Equal(t, spec.expectedString, actual, "case (%s) did not produce the right error details", name)
	}
}

func TestContextErrors(t *testing.T) {
	// Check sane translation of context errors
	assert.Equal(t, errors.Canceled, errors.Code(context.Canceled))
	assert.Equal(t, errors.DeadlineExceeded, errors.Code(context.DeadlineExceeded))
	assert.Equal(t, errors.Canceled, errors.Code(errors.ToGRPC(context.Canceled)))
	assert.Equal(t, context.Canceled.Error(), errors.Desc(errors.ToGRPC(context.Canceled)))
	assert.Equal(t, errors.DeadlineExceeded, errors.Code(errors.ToGRPC(context.DeadlineExceeded)))
	assert.Equal(t, context.DeadlineExceeded.Error(), errors.Desc(errors.ToGRPC(context.DeadlineExceeded)))
}

func TestGenericErrorList(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	l := struct {
		error
		Errors []error
	}{
		error: errors.New(nil, errors.DeadlineExceeded, "1"),
		Errors: []error{
			errors.New(nil, errors.Canceled, "2"),
			errors.New(nil, errors.AlreadyExists, "3"),
		},
	}
	d := strings.Split(errors.Details(l), "\n")
	assert.Equal(t, []string{
		"{code = Unknown details = {",
		"\t{code = Canceled desc = 2}",
		"\t{code = AlreadyExists desc = 3}}}",
	}, d)
}

func TestStackTrace(t *testing.T) {
	sharedflags.Set.Set("errors_include_debug_info", "false")
	sharedflags.Set.Set("errors_include_stack_trace", "true")

	parentErr := errors.New(nil, errors.DeadlineExceeded, "1")
	childErr := errors.Wrap(parentErr, "2")

	d := strings.Split(errors.Details(childErr), "\n")
	assert.Equal(t, "{code = DeadlineExceeded desc = 2}", d[0])
	assert.Equal(t, "{code = DeadlineExceeded desc = 1}", d[1])
	assert.Contains(t, d[2], "goroutine") // start of stack trace.
}
