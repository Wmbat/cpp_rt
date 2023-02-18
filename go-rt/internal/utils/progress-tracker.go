package utils

import "fmt"

type ProgressTracker struct {
	sampleCount      uint32
	completedSamples uint32
	progress         float64
}

const minProgress float64 = 5.0

func NewProgressTracker(sampleCount uint32) ProgressTracker {
	return ProgressTracker{sampleCount: sampleCount}
}

func (this *ProgressTracker) IncrementProgress() {
	this.completedSamples += 1
	this.progress = float64(this.completedSamples) / float64(this.sampleCount) * 100

	fmt.Printf("%.2f%% (%d / %d)\n", this.progress, this.completedSamples, this.sampleCount)
}
