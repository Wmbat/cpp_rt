package renderable

import "go_rt/maths"

type CollisionRecord struct {
    Position maths.Vec3
    Normal maths.Vec3
    Time float64
    FrontFace bool
}
