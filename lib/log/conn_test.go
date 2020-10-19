// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"context"
	"net"
	"testing"
	"time"

	"imp-billing-datalake/lib/errors"

	"github.com/fortytw2/leaktest"
	"github.com/jpillora/backoff"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestReconnectingWriter_WriteDoesNotReturnUntilWriteIsSuccessful(t *testing.T) {
	defer leaktest.CheckTimeout(t, 10*time.Second)()

	writerWithError := &mockWriterCloser{errOnWrite: errors.New(nil, errors.Internal, "some expected error")}
	writerWithoutError := &mockWriterCloser{}
	dialReturns := []net.Conn{
		writerWithError,
		writerWithoutError,
	}
	callCount := 0

	dial := func() (net.Conn, error) {
		if callCount >= len(dialReturns) {
			return nil, errors.New(nil, errors.Internal, "called more than call count")
		}
		ret := dialReturns[callCount]
		callCount++
		return ret, nil
	}

	errLogger, _ := test.NewNullLogger()

	writer := NewReconnectingWriter(
		dial,
		errLogger,
		logstashConnFailuresCounter,
		logstashConnCreationsCounter,
		logstashWriteFailuresCounter,
		2*time.Second,
	)
	defer func() { _ = writer.Close() }()

	var byteCount int
	var err error
	writeDone := make(chan struct{})
	body := []byte("some random data")
	go func() {
		byteCount, err = writer.Write(body)
		close(writeDone)
	}()

	select {
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for write to complete")
	case <-writeDone:
	}

	assert.Equal(t, 1, writerWithError.writeCallCount)
	assert.Equal(t, 1, writerWithoutError.writeCallCount)
	// verifies that the write is completed using the new writer
	assert.Equal(t, len(body), byteCount)
	assert.NoError(t, err, "must have successfully written a log after redial")
}

func TestReconnectingWriter_OnFirstWrite_DialCalledOnce(t *testing.T) {
	writerWithoutError := &mockWriterCloser{}
	dial := func() (net.Conn, error) {
		return writerWithoutError, nil
	}

	errLogger, _ := test.NewNullLogger()

	ctx, cancel := context.WithCancel(context.Background())
	writer := &ReconnectingWriter{
		ctx:    ctx,
		cancel: cancel,

		dialFn:    dial,
		errLogger: errLogger,
		backoff: &backoff.Backoff{
			Factor: 2,
			Min:    50 * time.Millisecond,
			Max:    2 * time.Second,
		},
		connFails:     logstashConnFailuresCounter,
		connCreations: logstashConnCreationsCounter,
		writeTimeout:  2 * time.Second,
		// Use the nil value for time, so that the first call to Write updates rotateConnAt
		rotateConnAt: time.Now(),
		connChan:     make(chan net.Conn, 1),
	}

	body := []byte("some random data")
	_, err := writer.Write(body)
	assert.NoError(t, err)

	// Did we dial on the first write?
	assert.Equal(t, 1, writerWithoutError.writeCallCount)
}

func TestReconnectingWriter_WhenPeriodExpires_DialNewConnection(t *testing.T) {
	writerWithoutError := &mockWriterCloser{writeCallCount: 0}
	dial := func() (net.Conn, error) {
		return writerWithoutError, nil
	}

	errLogger, _ := test.NewNullLogger()

	ctx, cancel := context.WithCancel(context.Background())
	now := time.Now()
	writer := &ReconnectingWriter{
		ctx:    ctx,
		cancel: cancel,

		dialFn:    dial,
		errLogger: errLogger,
		backoff: &backoff.Backoff{
			Factor: 2,
			Min:    50 * time.Millisecond,
			Max:    2 * time.Second,
		},
		connFails:     logstashConnFailuresCounter,
		connCreations: logstashConnCreationsCounter,
		writeTimeout:  2 * time.Second,
		// Ensure we will try and rotate connections after the period has expired
		rotateConnAt: now.Add(-time.Hour),
		connChan:     make(chan net.Conn, 1),
	}

	body := []byte("some random data")
	_, err := writer.Write(body)
	assert.NoError(t, err)

	// Simulate the case when the rotateConnAt time has passed
	writer.rotateConnAt = time.Now().Add(-time.Hour)

	_, err = writer.Write(body)
	assert.NoError(t, err)

	// Did we re-dial on the second write?
	assert.Equal(t, 2, writerWithoutError.writeCallCount)
}

type mockWriterCloser struct {
	writeCallCount int
	errOnWrite     error
	errOnClose     error
}

func (w *mockWriterCloser) Write(body []byte) (int, error) {
	w.writeCallCount += 1
	if w.errOnWrite != nil {
		return 0, w.errOnWrite
	}
	return len(body), nil
}

func (w *mockWriterCloser) Close() error {
	return w.errOnClose
}

func (w *mockWriterCloser) Read(b []byte) (n int, err error) {
	panic("implement me")
}

func (w *mockWriterCloser) LocalAddr() net.Addr {
	panic("implement me")
}

func (w *mockWriterCloser) RemoteAddr() net.Addr {
	panic("implement me")
}

func (w *mockWriterCloser) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (w *mockWriterCloser) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (w *mockWriterCloser) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestReconnectingWriter_DoesntLeak_WhenConnectionsExpire(t *testing.T) {
	mockFlakyWriter := &mockFlakyWriterCloser{}
	dialer := func() (net.Conn, error) {
		return mockFlakyWriter, nil
	}

	errLogger, _ := test.NewNullLogger()

	ctx, cancel := context.WithCancel(context.Background())
	writer := &ReconnectingWriter{
		ctx:    ctx,
		cancel: cancel,

		dialFn:    dialer,
		errLogger: errLogger,
		backoff: &backoff.Backoff{
			Factor: 2,
			Min:    50 * time.Millisecond,
			Max:    2 * time.Second,
		},
		connFails:     logstashConnFailuresCounter,
		connCreations: logstashConnCreationsCounter,
		writeFails:    logstashConnFailuresCounter,
		writeTimeout:  2 * time.Second,
		// Ensure we will try and rotate connections after the period has expired
		connChan: make(chan net.Conn),
	}
	body := []byte("some random data")

	// first write
	// we should create a new connection because it's always nil at start
	// the write succeeds
	_, err := writer.Write(body)
	assert.NoError(t, err)

	// force a connection recycle
	writer.rotateConnAt = time.Now()

	// second write
	// it's time to get a new connection so we schedule that
	// attempt to write, which will fail because the connection from first write has expired
	// it should block until the previously scheduled connection is acquired, and successfully write on that
	_, err = writer.Write(body)
	assert.NoError(t, err)

	// we should not have any inflight requests
	select {
	case <-writer.connChan:
		t.Fatal("found a hanging connection")
	case <-time.After(time.Second):
		return
	}
}

type mockFlakyWriterCloser struct {
	writeCallCount int
}

func (w *mockFlakyWriterCloser) Write(body []byte) (int, error) {
	w.writeCallCount += 1
	switch w.writeCallCount {
	case 2, 4:
		return 0, errors.New(nil, errors.Unknown, "server closed the connection")
	default:
		return len(body), nil
	}
}

func (w *mockFlakyWriterCloser) Close() error {
	return nil
}

func (w *mockFlakyWriterCloser) Read(b []byte) (n int, err error) {
	panic("implement me")
}

func (w *mockFlakyWriterCloser) LocalAddr() net.Addr {
	panic("implement me")
}

func (w *mockFlakyWriterCloser) RemoteAddr() net.Addr {
	panic("implement me")
}

func (w *mockFlakyWriterCloser) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (w *mockFlakyWriterCloser) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (w *mockFlakyWriterCloser) SetWriteDeadline(t time.Time) error {
	return nil
}
