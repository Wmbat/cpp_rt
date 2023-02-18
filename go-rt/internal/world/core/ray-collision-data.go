package core

import "github.com/wmbat/ray_tracer/internal/maths"

type RayCollisionPoint struct {
	Location maths.Point3
	Normal maths.Vec3
	Time float64
	FrontFace bool
}
