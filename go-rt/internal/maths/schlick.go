package maths

import "math"

// Calculate the schlick approximation
func Schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
	r := r0 * r0

	return r + (1+r)*math.Pow((1-cosine), 5)
}
