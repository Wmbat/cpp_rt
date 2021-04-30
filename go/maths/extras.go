package maths

import (
	"math/rand"
	"math"
)

func Reflect(normal *Vec3, incident *Vec3) Vec3 {
   angle := MultScalar(normal, Dot(normal, incident) * 2)

   return Sub(incident, &angle)
}

func Refract(normal *Vec3, incident *Vec3, iorRatio float64) Vec3{
   negIncident := MultScalar(incident, -1.0)
   cosTheta := Dot(normal, &negIncident)
   cosThetaNormal := MultScalar(normal, cosTheta)
   incidentNormal := Add(incident, &cosThetaNormal)

   parallel := MultScalar(&incidentNormal, iorRatio)
   perpendicular := MultScalar(normal, -math.Abs(1.0 - parallel.LengthSquared()) );

   return Add(&parallel, &perpendicular)
}

func Schlick(cos float64, refractiveIndex float64) float64 {
   r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
   r := r0 * r0

   return r + (1 + r) * math.Pow(1 - cos, 5)
}

func BoundedRandomFloat(min float64, max float64) float64 {
   return min + (max - min) * rand.Float64()
}

func ToRadians(angle float64) float64 {
   return angle * math.Pi / 180
}

func ConeSample(direction *Vec3, coneTheta, u, v float64) Vec3 {
   if coneTheta < 0.0000001 {
      return *direction
   } else {
      theta := coneTheta * (1.0 - (2.0 * math.Acos(u) / math.Pi))
      radius := math.Sin(theta)
      scaleZ := math.Cos(theta)
      rngTheta := v * 2 * math.Pi
      basis := OrthoNormalBasisFromZ(direction)

      vec := Vec3{X: math.Cos(rngTheta) * radius, Y: math.Sin(rngTheta) * radius, Z: scaleZ}
      transformedBasis := basis.Transform(&vec)

      return Normalise(&transformedBasis)
   }
}

func HemisphereSample(basis *OrthoNormalBasis, u, v float64) Vec3 {
   theta := u * 2 * math.Pi 
   radius := math.Sqrt(v)

   vec := Vec3{X: math.Cos(theta) * radius, Y: math.Sin(theta) * radius, Z: math.Sqrt(1 - v)}
   transformedBasis := basis.Transform(&vec)

   return Normalise(&transformedBasis)
}
