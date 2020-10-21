// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package reporting

import (
	"testing"

	blob "github.com/max1314/imp-billing-datalake"
	"github.com/max1314/imp-billing-datalake/internal/errors"
	"github.com/max1314/imp-billing-datalake/internal/testing/metrics"
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
