package core

import "fmt"

type TrueColour struct {
	Red uint32
	Green uint32
	Blue uint32
}

func (this TrueColour) String() string {
	return fmt.Sprintf("%d %d %d", this.Red, this.Green, this.Blue)
}
