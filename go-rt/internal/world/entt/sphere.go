package entt

import (
	"math"

	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/utils"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Sphere struct {
	Origin maths.Point3
	Radius float64
}

func (this Sphere) IsIntersectedByRay(ray core.Ray, timeBounds utils.TimeBoundaries) mo.Option[core.RayCollisionPoint] {
	a, b, c := this.GetQuadraticFactor(ray)

	discriminant := (b * b) - (a * c)
	if discriminant < 0 {
		return mo.None[core.RayCollisionPoint]()
	}

	time, isPresent := findNearestIntersectTime(a, b, discriminant, timeBounds).Get()
	if !isPresent {
		return mo.None[core.RayCollisionPoint]()
	}

	location := ray.At(time)
	normal := location.Sub(this.Origin).Scale(1 / this.Radius).ToVec3()

	if maths.DotProduct(ray.Direction, normal) > 0.0 {
		return mo.Some(core.RayCollisionPoint{
			Location:  location,
			Normal:    normal,
			Time:      time,
			FrontFace: true})
	} else {
		return mo.Some(core.RayCollisionPoint{
			Location:  location,
			Normal:    normal.Negate(),
			Time:      time,
			FrontFace: false})
	}
}

func (this Sphere) GetQuadraticFactor(ray core.Ray) (float64, float64, float64) {
	oc := ray.Origin.Sub(this.Origin).ToVec3()

	a := ray.Direction.LengthSquared()
	b := maths.DotProduct(oc, ray.Direction)
	c := oc.LengthSquared() - (this.Radius * this.Radius)

	return a, b, c
}

func findNearestIntersectTime(a, b, discriminant float64, timeBounds utils.TimeBoundaries) mo.Option[float64] {
	determinant := math.Sqrt(discriminant)

	intersectTimeOne := (-b - determinant) / a
	if timeBounds.IsTimeWithinBounds(intersectTimeOne) {
		return mo.Some(intersectTimeOne)
	}

	intersectTimeTwo := (-b + determinant) / a
	if timeBounds.IsTimeWithinBounds(intersectTimeTwo) {
		return mo.Some(intersectTimeTwo)
	}

	return mo.None[float64]()
}
