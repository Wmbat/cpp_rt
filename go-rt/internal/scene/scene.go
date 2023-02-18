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

	timeBounds := utils.TimeBoundaries{Min: 0, Max: math.Inf(1)}
	image := render.NewImage(config.ImageSize)

	for j := image.Height - 1; j >= 0; j-- {
		for i := int64(0); i < image.Width; i++ {
			for sampleIndex := 0; sampleIndex < config.SampleCount; sampleIndex++ {
				camTarget := maths.Point2{
					X: (float64(i) + rand.Float64()) / float64(image.Width-1),
					Y: (float64(j) + rand.Float64()) / float64(image.Height-1)}

				ray := cam.ShootRay(camTarget)

				image.AddSample(i, j, this.radiance(ray, this.hitables, timeBounds))
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

func (this Scene) radiance(ray core.Ray, hitables []hitable.Hitable, timeBounds utils.TimeBoundaries) render.Colour {
	for _, hitable := range hitables {
		record, isPresent := hitable.DoesIntersectWith(ray, timeBounds).Get()

		if isPresent {
			rawColour := record.Normal.Add(maths.Vec3{X: 1, Y: 1, Z: 1}).Scale(0.5)
			return render.ColourFromVec3(rawColour)
		}
	}

	return this.environment
}
