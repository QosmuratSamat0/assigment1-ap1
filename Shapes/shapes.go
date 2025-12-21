package Shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	A float64
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (s *Square) Area() float64 {
	return math.Pow(s.A, 2)
}

func (s *Square) Perimeter() float64 {
	return 4 * s.A
}

func (t *Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}
