package shapePeri

import "math"

type Shape interface {
	perim() float64
}

type Circle struct {
	r float64
}

func (c *Circle) SetRadius(radius float64) {
	c.r = radius
}

func (r *Rectangle) SetRectValues(height float64, width float64) {
	r.h = height
	r.w = width
}

type Rectangle struct {
	h, w float64
}

func (c *Circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perim() float64 {
	return 2 * (r.w + r.h)
}

func Perimeter(shape Shape) float64 {
	return shape.perim()
}
