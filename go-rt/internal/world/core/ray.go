package core

import "github.com/wmbat/ray_tracer/internal/maths"

// Represent a physical ray mean to interact with it's surroundings
type Ray struct {
	Origin    maths.Point3 // The starting point of the ray
	Direction maths.Vec3   // The direction of the ray
}

// Get the position of a ray at a specific distance along the line of it's direction
func (this Ray) At(distance float64) maths.Point3 {
	return this.Origin.Add(this.Direction.Scale(distance).ToPoint3())
}
