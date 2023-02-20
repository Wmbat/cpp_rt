package world

import "github.com/wmbat/ray_tracer/internal/maths"

type ImageRenderConfig struct {
	ImageSize maths.Size2[int]
	SampleCount int
	BounceDepth int
}
