package hitable

import (
	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal"
)

type Hitable interface {
	DoesIntersectWith(ray core.Ray, timeBound TimeBoundaries) mo.Option[HitRecord]
}
