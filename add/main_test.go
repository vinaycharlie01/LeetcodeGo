package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// a := []int{10, 2, 4, 5, 7}
	// b := []int{10, 2, 4, 10, 14}

	result := smallestEvenMultiple(10)
	if result != 10 {
		t.Errorf("got %q, wanted %q", result, 10)
	}
	// // a := smallestEvenMultiple(10)
	// for i, v := range a {
	// 	result := smallestEvenMultiple(v)
	// 	require.Equal(t, result, b[i])
	// }

}
