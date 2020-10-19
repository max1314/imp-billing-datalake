// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// We don't use an automatically generated mock here as it would introduce an import cycle
// that would be require significant splitting in the 'improbable.io/lib/log' package which
// is something that we are not ready to do. Furthermore use of the mock is extremely limited.
type mockErrLogger struct {
	wasCalledWithError bool
}

func (l *mockErrLogger) Warnf(format string, args ...interface{})  {}
func (l *mockErrLogger) Errorf(format string, args ...interface{}) { l.wasCalledWithError = true }

func (l *mockErrLogger) assertCalledWithError(t *testing.T) {
	assert.True(t, l.wasCalledWithError, "Should have been called with Errorf.")
}
func (l *mockErrLogger) assertNotCalledWithError(t *testing.T) {
	assert.False(t, l.wasCalledWithError, "Should have been called with Errorf.")
}

func TestWhenCloseIsCalledAndEntriesAreInBuffer_ShouldWriteEntries(t *testing.T) {
	fakeWriter, hook := setup(*fLogstashWriteBufferSize)
	mockLogger := &mockErrLogger{}
	defer mockLogger.assertNotCalledWithError(t)
	hook.errLogger = mockLogger
	expectedCount := 3
	expectedMessages := arrangeExpectations(expectedCount, hook)

	closeFinished := make(chan struct{})
	go func() {
		<-fakeWriter.blocked
		hook.Close()
		close(closeFinished)
	}()

	close(fakeWriter.blocked)
	select {
	case <-closeFinished:
	case <-time.After(5 * time.Second):
		t.Fatal("Timed out")
	}

	assert.Len(t, fakeWriter.written, expectedCount, "Should write the expected number of entries.")
	assert.Equal(t, expectedMessages, fakeWriter.written, "Should write in the correct order.")
}

func TestWhenBufferIsSmallAndCloseIsCalledAndEntriesAreInBuffer_ShouldWriteEntries(t *testing.T) {
	fakeWriter, hook := setup(1)
	mockLogger := &mockErrLogger{}
	defer mockLogger.assertCalledWithError(t)
	hook.errLogger = mockLogger
	_ = arrangeExpectations(6, hook)

	closeFinished := make(chan struct{})
	go func() {
		<-fakeWriter.blocked
		hook.Close()
		close(closeFinished)
	}()

	close(fakeWriter.blocked)
	select {
	case <-closeFinished:
	case <-time.After(5 * time.Second):
		t.Fatal("Timed out")
	}

	assert.True(t, len(fakeWriter.written) >= 1, "Should write at least buffer size entries.")
	assert.Equal(t, "0\n", fakeWriter.written[0], "Should write at least the first entry.")
}

func TestLogStashHook_CloseCanBeCalledMoreThanOnceWithNoPanic(t *testing.T) {
	_, hook := setup(3)

	assert.NotPanics(t, func() {
		hook.internalClose()
		hook.internalClose()
	})
}

func setup(bufferSize int) (*initiallyBlockedWriter, *LogStashHook) {
	fakeWriter := &initiallyBlockedWriter{
		blocked: make(chan struct{}),
	}
	hook := newLogStashHook(fakeWriter, &plainFormatter{}, NewNullLogger())
	hook.buffer = make(chan *Entry, bufferSize)
	go hook.start()
	return fakeWriter, hook
}

func arrangeExpectations(number int, hook *LogStashHook) []string {
	var expectedMessages []string
	for i := 0; i < number; i++ {
		message := fmt.Sprintf("%d", i)
		hook.Fire(&Entry{
			Message: message,
		})
		expectedMessages = append(expectedMessages, message+"\n")
	}
	return expectedMessages
}

type initiallyBlockedWriter struct {
	written []string
	blocked chan struct{}
}

func (f *initiallyBlockedWriter) Write(p []byte) (n int, err error) {
	<-f.blocked
	f.written = append(f.written, string(p))
	return len(p), nil
}

func (f *initiallyBlockedWriter) Close() error {
	return nil
}
