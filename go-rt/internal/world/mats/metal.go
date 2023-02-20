package mats

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Metal struct {
	Albedo render.Colour
}

func (this Metal) Scatter(info ScatterInfo) (ScatterResult, bool) {
	outputRay := core.Ray{
		Origin: info.Position, 
		Direction: info.Ray.Direction.Normalize().Reflect(info.Normal)}

	isReflected := maths.DotProduct(outputRay.Direction, info.Normal) > 0

	return ScatterResult{this.Albedo, outputRay}, isReflected
}
