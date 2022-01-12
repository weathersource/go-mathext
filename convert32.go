package mathext

// FToC32 converts the input from Fahrenheit to Celsius.
// f is the degrees in Fahrenheit to be converted.
// n is an option precision to apply to the result.
func FToC32(f float32, n ...int) float32 {
    return optRoundp32((f-32.0)*(5.0/9.0), n...)
}

// CToF32 converts the input from Celsius to Fahrenheit.
// c is the degrees in Celsius to be converted.
// n is an option precision to apply to the result.
func CToF32(c float32, n ...int) float32 {
    return optRoundp32(c*(9.0/5.0)+32.0, n...)
}

// FToK32 converts the input from Fahrenheit to Kelvin.
// f is the degrees in Fahrenheit to be converted.
// n is an option precision to apply to the result.
func FToK32(f float32, n ...int) float32 {
    return optRoundp32((f-32.0)*(5.0/9.0)+273.15, n...)
}

// KToF32 converts the input from Kelvin to Fahrenheit.
// k is the degrees in Kelvin to be converted.
// n is an option precision to apply to the result.
func KToF32(k float32, n ...int) float32 {
    return optRoundp32((k-273.15)*(9.0/5.0)+32.0, n...)
}

// CToK32 converts the input from Celsius to Kelvin.
// c is the degrees in Celsius to be converted.
// n is an option precision to apply to the result.
func CToK32(c float32, n ...int) float32 {
    return optRoundp32(c+273.15, n...)
}

// KToC32 converts the input from Kelvin to Celsius.
// k is the degrees in Kelvin to be converted.
// n is an option precision to apply to the result.
func KToC32(k float32, n ...int) float32 {
    return optRoundp32(k-273.15, n...)
}

// MiToKm32 converts the input from Miles to Kilometers.
// mi is the distance in Miles to be converted.
// n is an option precision to apply to the result.
func MiToKm32(mi float32, n ...int) float32 {
    return optRoundp32(mi*1.60934, n...)
}

// KmToMi32 converts the input from Kilometers to Miles.
// km is the distance in Kilometers to be converted.
// n is an option precision to apply to the result.
func KmToMi32(km float32, n ...int) float32 {
    return optRoundp32(km/1.60934, n...)
}

// FtToM32 converts the input from Feet to Meters.
// ft is the distance in Feet to be converted.
// n is an option precision to apply to the result.
func FtToM32(ft float32, n ...int) float32 {
    return optRoundp32(ft*.3048, n...)
}

// MToFt32 converts the input from Meters to Feet.
// cm is the distance in Centimeters to be converted.
// n is an option precision to apply to the result.
func MToFt32(m float32, n ...int) float32 {
    return optRoundp32(m/.3048, n...)
}

// InToCm32 converts the input from Inches to Centimeters.
// in is the distance in Inches to be converted.
// n is an option precision to apply to the result.
func InToCm32(in float32, n ...int) float32 {
    return optRoundp32(in*2.54, n...)
}

// CmToIn32 converts the input from Centimeters to Inches.
// cm is the distance in Centimeters to be converted.
// n is an option precision to apply to the result.
func CmToIn32(cm float32, n ...int) float32 {
    return optRoundp32(cm/2.54, n...)
}


// DegToRad32 converts an angle in Degrees to Radians
// deg is the angle in Degrees to be converted.
// n is an option precision to apply to the result.
func DegToRad32(deg float32, n ...int) float32 {
    return optRoundp32(float32(Deg2Rad(float64(deg))), n...)
}

// RadToDeg32 converts an angle in Radians to Degrees
// rad is the angle in Radians to be converted.
// n is an option precision to apply to the result.
func RadToDeg32(rad float32, n ...int) float32 {
    return optRoundp32(float32(Rad2Deg(float64(rad))), n...)
}

func optRoundp32(val float32, n ...int) float32{
    if len(n) > 0 {
        return Roundp32(val, n[0])
    }
    return val
}