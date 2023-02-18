package main

import (
	"fmt"
	"math"
	"os"

	"github.com/wmbat/ray_tracer/internal/core"
	"github.com/wmbat/ray_tracer/internal/hitable"
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/scene"
	"github.com/wmbat/ray_tracer/internal/utils"
)

const imageWidth int64 = 1280
const imageHeight int64 = 720
const aspectRatio float64 = float64(imageWidth) / float64(imageHeight)

func calculateSkyColour(ray core.Ray) render.Colour {
	t := 0.5 * (ray.Direction.Normalize().Y + 1.0)

	white := render.Colour{Red: 1.0, Green: 1.0, Blue: 1.0}
	gradient := render.Colour{Red: 0.5, Green: 0.7, Blue: 1.0}

	return white.Scale(1.0 - t).Add(gradient.Scale(t))
}

func calculatePixelRadience(ray core.Ray, hitables []hitable.Hitable, timeBounds utils.TimeBoundaries) render.Pixel {
	for _, hitable := range hitables {
		record, isPresent := hitable.DoesIntersectWith(ray, timeBounds).Get()

		if isPresent {
			rawColour := record.Normal.Add(maths.Vec3{X: 1, Y: 1, Z: 1}).Scale(0.5)
			return render.Pixel{Colour: render.ColourFromVec3(rawColour), SampleCount: 1}
		}
	}

	return render.Pixel{Colour: calculateSkyColour(ray), SampleCount: 1}
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

	image := render.NewImage(imageWidth, imageHeight)

	fmt.Printf("Creating an image of size <%d, %d>\n", image.Width, image.Height)

	// Camera
	viewport := maths.Size2{Width: aspectRatio * 2.0, Height: 2.0}
	camera := scene.NewCamera(maths.Point3{X: 0, Y: 0, Z: 0}, viewport, 1.0)

	// Render
	timeBounds := utils.TimeBoundaries{Min: 0, Max: math.Inf(1)}

	hitables := make([]hitable.Hitable, 0)
	hitables = append(hitables, hitable.Sphere{Origin: maths.Point3{X: 0, Y: 0, Z: 1}, Radius: 0.5})
	hitables = append(hitables, hitable.Sphere{Origin: maths.Point3{X: 0, Y: -100.5, Z: 1}, Radius: 100})

	for j := image.Height - 1; j >= 0; j-- {
		for i := int64(0); i < image.Width; i++ {
			camTarget := maths.Point2{
				X: float64(i) / float64(image.Width-1),
				Y: float64(j) / float64(image.Height-1)}

			ray := camera.ShootRay(camTarget)

			image.WritePixel(i, j, calculatePixelRadience(ray, hitables, timeBounds))
		}
	}

	image.SaveAsPPM(outputFile)
}
