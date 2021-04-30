package main

import "go_pt/maths"

type Camera struct {
    origin maths.Vec3
    horizontal maths.Vec3
    vertical maths.Vec3
    lower_left_corner maths.Vec3
    axis maths.OrthoNormalBasis
    lens_radius float64 
}

func (camera *Camera) ShootRay(u float64, v float64) Ray {
    return Ray{}
}
