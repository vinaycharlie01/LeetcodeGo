package main

import (
	"testing"
)

// func makeRandomInts(n int) []int {
// 	rand.Seed(10)
// 	ints := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		ints[i] = rand.Intn(n)
// 	}
// 	return ints
// }

const N = 20

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		resarr := MakeRandomInts(10)
		in := RandomInts(len(resarr))
		b.StopTimer()
		NumberOfEmployeesWhoMetTarget(resarr, in)
	}
}
