// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package errors

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func TestEtcd3ClientErrors(t *testing.T) {
	err := rpctypes.ErrDuplicateKey
	assert.Equal(t, InvalidArgument, Code(err))
}

// TestReadIOTimeoutCode tests that read IO timeouts are coded as DeadlineExceeded.
func TestReadIOTimeoutCode(t *testing.T) {
	// Listen on a random port.
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port

	// Connect to this port.
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	require.NoError(t, err)
	require.NoError(t, conn.SetDeadline(time.Now().Add(time.Second)))

	// Read data from conn and expect timeout.
	b := make([]byte, 1)
	i, err := conn.Read(b)
	require.Equal(t, 0, i)
	require.Error(t, err)

	assert.Equal(t, DeadlineExceeded, Code(err))
}

// timeoutError is copy/pasted from internal/poll/fd.go
type timeoutError struct{}

func (e *timeoutError) Error() string   { return "i/o timeout" }
func (e *timeoutError) Timeout() bool   { return true }
func (e *timeoutError) Temporary() bool { return true }

// TestDialIOTimeoutCode tests that dial i/o timeout is coded correctly.
func TestDialIOTimeoutCode(t *testing.T) {
	// Simulating dial I/O timeout properly is too complicated, see net/timeout_test.go.
	err := &net.OpError{Err: &timeoutError{}}
	assert.Equal(t, DeadlineExceeded, Code(err))
}

// TestDialConnResetByPeerCode tests that connection reset by peer is coded correctly.
func TestDialConnResetByPeerCode(t *testing.T) {
	listener, err := net.Listen("tcp", "")
	require.NoError(t, err)
	go func() {
		conn, connErr := listener.Accept()
		require.NoError(t, connErr)
		tcpConn := conn.(*net.TCPConn)
		connErr = tcpConn.SetLinger(0)
		require.NoError(t, connErr)
		// We need to wait until conn.Read call is made, so that it can be cancelled by calling Close().
		time.Sleep(100 * time.Millisecond)
		tcpConn.Close()
	}()
	conn, err := net.Dial(listener.Addr().Network(), listener.Addr().String())
	require.NoError(t, err)
	readBuf := make([]byte, 1)
	_, err = conn.Read(readBuf)
	assert.Equal(t, Unavailable.String(), Code(err).String(), Details(err))
}

// TestDNSNotFoundCode tests that DNS NotFound error is coded correctly.
func TestDNSNotFoundCode(t *testing.T) {
	// Start a DNS server locally.
	started := make(chan struct{})
	server := &dns.Server{
		Addr: "0.0.0.0:0",
		Net:  "udp",
		NotifyStartedFunc: func() {
			close(started)
		},
		Handler: dns.HandlerFunc(func(w dns.ResponseWriter, r *dns.Msg) {
			m := new(dns.Msg)
			m.SetRcode(r, dns.RcodeNameError)
			w.WriteMsg(m)
		}),
	}
	go func() {
		require.NoError(t, server.ListenAndServe())
	}()
	<-started
	defer func() {
		require.NoError(t, server.Shutdown())
	}()

	// Find out which port it is running on.
	hp := server.PacketConn.LocalAddr().String()
	_, port, err := net.SplitHostPort(hp)
	require.NoError(t, err, "invalid hostport %s", hp)

	// Create a dialer that resolves all hosts via our custom DNS server.
	dialer := net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("udp", fmt.Sprintf("127.0.0.1:%s", port))
			},
		},
	}

	_, err = dialer.Dial("tcp", "whatever.com:1234")
	assert.Equal(t, NotFound.String(), Code(err).String())
}

// TestDNSIOTimeoutCode tests that DNS i/o timeout error is coded correctly.
func TestDNSIOTimeoutCode(t *testing.T) {
	dialer := net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return nil, &net.OpError{Err: &timeoutError{}}
			},
		},
	}

	// TimeoutError should get wrapped into DNS error and be coded as 404.
	_, err := dialer.Dial("tcp", "whatever.com:1234")
	assert.Equal(t, DeadlineExceeded.String(), Code(err).String())
}
