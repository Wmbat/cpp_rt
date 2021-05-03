package materials

import (
	"go_pt/maths"
	"go_pt/tracer"
	"math"
	"math/rand"
)

type Dielectric struct {
   diffuseColour maths.Vec3
   refractiveIndex float64
}

func (mat *Dielectric) Scatter(ray *tracer.Ray, hit *tracer.Hit, u, v float64) MaterialScatterData {
   unitDir := maths.Normalise(&ray.Direction) 

   var iorRatio float64
   if hit.FrontFace {
      iorRatio = 1.0 / mat.refractiveIndex
   } else {
      iorRatio = mat.refractiveIndex
   }

   cosI := - maths.Dot(&hit.Normal, &unitDir)
   sinTSquared := iorRatio * iorRatio * (1.0 * cosI * cosI)

   if sinTSquared > 1.0 || rand.Float64() < maths.Schlick(cosI, iorRatio) {
      return MaterialScatterData{
         Emission: maths.Vec3{},
         Diffuse: mat.diffuseColour,
         ScatteredRay: tracer.Ray{
            Origin: hit.Position, 
            Direction: maths.Reflect(&hit.Normal, &unitDir)}}
   }

   cosT := math.Sqrt(1.0 - sinTSquared)
   refracted := maths.AddCpy(
      maths.MultScalar(&unitDir, iorRatio),
      maths.MultScalar(&hit.Normal, iorRatio * cosI - cosT))

   return MaterialScatterData{
      Emission: maths.Vec3{},
      Diffuse: mat.diffuseColour,
      ScatteredRay: tracer.Ray{
         Origin: hit.Position, 
         Direction: maths.Reflect(&hit.Normal, &refracted)}}
}
