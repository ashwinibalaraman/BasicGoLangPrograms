package main

import (
	"fmt"

	s "github.com/assignment1/shape/shapePeri"
)

func main() {
	c := s.Circle{}
	r := s.Rectangle{}
	c.SetRadius(6)
	r.SetRectValues(3, 7)

	CalcPerimeter(&c, &r)
}

func CalcPerimeter(shapes ...s.Shape) {
	for _, shape := range shapes {
		fmt.Println(s.Perimeter(shape))
	}
}
