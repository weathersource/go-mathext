package mathext

import (
	"math"
)

// Deg2Rad converts a degree to randian
func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}
