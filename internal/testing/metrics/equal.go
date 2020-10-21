// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package metrics

import (
	"fmt"
	"reflect"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

// CounterAssertEqual asserts that prometheus Counter metric value is equal to the expected float64.
//
//    CounterAssertEqual(t, float64(123), mySuperMetric.WithLabelValues("A"), "My super metric should equal 123")
//
// Returns whether the assertion was successful (true) or not (false).
func CounterAssertEqual(t assert.TestingT, expected float64, metric prometheus.Counter, msgAndArgs ...interface{}) bool {
	value := &dto.Metric{}
	_ = metric.Write(value)

	metricFullName := reflect.ValueOf(*metric.Desc()).FieldByName("fqName").String()

	if value.GetCounter() != nil {
		return isMetricEqual(t, expected, value.Counter.GetValue(), metricFullName, msgAndArgs...)
	}

	return assert.Fail(t, fmt.Sprintf("No value for counter metric %s", metricFullName), msgAndArgs...)
}

// GaugeAssertEqual asserts that prometheus Gauge metric value is equal to the expected float64.
//
//    GaugeAssertEqual(t, float64(123), mySuperMetric.WithLabelValues("A"), "My super metric should equal 123")
//
// Returns whether the assertion was successful (true) or not (false).
func GaugeAssertEqual(t assert.TestingT, expected float64, metric prometheus.Gauge, msgAndArgs ...interface{}) bool {
	value := &dto.Metric{}
	_ = metric.Write(value)

	metricFullName := reflect.ValueOf(*metric.Desc()).FieldByName("fqName").String()

	if value.GetGauge() != nil {
		return isMetricEqual(t, expected, value.Gauge.GetValue(), metricFullName, msgAndArgs...)
	}

	return assert.Fail(t, fmt.Sprintf("No value for gauge metric %s", metricFullName), msgAndArgs...)
}

func isMetricEqual(t assert.TestingT, expected interface{}, actual interface{}, metricFullName string, msgAndArgs ...interface{}) bool {
	if !assert.ObjectsAreEqual(expected, actual) {
		return assert.Fail(t, fmt.Sprintf("Metric %s value not equal: %#v (expected)\n"+
			"        != %#v (actual)", metricFullName, expected, actual), msgAndArgs...)
	}

	return true
}

// HistogramSamplesAssertEqual asserts that prometheus Histogram metric value is equal to the expected float64.
//
//    HistogramSamplesAssertEqual(t, float64(5), uint64(1), mySuperMetric.WithLabelValues("A"), "My super metric sample suml 123")
//
// Returns whether the assertion was successful (true) or not (false).
func HistogramSamplesAssertEqual(
	t assert.TestingT,
	expectedSum float64,
	expectedCount uint64,
	metricObserver prometheus.Observer,
	msgAndArgs ...interface{},
) bool {
	metric, ok := metricObserver.(prometheus.Metric)
	assert.True(t, ok, "metricObserver is not a prometheus.Metric")

	value := &dto.Metric{}
	_ = metric.Write(value)

	metricFullName := reflect.ValueOf(*metric.Desc()).FieldByName("fqName").String()

	if value.GetHistogram() != nil {
		if !assert.ObjectsAreEqual(expectedCount, value.Histogram.GetSampleCount()) {
			return assert.Fail(t, fmt.Sprintf("Metric %s Sample Count is not equal: %v (expected)\n"+
				"        != %v (actual)", metricFullName, expectedCount, value.Histogram.GetSampleCount()), msgAndArgs...)
		}

		if !assert.ObjectsAreEqual(expectedSum, value.Histogram.GetSampleSum()) {
			return assert.Fail(t, fmt.Sprintf("Metric %s Sample Sum is not equal: %#v (expected)\n"+
				"        != %#v (actual)", metricFullName, expectedSum, value.Histogram.GetSampleSum()), msgAndArgs...)
		}
		return true
	}

	return assert.Fail(t, fmt.Sprintf("No data for Histogram metric %s", metricFullName), msgAndArgs...)
}

// HistogramCountAssertEqual asserts that prometheus Histogram metric value is equal to the expected count.
//
//    HistogramCountAssertEqual(t, uint64(1), mySuperMetric.WithLabelValues("A"), "My super metric sample suml 123")
//
// Returns whether the assertion was successful (true) or not (false).
func HistogramCountAssertEqual(
	t assert.TestingT,
	expectedCount uint64,
	metricObserver prometheus.Observer,
	msgAndArgs ...interface{},
) bool {
	metric, ok := metricObserver.(prometheus.Metric)
	assert.True(t, ok, "metricObserver is not a prometheus.Metric")

	value := &dto.Metric{}
	_ = metric.Write(value)

	metricFullName := reflect.ValueOf(*metric.Desc()).FieldByName("fqName").String()

	if value.GetHistogram() != nil {
		if !assert.ObjectsAreEqual(expectedCount, value.Histogram.GetSampleCount()) {
			return assert.Fail(t, fmt.Sprintf("Metric %s Sample Count is not equal: %#v (expected)\n"+
				"        != %#v (actual)", metricFullName, expectedCount, value.Histogram.GetSampleCount()), msgAndArgs...)
		}
		return true
	}

	return assert.Fail(t, fmt.Sprintf("No data for Histogram metric %s", metricFullName), msgAndArgs...)
}
