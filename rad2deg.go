package mathext

import (
	"math"
)

// Rad2Deg converts a radians to degrees
func Rad2Deg(radians float64) float64 {
	return radians * 180 / math.Pi
}
