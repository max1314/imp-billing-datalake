// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogger_WhenLogStashHookIsClosed_DoesNotPanic(t *testing.T) {
	logger := NewNullLogger()
	hook := newLogStashHook(&fakeWriter{}, &plainFormatter{}, NewNullLogger())
	logger.AddHook(hook)
	go hook.start()

	logger.Info("amaze")
	closeFinished := make(chan struct{})
	go assert.NotPanics(t, func() {
		logger.Close()
		close(closeFinished)
	})
	select {
	case <-closeFinished:
	case <-time.After(5 * time.Second):
		assert.Fail(t, "Timed out, probably deadlocked")
	}
}

type fakeWriter struct{}

func (*fakeWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (*fakeWriter) Close() error {
	return nil
}

func TestLogger_WhenRemoteLoggerIsCleanedUp_NoBufferLeaks(t *testing.T) {
	logger := New("Test remote logger clean up")
	hook := newLogStashHook(&fakeWriter{}, &plainFormatter{}, New("error logger"))
	go hook.start()
	logger.AddHook(hook)
	logger.writer.remote = true

	logger.Warn("Some test message")

	// run garbage collector, should not clean up logger
	runtime.GC()

	logger.Warn("Some other test message")

	// run garbage collector, should clean up logger as it will not be referenced again
	runtime.GC()

	// channel should be closed
	_, ok := <-hook.buffer
	assert.False(t, ok)
}
