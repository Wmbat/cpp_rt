package world

import (
	"log"
	"math"
	"math/rand"
	"runtime"
	"time"

	"github.com/wmbat/ray_tracer/internal/maths"
	"github.com/wmbat/ray_tracer/internal/render"
	"github.com/wmbat/ray_tracer/internal/utils"
	"github.com/wmbat/ray_tracer/internal/world/core"
	"github.com/wmbat/ray_tracer/internal/world/entt"
	"github.com/wmbat/ray_tracer/internal/world/mats"
)

type Scene struct {
	Name string

	entities    []entt.Entity
	environment render.Colour
}

type imageRenderInfo struct {
	Tracker *utils.ProgressTracker
}

func NewScene(name string) Scene {
	return Scene{Name: name, entities: make([]entt.Entity, 0)}
}

func (this *Scene) AddEntity(entity entt.Entity) {
	this.entities = append(this.entities, entity)
}

func (this *Scene) AddEntities(entities []entt.Entity) {
	this.entities = append(this.entities, entities...)
}

func (this *Scene) SetEnvironmentColour(colour render.Colour) {
	this.environment = colour
}

func (this Scene) Render(cam Camera, config ImageRenderConfig) render.Image {
	log.Printf("[main] Start render of scene \"%s\"\n", this.Name)

	goroutineCount := runtime.NumCPU()
	batchCount := getGoroutineBatchCount(goroutineCount, config.SampleCount)
	leftoverSamples := config.SampleCount

	workerPool := utils.NewWorkerPool(goroutineCount)
	workerPool.Run()

	resultChannel := make(chan render.Image, config.SampleCount)

	tracker := utils.NewProgressTracker(config.SampleCount)
	image := render.NewImage(config.ImageSize)
	for batchIndex := 0; batchIndex < batchCount; batchIndex++ {
		sampleCount := maths.Min(leftoverSamples, goroutineCount)

		// Dispatch all the image render tasks
		for index := 0; index < sampleCount; index++ {
			workerPool.AddTask(func() {
				resultChannel <- this.RenderImage(cam, config)
			})
		}

		workerPool.Wait()

		// Collect all the sample images
		for index := 0; index < sampleCount; index++ {
			image.AddSampleImage(<-resultChannel)
		}

		tracker.IncrementProgress(sampleCount)

		leftoverSamples -= goroutineCount
	}

	close(resultChannel)
	workerPool.Close()

	log.Print("[main] Final image rendered")

	return image
}

func (this Scene) RenderImage(cam Camera, config ImageRenderConfig) render.Image {
	image := render.NewImage(config.ImageSize)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	records := make([]entt.IntersectRecord, len(this.entities))

	for j := image.Size.Height - 1; j >= 0; j-- {
		for i := 0; i < image.Size.Width; i++ {
			camTarget := maths.Point2[float64]{
				X: (float64(i) + rng.Float64()) / float64(image.Size.Width-1),
				Y: (float64(j) + rng.Float64()) / float64(image.Size.Height-1)}

			ray := cam.ShootRay(camTarget, rng)

			image.AddSample(i, j, this.radiance(ray, this.entities, records, rng, config.BounceDepth))
		}
	}

	return image
}

func (this Scene) radiance(ray core.Ray, entities []entt.Entity, records []entt.IntersectRecord, rng *rand.Rand, BounceDepth int) render.Colour {
	black := render.Colour{Red: 0.0, Green: 0.0, Blue: 0.0}

	if BounceDepth == 0 {
		return black
	}

	intersect, isIntersected := findNearestIntersectRecord(ray, entities, records)
	if isIntersected {
		scatterInfo := mats.ScatterInfo{
			Ray:         ray,
			Position:    intersect.Position,
			Normal:      intersect.Normal,
			IsFrontFace: intersect.IsFrontFace,
			Rng:         rng}

		scatter, isScattered := intersect.Material.Scatter(scatterInfo)
		if isScattered {
			return scatter.Attenuation.Mult(this.radiance(scatter.Ray, entities, records, rng, BounceDepth-1))
		} else {
			return black
		}
	}

	return this.environment
}

func findNearestIntersectRecord(ray core.Ray, entities []entt.Entity, records []entt.IntersectRecord) (*entt.IntersectRecord, bool) {
	var nearestRecord *entt.IntersectRecord = nil

	nearestDistance := math.Inf(1)
	for i, entity := range entities {
		isPresent := false
		records[i], isPresent = entity.IsIntersectedByRay(ray, nearestDistance)
		if isPresent {
			nearestDistance = records[i].Distance
			nearestRecord = &records[i]
		}
	}

	if nearestRecord == nil {
		return &entt.IntersectRecord{}, false
	} else {
		return nearestRecord, true
	}
}

func getGoroutineBatchCount(goroutineCount, sampleCount int) int {
	return int(math.Ceil(float64(sampleCount) / float64(goroutineCount)))
}
