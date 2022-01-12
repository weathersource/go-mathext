package mathext

import "math"

func Abs32(x float32) float32 {
    return float32(math.Abs(float64(x)))
}
func Acos32(x float32) float32 {
    return float32(math.Acos(float64(x)))
}
func Acosh32(x float32) float32 {
    return float32(math.Acosh(float64(x)))
}
func Asin32(x float32) float32 {
    return float32(math.Asin(float64(x)))
}
func Asinh32(x float32) float32 {
    return float32(math.Asinh(float64(x)))
}
func Atan32(x float32) float32 {
    return float32(math.Atan(float64(x)))
}
func Atan2_32(y, x float32) float32 {
    return float32(math.Atan2(float64(y), float64(x)))
}
func Atanh32(x float32) float32 {
    return float32(math.Atanh(float64(x)))
}
func Cbrt32(x float32) float32 {
    return float32(math.Cbrt(float64(x)))
}
func Ceil32(x float32) float32 {
    return float32(math.Ceil(float64(x)))
}
func Copysign32(x, y float32) float32 {
    return float32(math.Copysign(float64(x), float64(y)))
}
func Cos32(x float32) float32 {
    return float32(math.Cos(float64(x)))
}
func Cosh32(x float32) float32 {
    return float32(math.Cosh(float64(x)))
}
func Dim32(x, y float32) float32 {
    return float32(math.Dim(float64(x), float64(y)))
}
func Erf32(x float32) float32 {
    return float32(math.Erf(float64(x)))
}
func Erfc32(x float32) float32 {
    return float32(math.Erfc(float64(x)))
}
func Erfcinv32(x float32) float32 {
    return float32(math.Erfcinv(float64(x)))
}
func Erfinv32(x float32) float32 {
    return float32(math.Erfinv(float64(x)))
}
func Exp32(x float32) float32 {
    return float32(math.Exp(float64(x)))
}
func Exp2_32(x float32) float32 {
    return float32(math.Exp2(float64(x)))
}
func Expm1_32(x float32) float32 {
    return float32(math.Expm1(float64(x)))
}
func FMA32(x, y, z float32) float32 {
    return float32(math.FMA(float64(x), float64(y), float64(z)))
}
func Floor32(x float32) float32 {
    return float32(math.Floor(float64(x)))
}
func Frexp32(f float32) (frac float32, exp int) {
    frac64, exp0 := math.Frexp(float64(f))
    frac = float32(frac64)
    exp = exp0
    return
}
func Gamma32(x float32) float32 {
    return float32(math.Gamma(float64(x)))
}
func Hypot32(p, q float32) float32 {
    return float32(math.Hypot(float64(p), float64(q)))
}
func Ilogb32(x float32) int {
    return math.Ilogb(float64(x))
}
func Inf32(sign int) float32 {
    return float32(math.Inf(sign))
}
func IsInf32(f float32, sign int) bool {
    return math.IsInf(float64(f), sign)
}
func IsNaN32(f float32) (is bool) {
    is = math.IsNaN(float64(f))
    return
}
func J0_32(x float32) float32 {
    return float32(math.J0(float64(x)))
}
func J1_32(x float32) float32 {
    return float32(math.J1(float64(x)))
}
func Jn_32(n int, x float32) float32 {
    return float32(math.Jn(n, float64(x)))
}
func Ldexp32(frac float32, exp int) float32 {
    return float32(math.Ldexp(float64(frac), exp))
}
func Lgamma32(x float32) (lgamma float32, sign int) {
    lgamma64, sign0 := math.Lgamma(float64(x))
    lgamma = float32(lgamma64)
    sign = sign0
    return
}
func Log32(x float32) float32 {
    return float32(math.Log(float64(x)))
}
func Log10_32(x float32) float32 {
    return float32(math.Log10(float64(x)))
}
func Log1p32(x float32) float32 {
    return float32(math.Log1p(float64(x)))
}
func Log2_32(x float32) float32 {
    return float32(math.Log2(float64(x)))
}
func Logb32(x float32) float32 {
    return float32(math.Logb(float64(x)))
}
func Max32(x, y float32) float32 {
    return float32(math.Max(float64(x), float64(y)))
}
func Min32(x, y float32) float32 {
    return float32(math.Min(float64(x), float64(y)))
}
func Mod32(x, y float32) float32 {
    return float32(math.Mod(float64(x), float64(y)))
}
func Modf32(f float32) (int float32, frac float32) {
    int64, frac64 := math.Modf(float64(f))
    int = float32(int64)
    frac = float32(frac64)
    return
}
func NaN32() float32 {
    return float32(math.NaN())
}
func Nextafter32(x, y float32) (r float32) {
    r = math.Nextafter32(x, y)
    return
}
func Pow32(x, y float32) float32 {
    return float32(math.Pow(float64(x), float64(y)))
}
func Pow10_32(n int) float32 {
    return float32(math.Pow10(n))
}
func Remainder32(x, y float32) float32 {
    return float32(math.Remainder(float64(x), float64(y)))
}
func Round32(x float32) float32 {
    return float32(math.Round(float64(x)))
}
func RoundToEven32(x float32) float32 {
    return float32(math.RoundToEven(float64(x)))
}
func Signbit32(x float32) bool {
    return math.Signbit(float64(x))
}
func Sin32(x float32) float32 {
    return float32(math.Sin(float64(x)))
}
func Sincos32(x float32) (sin, cos float32) {
    sin64, cos64 := math.Sincos(float64(x))
    sin = float32(sin64)
    cos = float32(cos64)
    return
}
func Sinh32(x float32) float32 {
    return float32(math.Sinh(float64(x)))
}
func Sqrt32(x float32) float32 {
    return float32(math.Sqrt(float64(x)))
}
func Tan32(x float32) float32 {
    return float32(math.Tan(float64(x)))
}
func Tanh32(x float32) float32 {
    return float32(math.Tanh(float64(x)))
}
func Trunc32(x float32) float32 {
    return float32(math.Trunc(float64(x)))
}
func Y0_32(x float32) float32 {
    return float32(math.Y0(float64(x)))
}
func Y1_32(x float32) float32 {
    return float32(math.Y1(float64(x)))
}
func Yn32(n int, x float32) float32 {
    return float32(math.Yn(n, float64(x)))
}