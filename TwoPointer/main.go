package main

import (
	"fmt"
	"sort"
	"strconv"
)

type pair struct {
	val   int
	index int
}

func twoSum(nums []int, target int) []int {
	hash := make([]pair, len(nums))
	for i := 0; i < len(nums); i++ {
		hash[i] = pair{nums[i], i}
	}
	fmt.Println(hash)
	sort.Slice(hash, func(i, j int) bool {
		return hash[i].val < hash[j].val
	})
	fmt.Println(hash)
	start := 0
	end := len(nums) - 1
	for start < end {
		val := hash[start].val + hash[end].val
		if val == target {
			break
		} else if val < target {
			start++
		} else if val > target {
			end--
		}

	}
	return []int{hash[start].index, hash[end].index}
	// fmt.Println(hash)
}

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

func main() {

	// nums = [2,7,11,15], target = 9
	nums := []int{7, 11, 15, 2}

	// sort.Ints(nums)

	// b := sort.IntSlice(nums).
	// fmt.Println(b)
	target := 9
	twoSum(nums, target)
	// TwoSumPointer(nums, target)
	// Hash(nums, target)
	// fmt.Println(twoSum(nums, target))
}
