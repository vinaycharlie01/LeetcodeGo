package main

import (
	"fmt"
	"reflect"
)

func Revese(a []int) {

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	// for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	// 	a[i], a[j] = a[j], a[i]
	// }

}

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

// func (o *Orange) String() string {
// 	return fmt.Sprintf("%v", o.Quantity)
// }
func (o *Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}

func passByValue(x int) {
	x = x + 1
	fmt.Println("inside passByValue: x =", x)
}

func main() {
	// var a = []byte("a")
	// fmt.Println(a)
	// var count int
	// for _, v := range a {
	// 	count += int(v)
	// }
	// fmt.Println(count)

	// p := new(map[int]int) // allocate memory for an int, set it to 0, and return a pointer to it
	// fmt.Println(unsafe.Sizeof(p))
	var a interface{} = 10
	fmt.Println(reflect.TypeOf(a))

	// var orange Orange
	// orange.Increase(10)
	// orange.Decrease(5)
	// fmt.Println(orange)
	//runtime.NumGoroutine()

	b := struct {
		firstName, lastName string
		age                 int
	}{}
	// }{
	// 	firstName: "Diana",
	// 	lastName:  "Zimmermann",
	// }
	b.age = 12
	b.firstName = "vinay"
	b.lastName = "Kumar"
	// fmt.Println(b)

	// b := code{name: "vinay"}

	// a := []int{1, 2, 3, 4}
	// Revese(a)
	// fmt.Println(a)

	// a := map[string]bool{"A": true, "B": true}
	// b := map[string]bool{"C": true, "D": true}
	// check := a
	// a = b
	// fmt.Println(a, b, check)

}
