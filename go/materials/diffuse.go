package materials

import (
	"go_pt/maths"
	"go_pt/tracer"
)
   
type Diffuse struct {
   emissionColour maths.Vec3
   diffuseColour maths.Vec3
}

func NewDiffuseMaterial(emissionColour, diffuseColour *maths.Vec3) Diffuse {
   return Diffuse{emissionColour: *emissionColour, diffuseColour: *diffuseColour}
}

func (mat *Diffuse) Scatter(ray *tracer.Ray, hit *tracer.Hit, u, v float64) MaterialScatterData {
   basis := maths.OrthoNormalBasisFromZ(&hit.Normal)
   scatterDir := maths.HemisphereSample(&basis, u, v)

   return MaterialScatterData{
      Emission: mat.emissionColour, 
      Diffuse: mat.diffuseColour, 
      ScatteredRay: tracer.Ray{Origin: hit.Position, Direction: scatterDir}}
}
