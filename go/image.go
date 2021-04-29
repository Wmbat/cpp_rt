package go_pt

type Image struct {
    width int
    height int

    pixels []Pixel
}

func NewImage(width int, height int) Image {
    return Image{width: width, height: height, pixels: make([]Pixel, width * height)}
}
