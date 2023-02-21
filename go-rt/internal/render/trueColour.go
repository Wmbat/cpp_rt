package render

import "fmt"

type TrueColour struct {
	Red, Green, Blue uint8
}

func (this TrueColour) String() string {
	return fmt.Sprintf("%d %d %d", this.Red, this.Green, this.Blue)
}
