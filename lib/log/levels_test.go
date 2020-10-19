// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package log

import (
	"bytes"
	"testing"
	"time"

	"imp-billing-datalake/lib/sharedflags"

	"github.com/jpillora/backoff"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type plainFormatter struct{}

func (f plainFormatter) Format(e *logrus.Entry) ([]byte, error) {
	return append([]byte(e.Message), '\n'), nil
}

func updateLevelFlag(t *testing.T, logger *Logger, target Level) {
	flag := sharedflags.Set.Lookup("log_level")
	err := flag.Value.Set(target.String())
	assert.NoError(t, err)

	// notifications from dynflags are async, need to wait to get the update
	boff := backoff.Backoff{}
	for i := 0; i < 20; i++ {
		logger.writer.mu.RLock()
		level := logger.writer.Level
		logger.writer.mu.RUnlock()
		if level == logrus.Level(target) {
			break
		}
		time.Sleep(boff.Duration())
	}

	logger.writer.mu.RLock()
	defer logger.writer.mu.RUnlock()
	assert.Equal(t, logrus.Level(target), logger.writer.Level)
}

func TestLevelFlagChangeUpdatesLogger(t *testing.T) {
	logger := New()
	var b bytes.Buffer
	logger.writer.Out = &b
	logger.SetFormatter(Formatter(plainFormatter{}))

	logger.Info("this should be printed")
	// ensure the level will be reverted so other tests pass
	defer updateLevelFlag(t, logger, InfoLevel)
	// change flag for this test
	updateLevelFlag(t, logger, WarnLevel)

	logger.Info("this should not be printed")
	assert.Equal(t, "this should be printed\n", b.String())
}
