package main

import (
	"fmt"
	"math/big"
)

func Factorial(n int) {
	temp := n
	rev := 0
	for temp > 0 {
		rem := temp % 10
		rev = rev*10 + rem
		temp /= 10
	}
	if rev == n {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}

}

func HCF(a, b int) {
	hcf := 1
	min := 0
	if a > b {
		min = b
	} else {
		min = a
	}
	for i := 2; i < min+1; i++ {
		if a%i == 0 && b%i == 0 {
			hcf = i
		}
	}
	fmt.Println(hcf)
}

func lcm(a, b int64) int64 {
	gcd := new(big.Int).GCD(nil, nil, big.NewInt(int64(a)), big.NewInt(int64(b)))
	return (a * b) / gcd.Int64()
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func work(ch chan struct{}) {
	<-ch
	fmt.Println("Recived")
	close(ch)
}

func main() {

	ch := make(chan struct{})

	go work(ch)
	ch <- struct{}{}

	// ch <- struct{}{}

	//<-ch

	fmt.Println("Recived main")
	// lcm(200, 100)
	// fmt.Println(gcd(200, 100))

	// HCF(10, 2)
	// a := 122
	// Factorial(a)
}
