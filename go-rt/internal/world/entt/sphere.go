package entt

import (
	"math"

	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Sphere struct {
	Origin maths.Point3
	Radius float64
}

func (this Sphere) IsIntersectedByRay(ray core.Ray, nearestDistance float64) mo.Option[core.RayCollisionPoint] {
	a, b, c := this.GetQuadraticFactor(ray)

	discriminant := (b * b) - (a * c)
	if discriminant < 0 {
		return mo.None[core.RayCollisionPoint]()
	}

	distance, isPresent := findNearestIntersectTime(a, b, discriminant, nearestDistance).Get()
	if !isPresent {
		return mo.None[core.RayCollisionPoint]()
	}

	location := ray.At(distance)
	normal := location.Sub(this.Origin).Scale(1 / this.Radius).ToVec3()
	isFontFace := maths.DotProduct(ray.Direction, normal) > 0.0

	return mo.Some(core.RayCollisionPoint{
		Location:    location,
		Normal:      GetFrontFaceNormal(normal, isFontFace),
		Distance:    distance,
		IsFrontFace: isFontFace})
}

func (this Sphere) GetQuadraticFactor(ray core.Ray) (float64, float64, float64) {
	oc := ray.Origin.Sub(this.Origin).ToVec3()

	a := ray.Direction.LengthSquared()
	b := maths.DotProduct(oc, ray.Direction)
	c := oc.LengthSquared() - (this.Radius * this.Radius)

	return a, b, c
}

func GetFrontFaceNormal(normal maths.Vec3, isFrontFace bool) maths.Vec3 {
	if isFrontFace {
		return normal
	} else {
		return normal.Negate()
	}
}

func findNearestIntersectTime(a, b, discriminant float64, nearestDistance float64) mo.Option[float64] {
	determinant := math.Sqrt(discriminant)

	minT := (-b - determinant) / a
	if 0.001 <= minT && minT <= nearestDistance {
		return mo.Some(minT)
	}

	maxT := (-b + determinant) / a
	if 0.001 <= maxT && maxT <= nearestDistance {
		return mo.Some(maxT)
	}

	return mo.None[float64]()
}
