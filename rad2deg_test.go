package mathext

import (
	"math"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestRad2Deg(t *testing.T) {
	tests := []struct {
		value  float64
		result float64
	}{
		{
			value:  -1 * math.Pi,
			result: -180,
		},
		{
			value:  0,
			result: 0,
		},
		{
			value:  math.Pi / 2,
			result: 90,
		},
		{
			value:  math.Pi,
			result: 180,
		},
		{
			value:  (3 * math.Pi) / 2,
			result: 270,
		},
		{
			value:  2 * math.Pi,
			result: 360,
		},
		{
			value:  3 * math.Pi,
			result: 540,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Rad2Deg(test.value))
	}
}
