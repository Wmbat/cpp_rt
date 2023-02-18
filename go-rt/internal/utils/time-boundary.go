package utils

type TimeBoundaries struct {
	Min float64
	Max float64
}

func (this TimeBoundaries) IsTimeWithinBounds(time float64) bool {
	return this.Min <= time && time <= this.Max
}
