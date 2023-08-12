package main

import (
	b1 "myapp/random"
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
		resarr := b1.MakeRandomString(10)
		// fmt.Println(resarr)
		b.StopTimer()
		FinalString(resarr)
	}
}
