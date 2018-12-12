package eagle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenLabels(t *testing.T) {
	labels := []Label{Label{"test_name", "test_value"}}
	flatLabels := flattenLabels(labels)
	assert.Equal(t, 2, len(flatLabels))
}
