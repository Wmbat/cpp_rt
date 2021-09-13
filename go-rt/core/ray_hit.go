package core

import "go_rt/maths"

type RayHit struct {
    Position maths.Vec3
    Normal maths.Vec3
    Time float64
    FrontFace bool
}
