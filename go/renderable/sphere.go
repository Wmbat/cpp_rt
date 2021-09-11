package renderable

import (
	"go_rt/core"
	"go_rt/maths"
	"math"
)

type Sphere struct {
    Center maths.Vec3
    Radius float64
}

func (s Sphere) CheckRayCollision(ray *core.Ray, minTime float64, maxTime float64) RayCollisionResult {
    lineToSphere :=  maths.Sub(&ray.Origin, &s.Center)

    // We are computing the discriminant of the quadratic formula here
    // giving us the closest point of interection between the ray and the sphere

    a := ray.Direction.LengthSquared()
    halfB := maths.Vec3Dot(&lineToSphere, &ray.Direction)
    c := lineToSphere.LengthSquared() - s.Radius * s.Radius

    discriminant := halfB * halfB - a * c

    if discriminant < 0 {
        return RayCollisionResult{HasValue: false}
    }

    sqrtDiscriminant := math.Sqrt(discriminant)

    // Check the roots to figure out which is closer. It's hidden
    // behind if statements to reduce compute times by avoiding
    // unnecessarry computations
    root := (-halfB - sqrtDiscriminant) / a
    if root < minTime || maxTime < root {
        root = (-halfB + sqrtDiscriminant) / a
        if  root < minTime || maxTime < root {
            return RayCollisionResult{HasValue: false}    
        }
    }

    rayAlongT := ray.PositionAlong(root)
    fromCenter := maths.Sub(&rayAlongT, &s.Center)
    outwardNormal := maths.DivScalar(&fromCenter, s.Radius)

    frontFace := maths.Vec3Dot(&ray.Direction, &outwardNormal) < 0.0
    normal := maths.Vec3{}
    if frontFace {
        normal = outwardNormal
    } else {
        normal = maths.MultScalar(&outwardNormal, -1.0)
    }

    return RayCollisionResult{
        HasValue: true, 
        Record: CollisionRecord{
            Position: rayAlongT,
            Normal: normal,
            Time: root,
            FrontFace: frontFace}}
}
