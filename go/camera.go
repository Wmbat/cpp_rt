package main

import (
	"go_pt/maths"
	"math"
)

type Camera struct {
    origin maths.Vec3
    horizontal maths.Vec3
    vertical maths.Vec3
    lowerLeftCorner maths.Vec3
    axis maths.OrthoNormalBasis
    lensRadius float64 
}

type CameraCreateInfo struct {
    Eye maths.Vec3
    LookAt maths.Vec3
    Up maths.Vec3 
    VerticalFOV float64
    AspectRatio float64
    Aperture float64
    FocusDistance float64
}

func NewCamera(info *CameraCreateInfo) Camera {
    theta := maths.ToRadians(info.VerticalFOV)
    halfHeight := math.Tan(theta / 2.0)
    halfWidth := info.AspectRatio * halfHeight

    origin := info.Eye
    lensRadius := info.Aperture / 2.0

    axisZ := maths.NormaliseCpy(maths.Sub(&origin, &info.LookAt))
    axisX := maths.NormaliseCpy(maths.Cross(&info.Up, &axisZ))
    axisY := maths.Cross(&axisZ, &axisX)

    axis := maths.OrthoNormalBasis{X: axisX, Y: axisY, Z: axisZ}

    lowerLeftCorner := maths.SubCpy(
        maths.SubCpy(
            maths.SubCpy(
                origin, 
                maths.MultScalar(&axis.X, halfWidth * info.FocusDistance)),
            maths.MultScalar(&axis.Y, halfHeight * info.FocusDistance)),
        maths.MultScalar(&axis.Z, info.FocusDistance))

    horizontal := maths.MultScalar(&axis.X, 2.0 * halfHeight * info.FocusDistance)
    vertical := maths.MultScalar(&axis.Y, 2.0 * halfWidth * info.FocusDistance)

    return Camera{
        origin: origin, 
        horizontal: horizontal, 
        vertical: vertical, 
        lowerLeftCorner: lowerLeftCorner, 
        axis: axis, 
        lensRadius: lensRadius}
}

func (camera *Camera) ShootRay(u, v float64) Ray {
    unitDiskVec := maths.RandomVec3InUnitDisk()
    randomDisk := maths.MultScalar(&unitDiskVec, camera.lensRadius)
    offset := maths.AddCpy(
        maths.MultScalar(&camera.axis.X , randomDisk.X), 
        maths.MultScalar(&camera.axis.Y , randomDisk.Y))

    direction := maths.SubCpy(
        maths.AddCpy(
            maths.AddCpy(
                camera.lowerLeftCorner, 
                maths.MultScalar(&camera.horizontal, u)),
            maths.MultScalar(&camera.vertical, v)),
        maths.Sub(&camera.origin, &offset))


    return Ray{Origin: maths.Add(&camera.origin, &offset), Direction: direction}
}
