package main

func Hash(arr []int, target int) []int {
	hashmap := make(map[int]int)
	var a []int
	for i, v := range arr {
		val, ok := hashmap[target-v]
		if ok {
			a = append(a, val, i)
		}
		hashmap[v] = i

	}
	return a
}

/*
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:w
Input: nums = [1,1]
Output: [2]
*/

func findDisappearedNumbers(nums []int) []int {
	hashmap := make(map[int]struct{})
	var a []int
	for _, v := range nums {
		hashmap[v] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		_, ok := hashmap[i]
		if !ok {
			a = append(a, i)
		}
	}
	return a
}

func vowelStrings(words []string, left int, right int) int {
	hashmap := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	count := 0
	for i := left; i <= right; i++ {
		_, ok := hashmap[words[i][0]]
		_, ok1 := hashmap[words[i][len(words[i])-1]]
		if ok && ok1 {
			count++
		}

	}
	return count

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
