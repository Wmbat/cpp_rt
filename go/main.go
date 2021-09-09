package main

import (
	"bufio"
	"go_pt/maths"
	"log"
	"os"
)

func ComputeRayColour(ray *Ray) maths.Vec3 {
    unitVector := maths.Normalise(&ray.Direction)

    blendFactor := 0.5 * (unitVector.Y + 1.0) 
    start := maths.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
    end := maths.Vec3{X: 0.5, Y: 0.7, Z:1.0}

    return maths.Lerp(&start, &end, blendFactor)
}

func main() {
    aspectRatio := 16.0 / 9.0
    imageWidth := 400
    imageHeight := int(400 / aspectRatio)

    cameraCreateInfo := CameraCreateInfo{
        Origin: maths.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
        Height: 2.0,
        AspectRatio: aspectRatio,
        FocalLength: 1.0}

    camera := NewCamera(&cameraCreateInfo)

    image := NewImage(imageWidth, imageHeight)
    for y := image.Height - 1; y >= 0; y-- {
        for x := 0; x < image.Width; x++ {
            u := float64(x) / float64(imageWidth - 1)
            v := float64(y) / float64(imageHeight - 1)

            ray := camera.ShootRay(u, v)
            colour := ComputeRayColour(&ray)

            image.SetImagePixelColour(x, y, &colour)
        }
    }

    file, err := os.Create("result.ppm")
    if err != nil {
        log.Fatal(err)
    }

    writer := bufio.NewWriter(file)
    writer.WriteString(image.String()) 
    writer.Flush()
}
