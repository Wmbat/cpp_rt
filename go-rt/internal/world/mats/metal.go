package mats

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Metal struct {
	Albedo    render.Colour
	Roughness float64
}

func (this Metal) Scatter(info ScatterInfo) (ScatterResult, bool) {
	reflected := info.Ray.Direction.Normalize().Reflect(info.Normal)

	outputRay := core.Ray{
		Origin:    info.Position,
		Direction: reflected.Add(maths.RandVec3InUnitSphere(info.Rng).Scale(this.Roughness))}

	isReflected := maths.DotProduct(outputRay.Direction, info.Normal) > 0

	return ScatterResult{this.Albedo, outputRay}, isReflected
}
