package mathext

import (
    "math"
    "testing"

    assert "github.com/stretchr/testify/assert"
)

func TestAbs32(t *testing.T) {
    x := float32(60.12)
    r := Abs32(x)
    assert.Equal(t, float32(math.Abs(float64(x))), r)
}
func TestAcos32(t *testing.T) {
    x := float32(0.12)
    r := Acos32(x)
    assert.Equal(t, float32(math.Acos(float64(x))), r)
}
func TestAcosh32(t *testing.T) {
    x := float32(60.12)
    r := Acosh32(x)
    assert.Equal(t, float32(math.Acosh(float64(x))), r)
}
func TestAsin32(t *testing.T) {
    x := float32(0.12)
    r := Asin32(x)
    assert.Equal(t, float32(math.Asin(float64(x))), r)
}
func TestAsinh32(t *testing.T) {
    x := float32(60.12)
    r := Asinh32(x)
    assert.Equal(t, float32(math.Asinh(float64(x))), r)
}
func TestAtan32(t *testing.T) {
    x := float32(60.12)
    r := Atan32(x)
    assert.Equal(t, float32(math.Atan(float64(x))), r)
}
func TestAtan2_32(t *testing.T) {
    y := float32(10.62)
    x := float32(60.12)
    r := Atan2_32(y, x)
    assert.Equal(t, float32(math.Atan2(float64(y), float64(x))), r)
}
func TestAtanh32(t *testing.T) {
    x := float32(0.12)
    r := Atanh32(x)
    assert.Equal(t, float32(math.Atanh(float64(x))), r)
}
func TestCbrt32(t *testing.T) {
    x := float32(60.12)
    r := Cbrt32(x)
    assert.Equal(t, float32(math.Cbrt(float64(x))), r)
}
func TestCeil32(t *testing.T) {
    x := float32(60.12)
    r := Ceil32(x)
    assert.Equal(t, float32(math.Ceil(float64(x))), r)
}
func TestCopysign32(t *testing.T) {
    x := float32(60.12)
    y := float32(0.12)
    r := Copysign32(x, y)
    assert.Equal(t, float32(math.Copysign(float64(x), float64(y))), r)
}
func TestCos32(t *testing.T) {
    x := float32(60.12)
    r := Cos32(x)
    assert.Equal(t, float32(math.Cos(float64(x))), r)
}
func TestCosh32(t *testing.T) {
    x := float32(60.12)
    r := Cosh32(x)
    assert.Equal(t, float32(math.Cosh(float64(x))), r)
}
func TestDim32(t *testing.T) {
    x := float32(60.12)
    y := float32(0.12)
    r := Dim32(x, y)
    assert.Equal(t, float32(math.Dim(float64(x), float64(y))), r)
}
func TestErf32(t *testing.T) {
    x := float32(60.12)
    r := Erf32(x)
    assert.Equal(t, float32(math.Erf(float64(x))), r)
}
func TestErfc32(t *testing.T) {
    x := float32(60.12)
    r := Erfc32(x)
    assert.Equal(t, float32(math.Erfc(float64(x))), r)
}
func TestErfcinv32(t *testing.T) {
    x := float32(0.12)
    r := Erfcinv32(x)
    assert.Equal(t, float32(math.Erfcinv(float64(x))), r)
}
func TestErfinv32(t *testing.T) {
    x := float32(0.12)
    r := Erfinv32(x)
    assert.Equal(t, float32(math.Erfinv(float64(x))), r)
}
func TestExp32(t *testing.T) {
    x := float32(60.12)
    r := Exp32(x)
    assert.Equal(t, float32(math.Exp(float64(x))), r)
}
func TestExp2_32(t *testing.T) {
    x := float32(60.12)
    r := Exp2_32(x)
    assert.Equal(t, float32(math.Exp2(float64(x))), r)
}
func TestExpm1_32(t *testing.T) {
    x := float32(60.12)
    r := Expm1_32(x)
    assert.Equal(t, float32(math.Expm1(float64(x))), r)
}
func TestFMA32(t *testing.T) {
    x := float32(60.12)
    y := float32(60.12)
    z := float32(60.12)
    r := FMA32(x, y, z)
    assert.Equal(t, float32(math.FMA(float64(x), float64(y), float64(z))), r)
}
func TestFloor32(t *testing.T) {
    x := float32(60.12)
    r := Floor32(x)
    assert.Equal(t, float32(math.Floor(float64(x))), r)
}
func TestFrexp32(t *testing.T) {
    x := float32(60.12)
    r0, r1 := Frexp32(x)
    e0, e1 := math.Frexp(float64(x))
    assert.Equal(t, float32(e0), r0)
    assert.Equal(t, e1, r1)
}
func TestGamma32(t *testing.T) {
    x := float32(60.12)
    r := Gamma32(x)
    assert.Equal(t, float32(math.Gamma(float64(x))), r)
}
func TestHypot32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Hypot32(x, y)
    assert.Equal(t, float32(math.Hypot(float64(x), float64(y))), r)
}
func TestIlogb32(t *testing.T) {
    x := float32(60.12)
    r := Ilogb32(x)
    assert.Equal(t, math.Ilogb(float64(x)), r)
}

func TestInf32(t *testing.T) {
    x := int(1)
    r := Inf32(x)
    assert.Equal(t, float32(math.Inf(x)), r)
}
func TestIsInf32(t *testing.T) {
    x := float32(60.12)
    y := int(1)
    r := IsInf32(x, 1)
    assert.Equal(t, math.IsInf(float64(x), y), r)
}
func TestIsNaN32(t *testing.T) {
    x := float32(60.12)
    r := IsNaN32(x)
    assert.Equal(t, math.IsNaN(float64(x)), r)
}
func TestJ0_32(t *testing.T) {
    x := float32(60.12)
    r := J0_32(x)
    assert.Equal(t, float32(math.J0(float64(x))), r)
}
func TestJ1_32(t *testing.T) {
    x := float32(60.12)
    r := J1_32(x)
    assert.Equal(t, float32(math.J1(float64(x))), r)
}
func TestJn_32(t *testing.T) {
    x := int(1)
    y := float32(60.12)
    r := Jn_32(x, y)
    assert.Equal(t, float32(math.Jn(x, float64(y))), r)
}
func TestLdexp32(t *testing.T) {
    x := float32(60.12)
    y := int(1)
    r := Ldexp32(x, y)
    assert.Equal(t, float32(math.Ldexp(float64(x), y)), r)
}
func TestLgamma32(t *testing.T) {
    x := float32(60.12)
    r0, r1 := Lgamma32(x)
    e0, e1 := math.Lgamma(float64(x))
    assert.Equal(t, float32(e0), r0)
    assert.Equal(t, e1, r1)
}
func TestLog32(t *testing.T) {
    x := float32(60.12)
    r := Log32(x)
    assert.Equal(t, float32(math.Log(float64(x))), r)
}
func TestLog10_32(t *testing.T) {
    x := float32(60.12)
    r := Log10_32(x)
    assert.Equal(t, float32(math.Log10(float64(x))), r)
}
func TestLog1p32(t *testing.T) {
    x := float32(60.12)
    r := Log1p32(x)
    assert.Equal(t, float32(math.Log1p(float64(x))), r)
}
func TestLog2_32(t *testing.T) {
    x := float32(60.12)
    r := Log2_32(x)
    assert.Equal(t, float32(math.Log2(float64(x))), r)
}
func TestLogb32(t *testing.T) {
    x := float32(60.12)
    r := Logb32(x)
    assert.Equal(t, float32(math.Logb(float64(x))), r)
}
func TestMax32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Max32(x, y)
    assert.Equal(t, float32(math.Max(float64(x), float64(y))), r)
}
func TestMin32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Min32(x, y)
    assert.Equal(t, float32(math.Min(float64(x), float64(y))), r)
}
func TestMod32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Mod32(x, y)
    assert.Equal(t, float32(math.Mod(float64(x), float64(y))), r)
}
func TestModf32(t *testing.T) {
    x := float32(60.12)
    r0, r1 := Modf32(x)
    e0, e1 := math.Modf(float64(x))
    assert.Equal(t, float32(e0), r0)
    assert.Equal(t, float32(e1), r1)
}
func TestNaN32(t *testing.T) {
    r := NaN32()
    assert.Equal(t, math.IsNaN(math.NaN()), math.IsNaN(float64(r)))
}
func TestNextafter32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Nextafter32(x, y)
    assert.Equal(t, math.Nextafter32(x, y), r)
}
func TestPow32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Pow32(x, y)
    assert.Equal(t, float32(math.Pow(float64(x), float64(y))), r)
}
func TestPow10_32(t *testing.T) {
    x := int(4)
    r := Pow10_32(x)
    assert.Equal(t, float32(math.Pow10(x)), r)
}
func TestRemainder32(t *testing.T) {
    x := float32(60.12)
    y := float32(10.62)
    r := Remainder32(x, y)
    assert.Equal(t, float32(math.Remainder(float64(x), float64(y))), r)
}
func TestRound32(t *testing.T) {
    x := float32(60.12)
    r := Round32(x)
    assert.Equal(t, float32(math.Round(float64(x))), r)
}
func TestRoundToEven32(t *testing.T) {
    x := float32(60.12)
    r := RoundToEven32(x)
    assert.Equal(t, float32(math.RoundToEven(float64(x))), r)
}
func TestSignbit32(t *testing.T) {
    x := float32(60.12)
    r := Signbit32(x)
    assert.Equal(t, math.Signbit(float64(x)), r)
}
func TestSin32(t *testing.T) {
    x := float32(60.12)
    r := Sin32(x)
    assert.Equal(t, float32(math.Sin(float64(x))), r)
}
func TestSincos32(t *testing.T) {
    x := float32(60.12)
    r0, r1 := Sincos32(x)
    e0, e1 := math.Sincos(float64(x))
    assert.Equal(t, float32(e0), r0)
    assert.Equal(t, float32(e1), r1)
}
func TestSinh32(t *testing.T) {
    x := float32(60.12)
    r := Sinh32(x)
    assert.Equal(t, float32(math.Sinh(float64(x))), r)
}
func TestSqrt32(t *testing.T) {
    x := float32(60.12)
    r := Sqrt32(x)
    assert.Equal(t, float32(math.Sqrt(float64(x))), r)
}
func TestTan32(t *testing.T) {
    x := float32(60.12)
    r := Tan32(x)
    assert.Equal(t, float32(math.Tan(float64(x))), r)
}
func TestTanh32(t *testing.T) {
    x := float32(60.12)
    r := Tanh32(x)
    assert.Equal(t, float32(math.Tanh(float64(x))), r)
}
func TestTrunc32(t *testing.T) {
    x := float32(60.12)
    r := Trunc32(x)
    assert.Equal(t, float32(math.Trunc(float64(x))), r)
}
func TestY0_32(t *testing.T) {
    x := float32(60.12)
    r := Y0_32(x)
    assert.Equal(t, float32(math.Y0(float64(x))), r)
}
func TestY1_32(t *testing.T) {
    x := float32(60.12)
    r := Y1_32(x)
    assert.Equal(t, float32(math.Y1(float64(x))), r)
}
func TestYn32(t *testing.T) {
    x := int(2)
    y := float32(60.12)
    r := Yn32(x, y)
    assert.Equal(t, float32(math.Yn(x, float64(y))), r)
}
