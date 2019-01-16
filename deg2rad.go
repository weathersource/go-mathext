package mathext

import (
	"math"
)

// Deg2Rad converts a degrees to randian
func Deg2Rad(degrees float64) float64 {
	return degrees * math.Pi / 180
}
