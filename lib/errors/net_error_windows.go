// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import "syscall"

func isConnError(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.WSAECONNRESET || se == syscall.WSAECONNABORTED
	}
	return false
}
