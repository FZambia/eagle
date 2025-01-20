package eagle

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestFlattenLabels(t *testing.T) {
	labels := []metricLabel{
		{"test_name", "test_value"},
	}
	flatLabels := flattenLabels(labels)
	require.Equal(t, 2, len(flatLabels))
}

func TestMetricPrefixes(t *testing.T) {
	testCounter1 := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "white",
		Subsystem: "test",
		Name:      "test_counter",
		Help:      "test counter",
	})
	testCounter2 := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "black",
		Subsystem: "test",
		Name:      "test_counter",
		Help:      "test counter",
	})
	registry := prometheus.NewRegistry()
	registry.MustRegister(testCounter1, testCounter2)
	e := New(Config{
		Gatherer: registry,
		PrefixWhitelist: []string{
			"white",
		},
	})
	metrics, err := e.Export()
	require.NoError(t, err)
	flattened := metrics.Flatten("_")
	require.Len(t, flattened, 1)
	require.Contains(t, flattened, "white_test_test_counter")
}
