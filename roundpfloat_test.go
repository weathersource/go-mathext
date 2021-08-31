package mathext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestRoundpfloat(t *testing.T) {
	tests := []struct {
		value     float64
		precision float64
		result    float64
	}{
		{
			value:     9.999,
			precision: 0.5,
			result:    10,
		},
		{
			value:     1.311,
			precision: 0.5,
			result:    1.5,
		},
		{
			value:     5.725,
			precision: 0.25,
			result:    5.75,
		},
		{
			value:     55.955,
			precision: 1,
			result:    56,
		},
    {
			value:     55.501,
			precision: 0.5,
			result:    55.5,
		},
    {
			value:     55.497,
			precision: 0.5,
			result:    55.5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Roundpfloat(test.value, test.precision))
	}
}
