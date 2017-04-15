package shape

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	X float64 `tag:"test""`
	Y float64
	R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

type Ellipse struct {
	Circle
	R2 float64
}

func (e *Ellipse) Area() float64 {
	//return math.Pi * Circle{}

	return e.Circle.R * e.R2
}

type Retangle struct {
	W, H float64
}

func (r *Retangle) Area() float64 {
	return r.W * r.H
}
