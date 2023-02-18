package mats

import (
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Material interface {
	Scatter(ray core.Ray, intersect core.RayCollisionPoint) bool
}
