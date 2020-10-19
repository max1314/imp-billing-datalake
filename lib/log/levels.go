// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"imp-billing-datalake/lib/sharedflags"

	"github.com/mwitkow/go-flagz"
	"github.com/sirupsen/logrus"
)

// Level wraps a logrus.Level.
type Level logrus.Level

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with log.New().
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.logstash
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

var (
	fLogLevel = flagz.DynString(sharedflags.Set, "log_level", "", "default log level for all loggers. (default: info)")
	AllLevels = []Level{
		PanicLevel,
		FatalLevel,
		ErrorLevel,
		WarnLevel,
		InfoLevel,
		DebugLevel,
	}
)

func init() {
	fLogLevel.WithValidator(func(lvl string) error {
		_, err := ParseLevel(lvl)
		return err
	})
	// When the log level flag updates, we go through the list of root loggers to find all
	// underlying logrus loggers and update their log level.
	fLogLevel.WithNotifier(func(_ string, val string) {
		lvl, _ := ParseLevel(val)
		rootLoggersLock.Lock()
		defer rootLoggersLock.Unlock()
		for _, l := range rootLoggers {
			if !l.writer.levelOverride {
				l.writer.mu.Lock()
				l.writer.Level = logrus.Level(lvl)
				l.writer.mu.Unlock()
			}
		}
	})
}

func (l Level) String() string {
	return logrus.Level(l).String()
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
	logrusLevel, err := logrus.ParseLevel(lvl)
	return Level(logrusLevel), err
}
