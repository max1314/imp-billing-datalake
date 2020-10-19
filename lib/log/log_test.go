// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"testing"

	"imp-billing-datalake/lib/errors"

	"github.com/stretchr/testify/require"
)

func newLoggerForTest(b *bytes.Buffer) *Logger {
	l := New("test")
	l.SetOut(b)
	l.SetFormatter(&LogstashFormatter{})
	return l
}

func TestLoggingJSON(t *testing.T) {
	b := &bytes.Buffer{}
	l := newLoggerForTest(b)

	msg := "testing"
	l.Info(msg)

	t.Log(b.String())

	var proto map[string]interface{}
	err := json.Unmarshal(b.Bytes(), &proto)
	require.NoError(t, err, "unmarshal of %q should succeed", b.String())
	require.Equal(t, msg, proto["message"])
	require.Equal(t, []interface{}{"test"}, proto["@tags"])
}

func TestLogLevel(t *testing.T) {
	b := &bytes.Buffer{}
	l := newLoggerForTest(b)
	l.SetLevel(FatalLevel)

	l.Info("A walrus disappears")
	require.Len(t, b.String(), 0, "nothing should be logged on level InfoLevel")
}

func TestLogLevelForNullLoggers(t *testing.T) {
	l, h := NewNullLoggerWithHook()
	l.SetLevel(DebugLevel)
	l.Debug("A walrus disappears")
	require.Len(t, h.Entries, 1, "there should be one debug entry")
}

func TestLogWithFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := newLoggerForTest(b)

	l.WithFields(Fields{"animal": "walrus"}).Info("A walrus appears")

	var proto map[string]interface{}
	err := json.Unmarshal(b.Bytes(), &proto)
	require.NoError(t, err, "unmarshal of %q should succeed", b.String())
	require.Equal(t, "walrus", proto["animal"])
}

func TestLogWithField(t *testing.T) {
	b := &bytes.Buffer{}
	l := newLoggerForTest(b)
	l.WithField("animal", "walrus").Info("A walrus appears")

	var proto map[string]interface{}
	err := json.Unmarshal(b.Bytes(), &proto)
	require.NoError(t, err, "unmarshal of %q should succeed", b.String())
	require.Equal(t, "walrus", proto["animal"])
}

func TestLogWithError(t *testing.T) {
	b := &bytes.Buffer{}
	l := newLoggerForTest(b)
	msg := "oh noos"
	err := fmt.Errorf(msg)
	l.WithError(err).Info("A walrus appears")

	var proto map[string]interface{}
	err = json.Unmarshal(b.Bytes(), &proto)
	require.NoError(t, err, "unmarshal of %q should succeed", b.String())
	require.Equal(t, msg, proto["error"])
}

func testListener(t *testing.T, listPortChan chan int, ret chan []byte, errChan chan error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		errChan <- errors.Wrap(err, "error listening")
		return
	}
	_, port, _ := net.SplitHostPort(l.Addr().String())
	p, _ := strconv.Atoi(port)
	listPortChan <- p

	// 200 is enough for the current test.
	b := make([]byte, 200)

	for {
		c, err := l.Accept()
		if err != nil {
			errChan <- errors.Wrap(err, "error accepting")
			return
		}
		n, err := c.Read(b)
		require.NoError(t, err)

		if n > 0 {
			ret <- b[:n]
		}
		ret <- b[:n]
		c.Close()
	}
}

func TestLoggingRemoteHostAndPort(t *testing.T) {
	retChan := make(chan []byte)
	lisPortChan := make(chan int)
	errChan := make(chan error)
	buf := &bytes.Buffer{}
	msg := "This is a warning you walrus"

	go testListener(t, lisPortChan, retChan, errChan)

	sendLogMsg := make(chan struct{})
	go func() {
		l := newLoggerForTest(buf)
		p := <-lisPortChan
		l.remoteHostAndPort("127.0.0.1", p)
		for range sendLogMsg {
			l.Warnf(msg)
		}
	}()

	timestampRE, err := regexp.Compile("[0-9]+-[0-9]+-[0-9]+[Tt]+[0-9]+:[0-9]+:[0-9]+.[0-9]+Z")
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		sendLogMsg <- struct{}{}
		var remoteBytes []byte
		select {
		case err := <-errChan:
			t.Fatalf("test listener failed: %v", err)
		case remoteBytes = <-retChan:
		}

		var proto map[string]interface{}
		if err := json.Unmarshal(remoteBytes, &proto); err != nil {
			t.Fatal(err)

		}

		if proto["message"] != msg {
			t.Fatalf("remote log message not equal to local one\nexpected\n%s, got\n%s", msg, proto["message"])
		}

		// If the event occured at @timestamp with 000 millisecond suffix (chance of this happening is 1e-3), the suffix will be stripped and regexp won't match.
		// This is why we are doing 5 iterations, making the chance 1e-15.
		matched := timestampRE.MatchString(proto["@timestamp"].(string))
		if matched {
			break
		}
		if !matched && i == 4 {
			t.Fatalf("remote log message timestamp is missing milliseconds. Timestamp given: %s", proto["@timestamp"])
		}
	}
}

func TestDataRaceOnFields(t *testing.T) {
	buf := &bytes.Buffer{}

	l := newLoggerForTest(buf)
	go func() {
		l.WithField("bla", "bliep").Info("am I a walrus")
	}()

	l.WithField("bliep", "bla").Info("am I a shark")
}

func TestGolangLoggerKeepsFields(t *testing.T) {
	var b bytes.Buffer
	l := newLoggerForTest(&b).WithField("animal", "walrus")
	gl := l.ToGolangLogger()

	gl.Print("This is a golang logger print")
	t.Log("logged:", b.String())

	var proto map[string]interface{}
	err := json.Unmarshal(b.Bytes(), &proto)
	require.NoError(t, err, "unmarshal of %q should succeed", b.String())
	require.Equal(t, "walrus", proto["animal"])
}
