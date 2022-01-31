package mathext

import (
	"math"
)

// Deg2Rad converts a degrees to radian
func Deg2Rad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Rad2Deg converts a radians to degrees
func Rad2Deg(radians float64) float64 {
	return radians * 180 / math.Pi
}

// Sind is the sin trigonometric function expecting x in degrees
func Sind(x float64) float64 {
    return math.Sin(Deg2Rad(x))
}

// Cosd is the cos trigonometric function expecting x in degrees
func Cosd(x float64) float64 {
    return math.Cos(Deg2Rad(x))
}

// Tand is the tan trigonometric function expecting x in degrees
func Tand(x float64) float64 {
    return math.Tan(Deg2Rad(x))
}

// Asind is the asin trigonometric function returning result in degrees
func Asind(x float64) float64 {
    return Rad2Deg(math.Asin(x))
}

// Asind is the acos trigonometric function returning result in degrees
func Acosd(x float64) float64 {
    return Rad2Deg(math.Acos(x))
}

// Atand is the atan trigonometric function returning result in degrees
func Atand(x float64) float64 {
    return Rad2Deg(math.Atan(x))
}

// Atan2d is the atan2 trigonometric function returning result in degrees
func Atan2d(y, x float64) float64 {
    return Rad2Deg(math.Atan2(y, x))
}

// Sind32 is the sin trigonometric function expecting x in degrees as float32
func Sind32(x float32) float32 {
    return float32(Sind(float64(x)))
}

// Cosd32 is the cos trigonometric function expecting x in degrees as float32
func Cosd32(x float32) float32 {
    return float32(Cosd(float64(x)))
}

// Tand32 is the tan trigonometric function expecting x in degrees as float32
func Tand32(x float32) float32 {
    return float32(Tand(float64(x)))
}

// Asind32 is the asin trigonometric function returning result in degrees as float32
func Asind32(x float32) float32 {
    return float32(Asind(float64(x)))
}

// Asind32 is the acos trigonometric function returning result in degrees as float32
func Acosd32(x float32) float32 {
    return float32(Acosd(float64(x)))
}

// Atand32 is the atan trigonometric function returning result in degrees as float32
func Atand32(x float32) float32 {
    return float32(Atand(float64(x)))
}

// Atan2d32 is the atan2 trigonometric function returning result in degrees as float32
func Atan2d32(y, x float32) float32 {
    return float32(Atan2d(float64(y), float64(x)))
}
