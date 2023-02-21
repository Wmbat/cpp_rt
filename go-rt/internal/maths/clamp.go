package maths

import (
	"golang.org/x/exp/constraints"
)

// Clamp a value between a pair of boundary values
func Clamp[Type constraints.Ordered](val, min, max Type) Type {
	return Max(min, Min(val, max))
}
