package maths

import "fmt"

type Size2f struct {
	Width float64
	Height float64
}

type Size2i struct {
	Width int64
	Height int64
}

func (this Size2i) ToString() string {
	return fmt.Sprintf("<%d, %d>", this.Width, this.Height)
}
