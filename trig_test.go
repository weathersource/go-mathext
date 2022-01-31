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

func TestSind(t *testing.T) {
    tests := []struct {
        value  float64
        result float64
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  90,
            result: 1,
        },
        {
            value:  180,
            result: 0,
        },
        {
            value:  270,
            result: -1,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Sind(test.value), 10))
    }
}

func TestCosd(t *testing.T) {
    tests := []struct {
        value  float64
        result float64
    }{
        {
            value:  0,
            result: 1,
        },
        {
            value:  90,
            result: 0,
        },
        {
            value:  180,
            result: -1,
        },
        {
            value:  270,
            result: 0,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Cosd(test.value), 10))
    }
}

func TestTand(t *testing.T) {
    tests := []struct {
        value  float64
        result float64
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  45,
            result: 1,
        },
        {
            value:  135,
            result: -1,
        },
        {
            value:  180,
            result: 0,
        },
        {
            value:  225,
            result: 1,
        },
        {
            value:  315,
            result: -1,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Tand(test.value), 10))
    }
}

func TestAsind(t *testing.T) {
    tests := []struct {
        value  float64
        result float64
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  1,
            result: 90,
        },
        {
            value:  -1,
            result: -90,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Asind(test.value), 10))
    }
}

func TestAcosd(t *testing.T) {
    tests := []struct {
        value  float64
        result float64
    }{
        {
            value:  0,
            result: 90,
        },
        {
            value:  1,
            result: 0,
        },
        {
            value:  -1,
            result: 180,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Acosd(test.value), 10))
    }
}

func TestAtand(t *testing.T) {
    tests := []struct {
        value  float64
        result float64
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  1,
            result: 45,
        },
        {
            value:  -1,
            result: -45,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Atand(test.value), 10))
    }
}

func TestAtan2d(t *testing.T) {
    tests := []struct {
        y      float64
        x      float64
        result float64
    }{
        {
            y:      0,
            x:      0,
            result: 0,
        },
        {
            y:      1,
            x:      0,
            result: 90,
        },
        {
            y:      -1,
            x:      0,
            result: -90,
        },
        {
            y:      0,
            x:      1,
            result: 0,
        },
        {
            y:      1,
            x:      1,
            result: 45,
        },
        {
            y:      1,
            x:      -1,
            result: 135,
        },
        {
            y:      -1,
            x:      1,
            result: -45,
        },
        {
            y:      -1,
            x:      -1,
            result: -135,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp(Atan2d(test.y, test.x), 10))
    }
}

func TestSind32(t *testing.T) {
    tests := []struct {
        value  float32
        result float32
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  90,
            result: 1,
        },
        {
            value:  180,
            result: 0,
        },
        {
            value:  270,
            result: -1,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Sind32(test.value), 10))
    }
}

func TestCosd32(t *testing.T) {
    tests := []struct {
        value  float32
        result float32
    }{
        {
            value:  0,
            result: 1,
        },
        {
            value:  90,
            result: 0,
        },
        {
            value:  180,
            result: -1,
        },
        {
            value:  270,
            result: 0,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Cosd32(test.value), 10))
    }
}

func TestTand32(t *testing.T) {
    tests := []struct {
        value  float32
        result float32
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  45,
            result: 1,
        },
        {
            value:  135,
            result: -1,
        },
        {
            value:  180,
            result: 0,
        },
        {
            value:  225,
            result: 1,
        },
        {
            value:  315,
            result: -1,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Tand32(test.value), 10))
    }
}

func TestAsind32(t *testing.T) {
    tests := []struct {
        value  float32
        result float32
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  1,
            result: 90,
        },
        {
            value:  -1,
            result: -90,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Asind32(test.value), 10))
    }
}

func TestAcosd32(t *testing.T) {
    tests := []struct {
        value  float32
        result float32
    }{
        {
            value:  0,
            result: 90,
        },
        {
            value:  1,
            result: 0,
        },
        {
            value:  -1,
            result: 180,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Acosd32(test.value), 10))
    }
}

func TestAtand32(t *testing.T) {
    tests := []struct {
        value  float32
        result float32
    }{
        {
            value:  0,
            result: 0,
        },
        {
            value:  1,
            result: 45,
        },
        {
            value:  -1,
            result: -45,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Atand32(test.value), 10))
    }
}

func TestAtan2d32(t *testing.T) {
    tests := []struct {
        y      float32
        x      float32
        result float32
    }{
        {
            y:      0,
            x:      0,
            result: 0,
        },
        {
            y:      1,
            x:      0,
            result: 90,
        },
        {
            y:      -1,
            x:      0,
            result: -90,
        },
        {
            y:      0,
            x:      1,
            result: 0,
        },
        {
            y:      1,
            x:      1,
            result: 45,
        },
        {
            y:      1,
            x:      -1,
            result: 135,
        },
        {
            y:      -1,
            x:      1,
            result: -45,
        },
        {
            y:      -1,
            x:      -1,
            result: -135,
        },
    }
    for _, test := range tests {
        assert.Equal(t, test.result, Roundp32(Atan2d32(test.y, test.x), 10))
    }
}
