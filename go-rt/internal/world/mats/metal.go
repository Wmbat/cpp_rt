package mats

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Metal struct {
	Albedo render.Colour
}

func (this Metal) Scatter(ray core.Ray, info SurfaceInfo) (ScatterResult, bool) {
	outputRay := core.Ray{
		Origin: info.Position, 
		Direction: ray.Direction.Normalize().Reflect(info.Normal)}

	isReflected := maths.DotProduct(outputRay.Direction, info.Normal) > 0

	return ScatterResult{this.Albedo, outputRay}, isReflected
}
