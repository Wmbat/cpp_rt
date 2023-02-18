package rt

import "fmt"

type TrueColour struct {
	Red uint8
	Green uint8
	Blue uint8
}

func (this TrueColour) String() string {
	return fmt.Sprintf("%d %d %d", this.Red, this.Green, this.Blue)
}
