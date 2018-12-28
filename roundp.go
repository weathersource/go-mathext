package mathext

import (
	"math"
)

// Roundp rounds f to n decimal points of precision
func Roundp(f float64, n int) float64 {
	pow := math.Pow10(n)
	return math.Round(f*pow) / pow
}

// Roundp32 rounds f to n decimal points of precision
func Roundp32(f float32, n int) float32 {
	pow := math.Pow10(n)
	return float32(math.Round(float64(f)*pow) / pow)
}
