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

func (this Sphere) IsIntersectedByRay(ray core.Ray, closerThan float64) mo.Option[core.RayCollisionPoint] {
	const epsilon float64 = 0.001

	a, b, c := this.GetQuadraticFactor(ray)

	discriminant := (b * b) - (a * c)
	if discriminant < 0 {
		return mo.None[core.RayCollisionPoint]()
	}

	determinant := math.Sqrt(discriminant)

	distance := (-b - determinant) / a
	if distance < epsilon || closerThan < distance {
		distance = (-b + determinant) / a
		if distance < epsilon || closerThan < distance {
			return mo.None[core.RayCollisionPoint]()
		}
	}

	location := ray.At(distance)
	normal := location.Sub(this.Origin).Scale(1 / this.Radius).ToVec3()
	isFrontFace := maths.DotProduct(ray.Direction, normal) > 0.0

	return mo.Some(core.RayCollisionPoint{
		Location:    location,
		Normal:      GetFrontFaceNormal(normal, isFrontFace),
		Distance:    distance,
		IsFrontFace: isFrontFace})
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
