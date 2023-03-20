package main

import (
	"fmt"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	s := Pointer2(a)
	fmt.Println(s)
}
