package math

type Vec3 struct {
   x float32
   y float32
   z float32
}

func Add(lhs, rhs *Vec3) Vec3 {
   return Vec3{x: lhs.x + rhs.x, y: lhs.y + rhs.y, z: lhs.z + rhs.z}
}

func Sub(lhs, rhs *Vec3) Vec3 {
   return Vec3{x: lhs.x - rhs.x, y: lhs.y - rhs.y, z: lhs.z + rhs.z}
}

func Mult(lhs, rhs *Vec3) Vec3 {
   return Vec3{x: lhs.x * rhs.x, y: lhs.y * rhs.y, z: lhs.z * rhs.z}
}

func MultScalar(lhs *Vec3, scalar float32) Vec3 {
   return Vec3{x: lhs.x * scalar, y: lhs.y * scalar, z: lhs.z * scalar}
}

func Div(lhs, rhs *Vec3) Vec3 {
   return Vec3{x: lhs.x / rhs.x, y: lhs.y / rhs.y, z: lhs.z / rhs.z}
}

func DivScalar(lhs *Vec3, scalar float32) Vec3 {
   return Vec3{x: lhs.x / scalar, y: lhs.y / scalar, z: lhs.z / scalar}
}

func (lhs *Vec3) Add(rhs *Vec3) *Vec3 {
   lhs.x += rhs.x
   lhs.y += rhs.y
   lhs.z += rhs.z

   return lhs
}
