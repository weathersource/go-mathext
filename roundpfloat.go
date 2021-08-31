package mathext

import (
	"math"
)

// Roundpfloat rounds f to the nearest n points of precision
func Roundpfloat(f float64, n float64) float64 {
	return math.Round(f/n) * n
}
