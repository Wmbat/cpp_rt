package entities

import (
	"go_rt/core"
	"go_rt/materials"
)

type RayHitRecord struct {
    Hit core.RayHit
    Mat materials.Material
}

type RayHitResult struct {
    HasValue bool   
    Record RayHitRecord
}

type Entity interface {
    CheckRayHit(ray *core.Ray, minTime float64, maxTime float64) RayHitResult
}
