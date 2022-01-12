package mathext

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestFToC32(t *testing.T) {
    i32 := float32(60.12)
    o32 := FToC32(i32)
    assert.Equal(t, float32(15.622222), o32)
    o32 = FToC32(i32, 2)
    assert.Equal(t, float32(15.62), o32)
}

func TestCToF32(t *testing.T) {
    i32 := float32(15.62)
    o32 := CToF32(i32)
    assert.Equal(t, float32(60.115997), o32)
    o32 = CToF32(i32, 2)
    assert.Equal(t, float32(60.12), o32)
}

func TestFToK32(t *testing.T) {
    i32 := float32(60.12)
    o32 := FToK32(i32)
    assert.Equal(t, float32(288.77222), o32)
    o32 = FToK32(i32, 2)
    assert.Equal(t, float32(288.77), o32)
}

func TestKToF32(t *testing.T) {
    i32 := float32(288.77)
    o32 := KToF32(i32)
    assert.Equal(t, float32(60.11599), o32)
    o32 = KToF32(i32, 2)
    assert.Equal(t, float32(60.12), o32)
}

func TestCToK32(t *testing.T) {
    i32 := float32(15.62)
    o32 := CToK32(i32)
    assert.Equal(t, float32(288.77), o32)
    o32 = CToK32(i32, 1)
    assert.Equal(t, float32(288.8), o32)
}

func TestKToC32(t *testing.T) {
    i32 := float32(288.77)
    o32 := KToC32(i32)
    assert.Equal(t, float32(15.619995), o32)
    o32 = KToC32(i32, 2)
    assert.Equal(t, float32(15.62), o32)
}

func TestMiToKm32(t *testing.T) {
    i32 := float32(60.12)
    o32 := MiToKm32(i32)
    assert.Equal(t, float32(96.75352), o32)
    o32 = MiToKm32(i32, 2)
    assert.Equal(t, float32(96.75), o32)
}

func TestKmToM32(t *testing.T) {
    i32 := float32(96.75)
    o32 := KmToMi32(i32)
    assert.Equal(t, float32(60.117813), o32)
    o32 = KmToMi32(i32, 2)
    assert.Equal(t, float32(60.12), o32)
}

func TestFtToM32(t *testing.T) {
    i32 := float32(60.12)
    o32 := FtToM32(i32)
    assert.Equal(t, float32(18.324575), o32)
    o32 = FtToM32(i32, 2)
    assert.Equal(t, float32(18.32), o32)
}

func TestMToFt32(t *testing.T) {
    i32 := float32(18.32)
    o32 := MToFt32(i32)
    assert.Equal(t, float32(60.104984), o32)
    o32 = MToFt32(i32, 2)
    assert.Equal(t, float32(60.1), o32)
}

func TestInToCm32(t *testing.T) {
    i32 := float32(60.12)
    o32 := InToCm32(i32)
    assert.Equal(t, float32(152.70479), o32)
    o32 = InToCm32(i32, 2)
    assert.Equal(t, float32(152.7), o32)
}

func TestCmToIn32(t *testing.T) {
    i32 := float32(152.7)
    o32 := CmToIn32(i32)
    assert.Equal(t, float32(60.11811), o32)
    o32 = CmToIn32(i32, 2)
    assert.Equal(t, float32(60.12), o32)
}

func TestDegToRad32(t *testing.T) {
    i32 := float32(180)
    o32 := DegToRad32(i32)
    assert.Equal(t, float32(3.1415927), o32)
    o32 = DegToRad32(i32, 2)
    assert.Equal(t, float32(3.14), o32)
}

func TestRadToDeg32(t *testing.T) {
    i32 := float32(3.14)
    o32 := RadToDeg32(i32)
    assert.Equal(t, float32(179.90875), o32)
    o32 = RadToDeg32(i32, 2)
    assert.Equal(t, float32(179.91), o32)
}

