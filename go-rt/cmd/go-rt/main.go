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
const imageWidth int64 = 1280
const imageHeight int64 = int64((float64(imageWidth) / aspectRatio))

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

	viewport := maths.Size2f{Width: aspectRatio * 2.0, Height: 2.0}
	camera := world.NewCamera(maths.Point3{X: 0, Y: 0, Z: 0}, viewport, 1.0)

	sceneName := "Test Scene"

	mainScene := world.NewScene(sceneName)
	mainScene.SetEnvironmentColour(render.Colour{Red: 135 / 256.0, Green: 206 / 256.0, Blue: 235 / 256.0})
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 0, Y: 0, Z: 1},
		Radius:   0.5,
		Material: mats.Lambertian{Albedo: render.Colour{Red: 0.7, Green: 0.3, Blue: 0.3}}})
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 0, Y: -100.5, Z: 1},
		Radius:   100,
		Material: mats.Lambertian{Albedo: render.Colour{Red: 0.8, Green: 0.8, Blue: 0.0}}})
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: -1, Y: 0, Z: 1},
		Radius:   0.5,
		Material: mats.Metal{Albedo: render.Colour{Red: 0.8, Green: 0.8, Blue: 0.8}}})
	mainScene.AddEntity(entt.Sphere{
		Position: maths.Point3{X: 1, Y: 0, Z: 1},
		Radius:   0.5,
		Material: mats.Metal{Albedo: render.Colour{Red: 0.8, Green: 0.6, Blue: 0.2}}})

	image := mainScene.Render(camera, world.ImageRenderConfig{
		ImageSize:   maths.Size2i{Width: imageWidth, Height: imageHeight},
		SampleCount: 1000,
		BounceDepth: 50,
	})

	image.SaveAsPPM(sceneName)
}
