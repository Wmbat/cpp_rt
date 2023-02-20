package utils

import (
	"log"

	"github.com/wmbat/ray_tracer/internal/maths"
)

type ProgressTracker struct {
	sampleCount      int
	completedSamples int
	progress         float64
}

const minProgress float64 = 5.0

func NewProgressTracker(sampleCount int) ProgressTracker {
	return ProgressTracker{sampleCount: sampleCount}
}

func (this *ProgressTracker) IncrementProgress(count int) {
	this.completedSamples = maths.Min(this.sampleCount, this.completedSamples+count)
	this.progress = float64(this.completedSamples) / float64(this.sampleCount) * 100

	log.Printf("[main] %.2f%% (%d / %d)\n", this.progress, this.completedSamples, this.sampleCount)
}
