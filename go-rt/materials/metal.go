package materials

import (
	"go_rt/core"
	"go_rt/maths"
)

type Metal struct {
    Diffuse core.Colour

    Roughness float64
}

func (this Metal) Scatter(ray *core.Ray, hit *core.RayHit) ScatterData {
    unitDir := maths.Vec3Normalise(&ray.Direction)
    reflected := maths.Reflect(&hit.Normal, &unitDir)  

    random := maths.RandomVec3InUnitSphere()
    random.MultScalar(this.Roughness)

    scattered := core.Ray{Origin: hit.Position, Direction: maths.Add(&reflected, &random)}

    if maths.Vec3Dot(&scattered.Direction, &hit.Normal) > 0 {
        return ScatterData{
            Emission: core.Colour{},
            Diffuse: this.Diffuse,
            Ray: scattered}
    } else {
        return ScatterData{
            Emission: core.Colour{},
            Diffuse: core.Colour{},
            Ray: scattered}
    }
}
