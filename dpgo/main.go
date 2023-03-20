package main

import "fmt"

// func Max(a, b, c int) int {
// 	if a > b {

// 	}
// }

func main() {
	var a int

	go func() {
		a = 10
	}()
	fmt.Println(a)
	// fmt.Println(u)
	// a, b, c := 10, 40, 30
	// math.Max(float64(a),)
	// if a > b && a > c {
	// 	max = a
	// } else if b > a && b > c {
	// 	max = b
	// } else if c > a && c > b {
	// 	max = c
	// }
	// fmt.Println(max)

	// x, y, z := 0, 0, 1
	// max := math.Max(float64(x), math.Max(float64(y), float64(z)))
	// max2 := 0
	// if x >= y && x >= z {
	// 	max2 = x
	// } else if y >= x && y >= z {
	// 	max2 = y
	// } else if z >= y && z >= x {
	// 	max2 = z
	// }
	// var a interface{}
	// a = 10
	// fmt.Println(a)
	// var x complex64 = 5 + 7i
}
