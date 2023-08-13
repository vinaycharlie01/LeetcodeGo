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

const N = 10

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		resarr := b1.MakeRandomString(10)
		// fmt.Println(resarr)
		b.StopTimer()
		FinalString(resarr)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		resarr := b1.MakeRandomString(100)
		s := b1.RandomInts(10)
		// fmt.Println(resarr)
		b.StopTimer()
		ReverseStr(resarr, s)

	}
}
