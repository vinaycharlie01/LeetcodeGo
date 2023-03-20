package main

import "fmt"

type Vehical interface {
	Working(a interface{})
	Speedchek(a interface{}) any
}

type Car struct {
	Vehical
	name  string
	color string
	model string
	speed string
}

func (c *Car) Working() {
	fmt.Printf("%v\n %v\n %v\n", c.name, c.model, c.color)
}
func (c *Car) Speedchek() any {
	return fmt.Sprintf("%v\n", c.speed)
}

func main() {
	t := Car{
		name:  "Tesla",
		color: "RED",
		model: "1.2",
		speed: "120",
	}
}
