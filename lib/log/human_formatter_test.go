// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"imp-billing-datalake/lib/errors"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tatsushid/termdeco"
)

func TestMessages(t *testing.T) {
	startTime = time.Now()
	now := startTime

	type testcase struct {
		name     string
		after    time.Duration
		message  string
		level    Level
		fields   Fields
		expected string
		log      func(*Logger)
	}

	cases := []testcase{
		{
			name:     "simple-debug-log",
			log:      func(l *Logger) { l.Debug("a debug message") },
			expected: fmt.Sprint(termdeco.White("+000"), " a debug message"),
		},
		{
			name:     "simple-info-log",
			log:      func(l *Logger) { l.Info("an info message") },
			expected: fmt.Sprint(termdeco.Blue("+000"), " an info message"),
		},
		{
			name:     "simple-warning-log",
			log:      func(l *Logger) { l.Warn("warny mcwarnface") },
			expected: fmt.Sprintf("%v warny mcwarnface\n   %v %v=%v", termdeco.Yellow("WARN"), termdeco.Yellow(">"), termdeco.Bold("time"), now.Format(time.RFC3339)),
		},
		{
			name:     "simple-error-log",
			log:      func(l *Logger) { l.Error("badness") },
			expected: fmt.Sprintf("%v badness\n   %v %v=%v", termdeco.Red("ERRO"), termdeco.Red(">"), termdeco.Bold("time"), now.Format(time.RFC3339)),
		},
		{
			name:     "simple-fatal-log",
			message:  "fatality",
			level:    FatalLevel,
			expected: fmt.Sprintf("%v fatality\n   %v %v=%v", termdeco.Red("FATA"), termdeco.Red(">"), termdeco.Bold("time"), now.Format(time.RFC3339)),
		},
		{
			name: "multiline-info-log",
			log:  func(l *Logger) { l.Info("an info message\nwith more stuff\non other lines") },
			expected: fmt.Sprint(
				termdeco.Blue("+000"), " an info message\n     with more stuff\n     on other lines",
			),
		},
		{
			name:     "fields-info-log",
			log:      func(l *Logger) { l.WithField("key", "val").Info("something") },
			expected: fmt.Sprintf("%v something\n   %v %v=val", termdeco.Blue("+000"), termdeco.Blue(">"), termdeco.Bold("key")),
		},
		{
			name:     "http-error-code-log",
			log:      func(l *Logger) { l.WithField(ErrorCodeKey, 503).Info("an http error") },
			expected: fmt.Sprintf("%v an http error\n   %v %v=503", termdeco.Blue("+000"), termdeco.Blue(">"), termdeco.Bold("code")),
		},
		{
			name: "info-log-with-error",
			log: func(l *Logger) {
				l.WithError(errors.New(nil, errors.InvalidArgument, "an error")).Info("something went wrong")
			},
			expected: fmt.Sprintf("%v something went wrong\n   %v error (InvalidArgument): an error", termdeco.Blue("+000"), termdeco.Blue("!")),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			l := NewWithWriter(b)
			l.SetFormatter(humanFormatter{})
			l.SetLevel(DebugLevel)
			l.AddHook(&timeFixerHook{Now: &now})
			now = startTime.Add(c.after)

			if c.log != nil {
				c.log(l)
				assert.Equal(t, c.expected+"\n", b.String())
			}

			if c.message != "" {
				l = l.WithFields(c.fields)
				l.entry.Level = logrus.Level(c.level)
				l.entry.Message = c.message
				l.entry.Time = now
				b, err := l.entry.String()
				require.NoError(t, err)
				assert.Equal(t, c.expected+"\n", b)
			}
		})
	}
}

// timeFixerHook is used to set testable times on log entries
type timeFixerHook struct {
	Now *time.Time
}

func (h *timeFixerHook) Levels() []Level {
	return AllLevels
}

func (h *timeFixerHook) Fire(entry *Entry) error {
	entry.Time = *h.Now
	return nil
}
