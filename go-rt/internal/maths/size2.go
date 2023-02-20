package maths

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Size2[Type constraints.Integer | constraints.Float] struct {
	Width, Height Type
}

func (this Size2[Type]) ToString() string {
	return fmt.Sprintf("<%v, %v>", this.Width, this.Height)
}
