// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package reporting

import (
	"time"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/errors"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Operation string

const (
	namespace = "improbable"
	subsystem = "blob_client"

	Copy                                Operation = "copy"
	Delete                              Operation = "delete"
	Exists                              Operation = "exists"
	GetDirContent                       Operation = "getDirContent"
	GetContentWithCommonPrefix          Operation = "getContentWithCommonPrefix"
	GetSignedDownloadURL                Operation = "getSignedDownloadURL"
	GetSignedUploadURL                  Operation = "getSignedUploadURL"
	GetSignedResumableUploadURL         Operation = "getSignedResumableUploadURL"
	GetSignedCreateMultipartUploadURL   Operation = "getSignedCreateMultipartUploadURL"
	GetSignedUploadpartURL              Operation = "getSignedUploadPartURL"
	GetSignedCompleteMultipartUploadURL Operation = "getSignedCompleteMultipartUploadURL"
	NewReader                           Operation = "newReader"
	NewWriter                           Operation = "newWriter"
	Stat                                Operation = "stat"
	UpdateMeta                          Operation = "updateMeta"
	UpdateACL                           Operation = "updateACL"
)

var (
	operationCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "operations_started_total",
			Help:      "Total number of blob operations started.",
		}, []string{"operation", "backing_storage"})
	operationDurationSeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "operations_duration_seconds",
			Help:      "Bucketed histogram of handling times of successful backend operations.",
			Buckets:   prometheus.DefBuckets,
		}, []string{"operation", "backing_storage", "code"})
)

// Track reports invocations and latency of blob storage operations
// Usage
// func foo() (err error) {
// 	 defer Track(op, storage, &err)()
//   ...
// }
func Track(op Operation, storage blob.BackingStorage, err *error) func() {
	start := time.Now()
	operationCounter.WithLabelValues(string(op), string(storage)).Inc()

	return func() {
		operationDurationSeconds.WithLabelValues(
			string(op),
			string(storage),
			errors.Code(*err).String(),
		).Observe(time.Since(start).Seconds())
	}
}
