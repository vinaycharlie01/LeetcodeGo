package main

import (
	"sort"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var a bool = true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				sort.IntSlice(arr).Swap(j, j+1)
			}
		}
		if !a {
			break
		}
	}
}

func Merge(left []int, right []int) []int {
	L := 0
	R := 0
	var c []int
	for L < len(left) && R < len(right) {
		if left[L] < right[R] {
			c = append(c, left[L])
			L++
		} else {
			c = append(c, right[R])
			R++
		}
	}
	if L < len(left) {
		c = append(c, left[L:]...)
	}
	if R < len(right) {
		c = append(c, right[R:]...)
	}
	return c
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[mid:])
	right := MergeSort(arr[:mid])
	return Merge(left, right)
}

func thirdMax(nums []int) int {
	hashmap := make(map[int]int)
	for _, v := range nums {
		hashmap[v]++
	}
	var a []int
	for i, _ := range hashmap {
		a = append(a, i)
	}
	a = MergeSort(a)
	if len(a) < 3 {
		return a[len(a)-1]
	} else {
		return a[len(a)-3]
	}
}
func main() {
	a := []int{1, 3, 1}
	thirdMax(a)
	// a = MergeSort(a)

	//go BubbleSort(b)
	// fmt.Println(a)
}
