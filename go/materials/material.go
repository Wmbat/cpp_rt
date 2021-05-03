package materials

import (
	"go_pt/maths"
	"go_pt/tracer"
)

type MaterialScatterData struct {
   Emission maths.Vec3
   Diffuse maths.Vec3
   ScatteredRay tracer.Ray
}

type Material interface {
    Scatter(ray *tracer.Ray, hit *tracer.Hit, u, v float64) MaterialScatterData
}
