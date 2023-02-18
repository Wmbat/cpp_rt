package entt

import (
	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal/utils"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Entity interface {
	IsIntersectedByRay(ray core.Ray, timeBound utils.TimeBoundaries) mo.Option[core.RayCollisionPoint]
}
