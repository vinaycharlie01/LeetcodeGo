package random

import (
	"math/rand"
	"time"
)

func MakeRandomInts(n int) []int {
	rand.Seed(10)
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(n)
	}
	return ints
}

func RandomInts(n int) int {
	return rand.Intn(n)
}

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func MakeRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = Charset[rand.Intn(len(Charset))]
	}
	return string(result)
}
