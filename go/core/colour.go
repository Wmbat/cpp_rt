package core

import "go_rt/maths"

type Colour struct {
    R float64
    G float64
    B float64
}

func (lhs *Colour) Add(rhs *Colour) *Colour {
   lhs.R += rhs.R
   lhs.G += rhs.G
   lhs.B += rhs.B

   return lhs
}

func (c *Colour) AddScalar(scalar float64) *Colour {
    c.R += scalar 
    c.G += scalar 
    c.B += scalar 

    return c;
}

func (c *Colour) MultScalar(scalar float64) *Colour {
    c.R *= scalar 
    c.G *= scalar 
    c.B *= scalar 

    return c;
}

func Vec3ToColour(vec *maths.Vec3) Colour {
    return Colour{R: vec.X, G: vec.Y, B: vec.Z}
}
