package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Height: 5.0, Width: 10.0}
	got := Perimeter(rectangle)
	want := 30.0
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{5.0, 10.0}, 50.0},
		{"Circle", Circle{8.0}, math.Pi * math.Pow(8.0, 2.0)},
		{"Triangle", Triangle{11.0, 3.0}, 16.5},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want
			if got != want {
				t.Errorf("shape: %#v, got: %v, want: %v", tt.shape, got, want)
			}
		})
	}
}
