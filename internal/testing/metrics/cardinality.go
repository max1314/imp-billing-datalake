// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// VectorCardinality returns the number of metrics in a CounterVec.
// Note that calling With/WithLabelValues on a Vec creates a metric immediately, even if Inc/Dec/Set are not called on
// the metric.
func VectorCardinality(metric interface {
	Collect(chan<- prometheus.Metric)
}) int {
	metricsCh := make(chan prometheus.Metric)
	go func() {
		metric.Collect(metricsCh)
		close(metricsCh)
	}()
	var count int
	for range metricsCh {
		count++
	}

	return count
}
