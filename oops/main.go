package main

import (
	"fmt"
	pk "myapp/oops/sq"
)

// main package
func main() {
	var a pk.Bank = &pk.OnlinePay{}
	fmt.Println(pk.Phonepay(a, 10, "NDNDN"))
}
