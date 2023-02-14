package main

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	L := 0
	R := len(nums) - 1
	var a int
	for L <= R {
		if L == R {
			a += nums[L]
		} else {
			val := strconv.Itoa(nums[L]) + strconv.Itoa(nums[R])
			s, _ := strconv.Atoi(val)
			a += s
		}
		L++
		R--

	}
	return int64(a)
}
