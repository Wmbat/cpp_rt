package main

import (
	"bufio"
	"fmt"
	"go_rt/core"
	"go_rt/entities"
	"go_rt/materials"
	"go_rt/maths"
	"log"
	"math"
	"math/rand"
	"os"
)

func FindClosestCollision(ray *core.Ray, sceneEntities []entities.Entity) entities.RayHitResult {
    result := entities.RayHitResult{HasValue: false}
    closestTime := math.MaxFloat64
    
    for _, renderable := range sceneEntities{
        localResult := renderable.CheckRayHit(ray, 0.001, closestTime)

        if localResult.HasValue {
            result.HasValue = true
            result.Record = localResult.Record

            closestTime = result.Record.Hit.Time
        }
    }

    return result
}

func Radiance(ray *core.Ray, sceneEntities []entities.Entity, bounceDepth uint64, 
    scatterRayCount uint64) core.Colour {    

    if bounceDepth == 0 {
        return core.Colour{R: 0.0, G: 0.0, B: 0.0}
    }

    collisionResult := FindClosestCollision(ray, sceneEntities)
    if collisionResult.HasValue {
        mat := collisionResult.Record.Mat

        scatterData := mat.Scatter(ray, &collisionResult.Record.Hit)

        result := core.Colour{}
        
        for i := uint64(0); i < scatterRayCount; i++ {
            localResult := Radiance(&scatterData.Ray, sceneEntities, bounceDepth - 1, scatterRayCount)
            localResult.Mult(&scatterData.Diffuse)
            localResult.Add(&scatterData.Emission)

            result.Add(&localResult)
        }

        result.MultScalar(1 / float64(scatterRayCount))

        return result;
    } else {
        unitVector := maths.Vec3Normalise(&ray.Direction)

        blendFactor := 0.5 * (unitVector.Y + 1.0) 
        start := maths.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
        end := maths.Vec3{X: 0.5, Y: 0.7, Z:1.0}
 
        lerp := maths.Lerp(&start, &end, blendFactor)

        return core.Vec3ToColour(&lerp)
    }
}

func RenderSampleImage(camera *Camera, sceneEntities []entities.Entity, settings *core.Settings) core.Image {    
    image := core.NewImage(settings.ImageWidth, settings.ImageHeight)
    for y := image.Height - 1; y >= 0; y-- {
        for x := 0; x < image.Width; x++ {
            u := (float64(x) + rand.Float64()) / float64(image.Width - 1)
            v := (float64(y) + rand.Float64()) / float64(image.Height - 1)

            ray := camera.ShootRay(u, v)

            colour := Radiance(&ray, sceneEntities, settings.BounceDepth, settings.ScatterRayCount)

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
        SampleCount: 24,
        BounceDepth: 5,
        ScatterRayCount: 4}

    cameraCreateInfo := CameraCreateInfo{
        Origin: maths.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
        Height: 2.0,
        AspectRatio: aspectRatio,
        FocalLength: 1.0}

    camera := NewCamera(&cameraCreateInfo)

    entityMaterials := make([]materials.Material, 0)
    entityMaterials = append(entityMaterials, materials.Lambertian{
        Diffuse: core.Colour{R: 0.7, G: 0.3, B: 0.3}})
    entityMaterials = append(entityMaterials, materials.Lambertian{
        Diffuse: core.Colour{R: 0.8, G: 0.8, B: 0.0}})
    entityMaterials = append(entityMaterials, materials.Metal{
        Diffuse: core.Colour{R: 0.8, G: 0.8, B: 0.8},
        Roughness: 0.3,})
    entityMaterials = append(entityMaterials, materials.Metal{
        Diffuse: core.Colour{R: 0.8, G: 0.6, B: 0.2},
        Roughness: 1.0,})

    sceneEntities := make([]entities.Entity, 0)
    sceneEntities = append(sceneEntities, entities.Sphere{
        Center: maths.Vec3{X: 0.0, Y: 0.0, Z: -1.0},
        Radius: 0.5,
        Mat: entityMaterials[0],})
    sceneEntities = append(sceneEntities, entities.Sphere{
        Center: maths.Vec3{X: -1.0, Y: 0.0, Z: -1.0},
        Radius: 0.5,
        Mat: entityMaterials[2],})
    sceneEntities = append(sceneEntities, entities.Sphere{
        Center: maths.Vec3{X: 1.0, Y: 0.0, Z: -1.0},
        Radius: 0.5,
        Mat: entityMaterials[3],})
    sceneEntities = append(sceneEntities, entities.Sphere{
        Center: maths.Vec3{X: 0.0, Y: -100.5, Z: -1.0},
        Radius: 100.0,
        Mat: entityMaterials[1]})

    finalImage := core.NewImage(settings.ImageWidth, settings.ImageHeight)
    for i := uint64(0); i < settings.SampleCount; i++ {
        localImage := RenderSampleImage(&camera, sceneEntities, &settings)

        finalImage.AddSampleImage(&localImage)

        fmt.Println(fmt.Sprintf("Render Status: %3.2f%%", float64(i) / float64(settings.SampleCount) * 100))
    }


    file, err := os.Create("result.ppm")
    if err != nil {
        log.Fatal(err)
    }

    writer := bufio.NewWriter(file)
    writer.WriteString(finalImage.String()) 
    writer.Flush()
}
