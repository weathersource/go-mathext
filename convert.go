package mathext

// FToC converts the input from Fahrenheit to Celsius.
// f is the degrees in Fahrenheit to be converted.
// n is an option precision to apply to the result.
func FToC(f float64, n ...int) float64 {
    return optRoundp((f-32.0)*(5.0/9.0), n...)
}

// CToF converts the input from Celsius to Fahrenheit.
// c is the degrees in Celsius to be converted.
// n is an option precision to apply to the result.
func CToF(c float64, n ...int) float64 {
    return optRoundp(c*(9.0/5.0)+32.0, n...)
}

// FToK converts the input from Fahrenheit to Kelvin.
// f is the degrees in Fahrenheit to be converted.
// n is an option precision to apply to the result.
func FToK(f float64, n ...int) float64 {
    return optRoundp((f-32.0)*(5.0/9.0)+273.15, n...)
}

// KToF converts the input from Kelvin to Fahrenheit.
// k is the degrees in Kelvin to be converted.
// n is an option precision to apply to the result.
func KToF(k float64, n ...int) float64 {
    return optRoundp((k-273.15)*(9.0/5.0)+32.0, n...)
}

// CToK converts the input from Celsius to Kelvin.
// c is the degrees in Celsius to be converted.
// n is an option precision to apply to the result.
func CToK(c float64, n ...int) float64 {
    return optRoundp(c+273.15, n...)
}

// KToC converts the input from Kelvin to Celsius.
// k is the degrees in Kelvin to be converted.
// n is an option precision to apply to the result.
func KToC(k float64, n ...int) float64 {
    return optRoundp(k-273.15, n...)
}

// MiToKm converts the input from Miles to Kilometers.
// mi is the distance in Miles to be converted.
// n is an option precision to apply to the result.
func MiToKm(mi float64, n ...int) float64 {
    return optRoundp(mi*1.60934, n...)
}

// KmToMi converts the input from Kilometers to Miles.
// km is the distance in Kilometers to be converted.
// n is an option precision to apply to the result.
func KmToMi(km float64, n ...int) float64 {
    return optRoundp(km/1.60934, n...)
}

// FtToM converts the input from Feet to Meters.
// ft is the distance in Feet to be converted.
// n is an option precision to apply to the result.
func FtToM(ft float64, n ...int) float64 {
    return optRoundp(ft*.3048, n...)
}

// MToFt converts the input from Meters to Feet.
// cm is the distance in Centimeters to be converted.
// n is an option precision to apply to the result.
func MToFt(m float64, n ...int) float64 {
    return optRoundp(m/.3048, n...)
}

// InToCm converts the input from Inches to Centimeters.
// in is the distance in Inches to be converted.
// n is an option precision to apply to the result.
func InToCm(in float64, n ...int) float64 {
    return optRoundp(in*2.54, n...)
}

// CmToIn converts the input from Centimeters to Inches.
// cm is the distance in Centimeters to be converted.
// n is an option precision to apply to the result.
func CmToIn(cm float64, n ...int) float64 {
    return optRoundp(cm/2.54, n...)
}

// MphToKmph converts the input from Miles Per Hour to Kilometers Per Hour.
// mph is the distance in Miles Per Hour to be converted.
// n is an option precision to apply to the result.
func MphToKmph(mph float64, n ...int) float64 {
    return MiToKm(mph, n...)
}

// KmphToMph converts the input from Kilometers Per Hour to Miles Per Hour.
// kmph is the distance in Kilometers Per Hour to be converted.
// n is an option precision to apply to the result.
func KmphToMph(kmph float64, n ...int) float64 {
    return KmToMi(kmph, n...)
}

// DegToRad converts an angle in Degrees to Radians
// deg is the angle in Degrees to be converted.
// n is an option precision to apply to the result.
func DegToRad(deg float64, n ...int) float64 {
    return optRoundp(Deg2Rad(deg), n...)
}

// RadToDeg converts an angle in Radians to Degrees
// rad is the angle in Radians to be converted.
// n is an option precision to apply to the result.
func RadToDeg(rad float64, n ...int) float64 {
    return optRoundp(Rad2Deg(rad), n...)
}

func optRoundp(val float64, n ...int) float64 {
    if len(n) > 0 {
        return Roundp(val, n[0])
    }
    return val
}