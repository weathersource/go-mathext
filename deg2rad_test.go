package mathext

import (
	"math"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestDeg2Rad(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		result    float64
	}{
		{
			value:  -180,
			result: -1 * math.Pi,
		},
		{
			value:  0,
			result: 0,
		},
		{
			value:  90,
			result: math.Pi / 2,
		},
		{
			value:  180,
			result: math.Pi,
		},
		{
			value:  270,
			result: (3 * math.Pi) / 2,
		},
		{
			value:  360,
			result: 2 * math.Pi,
		},
		{
			value:  540,
			result: 3 * math.Pi,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Deg2Rad(test.value))
	}
}
