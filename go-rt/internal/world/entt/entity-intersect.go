package entt

import (
	"github.com/wmbat/ray_tracer/internal/world/core"
	"github.com/wmbat/ray_tracer/internal/world/mats"
)

type EntityIntesect struct {
	RayCollision core.RayCollisionPoint
	Material mats.Material
}
