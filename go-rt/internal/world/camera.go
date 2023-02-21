package world

import (
	"math"
	"math/rand"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Camera struct {
	origin          maths.Point3
	horizontal      maths.Vec3
	vertical        maths.Vec3
	lowerLeftCorner maths.Point3
	u, v, w         maths.Vec3
	lensRadius      float64
}

type CameraCreateInfo struct {
	LookFrom, LookAt maths.Point3
	Up               maths.Vec3
	Fov              maths.Degree
	AspectRatio      float64
	Aperture         float64
	FocusDistance    float64
}

func NewCamera(info CameraCreateInfo) Camera {
	theta := maths.DegreesToRadians(info.Fov)
	h := math.Tan(float64(theta) / 2)
	viewportHeight := 2.0 * h
	viewportWidth := info.AspectRatio * viewportHeight

	w := info.LookFrom.Sub(info.LookAt).ToVec3().Normalize()
	u := maths.CrossProduct(info.Up, w).Normalize()
	v := maths.CrossProduct(w, u)

	horizontal := u.Scale(viewportWidth * info.FocusDistance)
	vertical := v.Scale(viewportHeight * info.FocusDistance)

	horizontalMidpoint := horizontal.Scale(0.5).ToPoint3()
	verticalMidpoint := vertical.Scale(0.5).ToPoint3()
	focusedW := w.Scale(info.FocusDistance).ToPoint3()

	lowerLeftCorner := info.LookFrom.Sub(horizontalMidpoint).Sub(verticalMidpoint).Sub(focusedW)
	lensRadius := info.Aperture / 2

	return Camera{
		origin:          info.LookFrom,
		horizontal:      horizontal,
		vertical:        vertical,
		lowerLeftCorner: lowerLeftCorner,
		u:               u,
		v:               v,
		w:               w,
		lensRadius:      lensRadius}
}

func (this Camera) ShootRay(camTarget maths.Point2[float64], rng *rand.Rand) core.Ray {
	rd := maths.RandVec3InUnitDisk(rng).Scale(this.lensRadius)
	offset := this.u.Scale(rd.X).Add(this.v.Scale(rd.Y))

	scaledHorizontal := this.horizontal.Scale(camTarget.X)
	scaledVertical := this.vertical.Scale(camTarget.Y)

	originVec := this.origin.ToVec3()
	cornerVec := this.lowerLeftCorner.ToVec3()

	return core.Ray{
		Origin:    this.origin.Add(offset.ToPoint3()),
		Direction: cornerVec.Add(scaledHorizontal).Add(scaledVertical).Sub(originVec).Sub(offset)}
}
