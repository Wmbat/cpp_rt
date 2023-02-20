package maths

import "golang.org/x/exp/constraints"

func Min[Type constraints.Ordered] (x, y Type) Type {
	if x < y {
		return x
	} else {
		return y
	}
}
