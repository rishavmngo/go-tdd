package structs

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Reactangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2.0
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Reactangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
