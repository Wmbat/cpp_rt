package maths

import "math/rand"

func RandVec3() Vec3 {
	return Vec3{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64()}
}

func RandVec3InRange(min, max float64) Vec3 {

	return Vec3{
		X: min + (rand.Float64() * (max - min)),
		Y: min + (rand.Float64() * (max - min)),
		Z: min + (rand.Float64() * (max - min))}
}

func RandVec3InUnitSphere() Vec3 {
	for {
		vec := RandVec3InRange(-1, 1)
		if vec.LengthSquared() >= 1 {
			continue
		}
		return vec
	}
}

func RandVec3InHemisphere(normal Vec3) Vec3 {
	unitSphere := RandVec3InUnitSphere()
	if DotProduct(unitSphere, normal) > 0.0 {
		return unitSphere
	} else {
		return unitSphere.Negate()
	}
}
