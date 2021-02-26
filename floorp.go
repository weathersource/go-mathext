package mathext

import (
	"math"
)

// Floorp Floors f to n decimal points of precision
func Floorp(f float64, n int) float64 {
	pow := math.Pow10(n)
	return math.Floor(f*pow) / pow
}

// Floorp32 Floors f to n decimal points of precision
func Floorp32(f float32, n int) float32 {
	pow := math.Pow10(n)
	return float32(math.Floor(float64(f)*pow) / pow)
}
