// package main

// import "fmt"

// //call by val
// func Hello(a int) int {
// 	return a + 1
// }

// //cal by reference
// func Hello2(a []int) {
// 	a[0] += 1
// }

// var a int = 290
// func Call2() {
// 	a++
// }

// func main() {
// 	fmt.Println(a)
// 	// a := 10
// 	// fmt.Println(a)
// 	// fmt.Println(Hello(a))
// 	// a := []int{1, 2}
// 	// Hello2(a)
// 	// fmt.Println(a)

// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// var a int

// func localvar() {
// 	a++
// }

// func Coding() {

// }
// func Up(s string) {

// }

// func main() {
// 	// go localvar()
// 	// fmt.Println(a)

// 	a := []string{"Hi", "Hello", "Workd"}
// 	d := map[string]int{}
// 	for i := 0; i < len(a); i++ {
// 		d[strings.ToUpper(string(a[i]))] = len(string(a[i]))
// 	}
// 	fmt.Println(d)

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	stackVariable := 10
// 	fmt.Println("Address of stackVariable:", &stackVariable)
// 	printStackMemory()

// 	heapVariable := new(int)
// 	*heapVariable = 20
// 	fmt.Println("Address of heapVariable:", heapVariable)
// 	printHeapMemory()
// }

// func printStackMemory() {
// 	var stackVariable int
// 	fmt.Println("Address of stackVariable:", &stackVariable)
// }

// func printHeapMemory() {
// 	heapVariable := new(int)
// 	fmt.Println("Address of heapVariable:", heapVariable)

// 	// runtime.GC()
// }

package main

import "fmt"

func greet() string {
	return "Hello World!"
}

func main() {
	greeting := greet
	fmt.Println(greeting())

	// greet = func() string {
	// 	return "Goodbye World!"
	// }

	fmt.Println(greeting())
}
