package maths

type QuadraticFormula struct {
	A float64
	HalfB float64
	C float64
}

func (this QuadraticFormula) ComputeDiscriminant() float64 {
	return this.HalfB * this.HalfB - this.A * this.C
}
