// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"context"
	"net"
	"time"

	"imp-billing-datalake/lib/errors"

	"github.com/jpillora/backoff"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	logWritePeriod = 15 * time.Second
)

type DialFunc func() (net.Conn, error)

// ReconnectingWriter wraps a net.Conn such that it retries each write until Close is invoked.
// It also handles too long writes and reconnections.
type ReconnectingWriter struct {
	ctx    context.Context
	cancel context.CancelFunc

	underlyingConn     net.Conn
	rotateConnAt       time.Time
	dialFn             DialFunc
	backoff            *backoff.Backoff
	errLogger          ErrLogger
	connFails          prometheus.Counter
	connCreations      prometheus.Counter
	writeFails         prometheus.Counter
	writeTimeout       time.Duration
	connChan           chan net.Conn
	ongoingConnRequest bool
}

func NewReconnectingWriter(
	dialFn DialFunc,
	errLogger ErrLogger,
	connFails prometheus.Counter,
	connCreations prometheus.Counter,
	writeFails prometheus.Counter,
	writeTimeout time.Duration,
) *ReconnectingWriter {
	ctx, cancel := context.WithCancel(context.Background())
	return &ReconnectingWriter{
		ctx:    ctx,
		cancel: cancel,

		dialFn:    dialFn,
		errLogger: errLogger,
		backoff: &backoff.Backoff{
			Factor: 2,
			Min:    50 * time.Millisecond,
			Max:    2 * time.Second,
		},
		connFails:     connFails,
		connCreations: connCreations,
		writeFails:    writeFails,
		writeTimeout:  writeTimeout,
		// Use the nil value for time, so that the first call to Write updates rotateConnAt
		rotateConnAt:       time.Now().Add(logWritePeriod),
		connChan:           make(chan net.Conn),
		ongoingConnRequest: false,
	}
}

// setConn is responsible for two things:
// 1. getting and setting the connection that our writer will use
// 2. cycling the connection when needed
func (c *ReconnectingWriter) setConn() error {
	now := time.Now()
	// Do I need to request a new connection?
	// If a request is already inflight, no need to fire off another one
	// If there is no request in flight, attempt to create if _either_ we currently have no connection _or_ it is time to force a cycle
	if !c.ongoingConnRequest && (c.underlyingConn == nil || now.After(c.rotateConnAt)) {
		// Ensure we only request one new connection at a time
		c.ongoingConnRequest = true
		go func() {
			for {
				conn, err := c.dialFn()
				if err != nil {
					c.errLogger.Errorf("attempted to create connection but failed: %s", err)
					// Don't absolutely hammer the retries.
					time.Sleep(c.backoff.Duration())
					c.connFails.Inc()
					continue
				}
				c.connCreations.Inc()
				c.connChan <- conn
				return
			}
		}()
	}
	// Do we need to block and wait for a connection? Either...
	// 1. Nothing is there, yes block & wait.
	// 2. New healthy connection is ready to be used - so use it.
	if c.underlyingConn == nil || len(c.connChan) > 0 {
		select {
		// Block until we either have a connection to use
		case conn := <-c.connChan:
			c.resetConn()
			c.rotateConnAt = time.Now().Add(logWritePeriod)
			c.ongoingConnRequest = false
			c.underlyingConn = conn
		// Or the context is cancelled
		case <-c.ctx.Done():
			return c.ctx.Err()
		}
	}
	return nil
}

// Write will run the following retry all of the following in a loop until it manages one successful write:
// * attempt to re-establish a connection if a connection is not available
// * attempt to write the given bytes if a connection is available and exit if it succeeds
// * reset the underlying connection if write fails
//
// Deadline is set for all writes as specified in the flag.
func (c *ReconnectingWriter) Write(b []byte) (int, error) {
	for c.ctx.Err() == nil {
		err := c.setConn()
		if err != nil {
			return 0, err
		}
		n, err := c.send(b)
		if err == nil {
			return n, nil
		}
		c.resetConn()
		c.writeFails.Inc()
		c.errLogger.Warnf("failed to write: %v", errors.Details(err))
	}
	return 0, c.ctx.Err()
}

func (c *ReconnectingWriter) send(b []byte) (int, error) {
	// No way to propagate context, but timeout will ensure we are not there forever.
	if err := c.underlyingConn.SetWriteDeadline(time.Now().Add(c.writeTimeout)); err != nil {
		c.resetConn()
		return 0, errors.Wrap(err, "could not set deadline for write")
	}
	n, err := c.underlyingConn.Write(b)
	if err != nil {
		c.resetConn()
		return 0, errors.Wrap(err, "could not write to external connection")
	}
	return n, nil
}

func (c *ReconnectingWriter) resetConn() {
	if c.underlyingConn == nil {
		return
	}
	if err := c.underlyingConn.Close(); err != nil {
		c.errLogger.Errorf("failed to close connection %v. Resource leak will happen.", errors.Details(err))
	}
	c.underlyingConn = nil
	c.backoff.Reset()
}

func (c *ReconnectingWriter) Close() error {
	c.cancel()
	return nil
}
