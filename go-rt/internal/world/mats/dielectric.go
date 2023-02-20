package mats

import (
	"math"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Dielectric struct {
	Diffuse         render.Colour
	RefractionIndex float64
}

func (this Dielectric) Scatter(info ScatterInfo) (ScatterResult, bool) {
	unitDir := info.Ray.Direction.Normalize()
	refractionRatio := getRefractiveRatio(this.RefractionIndex, info.IsFrontFace)
	cosI := -maths.DotProduct(unitDir, info.Normal)
	sinThetaSq := refractionRatio * refractionRatio * (1.0 - (cosI * cosI))

	// total internal reflection
	if sinThetaSq > 1.0 {
		ray := core.Ray{Origin: info.Position, Direction: unitDir.Reflect(info.Normal)}
		return ScatterResult{Attenuation: this.Diffuse, Ray: ray}, true
	}

	if info.Rng.Float64() < maths.Schlick(cosI, refractionRatio) {
		ray := core.Ray{Origin: info.Position, Direction: unitDir.Reflect(info.Normal)}
		return ScatterResult{Attenuation: this.Diffuse, Ray: ray}, true
	}

	cosTheta := math.Sqrt(1.0 - sinThetaSq)
	refracted := unitDir.Scale(refractionRatio).Add(info.Normal.Scale(refractionRatio*cosI - cosTheta))

	ray := core.Ray{Origin: info.Position, Direction: refracted}
	return ScatterResult{Attenuation: this.Diffuse, Ray: ray}, true
}

func getRefractiveRatio(refractionIndex float64, isFrontFace bool) float64 {
	if isFrontFace {
		return 1.0 / refractionIndex
	} else {
		return refractionIndex
	}
}
