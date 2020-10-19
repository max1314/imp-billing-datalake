// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package metrics

import(
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

var (
	testCounterVecMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "SomeNamespace",
			Subsystem: "SomeSubsys",
			Name:      "test_counter_vec",
			Help:      "Test help.",
			ConstLabels: prometheus.Labels{
				"some_const_label": "some_value",
			},
		}, []string{"label1", "label2"})
	testGaugeVecMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "SomeNamespace",
			Subsystem: "SomeSubsys",
			Name:      "test_gauge_vec",
			Help:      "Test help.",
			ConstLabels: prometheus.Labels{
				"some_const_label": "some_value",
			},
		}, []string{"label1", "label2"})
	testHistogramVecMetric = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "SomeNamespace",
			Subsystem: "SomeSubsys",
			Name:      "test_histogram_vec",
			Help:      "Test help.",
			ConstLabels: prometheus.Labels{
				"some_const_label": "some_value",
			},
		}, []string{"label1", "label2"})
)

func TestCounterAssertEqual_OK(t *testing.T) {
	testCounterVecMetric.Reset()

	testCounterVecMetric.WithLabelValues("A", "B").Inc()
	testCounterVecMetric.WithLabelValues("A", "B").Inc()

	testCounterVecMetric.WithLabelValues("A", "B").Inc()

	// Unrelated values for testing A, B.
	testCounterVecMetric.WithLabelValues("A", "X").Add(float64(100))
	testCounterVecMetric.WithLabelValues("X", "B").Add(float64(101))
	testCounterVecMetric.WithLabelValues("X", "Y").Add(float64(102))

	assert.True(t,
		CounterAssertEqual(t, float64(3), testCounterVecMetric.WithLabelValues("A", "B"), "metric for labels A, B is equal to 3"))
}

func TestGaugeAssertEqual_OK(t *testing.T) {
	testGaugeVecMetric.Reset()

	testGaugeVecMetric.WithLabelValues("A", "B").Inc()
	testGaugeVecMetric.WithLabelValues("A", "B").Inc()

	// Set should reset the value.
	testGaugeVecMetric.WithLabelValues("A", "B").Set(float64(4))
	testGaugeVecMetric.WithLabelValues("A", "B").Inc()

	// Unrelated values for testing A, B.
	testGaugeVecMetric.WithLabelValues("A", "X").Set(float64(100))
	testGaugeVecMetric.WithLabelValues("X", "B").Set(float64(101))
	testGaugeVecMetric.WithLabelValues("X", "Y").Set(float64(102))

	assert.True(t,
		GaugeAssertEqual(t, float64(5), testGaugeVecMetric.WithLabelValues("A", "B"), "metric for labels A, B is equal to 5"))
}

func TestHistogramSamplesAssertEqual_OK(t *testing.T) {
	testHistogramVecMetric.Reset()

	testHistogramVecMetric.WithLabelValues("A", "B").Observe(float64(2))
	testHistogramVecMetric.WithLabelValues("A", "B").Observe(float64(6))

	// Unrelated values for testing A, B.
	testHistogramVecMetric.WithLabelValues("A", "X").Observe(float64(100))
	testHistogramVecMetric.WithLabelValues("X", "B").Observe(float64(101))
	testHistogramVecMetric.WithLabelValues("X", "Y").Observe(float64(102))

	assert.True(t,
		HistogramSamplesAssertEqual(
			t, float64(8), uint64(2),
			testHistogramVecMetric.WithLabelValues("A", "B"),
			"histogram samples sum for labels A, B is equal to 8, count is equal 2"))
}

type fakeT struct {
	errors int
}

func (f *fakeT) Errorf(format string, args ...interface{}) {
	f.errors++
}

func TestCounterAssertEqual_Fail(t *testing.T) {
	testCounterVecMetric.Reset()

	testCounterVecMetric.WithLabelValues("A", "B").Inc()
	testCounterVecMetric.WithLabelValues("A", "B").Inc()

	// Set should reset the value.
	testCounterVecMetric.WithLabelValues("A", "B").Add(float64(104))
	testCounterVecMetric.WithLabelValues("A", "B").Inc()

	// Unrelated values for testing A, B.
	testCounterVecMetric.WithLabelValues("A", "X").Add(float64(100))
	testCounterVecMetric.WithLabelValues("X", "B").Add(float64(101))
	testCounterVecMetric.WithLabelValues("X", "Y").Add(float64(102))

	fake := &fakeT{}

	assert.False(t,
		CounterAssertEqual(fake, float64(5), testCounterVecMetric.WithLabelValues("A", "B"), "metric for labels A, B is equal to 5"))
	assert.Equal(t, 1, fake.errors, "CounterAssertEqual should report one error.")
}

func TestGaugeAssertEqual_Fail(t *testing.T) {
	testGaugeVecMetric.Reset()

	testGaugeVecMetric.WithLabelValues("A", "B").Inc()
	testGaugeVecMetric.WithLabelValues("A", "B").Inc()

	// Set should reset the value.
	testGaugeVecMetric.WithLabelValues("A", "B").Set(float64(104))
	testGaugeVecMetric.WithLabelValues("A", "B").Inc()

	// Unrelated values for testing A, B.
	testGaugeVecMetric.WithLabelValues("A", "X").Set(float64(100))
	testGaugeVecMetric.WithLabelValues("X", "B").Set(float64(101))
	testGaugeVecMetric.WithLabelValues("X", "Y").Set(float64(102))

	fake := &fakeT{}

	assert.False(t,
		GaugeAssertEqual(fake, float64(5), testGaugeVecMetric.WithLabelValues("A", "B"), "metric for labels A, B is equal to 5"))
	assert.Equal(t, 1, fake.errors, "GaugeAssertEqual should report one error.")
}

func TestHistogramSamplesAssertEqual_CountMismatch(t *testing.T) {
	testHistogramVecMetric.Reset()

	testHistogramVecMetric.WithLabelValues("A", "B").Observe(float64(2))

	// Unrelated values for testing A, B.
	testHistogramVecMetric.WithLabelValues("A", "X").Observe(float64(100))
	testHistogramVecMetric.WithLabelValues("X", "B").Observe(float64(101))
	testHistogramVecMetric.WithLabelValues("X", "Y").Observe(float64(102))

	fake := &fakeT{}

	assert.False(t,
		HistogramSamplesAssertEqual(
			fake, float64(8), uint64(1),
			testHistogramVecMetric.WithLabelValues("A", "B"),
			"histogram samples sum for labels A, B is equal to 8, count is equal 2"))
	assert.Equal(t, 1, fake.errors, "HistogramSamplesAssertEqual should report one error.")
}

func TestHistogramSamplesAssertEqual_SumMismatch(t *testing.T) {
	testHistogramVecMetric.Reset()

	testHistogramVecMetric.WithLabelValues("A", "B").Observe(float64(2))
	testHistogramVecMetric.WithLabelValues("A", "B").Observe(float64(2))

	// Unrelated values for testing A, B.
	testHistogramVecMetric.WithLabelValues("A", "X").Observe(float64(100))
	testHistogramVecMetric.WithLabelValues("X", "B").Observe(float64(101))
	testHistogramVecMetric.WithLabelValues("X", "Y").Observe(float64(102))

	fake := &fakeT{}

	assert.False(t,
		HistogramSamplesAssertEqual(
			fake, float64(8), uint64(2),
			testHistogramVecMetric.WithLabelValues("A", "B"),
			"histogram samples sum for labels A, B is equal to 8, count is equal 2"))
	assert.Equal(t, 1, fake.errors, "HistogramSamplesAssertEqual should report one error.")
}
