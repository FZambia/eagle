package eagle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlattenLabels(t *testing.T) {
	labels := []metricLabel{metricLabel{"test_name", "test_value"}}
	flatLabels := flattenLabels(labels)
	require.Equal(t, 2, len(flatLabels))
}
