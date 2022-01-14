package mathext

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestFToC(t *testing.T) {
    i64 := float64(60.12)
    o64 := FToC(i64)
    assert.Equal(t, float64(15.622222222222222), o64)
    o64 = FToC(i64, 2)
    assert.Equal(t, float64(15.62), o64)
}

func TestCToF(t *testing.T) {
    i64 := float64(15.62)
    o64 := CToF(i64)
    assert.Equal(t, float64(60.116), o64)
    o64 = CToF(i64, 2)
    assert.Equal(t, float64(60.12), o64)
}

func TestFToK(t *testing.T) {
    i64 := float64(60.12)
    o64 := FToK(i64)
    assert.Equal(t, float64(288.7722222222222), o64)
    o64 = FToK(i64, 2)
    assert.Equal(t, float64(288.77), o64)
}

func TestKToF(t *testing.T) {
    i64 := float64(288.77)
    o64 := KToF(i64)
    assert.Equal(t, float64(60.116000000000014), o64)
    o64 = KToF(i64, 2)
    assert.Equal(t, float64(60.12), o64)
}

func TestCToK(t *testing.T) {
    i64 := float64(15.62)
    o64 := CToK(i64)
    assert.Equal(t, float64(288.77), o64)
    o64 = CToK(i64, 1)
    assert.Equal(t, float64(288.8), o64)
}

func TestKToC(t *testing.T) {
    i64 := float64(288.77)
    o64 := KToC(i64)
    assert.Equal(t, float64(15.620000000000005), o64)
    o64 = KToC(i64, 2)
    assert.Equal(t, float64(15.62), o64)
}

func TestMiToKm(t *testing.T) {
    i64 := float64(60.12)
    o64 := MiToKm(i64)
    assert.Equal(t, float64(96.75352079999999), o64)
    o64 = MiToKm(i64, 2)
    assert.Equal(t, float64(96.75), o64)
}

func TestKmToMi(t *testing.T) {
    i64 := float64(96.75)
    o64 := KmToMi(i64)
    assert.Equal(t, float64(60.117812270868804), o64)
    o64 = KmToMi(i64, 2)
    assert.Equal(t, float64(60.12), o64)
}

func TestFtToM(t *testing.T) {
    i64 := float64(60.12)
    o64 := FtToM(i64)
    assert.Equal(t, float64(18.324576), o64)
    o64 = FtToM(i64, 2)
    assert.Equal(t, float64(18.32), o64)
}

func TestMToFt(t *testing.T) {
    i64 := float64(18.32)
    o64 := MToFt(i64)
    assert.Equal(t, float64(60.10498687664042), o64)
    o64 = MToFt(i64, 2)
    assert.Equal(t, float64(60.1), o64)
}

func TestInToCm(t *testing.T) {
    i64 := float64(60.12)
    o64 := InToCm(i64)
    assert.Equal(t, float64(152.7048), o64)
    o64 = InToCm(i64, 2)
    assert.Equal(t, float64(152.7), o64)
}

func TestCmToIn(t *testing.T) {
    i64 := float64(152.7)
    o64 := CmToIn(i64)
    assert.Equal(t, float64(60.11811023622047), o64)
    o64 = CmToIn(i64, 2)
    assert.Equal(t, float64(60.12), o64)
}

func TestMphToKmph(t *testing.T) {
    i64 := float64(60.12)
    o64 := MphToKmph(i64)
    assert.Equal(t, float64(96.75352079999999), o64)
    o64 = MphToKmph(i64, 2)
    assert.Equal(t, float64(96.75), o64)
}

func TestKmphToMph(t *testing.T) {
    i64 := float64(96.75)
    o64 := KmphToMph(i64)
    assert.Equal(t, float64(60.117812270868804), o64)
    o64 = KmphToMph(i64, 2)
    assert.Equal(t, float64(60.12), o64)
}

func TestDegToRad(t *testing.T) {
    i64 := float64(180)
    o64 := DegToRad(i64)
    assert.Equal(t, float64(3.141592653589793), o64)
    o64 = DegToRad(i64, 2)
    assert.Equal(t, float64(3.14), o64)
}

func TestRadToDeg(t *testing.T) {
    i64 := float64(3.14)
    o64 := RadToDeg(i64)
    assert.Equal(t, float64(179.90874767107852), o64)
    o64 = RadToDeg(i64, 2)
    assert.Equal(t, float64(179.91), o64)
}

