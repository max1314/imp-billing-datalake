// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

// Monitoring around logged messages and dropped stuff.

package log

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	loggedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "improbable",
			Subsystem: "logging",
			Name:      "entries_total",
			Help:      "Number of all entries incoming to the logger by log level.",
		}, []string{"level"})

	remoteCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "improbable",
			Subsystem: "logging",
			Name:      "remote_entries_total",
			Help:      "Number of successfully logged remote entries by log level.",
		}, []string{"level"})

	droppedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "improbable",
			Subsystem: "logging",
			Name:      "dropped_entries_total",
			Help:      "Number of dropped entries by log level.",
		}, []string{"level", "reason"})

	logstashConnFailuresCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "improbable",
			Subsystem: "logging",
			Name:      "logstash_connection_failures_total",
			Help:      "Number of failed attempts to connect to logstash.",
		})

	logstashConnCreationsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "improbable",
			Subsystem: "logging",
			Name:      "logstash_connection_creations_total",
			Help:      "Number of time we have successfully created a connection to logstash.",
		})
	logstashWriteFailuresCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "improbable",
			Subsystem: "logging",
			Name:      "logstash_connection_write_failures_total",
			Help:      "Number of time we failed to write to a connection to logstash.",
		})
)

// dropReason for logs being dropped.
type dropReason string

const (
	BufferFull    dropReason = "buffer_full"
	FailedToWrite dropReason = "failed_to_write"
	BadFormat     dropReason = "bad_format"
)

func init() {
	prometheus.MustRegister(loggedCounter)
	prometheus.MustRegister(remoteCounter)
	prometheus.MustRegister(droppedCounter)
	prometheus.MustRegister(logstashConnFailuresCounter)
	prometheus.MustRegister(logstashConnCreationsCounter)
	prometheus.MustRegister(logstashWriteFailuresCounter)
	registerAllLabels()
}

func registerAllLabels() {
	for _, level := range logrus.AllLevels {
		loggedCounter.WithLabelValues(level.String())
		remoteCounter.WithLabelValues(level.String())
		droppedCounter.WithLabelValues(level.String(), string(BufferFull))
		droppedCounter.WithLabelValues(level.String(), string(FailedToWrite))
		droppedCounter.WithLabelValues(level.String(), string(BadFormat))
	}
}

func reportEntry(level logrus.Level) {
	loggedCounter.WithLabelValues(level.String()).Inc()
}

func reportRemoteSuccess(level logrus.Level) {
	remoteCounter.WithLabelValues(level.String()).Inc()
}

func reportDropped(level logrus.Level, r dropReason) {
	droppedCounter.WithLabelValues(level.String(), string(r)).Inc()
}
