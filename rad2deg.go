package mathext

import (
	"math"
)

// Rad2Deg converts a randian to degree
func Rad2Deg(radian float64) float64 {
	return radian * 180 / math.Pi
}
