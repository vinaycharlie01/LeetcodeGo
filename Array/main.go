package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func intersection(nums1 []int, nums2 []int) []int {
	Hashmap := make(map[int]int)
	for i, val := range nums2 {
		Hashmap[val] = i
	}
	hashmap2 := make(map[int]int)
	for _, val := range nums1 {
		_, ok := Hashmap[val]
		if ok {
			hashmap2[val]++
		}
	}
	var a []int
	for i, _ := range hashmap2 {
		a = append(a, i)
	}
	return a
}

func containsDuplicate(nums []int) bool {
	hashmap := make(map[int]int)
	for _, v := range nums {
		_, ok := hashmap[v]
		if ok {
			return true
		}
		hashmap[v]++
	}
	return false
}

func findKthLargest(nums []int, k int) int {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sort.Ints(nums)
	}()
	wg.Wait()
	return nums[len(nums)-k]
}

func wiggleSort(nums []int) {
	// mid := len(nums) / 2
	// l1 := nums[:mid]
	// sort.Sort(sort.Reverse(sort.IntSlice(l1)))
	// r1 := nums[mid:]
	// sort.Sort(sort.Reverse(sort.IntSlice(r1)))
	// fmt.Println(l1, r1)
	// L := 0
	// R := 0
	// var a []int
	// for L < len(l1) && R < len(r1) {
	// 	a = append(a, r1[R])
	// 	R++
	// 	a = append(a, l1[L])
	// 	L++
	// }
	// copy(nums, a)
	// fmt.Println(a)
}

func wiggleSort2(nums []int) {
	mid := len(nums) / 2
	l1 := nums[:mid]
	sort.Sort(sort.Reverse(sort.IntSlice(l1)))
	r1 := nums[mid:]
	sort.Sort(sort.Reverse(sort.IntSlice(r1)))
	fmt.Println(l1, r1)
	L := 0
	R := 0
	var a []int
	for L < len(l1) && R < len(r1) {
		a = append(a, r1[R])
		R++
		a = append(a, l1[L])
		L++
	}
	copy(nums, a)
	fmt.Println(a)
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

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
		if len(result) > 0 {
			break
		}
	}
	return result

}

func findMaxConsecutiveOnes(nums []int) {
	// hashmap := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		}
	}
	fmt.Println(count)
}

func searchRange(nums []int, target int) []int {
	hashmap := make(map[int]int)
	var a []int
	for i, v := range nums {
		hashmap[v] = i
		val, ok := hashmap[target]
		if ok {
			a = append(a, val)
		}

	}
	if len(a) == 0 {
		return []int{-1, -1}
	} else {
		return []int{a[0], a[len(a)-1]}
	}
}

func searchRange2(nums []int, target int) []int {
	L := 0
	var a []int
	for L < len(nums) {
		if nums[L] == target {
			a = append(a, L)
			break
		}
		L++
	}
	R := len(a) - 1
	for R < len(nums) {
		if nums[R] == target {
			a = append(a, R)
			break
		}
		R--
	}
	return a

}

func main() {
	// nums = [5,7,7,8,8,10], target = 8
	a := []int{5, 7, 7, 8, 8, 10}
	t := 8
	searchRange2(a, t)
}

// func main() {

// 	a := []int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3}
// 	// b := []int{1, 3}
// 	// wiggleSort2(a)
// 	// wiggleSort(b)

// }
