package entt

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/world/mats"
)

type IntersectRecord struct {
	Position maths.Point3
	Normal maths.Vec3
	Distance float64
	IsFrontFace bool
	Material *mats.Material
}
