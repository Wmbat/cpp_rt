package tracer

import "go_pt/maths"

type Hit struct {
   Position maths.Vec3
   Normal maths.Vec3
   Distance float64
   FrontFace bool
}
