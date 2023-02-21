package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/world"
	"github.com/wmbat/ray_tracer/internal/world/entt"
	"github.com/wmbat/ray_tracer/internal/world/mats"
)

const aspectRatio float64 = 16.0 / 9.0
const imageWidth int = 1280
const imageHeight int = int((float64(imageWidth) / aspectRatio))

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	lookFrom := maths.Point3{X: -3, Y: 3, Z: 2}
	lookAt := maths.Point3{X: 0, Y: 0, Z: -1}
	cameraInfo := world.CameraCreateInfo{
		LookFrom:    lookFrom,
		LookAt:      lookAt,
		Up:          maths.Vec3{X: 0, Y: 1, Z: 0},
		Fov:         20,
		AspectRatio: aspectRatio,
		Aperture:    2.0,
		FocusDistance: lookFrom.Sub(lookAt).ToVec3().Length(),}

	camera := world.NewCamera(cameraInfo)

	sceneName := "Test Scene"

	mainScene := world.NewScene(sceneName)
	mainScene.SetEnvironmentColour(render.Colour{Red: 135 / 256.0, Green: 206 / 256.0, Blue: 235 / 256.0})
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 0, Y: 0, Z: -1},
		Radius:   0.5,
		Material: mats.Lambertian{Albedo: render.Colour{Red: 0.1, Green: 0.2, Blue: 0.5}}})

	// Floor
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 0, Y: -100.5, Z: -1},
		Radius:   100,
		Material: mats.Lambertian{Albedo: render.Colour{Red: 0.8, Green: 0.8, Blue: 0.0}}})

	// Glass
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: -1, Y: 0, Z: -1},
		Radius:   0.5,
		Material: mats.Dielectric{
			Diffuse:         render.Colour{Red: 1.0, Green: 1.0, Blue: 1.0},
			RefractionIndex: 1.5}})

	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: -1, Y: 0, Z: -1},
		Radius:   -0.4,
		Material: mats.Dielectric{
			Diffuse:         render.Colour{Red: 1.0, Green: 1.0, Blue: 1.0},
			RefractionIndex: 1.5}})

	// Metal
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 1, Y: 0, Z: -1},
		Radius:   0.5,
		Material: mats.Metal{
			Albedo:    render.Colour{Red: 0.8, Green: 0.6, Blue: 0.2},
			Roughness: 0.0}})

	image := mainScene.Render(camera, world.ImageRenderConfig{
		ImageSize:   maths.Size2[int]{Width: imageWidth, Height: imageHeight},
		SampleCount: 1000,
		BounceDepth: 100,
	})

	image.SaveAsPPM(sceneName)
}
