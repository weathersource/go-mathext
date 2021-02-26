package mathext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestFloorp(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		result    float64
	}{
		{
			value:     9.999,
			precision: 2,
			result:    9.99,
		},
		{
			value:     1.111,
			precision: 1,
			result:    1.1,
		},
		{
			value:     5.555,
			precision: 0,
			result:    5,
		},
		{
			value:     55.555,
			precision: -1,
			result:    50,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Floorp(test.value, test.precision))
	}
}

func TestFloorp32(t *testing.T) {
	tests := []struct {
		value     float32
		precision int
		result    float32
	}{
		{
			value:     9.999,
			precision: 2,
			result:    9.99,
		},
		{
			value:     1.111,
			precision: 1,
			result:    1.1,
		},
		{
			value:     5.555,
			precision: 0,
			result:    5,
		},
		{
			value:     55.555,
			precision: -1,
			result:    50,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Floorp32(test.value, test.precision))
	}
}
