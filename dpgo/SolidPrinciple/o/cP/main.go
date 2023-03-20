package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

type AreaCalculator struct {
	shapes []Shape
}

func (ac *AreaCalculator) AddShape(shape Shape) {
	ac.shapes = append(ac.shapes, shape)
}

func (ac *AreaCalculator) TotalArea() float64 {
	total := 0.0
	for _, shape := range ac.shapes {
		total += shape.Area()
	}
	return total
}

func main() {

	calculator := AreaCalculator{}
	calculator.AddShape(&Rectangle{width: 2, height: 3})
	calculator.AddShape(&Circle{radius: 1})
	calculator.AddShape(&Triangle{base: 4, height: 5})
	totalArea := calculator.TotalArea()
	fmt.Println(totalArea)
	fmt.Println(calculator.shapes[0].Area())
	fmt.Println(calculator.shapes[1].Area())
	fmt.Println(calculator.shapes[2].Area())

}
