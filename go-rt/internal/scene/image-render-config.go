package scene

import "github.com/wmbat/ray_tracer/internal/maths"

type ImageRenderConfig struct {
	ImageSize maths.Size2i
	SampleCount uint
	BounceDepth uint
}
