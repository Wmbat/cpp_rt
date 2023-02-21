package maths

import "math"

type Radian float64
type Degree float64

func DegreesToRadians(angle Degree) Radian {
	return Radian(angle * math.Pi / 180)
}
