// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"context"
)

// DropCtxErr ignores the original err if the context has the same error code.
// This can be used in functions that run indefinitely until canceled, in which case the cancel is not an error.
// Example:
//   func(ctx context.Context) error {
//     err := WatchChanges(ctx, params...)
//     return DropCtxErr(ctx, err)
//   }
func DropCtxErr(ctx context.Context, err error) error {
	if Code(err) == Code(ctx.Err()) {
		return nil
	}
	return err
}

// CheckLocalDeadline returns an Unavailable error if the given error was caused by a local context's deadline.
// This can be used in long-running RPCs (e.g. streams, watches) when individual backend calls need a shorter deadline.
// In such a case, the backend call may return DeadlineExceeded which cannot be returned to the caller as the RPC's
// deadline may not have been exceeded.
// Example:
//   func(ctx context.Context, params...) error {
//     localCtx, cancel := context.WithTimeout(10 * time.Second)
//     defer cancel()
//     err := someBackend(localCtx, args...)
//     if err2 := CheckLocalDeadline(localCtx, ctx, err, "backend call stalled"); err2 != nil {
//       return err2
//     }
//     ...
//   }
func CheckLocalDeadline(localCtx context.Context, parentCtx context.Context, err error, message string) error {
	if localCtx.Err() != context.DeadlineExceeded {
		return nil
	}
	if parentCtx.Err() == context.DeadlineExceeded {
		return nil
	}
	return New(err, Unavailable, message)
}
