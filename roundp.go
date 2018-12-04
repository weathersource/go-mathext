package mathext

import (
	"math"
)

// Roundp rounds f to n decimal points of precision
func Roundp(f float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return float64(math.Round(f*pow)) / pow
}
