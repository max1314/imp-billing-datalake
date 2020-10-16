// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package reporting

import (
	"testing"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/errors"
	"improbable.io/lib/testing/metrics"
)

func TestTrack(t *testing.T) {
	operationCounter.Reset()
	operationDurationSeconds.Reset()

	_ = func() (err error) {
		defer Track(UpdateACL, blob.GCS, &err)()

		err = errors.New(nil, errors.InvalidArgument, "test")
		return err
	}()

	metrics.CounterAssertEqual(t, 1, operationCounter.WithLabelValues(string(UpdateACL), string(blob.GCS)))
	metrics.HistogramCountAssertEqual(t, 1, operationDurationSeconds.WithLabelValues(string(UpdateACL), string(blob.GCS), errors.InvalidArgument.String()))
}
