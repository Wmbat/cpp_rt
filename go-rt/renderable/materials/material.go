package materials

import (
	"go_rt/core"
	"go_rt/maths"
	"go_rt/renderable"
)

type ScatterData struct {
    Emission maths.Vec3
    Diffuse maths.Vec3

    Ray core.Ray
}

type Material interface {
    Scatter(ray *core.Ray, hit *renderable.CollisionRecord, u float64, v float64) ScatterData
}
