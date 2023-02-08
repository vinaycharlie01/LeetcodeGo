package main

import "fmt"

func smallestEvenMultiple(n int) int {
	if n%2 == 1 {
		return n * 2
	} else {
		return n
	}

}
func main() {
	var n int
	fmt.Println("enter the number: ")
	fmt.Scanln(&n)
	fmt.Print(smallestEvenMultiple(n))
}
