// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"imp-billing-datalake/lib/sharedflags"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	FLogstashHost         = sharedflags.Set.String("log_logstash_host", "logstash.infra.svc", "host to be used to reach logshipper")
	fLogstashPort         = sharedflags.Set.Int("log_logstash_port", 10105, "port to reach logshipper")
	fRemoteLoggingEnabled = sharedflags.Set.Bool("log_remote_logging", false, "log to logstash if set to true")
)

func init() {
	cobra.OnInitialize(func() {
		// this is called after cobra has parsed args and before it is about to start running things
		if *fRemoteLoggingEnabled {
			// start remote-logging if it is configured.
			rootLoggersLock.Lock()
			defer rootLoggersLock.Unlock()
			for _, l := range rootLoggers {
				l.enableRemote()
			}
		}
	})
}

// enableRemote makes l a remote logger according to the flags. See remoteHostAndPort for more information.
func (l *Logger) enableRemote() {
	if !l.writer.remote {
		l.remoteHostAndPort(*FLogstashHost, *fLogstashPort)
	}
}

// remoteHostAndPort makes l a remote logger. It will log in JSON (via TCP) and send to LogStash. If port is 0 we use
// the default port for LogStash (=10104). If the LogStash host cannot be found (NXDOMAIN) an error is returned.
// If there are connection problems with Logstash, it will try to reconnect. If that still fails we start an exponential
// backoff. During the backoff, all logs are dropped. This solves two major problems: hammering logstash with reconnects
// and blocking the logging of the application.
func (l *Logger) remoteHostAndPort(host string, port int) {
	if port == 0 {
		port = *fLogstashPort
	}
	hook, err := NewLogStashHook(host, port, logrus.New())
	if err != nil {
		l.Warnf("failed to set up remote logging. will retry later. err: %v", err)
	}
	l.AddHook(hook)
	l.writer.remote = true
}
