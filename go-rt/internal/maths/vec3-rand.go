package maths

import (
	"math/rand"
)

func RandVec3(rng *rand.Rand) Vec3 {
	return Vec3{
		X: rng.Float64(),
		Y: rng.Float64(),
		Z: rng.Float64()}
}

func RandVec3InRange(rng *rand.Rand, min, max float64) Vec3 {

	return Vec3{
		X: min + (rng.Float64() * (max - min)),
		Y: min + (rng.Float64() * (max - min)),
		Z: min + (rng.Float64() * (max - min))}
}

func RandVec3InUnitSphere(rng *rand.Rand) Vec3 {
	for {
		vec := RandVec3InRange(rng, -1, 1)
		if vec.LengthSquared() >= 1 {
			continue
		}
		return vec
	}
}

func RandVec3InHemisphere(rng *rand.Rand, normal Vec3) Vec3 {
	unitSphere := RandVec3InUnitSphere(rng)
	if DotProduct(unitSphere, normal) > 0.0 {
		return unitSphere
	} else {
		return unitSphere.Negate()
	}
}
