package maths

import "math/rand"

func Lerp(start *Vec3, end *Vec3, blendFactor float64)  Vec3 {
    startBlend := MultScalar(start, (1.0 - blendFactor))
    endBlend := MultScalar(end, blendFactor)

    return Add(&startBlend, &endBlend)
}

func Clamp(target float64, min float64, max float64) float64 {
    if target < min {
        return min
    }

    if target > max {
        return max
    }

    return target
}

func Reflect(normal *Vec3, incident *Vec3) Vec3 {
    rhs := MultScalar(normal, 2 * Vec3Dot(normal, incident))

    return Sub(incident, &rhs)
}

func RandomFloat64(min float64, max float64) float64 {
    return min + (max - min) * rand.Float64()
}
