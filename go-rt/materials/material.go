package materials

import (
	"go_rt/core"
)

type ScatterData struct {
    Emission core.Colour
    Diffuse core.Colour

    Ray core.Ray
}

type Material interface {
    Scatter(ray *core.Ray, hit *core.RayHit) ScatterData
}
