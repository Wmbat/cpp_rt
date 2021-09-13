package materials

import (
	"go_rt/core"
	"go_rt/maths"
)

type Lambertian struct {
    Diffuse core.Colour 
}

func (this Lambertian) Scatter(ray *core.Ray, hit *core.RayHit) ScatterData {
    random := maths.RandomNormalizedVec3()
    scatterDir := maths.Add(&hit.Normal, &random)

    if scatterDir.IsNearZero() {
        scatterDir = hit.Normal
    }

    return ScatterData{
        Emission: core.Colour{}, 
        Diffuse: this.Diffuse, 
        Ray: core.Ray{Origin: hit.Position ,Direction: scatterDir}}
}
