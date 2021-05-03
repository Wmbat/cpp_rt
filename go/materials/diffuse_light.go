package materials

import (
	"go_pt/maths"
	"go_pt/tracer"
)

type DiffuseLight struct {
   emissionColour maths.Vec3
}

func NewDiffuseLightMaterial(emissionColour *maths.Vec3) DiffuseLight {
   return DiffuseLight{emissionColour: *emissionColour}
}

func (mat *DiffuseLight) Scatter(ray *tracer.Ray, hit *tracer.Hit, u, v float64) MaterialScatterData {
   basis := maths.OrthoNormalBasisFromZ(&hit.Normal)
   scatterDir := maths.HemisphereSample(&basis, u, v)

   return MaterialScatterData{
      Emission: mat.emissionColour, 
      Diffuse: maths.Vec3{}, 
      ScatteredRay: tracer.Ray{Origin: hit.Position, Direction: scatterDir}}
}
