package structs_methods_interfaces

import "math"

type Shape interface{
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct{
	radius float64
}

type Triangle struct{
	Base float64
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

func (t Triangle) Area() float64{
	return (t.Height * t.Base)/2
}


