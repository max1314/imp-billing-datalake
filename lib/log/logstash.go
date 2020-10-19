// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Hook for talking to logstash remotely.

package log

import (
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"imp-billing-datalake/lib/errors"
	"imp-billing-datalake/lib/sharedflags"

	"github.com/sirupsen/logrus"
)

var (
	fLogstashWriteTimeout = sharedflags.Set.Duration(
		"logstash_write_timeout",
		2*time.Second,
		"time to wait for a succesful write to logstash before retrying",
	)
	fLogstashWriteBufferSize = sharedflags.Set.Int(
		"logstash_write_buffer_size",
		5000,
		"size of the buffer for log entries for the async writer",
	)
)

// ErrLogger is used to report error messages about connection errors and failures to send logs to logstash.
type ErrLogger interface {
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// LogStashHook sends logs to LogStash.
type LogStashHook struct {
	conn                io.Writer
	buffer              chan *Entry
	errLogger           ErrLogger
	formatter           Formatter
	closing             chan struct{}
	closed              chan struct{}
	pending             sync.WaitGroup
	lastBufferFullFired time.Time
}

// NewLogStashHook creates a hook to be added to an instance of logger. It connects
// to LogStash on host via TCP.
func NewLogStashHook(host string, port int, errLogger ErrLogger) (*LogStashHook, error) {
	dial := func() (net.Conn, error) {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), 2*time.Second)
		if err != nil {
			return nil, errors.Wrap(err, "could not establish logstash connection")
		}

		return conn, nil
	}

	// we ignore error because we will try to reconnect inside Fire() if there is no connection
	hostname, _ := os.Hostname()
	formatter := &LogstashFormatter{Hostname: hostname}
	hook := newLogStashHook(
		NewReconnectingWriter(dial, errLogger, logstashConnFailuresCounter, logstashConnCreationsCounter, logstashWriteFailuresCounter, *fLogstashWriteTimeout),
		formatter,
		errLogger,
	)

	go hook.start()

	return hook, nil
}

func newLogStashHook(writer io.Writer, formatter Formatter, errLogger ErrLogger) *LogStashHook {
	if errLogger == nil {
		errLogger = NewNullLogger()
	}
	return &LogStashHook{
		conn:      writer,
		buffer:    make(chan *Entry, *fLogstashWriteBufferSize),
		closing:   make(chan struct{}),
		closed:    make(chan struct{}),
		pending:   sync.WaitGroup{},
		formatter: formatter,
		errLogger: errLogger,
	}
}

// Fire implements the Fire method from the Hook interface. It writes to its local buffer for the Entry to be sent
// off asynchronously. If the buffer is full, it will drop the entry.
func (hook *LogStashHook) Fire(entry *Entry) error {
	hook.pending.Add(1)
	defer hook.pending.Done()
	select {
	// note: even though the channel might be closed, we can still read from it.
	case <-hook.closing:
		return errors.New(nil, errors.Unavailable, "cannot fire log entry when the hook is closed")
	default:
	}

	reportEntry(entry.Level)
	select {
	case hook.buffer <- entry:
	default:
		// Only log bufferFull errors once every 5 seconds, to avoid spamming our output (which renders it useless).
		if time.Since(hook.lastBufferFullFired).Seconds() > 5 {
			hook.lastBufferFullFired = time.Now()
			hook.errLogger.Errorf("Failed to write log message due to full buffer.")
		}
		reportDropped(entry.Level, BufferFull)
	}
	return nil
}

// Levels implements the Levels method from the Hook interface.
func (hook *LogStashHook) Levels() []Level {
	return []Level{
		PanicLevel,
		FatalLevel,
		ErrorLevel,
		WarnLevel,
		InfoLevel,
		DebugLevel,
	}
}

func (hook *LogStashHook) Close() {
	hook.internalClose()
	<-hook.closed
}

func (hook *LogStashHook) start() {
	defer close(hook.closed)
	for entry := range hook.buffer {
		hook.sendMsg(entry)
	}
}

func (hook *LogStashHook) sendMsg(entry *Entry) {
	payload, err := hook.formatter.Format((*logrus.Entry)(entry))
	if err != nil {
		hook.errLogger.Errorf("Failed to write log message due to bad format: %v", errors.Details(err))
		reportDropped(entry.Level, BadFormat)
		return
	}
	_, err = hook.conn.Write(payload)
	if err != nil {
		reportDropped(entry.Level, FailedToWrite)
		hook.errLogger.Errorf("Failed to write log message to logstash due to connection issues: %v", errors.Details(err))
		return
	}
	reportRemoteSuccess(entry.Level)
}

func (hook *LogStashHook) internalClose() {
	select {
	// note: even though the channel might be closed, we can still read from it.
	case <-hook.closing:
		<-hook.closed
		return
	default:
	}
	// Signal that we're closing, so we accept no new messages.
	close(hook.closing)

	// Wait until all pending fires put all items into buffer.
	hook.pending.Wait()

	// At that point we close potentially partially filled channel. Start loop will be able to read
	// all the items until empty and finished channel will be closed notifying that graceful shutdown finished.
	close(hook.buffer)

	<-hook.closed
}
