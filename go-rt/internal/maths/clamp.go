package maths

import (
	"golang.org/x/exp/constraints"
)

func Clamp[Type constraints.Ordered](val, min, max Type) Type {
	return Max(min, Min(val, max))
}
