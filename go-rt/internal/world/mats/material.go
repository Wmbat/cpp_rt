package mats

import (
	"math/rand"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

// Represent the resultant of a ray scatter after encountering an entity's material properties
type ScatterResult struct {
	Attenuation render.Colour // The colour attenuation produced by the scatter
	Ray         core.Ray      // The output ray produced by the material scatter
}

// Information about where the ray hit
type ScatterInfo struct {
	Ray         core.Ray     // The ray that hit the entity
	Position    maths.Point3 // The position where the material surface was hit
	Normal      maths.Vec3   // The normal at the position on thu material surface
	IsFrontFace bool
	Rng         *rand.Rand // The random number generation engine used for random numbers
}

// Represents the abstract concept of a material which defines the way a ray should bounce off
// and scatter when intersecting with an entity
type Material interface {
	Scatter(info ScatterInfo) (ScatterResult, bool)
}
