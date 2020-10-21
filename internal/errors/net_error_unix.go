// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package errors

import "syscall"

func isConnError(err error) bool {
	if se, ok := err.(syscall.Errno); ok {
		return se == syscall.ECONNRESET || se == syscall.ECONNABORTED
	}
	return false
}
