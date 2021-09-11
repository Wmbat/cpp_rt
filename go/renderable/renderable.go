package renderable

import "go_rt/core"

type RayCollisionResult struct {
    HasValue bool
    Record CollisionRecord
}

type Renderable interface {
    CheckRayCollision(ray *core.Ray, minTime float64, maxTime float64) RayCollisionResult
}
