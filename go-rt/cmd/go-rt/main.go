package main

import (
	"fmt"
	"os"

	"github.com/wmbat/ray_tracer/internal"
	"github.com/wmbat/ray_tracer/internal/maths"
)

const imageWidth int = 720
const imageHeight int = 480
const aspectRatio float64 = float64(imageWidth) / float64(imageHeight)

func calculatePixel(ray *internal.Ray) internal.Pixel {
	unitDir := maths.UnitVector(&ray.Direction)
	t := 0.5 * (unitDir.Y + 1.0)

	gradient := maths.Vec3{X: 0.5, Y: 0.7, Z: 1.0}.Scale(t)
	colour := maths.Vec3{X: 1.0, Y: 1.0, Z: 1.0}.Scale(1.0 - t).Add(&gradient)

	return internal.Pixel{Red: colour.X, Green: colour.Y, Blue: colour.Z}
}

func main() {
	imageFile := "image.ppm"

	if internal.DoesFileExist(imageFile) {
		os.Remove(imageFile)
	}

	f, err := os.Create("image.ppm")
	if err != nil {
		fmt.Println("Failed to open file:", err)
	}
	defer f.Close()

	image := internal.Image{}
	image.Init(imageWidth, imageHeight)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := maths.Point3{X: 0, Y: 0, Z: 0}
	horizontal := maths.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := maths.Vec3{X: 0, Y: viewportHeight, Z: 0}

	lowerLeftCorner := internal.CalculateLowerLeftCorner(&origin, &horizontal, &vertical, focalLength)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)

			scaledHorizontal := horizontal.Scale(u)
			scaledVertical := vertical.Scale(v)
			rayDir := lowerLeftCorner.Sub(&origin).ToVec3().Add(&scaledHorizontal).Add(&scaledVertical)

			ray := internal.Ray{Origin: origin, Direction: rayDir}

			pixel := calculatePixel(&ray)

			image.WritePixel(i, j, &pixel)
		}
	}

	fmt.Fprintln(f, image)
}
