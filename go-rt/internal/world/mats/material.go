package mats

import (
	"math/rand"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type ScatterResult struct {
	Attenuation render.Colour // The colour attenuation produced by the scatter
	Ray         core.Ray      // The output ray produced by the material scatter
}

// Information about where the ray hit
type ScatterInfo struct {
	Ray      core.Ray
	Position maths.Point3 // The position where the material surface was hit
	Normal   maths.Vec3   // The normal at the position on thu material surface
	Rng      *rand.Rand
}

type Material interface {
	Scatter(info ScatterInfo) (ScatterResult, bool)
}
