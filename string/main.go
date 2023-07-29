package main

import (
	"fmt"
	"math/big"
)

func example() {
	a := 0
	defer func(a int) {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
		fmt.Println("Succesfully Completed", a)
	}(a)
	fmt.Println("Starting function")

	if a == 0 {
		panic("Something went wrong! is is Not Zero")
	}
	fmt.Println("This line will never be executed")
}

type Point struct {
	x, y int
}

func (p *Point) Move(dx, dy int) {
	p.x += dx
	p.y += dy
}
func main() {

	p := Point{x: 1, y: 2}
	p.Move(2, 3)

	// example()
	// a := []string{"hey", "aeo", "mu", "ooo", "artro"}
	// fmt.Println(vowelStrings(a, 1, 4))
	// Example using bytes
	// message := []byte("ಹಾಯ್ ನನ್ನಾ")
	// for i := 0; i < len(message); i++ {
	// 	fmt.Printf("%c\n", message[i])
	// }

	// // Example using runes
	// message2 := []rune("ಹಾಯ್ ನನ್ನಾ")
	// for i := 0; i < len(message2); i++ {
	// 	fmt.Printf("%k", message2[i])
	// }

	a := new(big.Int)
	_, _ = a.SetString("123455111111111111111111111111111111111111111111111111111112345511111111111111111111111111111111111111111111111111", 10)

	fmt.Println(a)
	rev := new(big.Int)

	number := big.NewInt(0)
	rem := big.NewInt(10)
	for a.Cmp(number) > 0 {
		dive := new(big.Int)
		dive.Mod(a, rem)
		rev.Mul(rev, rem)
		rev.Add(rev, dive)
		a.Div(a, rem)
	}

	fmt.Println(rev)
	// var a interface{}
	// b := 102
	// *a = b

	// a := Sub(10)
	// fmt.Println(a(10))

}
func hello() {

}

func Sub(a int) func(int) int {
	return func(i int) int {
		return a + i
	}
}
