package mats

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type ScatterResult struct {
	Attenuation render.Colour // The colour attenuation produced by the scatter
	Ray         core.Ray      // The output ray produced by the material scatter
}

// Information about where the ray hit
type SurfaceInfo struct {
	Position maths.Point3 // The position where the material surface was hit
	Normal   maths.Vec3   // The normal at the position on thu material surface
}

type Material interface {
	Scatter(ray core.Ray, info SurfaceInfo) (ScatterResult, bool)
}
