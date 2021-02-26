package mathext

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestCeilp(t *testing.T) {
	tests := []struct {
		value     float64
		precision int
		result    float64
	}{
		{
			value:     9.999,
			precision: 2,
			result:    10,
		},
		{
			value:     1.111,
			precision: 1,
			result:    1.2,
		},
		{
			value:     5.555,
			precision: 0,
			result:    6,
		},
		{
			value:     55.555,
			precision: -1,
			result:    60,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Ceilp(test.value, test.precision))
	}
}

func TestCeilp32(t *testing.T) {
	tests := []struct {
		value     float32
		precision int
		result    float32
	}{
		{
			value:     9.999,
			precision: 2,
			result:    10,
		},
		{
			value:     1.111,
			precision: 1,
			result:    1.2,
		},
		{
			value:     5.555,
			precision: 0,
			result:    6,
		},
		{
			value:     55.555,
			precision: -1,
			result:    60,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, Ceilp32(test.value, test.precision))
	}
}
