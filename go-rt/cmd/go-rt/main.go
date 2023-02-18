package main

import (
	"fmt"
	"math"
	"os"

	"github.com/wmbat/ray_tracer/internal"
	"github.com/wmbat/ray_tracer/internal/hitable"
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/utils"
)

const imageWidth int64 = 1080
const imageHeight int64 = 720
const aspectRatio float64 = float64(imageWidth) / float64(imageHeight)

func calculateSkyColour(ray rt.Ray) rt.Colour {
	t := 0.5 * (ray.Direction.Normalize().Y + 1.0)

	white := rt.Colour{Red: 1.0, Green: 1.0, Blue: 1.0}
	gradient := rt.Colour{Red: 0.5, Green: 0.7, Blue: 1.0}

	return white.Scale(1.0 - t).Add(gradient.Scale(t))
}

func calculatePixelRadience(ray rt.Ray, hitables []hitable.Hitable, timeBounds rt.TimeBoundaries) rt.Pixel {
	for _, hitable := range hitables {
		record, isPresent := hitable.DoesIntersectWith(ray, timeBounds).Get()

		if isPresent {
			rawColour := record.Normal.Add(maths.Vec3{X: 1, Y: 1, Z: 1}).Scale(0.5)
			return rt.Pixel{Colour: rt.ColourFromVec3(rawColour), SampleCount: 1}
		}
	}

	return rt.Pixel{Colour: calculateSkyColour(ray), SampleCount: 1}
}

func main() {
	imageFile := "image.ppm"

	if utils.DoesFileExist(imageFile) {
		os.Remove(imageFile)
	}

	outputFile, err := os.Create("image.ppm")
	if err != nil {
		fmt.Println("Failed to open file:", err)
	}
	defer outputFile.Close()

	image := rt.NewImage(imageWidth, imageHeight)

	fmt.Printf("Creating an image of size <%d, %d>\n", image.Width, image.Height)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := maths.Point3{X: 0, Y: 0, Z: 0}
	horizontal := maths.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical := maths.Vec3{X: 0, Y: viewportHeight, Z: 0}

	horizontalMidpoint := horizontal.Scale(0.5).ToPoint3()
	verticalMidpoint := vertical.Scale(0.5).ToPoint3()
	depth := maths.Point3{X: 0, Y: 0, Z: focalLength}

	lowerLeftCorner := origin.Sub(horizontalMidpoint).Sub(verticalMidpoint).Sub(depth)

	// Render

	timeBounds := rt.TimeBoundaries{Min: 0, Max: math.Inf(1)}

	hitables := make([]hitable.Hitable, 0)
	hitables = append(hitables, hitable.Sphere{Origin: maths.Point3{X: 0, Y: 0, Z: 1}, Radius: 0.5})
	hitables = append(hitables, hitable.Sphere{Origin: maths.Point3{X: 0, Y: -100.5, Z: 1}, Radius: 100})

	for j := image.Height - 1; j >= 0; j-- {
		for i := int64(0); i < image.Width; i++ {
			u := float64(i) / float64(image.Width-1)
			v := float64(j) / float64(image.Height-1)

			scaledHorizontal := horizontal.Scale(u)
			scaledVertical := vertical.Scale(v)
			rayDir := lowerLeftCorner.ToVec3().Add(scaledHorizontal).Add(scaledVertical).Sub(origin.ToVec3())

			ray := rt.Ray{Origin: origin, Direction: rayDir}

			image.WritePixel(i, j, calculatePixelRadience(ray, hitables, timeBounds))
		}
	}

	image.SaveAsPPM(outputFile)
}
