package hitable

import (
	"math"

	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal"
	"github.com/wmbat/ray_tracer/internal/maths"
)

type Sphere struct {
	Origin maths.Point3
	Radius float64
}

func (this Sphere) DoesIntersectWith(ray core.Ray, timeBounds TimeBoundaries) mo.Option[HitRecord] {
	oc := ray.Origin.Sub(this.Origin).ToVec3()

	quadEquation := maths.QuadraticEquation{
		A:     ray.Direction.LengthSquared(),
		HalfB: maths.DotProduct(oc, ray.Direction),
		C:     oc.LengthSquared() - (this.Radius * this.Radius)}

	discriminant := quadEquation.ComputeDiscriminant()
	if discriminant < 0 {
		return mo.None[HitRecord]()
	}

	time, isPresent := findNearestIntersectTime(quadEquation, timeBounds).Get()
	if !isPresent {
		return mo.None[HitRecord]()
	}

	location := ray.At(time)
	normal := location.Sub(this.Origin).Scale(1 / this.Radius).ToVec3()

	if maths.DotProduct(ray.Direction, normal) > 0.0 {
		return mo.Some(HitRecord{
			Location:  location,
			Normal:    normal.Negate(),
			Time:      time,
			FrontFace: false})
	} else {
		return mo.Some(HitRecord{
			Location:  location,
			Normal:    normal,
			Time:      time,
			FrontFace: true})
	}
}

func findNearestIntersectTime(quadEq maths.QuadraticEquation, timeBounds TimeBoundaries) mo.Option[float64] {
	sqrtD := math.Sqrt(quadEq.ComputeDiscriminant())

	intersectTimeOne := (-quadEq.HalfB - sqrtD) / quadEq.A
	if timeBounds.IsTimeWithinBounds(intersectTimeOne) {
		return mo.Some(intersectTimeOne)
	}

	intersectTimeTwo := (-quadEq.HalfB + sqrtD) / quadEq.A
	if timeBounds.IsTimeWithinBounds(intersectTimeTwo) {
		return mo.Some(intersectTimeTwo)
	}

	return mo.None[float64]()
}
