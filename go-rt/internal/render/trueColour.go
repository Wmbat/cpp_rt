package render

import "fmt"

// Represent a 24bit colour element where each value (Red, Green, and Blue) contain
// a number between [0, 255]
type TrueColour struct {
	Red, Green, Blue uint8
}

func (this TrueColour) String() string {
	return fmt.Sprintf("%d %d %d", this.Red, this.Green, this.Blue)
}
