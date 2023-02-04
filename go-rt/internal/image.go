package internal

import (
	"fmt"
	"reflect"
	"strings"
)

type Image struct {
	Height int
	Width  int

	Data [][]Pixel
}

func (this *Image) Init(width int, height int) {
	this.Width = width
	this.Height = height
	this.Data = make([][]Pixel, this.Width)
	for i := 0; i < this.Width; i++ {
		this.Data[i] = make([]Pixel, this.Height)
	}
}

func (this *Image) WritePixel(x int, y int, value *Pixel) {
	this.Data[x][y] = *value
}

func (this Image) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", this.Width, this.Height))
	builder.Grow(int(reflect.TypeOf(this.Data[0][0]).Size()))

	for j := this.Height - 1; j >= 0; j-- {
		for i := 0; i < this.Width; i++ {
			builder.WriteString(fmt.Sprintln(this.Data[i][j].String()))
		}
	}

	return builder.String()
}
