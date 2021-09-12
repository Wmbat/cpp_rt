package main

import (
	"bufio"
	"go_rt/core"
	"go_rt/maths"
	"go_rt/renderable"
	"log"
	"math"
	"math/rand"
	"os"
)

func FindClosestCollision(ray *core.Ray, renderables []renderable.Renderable) renderable.RayCollisionResult {
    result := renderable.RayCollisionResult{HasValue: false}
    closestTime := math.MaxFloat64
    
    for _, renderable := range renderables {
        localResult := renderable.CheckRayCollision(ray, 0.001, closestTime)

        if localResult.HasValue {
            result.HasValue = true
            result.Record = localResult.Record

            closestTime = result.Record.Time
        }
    }

    return result
}

func Radiance(ray *core.Ray, 
    renderables []renderable.Renderable, 
    bounceDepth uint64, 
    scatterRayCount uint64) core.Colour {    

    if bounceDepth == 0 {
        return core.Colour{R: 0.0, G: 0.0, B: 0.0}
    }

    collisionResult := FindClosestCollision(ray, renderables)
    if collisionResult.HasValue {
        randomVec := maths.RandomVec3InHemisphere(&collisionResult.Record.Normal)

        target := maths.Add(&collisionResult.Record.Position, &collisionResult.Record.Normal)
        target.Add(&randomVec)
        
        outRay := core.Ray{
            Origin: collisionResult.Record.Position,
            Direction: maths.Sub(&target, &collisionResult.Record.Position)}  

        result := Radiance(&outRay, renderables, bounceDepth - 1, scatterRayCount)
        result.MultScalar(0.5)

        return result;
    }

    unitVector := maths.Vec3Normalise(&ray.Direction)

    blendFactor := 0.5 * (unitVector.Y + 1.0) 
    start := maths.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
    end := maths.Vec3{X: 0.5, Y: 0.7, Z:1.0}
 
    lerp := maths.Lerp(&start, &end, blendFactor)

    return core.Vec3ToColour(&lerp)
}

func RenderSampleImage(camera *Camera, renderables []renderable.Renderable, settings *core.Settings) core.Image {    
    image := core.NewImage(settings.ImageWidth, settings.ImageHeight)
    for y := image.Height - 1; y >= 0; y-- {
        for x := 0; x < image.Width; x++ {
            u := (float64(x) + rand.Float64()) / float64(image.Width - 1)
            v := (float64(y) + rand.Float64()) / float64(image.Height - 1)

            ray := camera.ShootRay(u, v)

            colour := Radiance(&ray, renderables, settings.BounceDepth, settings.ScatterRayCount)

            image.AddSamples(x, y, &colour, 1)
        }
    }

    return image;
}

func main() {
    aspectRatio := 16.0 / 9.0
    imageWidth := 400

    settings := core.Settings{
        ImageWidth: imageWidth,
        ImageHeight: int(float64(imageWidth) / aspectRatio),
        SampleCount: 100,
        BounceDepth: 50,
        ScatterRayCount: 2}

    cameraCreateInfo := CameraCreateInfo{
        Origin: maths.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
        Height: 2.0,
        AspectRatio: aspectRatio,
        FocalLength: 1.0}

    camera := NewCamera(&cameraCreateInfo)
    renderables := make([]renderable.Renderable, 0)
    renderables = append(renderables, renderable.Sphere{
        Center: maths.Vec3{X: 0.0, Y: 0.0, Z: -1.0},
        Radius: 0.5})
    renderables = append(renderables, renderable.Sphere{
        Center: maths.Vec3{X: 0.0, Y: -100.5, Z: -1.0},
        Radius: 100.0})

    finalImage := core.NewImage(settings.ImageWidth, settings.ImageHeight)
    for i := uint64(0); i < settings.SampleCount; i++ {
        localImage := RenderSampleImage(&camera, renderables, &settings)

        finalImage.AddSampleImage(&localImage)
    }


    file, err := os.Create("result.ppm")
    if err != nil {
        log.Fatal(err)
    }

    writer := bufio.NewWriter(file)
    writer.WriteString(finalImage.String()) 
    writer.Flush()
}
