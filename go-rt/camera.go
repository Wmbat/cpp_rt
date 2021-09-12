package main

import (
	"go_rt/core"
	"go_rt/maths"
)

type Camera struct {
    Origin maths.Vec3
    Horizontal maths.Vec3
    Vertical maths.Vec3
    LowerLeftCorner maths.Vec3 
}

type CameraCreateInfo struct {
    Origin maths.Vec3   
    Height float64
    AspectRatio float64
    FocalLength float64
}

func ComputeLowerLeftCorner(origin *maths.Vec3, horizontal *maths.Vec3, 
    vertical *maths.Vec3, focalLength float64) maths.Vec3 { 

    focal := maths.Vec3{X: 0.0, Y: 0.0, Z: focalLength}
    halfHorizontal := maths.DivScalar(horizontal, 2.0)
    halfVertical := maths.DivScalar(vertical, 2.0)

    result := *origin
    result.Sub(&halfHorizontal)
    result.Sub(&halfVertical)
    result.Sub(&focal)

    return result
}

func NewCamera(info *CameraCreateInfo) Camera {
    horizontal := maths.Vec3{X: info.Height * info.AspectRatio, Y: 0.0, Z: 0.0}
    vertical := maths.Vec3{X: 0.0, Y: info.Height, Z: 0.0}

    return Camera{
        Origin: info.Origin,
        Horizontal: horizontal,
        Vertical: vertical,
        LowerLeftCorner: ComputeLowerLeftCorner(&info.Origin, &horizontal,&vertical, info.FocalLength)}
}

func (camera *Camera) ShootRay(u float64, v float64) core.Ray {
    ur := maths.MultScalar(&camera.Horizontal, u)
    vr := maths.MultScalar(&camera.Vertical, v)

    dir := maths.Add(&camera.LowerLeftCorner, &ur)
    dir.Add(&vr)
    dir.Sub(&camera.Origin)

    return core.Ray{Origin: camera.Origin, Direction: dir}
}
