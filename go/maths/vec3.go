package maths

import (
	"fmt"
	"math"
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

func Normalise(vec *Vec3) Vec3 {
    return DivScalar(vec, vec.Length())
}
