package maths

import (
	"fmt"
	"math"
	"math/rand"
)

type Vec3 struct {
    X float64
    Y float64
    Z float64
}

func (lhs *Vec3) Add(rhs *Vec3) *Vec3 {
   lhs.X += rhs.X
   lhs.Y += rhs.Y
   lhs.Z += rhs.Z

   return lhs
}

func (lhs *Vec3) AddScalar(scalar float64) *Vec3 {
   lhs.X += scalar 
   lhs.Y += scalar 
   lhs.Z += scalar 

   return lhs
}

func (lhs *Vec3) Sub(rhs *Vec3) *Vec3 {
   lhs.X -= rhs.X
   lhs.Y -= rhs.Y
   lhs.Z -= rhs.Z

   return lhs 
}

func (lhs *Vec3) Mult(rhs *Vec3) *Vec3 {
   lhs.X *= rhs.X
   lhs.Y *= rhs.Y
   lhs.Z *= rhs.Z

   return lhs  
}

func (lhs *Vec3) MultScalar(scalar float64) *Vec3 {
   lhs.X *= scalar
   lhs.Y *= scalar
   lhs.Z *= scalar

   return lhs  
}

func (lhs *Vec3) DivScalar(scalar float64) *Vec3 {
   return lhs.MultScalar(1 / scalar)
}

func (lhs *Vec3) LengthSquared() float64 {
    return lhs.X * lhs.X + lhs.Y * lhs.Y + lhs.Z * lhs.Z
}

func (lhs *Vec3) Length() float64 {
    return math.Sqrt(lhs.LengthSquared())
}

func (lhs *Vec3) Normalise() *Vec3 {
    return lhs.DivScalar(lhs.Length()) 
}

func (lhs *Vec3) String() string {
    return fmt.Sprintf("%f %f %f", lhs.X, lhs.Y, lhs.Z)
}

func (this Vec3) IsNearZero() bool {
    epsilon := 1e-8
    return (math.Abs(this.X) < epsilon) &&
        (math.Abs(this.Y) < epsilon) &&
        (math.Abs(this.Z) < epsilon)
}

func Add(lhs *Vec3, rhs *Vec3) Vec3 {
    return Vec3{X: lhs.X + rhs.X, Y: lhs.Y + rhs.Y, Z: lhs.Z + rhs.Z}
}

func Sub(lhs *Vec3, rhs *Vec3) Vec3 {
    return Vec3{X: lhs.X - rhs.X, Y: lhs.Y - rhs.Y, Z: lhs.Z - rhs.Z}
}

func MultScalar(lhs *Vec3, rhs float64) Vec3 {
    return Vec3{X: lhs.X * rhs, Y: lhs.Y * rhs, Z: lhs.Z * rhs}
}

func DivScalar(lhs *Vec3, rhs float64) Vec3 {
    return MultScalar(lhs, 1 / rhs)
}

func Vec3Normalise(vec *Vec3) Vec3 {
    return DivScalar(vec, vec.Length())
}

func Vec3Dot(lhs *Vec3, rhs *Vec3) float64 {
    return lhs.X * rhs.X + lhs.Y * rhs.Y + lhs.Z * rhs.Z
}

func Vec3Cross(lhs *Vec3, rhs *Vec3) Vec3 {
    return Vec3{
        X: lhs.Y * rhs.Z - lhs.Z * rhs.Y,
        Y: lhs.Z * rhs.X - lhs.X * rhs.Z,
        Z: lhs.X * rhs.Y - lhs.Y * rhs.X} 
}    

func RandomVec3() Vec3 {
    return Vec3{X: rand.Float64(), Y: rand.Float64(), Z: rand.Float64()}
}

func RandomBoundedVec3(min float64, max float64) Vec3 {
    return Vec3{
        X: RandomFloat64(min, max),
        Y: RandomFloat64(min, max),
        Z: RandomFloat64(min, max)}
}

func RandomVec3InUnitSphere() Vec3 {
     for {
        vec := RandomBoundedVec3(-1.0, 1.0)

        if vec.LengthSquared() < 1.0 {
            return vec 
        }
     }
}

func RandomNormalizedVec3() Vec3 {
    vec := RandomVec3InUnitSphere();
    vec.Normalise()

    return vec;
}

func RandomVec3InHemisphere(normal *Vec3) Vec3 {
    normalized := RandomNormalizedVec3()
    if Vec3Dot(&normalized, normal) > 0.0 {
        return normalized
    } else {
        return MultScalar(&normalized, -1.0)
    }
}

