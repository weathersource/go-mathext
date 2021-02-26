package mathext

import (
	"math"
)

// Ceilp Ceils f to n decimal points of precision
func Ceilp(f float64, n int) float64 {
	pow := math.Pow10(n)
	return math.Ceil(f*pow) / pow
}

// Ceilp32 Ceils f to n decimal points of precision
func Ceilp32(f float32, n int) float32 {
	pow := math.Pow10(n)
	return float32(math.Ceil(float64(f)*pow) / pow)
}
