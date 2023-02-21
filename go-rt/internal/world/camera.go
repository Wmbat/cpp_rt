package world

import (
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/world/core"
)

type Camera struct {
	Origin          maths.Point3
	Horizontal      maths.Vec3
	Vertical        maths.Vec3
	LowerLeftCorner maths.Point3
	FocalLength     float64
}

/*
func Test(lookFrom, lookAt maths.Point3, up maths.Vec3, fov maths.Degree, aspectRatio float64) Camera {
	theta := maths.DegreesToRadians(fov)
	h := math.Tan(float64(theta) / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.Sub(lookAt)

}
*/

func NewCamera(origin maths.Point3, viewport maths.Size2[float32], focalLength float64) Camera {
	horizontal := maths.Vec3{X: float64(viewport.Width), Y: 0, Z: 0}
	vertical := maths.Vec3{X: 0, Y: float64(viewport.Height), Z: 0}

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

func (this Camera) ShootRay(camTarget maths.Point2[float64]) core.Ray {
	scaledHorizontal := this.Horizontal.Scale(camTarget.X)
	scaledVertical := this.Vertical.Scale(camTarget.Y)

	originVec := this.Origin.ToVec3()
	cornerVec := this.LowerLeftCorner.ToVec3()

	return core.Ray{
		Origin:    this.Origin,
		Direction: cornerVec.Add(scaledHorizontal).Add(scaledVertical).Sub(originVec)}
}
