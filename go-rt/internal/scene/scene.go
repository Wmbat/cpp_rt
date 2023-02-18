package scene

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/wmbat/ray_tracer/internal/core"
	"github.com/wmbat/ray_tracer/internal/hitable"
	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/utils"
)

type Scene struct {
	Name string

	hitables    []hitable.Hitable
	environment render.Colour
}

func NewScene(name string) Scene {
	return Scene{Name: name, hitables: make([]hitable.Hitable, 0)}
}

func (this Scene) Render(cam Camera, config ImageRenderConfig) render.Image {
	fmt.Printf("Start render of scene \"%s\"\n", this.Name)

	image := render.NewImage(config.ImageSize)

	for j := image.Height - 1; j >= 0; j-- {
		for i := int64(0); i < image.Width; i++ {
			for sampleIndex := uint(0); sampleIndex < config.SampleCount; sampleIndex++ {
				camTarget := maths.Point2{
					X: (float64(i) + rand.Float64()) / float64(image.Width-1),
					Y: (float64(j) + rand.Float64()) / float64(image.Height-1)}

				ray := cam.ShootRay(camTarget)

				image.AddSample(i, j, this.radiance(ray, this.hitables, int(config.BounceDepth)))
			}
		}
	}

	return image
}

func (this *Scene) AddHitable(hitable hitable.Hitable) {
	this.hitables = append(this.hitables, hitable)
}

func (this *Scene) AddHitables(hitables []hitable.Hitable) {
	this.hitables = append(this.hitables, hitables...)
}

func (this *Scene) SetEnvironmentColour(colour render.Colour) {
	this.environment = colour
}

func (this Scene) radiance(ray core.Ray, hitables []hitable.Hitable, BounceDepth int) render.Colour {
	if BounceDepth <= 0 {
		return render.Colour{Red: 0.0, Green: 0.0, Blue: 0.0}
	}

	timeBounds := utils.TimeBoundaries{Min: 0.001, Max: math.Inf(1)}

	for _, hitable := range hitables {
		record, isPresent := hitable.DoesIntersectWith(ray, timeBounds).Get()

		if isPresent {
			locationVec := record.Location.ToVec3()
			target := locationVec.Add(record.Normal).Add(maths.RandomVec3InUnitSphere())
			newRay := core.Ray{Origin: record.Location, Direction: target.Sub(record.Location.ToVec3())}
			return this.radiance(newRay, hitables, BounceDepth-1).Scale(0.5)
		}
	}

	return this.environment
}
