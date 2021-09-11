package core

import "go_rt/maths"

type Ray struct {
    Origin maths.Vec3
    Direction maths.Vec3
}

func (ray *Ray) PositionAlong(time float64) maths.Vec3 {
    timeOffset := maths.MultScalar(&ray.Direction, time);

    return maths.Add(&ray.Origin, &timeOffset);
}
