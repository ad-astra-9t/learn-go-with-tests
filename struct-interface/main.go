package main

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radicus float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radicus, 2.0)
}

func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2.0
}
