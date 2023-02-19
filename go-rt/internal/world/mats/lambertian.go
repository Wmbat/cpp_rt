package mats

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Lambertian struct {
	Albedo render.Colour
}

func (this Lambertian) Scatter(ray core.Ray, info SurfaceInfo) (ScatterResult, bool) {
	scatterDir := info.Normal.Add(maths.RandVec3InUnitSphere().Normalize())
	outputRay := core.Ray{Origin: info.Position, Direction: scatterDir}

	if scatterDir.IsNearZero() {
		scatterDir = info.Normal
	}

	return ScatterResult{this.Albedo, outputRay}, true
}
