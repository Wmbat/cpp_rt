package main

import (
	"fmt"
	"os"

	core "github.com/wmbat/ray_tracer/internal"
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/utils"
)

const imageWidth int64 = 720
const imageHeight int64 = 480
const aspectRatio float64 = float64(imageWidth) / float64(imageHeight)

func calculatePixel(ray core.Ray) core.Pixel {
	t := 0.5 * (ray.Direction.Normalize().Y + 1.0)

	white := core.Colour{Red: 1.0, Green: 1.0, Blue: 1.0}
	gradient := core.Colour{Red: 0.5, Green: 0.7, Blue: 1.0}

	blendedColour := white.Scale(1.0 - t).Add(gradient.Scale(t))

	return core.Pixel{Colour: blendedColour, SampleCount: 1}
}

func main() {
	imageFile := "image.ppm"

	if utils.DoesFileExist(imageFile) {
		os.Remove(imageFile)
	}

	f, err := os.Create("image.ppm")
	if err != nil {
		fmt.Println("Failed to open file:", err)
	}
	defer f.Close()

	image := core.NewImage(imageWidth, imageHeight)

	fmt.Printf("Creating an image of size <%d, %d>\n", image.Width, image.Height)

	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := maths.Point3{X: 0, Y: 0, Z: 0}
	horizontal := maths.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := maths.Vec3{X: 0, Y: viewportHeight, Z: 0}

	horizontalMidpoint := horizontal.Scale(1 / 2).ToPoint3()
	verticalMidpoint := vertical.Scale(1 / 2).ToPoint3()
	depth := maths.Point3{X: 0, Y: 0, Z: focalLength}

	lowerLeftCorner := origin.Sub(horizontalMidpoint).Sub(verticalMidpoint).Sub(depth)

	for j := image.Height - 1; j >= 0; j-- {
		for i := int64(0); i < image.Width; i++ {
			u := float64(i) / float64(image.Width-1)
			v := float64(j) / float64(image.Height-1)

			scaledHorizontal := horizontal.Scale(u)
			scaledVertical := vertical.Scale(v)
			rayDir := lowerLeftCorner.Sub(origin).ToVec3().Add(scaledHorizontal).Add(scaledVertical)

			ray := core.Ray{Origin: origin, Direction: rayDir}

			image.WritePixel(i, j, calculatePixel(ray))
		}
	}

	image.SaveToFile(f)
}
