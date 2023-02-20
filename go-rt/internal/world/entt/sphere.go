package entt

import (
	"math"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/world/core"
	"github.com/wmbat/ray_tracer/internal/world/mats"
)

type Sphere struct {
	Position maths.Point3 // The position fo the sphere in the world
	Radius   float64      // The radius of the sphere
	Material mats.Material
}

func (this Sphere) IsIntersectedByRay(ray core.Ray, closerThan float64) (IntersectRecord, bool) {
	const epsilon float64 = 0.001

	a, b, c := this.GetQuadraticFactor(ray)

	discriminant := (b * b) - (a * c)
	if discriminant < 0 {
		return IntersectRecord{}, false
	}

	determinant := math.Sqrt(discriminant)

	distance := (-b - determinant) / a
	if distance < epsilon || closerThan < distance {
		distance = (-b + determinant) / a
		if distance < epsilon || closerThan < distance {
			return IntersectRecord{}, false
		}
	}

	location := ray.At(distance)
	normal := location.Sub(this.Position).Scale(1 / this.Radius).ToVec3()
	isFrontFace := maths.DotProduct(ray.Direction, normal) < 0.0

	record := IntersectRecord{
		Position:    location,
		Normal:      GetFrontFaceNormal(normal, isFrontFace),
		Distance:    distance,
		IsFrontFace: isFrontFace,
		Material:    this.Material}

	return record, true
}

func (this Sphere) GetQuadraticFactor(ray core.Ray) (float64, float64, float64) {
	oc := ray.Origin.Sub(this.Position).ToVec3()

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
