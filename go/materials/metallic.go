package materials

import (
	"go_pt/maths"
	"go_pt/tracer"
)

type Metallic struct {
   emissionColour maths.Vec3
   diffuseColour maths.Vec3
   roughness float64
}

func NewMetallicMaterial(emissionColour, diffuseColour *maths.Vec3, roughness float64) Metallic {
   return Metallic{
      emissionColour: *emissionColour,
      diffuseColour: *diffuseColour,
      roughness: roughness} 
}

func (mat *Metallic) Scatter(ray *tracer.Ray, hit *tracer.Hit, u, v float64) MaterialScatterData {
   normalisedRayDir := maths.Normalise(&ray.Direction)
   reflectedDir := maths.Reflect(&hit.Normal, &normalisedRayDir)

   randomDir := maths.RandomVec3InUnitSphere()

   scatteredRay := tracer.Ray{
      Origin: hit.Position, 
      Direction: maths.AddCpy(
         reflectedDir, 
         maths.MultScalar(&randomDir, mat.roughness))}

   if maths.Dot(&scatteredRay.Direction, &hit.Normal) > 0 {
      return MaterialScatterData{
         Emission: mat.emissionColour, 
         Diffuse: mat.diffuseColour,
         ScatteredRay: scatteredRay}
   } else {
      return MaterialScatterData{
         Emission: mat.emissionColour, 
         Diffuse: maths.Vec3{},
         ScatteredRay: scatteredRay}
   }
}
