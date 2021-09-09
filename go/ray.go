package main

import "go_pt/maths"

type Ray struct {
    Origin maths.Vec3
    Direction maths.Vec3
}

func (ray *Ray) position_along(time float64) maths.Vec3 {
    timeOffset := maths.MultScalar(&ray.Direction, time);

    return maths.Add(&ray.Origin, &timeOffset);
}
