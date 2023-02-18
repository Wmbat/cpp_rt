package hitable

import (
	"github.com/samber/mo"
	"github.com/wmbat/ray_tracer/internal"
)

type Hitable interface {
	DoesIntersectWith(ray rt.Ray, timeBound rt.TimeBoundaries) mo.Option[HitRecord]
}
