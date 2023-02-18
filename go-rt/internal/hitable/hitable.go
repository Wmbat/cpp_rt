package hitable

import (
	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal/core"
	"github.com/wmbat/ray_tracer/internal/utils"
)

type Hitable interface {
	DoesIntersectWith(ray core.Ray, timeBound utils.TimeBoundaries) mo.Option[HitRecord]
}
