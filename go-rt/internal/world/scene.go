package world

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/utils"
	"github.com/wmbat/ray_tracer/internal/world/core"
	"github.com/wmbat/ray_tracer/internal/world/entt"
)

type Scene struct {
	Name string

	hitables    []entt.Entity
	environment render.Colour
}

func NewScene(name string) Scene {
	return Scene{Name: name, hitables: make([]entt.Entity, 0)}
}

func (this Scene) Render(cam Camera, config ImageRenderConfig) render.Image {
	fmt.Printf("Start render of scene \"%s\"\n", this.Name)

	image := render.NewImage(config.ImageSize)

	tracker := utils.NewProgressTracker(config.SampleCount)
	for sampleIndex := uint32(0); sampleIndex < config.SampleCount; sampleIndex++ {
		for j := image.Height - 1; j >= 0; j-- {
			for i := int64(0); i < image.Width; i++ {
				camTarget := maths.Point2{
					X: (float64(i) + rand.Float64()) / float64(image.Width-1),
					Y: (float64(j) + rand.Float64()) / float64(image.Height-1)}

				ray := cam.ShootRay(camTarget)

				image.AddSample(i, j, this.radiance(ray, this.hitables, config.BounceDepth))
			}
		}

		tracker.IncrementProgress()
	}

	return image
}

func (this *Scene) AddEntity(entity entt.Entity) {
	this.hitables = append(this.hitables, entity)
}

func (this *Scene) AddEntities(entities []entt.Entity) {
	this.hitables = append(this.hitables, entities...)
}

func (this *Scene) SetEnvironmentColour(colour render.Colour) {
	this.environment = colour
}

func (this Scene) radiance(ray core.Ray, entities []entt.Entity, BounceDepth uint32) render.Colour {
	if BounceDepth == 0 {
		return render.Colour{Red: 0.0, Green: 0.0, Blue: 0.0}
	}

	maxDistance := math.Inf(1)

	for _, entity := range entities {
		record, isPresent := entity.IsIntersectedByRay(ray, maxDistance).Get()

		if isPresent {
			locationVec := record.Location.ToVec3()
			target := locationVec.Add(maths.RandVec3InHemisphere(record.Normal))
			newRay := core.Ray{Origin: record.Location, Direction: target.Sub(record.Location.ToVec3())}
			return this.radiance(newRay, entities, BounceDepth-1).Scale(0.5)
		}
	}

	return this.environment
}
