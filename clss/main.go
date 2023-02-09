// package main

// import "fmt"

// type Car struct {
// 	name  string
// 	color string
// 	model string
// 	speed int
// }
// type bike struct {
// 	name  string
// 	color string
// 	model string
// 	speed int
// }

// type Working interface {
// 	Speedcontrol(speed int) string
// }

// func (c *Car) Speedcontrol(speed int) string {
// 	if speed > c.speed {
// 		return "stop"
// 	} else {
// 		return "good speed"
// 	}
// }

// func (c *bike) Speedcontrol(speed int) string {
// 	if c.speed > speed {
// 		return "stop"
// 	} else {
// 		return "good speed"
// 	}
// }
// func main() {

// 	a := []Working{
// 		&bike{
// 			name:  "tesla",
// 			color: "red",
// 			speed: 120,
// 			model: "1.0",
// 		}}
// 	for _, v := range a {
// 		fmt.Println(v.Speedcontrol(200))
// 	}
// 	// fmt.Println(a.Speedcontrol(100))
// 	// fmt.Println(b.Speedcontrol(200))
// }

// package main

// import "fmt"

// // Shape is an interface that defines the behavior of shapes.
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// // Rectangle represents a rectangle shape.
// type Rectangle struct {
// 	Width, Height float64
// }

// // Area returns the area of the rectangle.
// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// // Perimeter returns the perimeter of the rectangle.
// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// // Circle represents a circle shape.
// type Circle struct {
// 	Radius float64
// }

// // Area returns the area of the circle.
// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// // Perimeter returns the perimeter of the circle.
// func (c Circle) Perimeter() float64 {
// 	return 2 * 3.14 * c.Radius
// }

// // PrintShapeInfo takes a shape and prints its area and perimeter.
// func PrintShapeInfo(s Shape) {
// 	fmt.Println("Area:", s.Area())
// 	fmt.Println("Perimeter:", s.Perimeter())
// }

// func main() {
// 	r := Rectangle{Width: 10, Height: 5}
// 	c := Circle{Radius: 5}
// 	PrintShapeInfo(r)
// 	PrintShapeInfo(c)
// }

// package main

// import (
// 	"fmt"
// )

// type Payment interface {
// 	pay(a int) string
// }
// type cash struct {
// 	money int
// }
// type onlinepay struct {
// 	money int
// }
// type usadoller struct {
// 	money int
// }

// func (c *cash) pay(a int) string {
// 	return "sucessful your cash payment"
// }
// func (o *onlinepay) pay(a int) string {
// 	return "sucessful your onlince payment"
// }
// func (u *usadoller) pay(a int) string {
// 	return "sucessful your usadoller payment"
// }

// func Interconnection(p Payment) string {
// 	return p.pay(10)
// }

// func main() {

// 	type a comparable
// 	a := onlinepay{}
// 	var s []Payment
// 	s = append(s, &a)
// 	fmt.Println(Interconnection(s[0]))
// 	fmt.Println("Done")
// }
package main

import (
	"fmt"
	"reflect"
)

type Payment interface {
	pay(a int) string
}
type cash struct {
	money int
}
type onlinepay struct {
	money int
}
type usadoller struct {
	money int
}

func (c *cash) pay(a int) string {
	return "sucessful your cash payment"
}
func (o *onlinepay) pay(a int) string {
	return "sucessful your onlince payment"
}
func (u *usadoller) pay(a int) string {
	return "sucessful your usadoller payment"
}

func Interconnection(p Payment) string {
	return p.pay(10)
}

func main() {
	a := onlinepay{}
	var s []Payment
	b := cash{}
	c := usadoller{}
	s = append(s, &a, &b, &c)
	for _, j := range s {
		fmt.Println(reflect.TypeOf(Interconnection(j)))
		///fmt.Println((Interconnection(j)))
	}
	fmt.Println("Done")
}
