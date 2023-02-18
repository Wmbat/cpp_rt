package scene

import "github.com/wmbat/ray_tracer/internal/hitable"

type Scene struct {
	Hitables []hitable.Hitable
}
