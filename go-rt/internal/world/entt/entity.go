package entt

import (
	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Entity interface {
	IsIntersectedByRay(ray core.Ray, closerThan float64) mo.Option[core.RayCollisionPoint]
}
