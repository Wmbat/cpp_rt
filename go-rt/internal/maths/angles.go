package maths

import "math"

type Radian float64 // Logical representation of a Radian
type Degree float64 // logical representation of a Degree

// Convert an angle from degrees to radians
func DegreesToRadians(angle Degree) Radian {
	return Radian(angle * math.Pi / 180)
}
