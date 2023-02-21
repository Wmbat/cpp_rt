package main

import (
	"math/rand"
	"time"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world"
	"github.com/wmbat/ray_tracer/internal/world/entt"
	"github.com/wmbat/ray_tracer/internal/world/mats"
)

const aspectRatio float64 = 16.0 / 9.0
const imageWidth int = 1280
const imageHeight int = int((float64(imageWidth) / aspectRatio))

func main() {
	lookFrom := maths.Point3{X: 13, Y: 2, Z: 3}
	lookAt := maths.Point3{X: 0, Y: 0, Z: 0}
	cameraInfo := world.CameraCreateInfo{
		LookFrom:      lookFrom,
		LookAt:        lookAt,
		Up:            maths.Vec3{X: 0, Y: 1, Z: 0},
		Fov:           20,
		AspectRatio:   aspectRatio,
		Aperture:      0.1,
		FocusDistance: 10}

	camera := world.NewCamera(cameraInfo)

	scene := GenerateRandomScene()

	image := scene.Render(camera, world.ImageRenderConfig{
		ImageSize:   maths.Size2[int]{Width: imageWidth, Height: imageHeight},
		SampleCount: 500,
		BounceDepth: 50,
	})

	image.SaveAsPPM(scene.Name)
}

func GenerateRandomScene() world.Scene {
	scene := world.NewScene("RandomScene")
	scene.SetEnvironmentColour(render.Colour{Red: 135 / 256.0, Green: 206 / 256.0, Blue: 235 / 256.0})

	// Floor
	scene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 0, Y: -1000, Z: -1},
		Radius:   1000,
		Material: mats.Lambertian{Albedo: render.Colour{Red: 0.5, Green: 0.5, Blue: 0.5}}})

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := -11; i < 11; i++ {
		for j := -11; j < 11; j++ {
			center := maths.Point3{
				X: float64(i) + (0.9 * rng.Float64()),
				Y: 0.2,
				Z: float64(j) + (0.9 * rng.Float64())}

			if center.Sub(maths.Point3{X: 4, Y: 0.2, Z: 0}).ToVec3().Length() > 0.9 {
				chooseMat := rng.Float64()

				if chooseMat < 0.8 {
					scene.AddEntity(entt.Sphere{
						Position: center,
						Radius:   0.2,
						Material: mats.Lambertian{Albedo: render.NewRandColour(rng)}})
				} else if chooseMat < 0.95 {
					scene.AddEntity(entt.Sphere{
						Position: center,
						Radius:   0.2,
						Material: mats.Metal{
							Albedo:    render.NewRandColour(rng),
							Roughness: maths.Min(rng.Float64(), 0.7)}})
				} else {
					scene.AddEntity(entt.Sphere{
						Position: center,
						Radius:   0.2,
						Material: mats.Dielectric{
							Diffuse:         render.Colour{Red: 1.0, Green: 1.0, Blue: 1.0},
							RefractionIndex: 1.5}})
				}
			}
		}
	}

	scene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 0, Y: 1, Z: 0},
		Radius:   1.0,
		Material: mats.Dielectric{
			Diffuse:         render.Colour{Red: 1.0, Green: 1.0, Blue: 1.0},
			RefractionIndex: 1.5}})

	scene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: -4, Y: 1, Z: 0},
		Radius:   1.0,
		Material: mats.Lambertian{Albedo: render.Colour{Red: 0.4, Green: 0.2, Blue: 0.1}}})

	scene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 4, Y: 1, Z: 0},
		Radius:   1.0,
		Material: mats.Metal{
			Albedo:    render.Colour{Red: 0.7, Green: 0.6, Blue: 0.5},
			Roughness: 0}})

	return scene
}
