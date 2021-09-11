package maths

import "math/rand"

func Lerp(start *Vec3, end *Vec3, blendFactor float64)  Vec3 {
    startBlend := MultScalar(start, (1.0 - blendFactor))
    endBlend := MultScalar(end, blendFactor)

    return Add(&startBlend, &endBlend)
}

func RandomFloat64(min float64, max float64) float64 {
    return min + (max - min) * rand.Float64()
}
