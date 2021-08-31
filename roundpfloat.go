package mathext

import (
	"math"
)

// Roundpfloat rounds f to the nearest n points of precision
func Roundpfloat(f float64, n float64) float64 {
	return math.Round(f/n) * n
}

// Roundpfloat rounds f to the nearest n points of precision
func Roundpfloat32(f float32, n float32) float32 {
	f64 := float64(f)
	n64 := float64(n)
	return float32(math.Round(f64/n64) * n64)
}
