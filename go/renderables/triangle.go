package renderables

import "go_pt/maths"

type Triangle struct {
   data [3]maths.Vec3
}

func NewTriangle(v0, v1, v2 maths.Vec3) Triangle {
   return Triangle{data: [3]maths.Vec3{v0, v1, v2}}
}

func (tri *Triangle) Vertex(index int32) *maths.Vec3 {
   return &tri.data[index]
}

func (tri *Triangle) U() maths.Vec3 {
   return maths.Sub(tri.Vertex(1), tri.Vertex(0))
}

func (tri *Triangle) V() maths.Vec3 {
   return maths.Sub(tri.Vertex(2), tri.Vertex(0))
}

func (tri *Triangle) Normal() maths.Vec3 {
   u := tri.U()
   v := tri.V()

   return maths.NormaliseCpy(maths.Cross(&u, &v))
}
