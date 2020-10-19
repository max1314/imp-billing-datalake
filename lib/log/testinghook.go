// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// TestingHook is a Logrus hook that logs using `*testing.T`

package log

import (
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
)

var defaultLevels = []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel}

type TestingHook struct {
	t             *testing.T
	l             *Logger
	msgs          []*Entry
	messagesMutex sync.Mutex
	level         Level
}

func NewTestingHook(t *testing.T) *TestingHook {
	return &TestingHook{t: t, level: DebugLevel}
}

func NewTestingHookWithLevel(t *testing.T, level Level) *TestingHook {
	return &TestingHook{t: t, level: level}
}

func (h *TestingHook) Levels() []Level {
	var lvls []Level
	for i := int(h.level); i >= 0; i-- {
		lvls = append(lvls, Level(i))
	}

	return lvls
}

func (h *TestingHook) Fire(e *Entry) error {
	if h.t == nil {
		return fmt.Errorf("testing hook not registered")
	}
	h.messagesMutex.Lock()
	defer h.messagesMutex.Unlock()
	h.msgs = append(h.msgs, e)
	h.t.Logf("lg: %7s | %s | %v", e.Level, e.Message, e.Data)
	return nil
}

func (h *TestingHook) AllEntries() []*Entry {
	return h.msgs
}

func (h *TestingHook) Reset() {
	h.msgs = make([]*Entry, 0)
}

// SetT should be called on each test, to make sure that the log entries are tied to the appropriate test.
func (h *TestingHook) SetT(t *testing.T) {
	h.t = t
}

// Logger returns an instance of Logger that is tied to this TestingHook.
func (h *TestingHook) Logger() *Logger {
	return h.LoggerWithLevel(defaultLevels[0])
}

func (h *TestingHook) LoggerWithLevel(logLevel Level) *Logger {
	if h.l != nil {
		return h.l
	}
	h.l = NewWithWriter(ioutil.Discard)
	h.Reset()
	h.l.SetLevel(logLevel)
	h.l.AddHook(h)
	return h.l
}
