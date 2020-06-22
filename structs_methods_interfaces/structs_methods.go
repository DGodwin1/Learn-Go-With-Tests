package structs_methods_interfaces

import "math"

func Area(b, h float64) float64{
	return b*h
}

type Rectangle struct {
	Width float64
	Height float64

}

func (r Rectangle) Perimeter() float64{
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64{
	return r.Height * r.Width
}

func (c Circle) Area() float64{
	return c.radius * c.radius * math.Pi
}

type Circle struct{
	radius float64
}

