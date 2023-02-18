package scene

import (
	"github.com/wmbat/ray_tracer/internal/core"
	"github.com/wmbat/ray_tracer/internal/maths"
)

type Camera struct {
	Origin          maths.Point3
	Horizontal      maths.Vec3
	Vertical        maths.Vec3
	LowerLeftCorner maths.Point3
	FocalLength     float64
}

func NewCamera(origin maths.Point3, viewport maths.Size2f, focalLength float64) Camera {
	horizontal := maths.Vec3{X: viewport.Width, Y: 0, Z: 0}
	vertical := maths.Vec3{X: 0, Y: viewport.Height, Z: 0}

	horizontalMidpoint := horizontal.Scale(0.5).ToPoint3()
	verticalMidpoint := vertical.Scale(0.5).ToPoint3()
	depth := maths.Point3{X: 0, Y: 0, Z: focalLength}

	lowerLeftCorner := origin.Sub(horizontalMidpoint).Sub(verticalMidpoint).Sub(depth)

	return Camera{
		Origin:          origin,
		Horizontal:      horizontal,
		Vertical:        vertical,
		LowerLeftCorner: lowerLeftCorner,
		FocalLength:     focalLength}
}

func (this Camera) ShootRay(camTarget maths.Point2) core.Ray {
	scaledHorizontal := this.Horizontal.Scale(camTarget.X)
	scaledVertical := this.Vertical.Scale(camTarget.Y)

	originVec := this.Origin.ToVec3()
	cornerVec := this.LowerLeftCorner.ToVec3()

	return core.Ray{
		Origin:    this.Origin,
		Direction: cornerVec.Add(scaledHorizontal).Add(scaledVertical).Sub(originVec)}
}
