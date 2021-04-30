package main

import (
	"bufio"
	"go_pt/maths"
	"log"
	"os"
)

func main() {
    image := NewImage(256, 256)
    
    for y := image.height - 1; y >= 0; y-- {
        for x := 0; x < image.width; x++ {
            r := (float64(y) / float64(image.width - 1))
            g := (float64(x) / float64(image.height - 1))
            b := 0.25

            colour := maths.Vec3{X: r, Y: g, Z: b}
            image.AddSample(x, y, &colour, 1)
        }
    }
 
    file, err := os.Create("image.ppm")
    if err != nil {
        log.Fatal(err)
    }

    writer := bufio.NewWriter(file)
    writer.WriteString(image.String())
    writer.Flush()
}
